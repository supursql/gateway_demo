package main

import ()

var addr = "127.0.0.1:2002"

func main() {
	//reverseProxy := func(c *middleware.SliceRouterContext) http.Handler {
	//	rs1 := "http://127.0.0.1:2003/base"
	//	url1, err1 := url.Parse(rs1)
	//	if err1 != nil {
	//		log.Println(err1)
	//	}
	//
	//	rs2 := "http://127.0.0.1:2004/base"
	//	url2, err2 := url.Parse(rs2)
	//	if err2 != nil {
	//		log.Println(err2)
	//	}
	//
	//	urls := []*url.URL{url1, url2}
	//	//return proxy.NewMultipleHostsReverseProxy(c, urls)
	//}
	//log.Println("Starting httpserver at " + addr)
}
