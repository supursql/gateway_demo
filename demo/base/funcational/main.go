package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type HandlerFunc func(writer http.ResponseWriter, request *http.Request)

func (f HandlerFunc) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	f(writer, request)
}

func main() {
	hf := HandlerFunc(HelloHandler)

	/**
	func NewRecorder() *ResponseRecorder
	NewRecorder返回一个初始化了的ResponseRecorder.

	type ResponseRecorder struct {
	    Code      int           // HTTP回复的状态码
	    HeaderMap http.Header   // HTTP回复的头域
	    Body      *bytes.Buffer // 如非nil，会将Write方法写入的数据写入bytes.Buffer
	    Flushed   bool
	    // 内含隐藏或非导出字段
	}
	ResponseRecorder实现了http.ResponseWriter接口，它记录了其修改，用于之后的检查。
	*/
	resp := httptest.NewRecorder()

	/**
	func NewRequest(method, target string, body io.Reader) *http.Request

	NewRequest 返回一个新的服务器访问请求，这个请求可以传递给 http.Handler 以便进行测试。
	target 参数的值为 RFC 7230 中提到的“请求目标”（request-target)： 它可以是一个路径或者一个绝对 URL。如果 target 是一个绝对 URL，那么 URL 中的主机名（host name）将被使用；否则主机名将为 example.com。
	当 target 的模式为 https 时，TLS 字段的值将被设置为一个非 nil 的随意值（dummy value）。
	Request.Proto 总是为 HTTP/1.1。
	如果 method 参数的值为空， 那么使用 GET 方法作为默认值。
	body 参数的值可以为 nil；另一方面，如果 body 参数的值为 *bytes.Reader 类型、 *strings.Reader 类型或者 *bytes.Buffer 类型，那么 Request.ContentLength 将被设置。
	为了使用的方便，NewRequest 将在 panic 可以被接受的情况下，使用 panic 代替错误。
	如果你想要生成的不是服务器访问请求，而是一个客户端 HTTP 请求，那么请使用 net/http 包中的 NewRequest 函数。
	*/
	req := httptest.NewRequest("GET", "/", bytes.NewBuffer([]byte("test")))

	hf.ServeHTTP(resp, req)
	bts, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bts))
}

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}
