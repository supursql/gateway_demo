package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type RealServer struct {
	Addr string
}

func (r *RealServer) Run() {
	log.Println("Starting httpserver at " + r.Addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.HelloHandler)
	mux.HandleFunc("/base/error", r.ErrorHandler)

	server := &http.Server{
		Addr:         r.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}

func (r *RealServer) HelloHandler(writer http.ResponseWriter, request *http.Request) {
	upath := fmt.Sprintf("http://%s%s\n", r.Addr, request.URL.Path)
	realIP := fmt.Sprintf("RemoteAddr=%s,X-Forwarded-For=%v,X-Real-Ip=%v\n", request.RemoteAddr, request.Header.Get("X-Forwarded-For"), request.Header.Get("X-Real-Ip"))
	io.WriteString(writer, upath)
	io.WriteString(writer, realIP)
}

func (r *RealServer) ErrorHandler(writer http.ResponseWriter, request *http.Request) {
	upath := "error handler"
	writer.WriteHeader(500)
	io.WriteString(writer, upath)
}

func main() {
	rs1 := &RealServer{Addr: "127.0.0.1:2003"}
	rs1.Run()

	rs2 := &RealServer{Addr: "127.0.0.1:2004"}
	rs2.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
