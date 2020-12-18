package main

import (
	"log"
	"net/http"
	"time"
)

var (
	Addr = ":1210"
)

func main() {

	/**
	func NewServeMux() *ServeMux

	NewServeMux创建并返回一个新的*ServeMux

	type ServeMux struct {
	    // 内含隐藏或非导出字段
	}

	ServeMux类型是HTTP请求的多路转接器。它会将每一个接收的请求的URL与一个注册模式的列表进行匹配，并调用和URL最匹配的模式的处理器。
	模式是固定的、由根开始的路径，如"/favicon.ico"，或由根开始的子树，如"/images/"（注意结尾的斜杠）。较长的模式优先于较短的模式，因此如果模式"/images/"和"/images/thumbnails/"都注册了处理器，后一个处理器会用于路径以"/images/thumbnails/"开始的请求，前一个处理器会接收到其余的路径在"/images/"子树下的请求。
	注意，因为以斜杠结尾的模式代表一个由根开始的子树，模式"/"会匹配所有的未被其他注册的模式匹配的路径，而不仅仅是路径"/"。
	模式也能（可选地）以主机名开始，表示只匹配该主机上的路径。指定主机的模式优先于一般的模式，因此一个注册了两个模式"/codesearch"和"codesearch.google.com/"的处理器不会接管目标为"http://www.google.com/"的请求。
	ServeMux还会注意到请求的URL路径的无害化，将任何路径中包含"."或".."元素的请求重定向到等价的没有这两种元素的URL。（参见path.Clean函数）
	*/
	mux := http.NewServeMux()

	/**
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	HandleFunc注册一个处理器函数handler和对应的模式pattern。
	*/
	mux.HandleFunc("/bye", sayBye)

	server := http.Server{
		Addr:         Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 3,
	}

	log.Println("Starting httpserver at" + Addr)
	log.Fatal(server.ListenAndServe())
}

func sayBye(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(1 * time.Second)
	writer.Write([]byte("bye bye, this is httpServer"))
}
