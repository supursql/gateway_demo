package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {

	/**
	type Transport struct {
	    // Proxy指定一个对给定请求返回代理的函数。
	    // 如果该函数返回了非nil的错误值，请求的执行就会中断并返回该错误。
	    // 如果Proxy为nil或返回nil的*URL置，将不使用代理。
	    Proxy func(*Request) (*url.URL, error)
	    // Dial指定创建TCP连接的拨号函数。如果Dial为nil，会使用net.Dial。
	    Dial func(network, addr string) (net.Conn, error)
	    // TLSClientConfig指定用于tls.Client的TLS配置信息。
	    // 如果该字段为nil，会使用默认的配置信息。
	    TLSClientConfig *tls.Config
	    // TLSHandshakeTimeout指定等待TLS握手完成的最长时间。零值表示不设置超时。
	    TLSHandshakeTimeout time.Duration
	    // 如果DisableKeepAlives为真，会禁止不同HTTP请求之间TCP连接的重用。
	    DisableKeepAlives bool
	    // 如果DisableCompression为真，会禁止Transport在请求中没有Accept-Encoding头时，
	    // 主动添加"Accept-Encoding: gzip"头，以获取压缩数据。
	    // 如果Transport自己请求gzip并得到了压缩后的回复，它会主动解压缩回复的主体。
	    // 但如果用户显式的请求gzip压缩数据，Transport是不会主动解压缩的。
	    DisableCompression bool
	    // 如果MaxIdleConnsPerHost!=0，会控制每个主机下的最大闲置连接。
	    // 如果MaxIdleConnsPerHost==0，会使用DefaultMaxIdleConnsPerHost。
	    MaxIdleConnsPerHost int
	    // ResponseHeaderTimeout指定在发送完请求（包括其可能的主体）之后，
	    // 等待接收服务端的回复的头域的最大时间。零值表示不设置超时。
	    // 该时间不包括获取回复主体的时间。
	    ResponseHeaderTimeout time.Duration
	    // 内含隐藏或非导出字段
	}
	Transport类型实现了RoundTripper接口，支持http、https和http/https代理。Transport类型可以缓存连接以在未来重用。

	*/
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, //连接超时
			KeepAlive: 30 * time.Second, //长连接超时时间
		}).DialContext,
		MaxIdleConns:          100,              //最大空闲连接数
		IdleConnTimeout:       90 * time.Second, //空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, //tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,
	}
	//创建客户端

	/**
	type Client struct {
	    // Transport指定执行独立、单次HTTP请求的机制。
	    // 如果Transport为nil，则使用DefaultTransport。
	    Transport RoundTripper
	    // CheckRedirect指定处理重定向的策略。
	    // 如果CheckRedirect不为nil，客户端会在执行重定向之前调用本函数字段。
	    // 参数req和via是将要执行的请求和已经执行的请求（切片，越新的请求越靠后）。
	    // 如果CheckRedirect返回一个错误，本类型的Get方法不会发送请求req，
	    // 而是返回之前得到的最后一个回复和该错误。（包装进url.Error类型里）
	    //
	    // 如果CheckRedirect为nil，会采用默认策略：连续10此请求后停止。
	    CheckRedirect func(req *Request, via []*Request) error
	    // Jar指定cookie管理器。
	    // 如果Jar为nil，请求中不会发送cookie，回复中的cookie会被忽略。
	    Jar CookieJar
	    // Timeout指定本类型的值执行请求的时间限制。
	    // 该超时限制包括连接时间、重定向和读取回复主体的时间。
	    // 计时器会在Head、Get、Post或Do方法返回后继续运作并在超时后中断回复主体的读取。
	    //
	    // Timeout为零值表示不设置超时。
	    //
	    // Client实例的Transport字段必须支持CancelRequest方法，
	    // 否则Client会在试图用Head、Get、Post或Do方法执行请求时返回错误。
	    // 本类型的Transport字段默认值（DefaultTransport）支持CancelRequest方法。
	    Timeout time.Duration
	}
	*/
	client := &http.Client{
		Timeout:   time.Second * 30, //请求超时时间
		Transport: transport,
	}

	resp, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bds, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bds))
}
