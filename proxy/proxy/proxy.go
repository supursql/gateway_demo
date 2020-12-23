package proxy

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"gateway_demo/proxy/middleware"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

func NewMultipleHostsReverseProxy(c *middleware.SliceRouterContext, targets []*url.URL) *httputil.ReverseProxy {
	//请求协调者
	director := func(req *http.Request) {

		targetIndex := rand.Intn(len(targets))
		target := targets[targetIndex]
		targetQuery := target.RawQuery

		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path, req.URL.RawPath = joinURLPath(target, req.URL)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

	modifyFunc := func(res *http.Response) error {
		if strings.Contains(res.Header.Get("Connection"), "Upgrade") {
			return nil
		}
		var payload []byte
		var readErr error

		if strings.Contains(res.Header.Get("Content-Encoding"), "gzip") {
			gr, err := gzip.NewReader(res.Body)
			if err != nil {
				return err
			}
			payload, readErr = ioutil.ReadAll(gr)
			res.Header.Del("Content-Encoding")
		} else {
			payload, readErr = ioutil.ReadAll(res.Body)
		}

		if readErr != nil {
			return readErr
		}

		if res.StatusCode != 200 {
			payload = []byte("StatusCode error:" + string(payload))
		}
		c.Set("status_code", res.StatusCode)
		c.Set("payload", payload)
		res.Body = ioutil.NopCloser(bytes.NewBuffer(payload))
		res.ContentLength = int64(len(payload))
		res.Header.Set("Content-Length", strconv.FormatInt(int64(len(payload)), 10))

		return nil
	}

	errFunc := func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Println(err)
	}
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc, Transport: transport, ErrorHandler: errFunc}
}

func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
