package main

import (
	"gateway_demo/proxy/load_balance"
	"gateway_demo/proxy/middleware"
	"gateway_demo/proxy/proxy"
	"log"
	"net/http"
)

var (
	addr = "127.0.0.1:2002"
)

func main() {
	coreFunc := func(c *middleware.SliceRouterContext) http.Handler {
		rb := load_balance.LoadBalanceFactory(load_balance.LbWeightRoundRobin)
		rb.Add("http://127.0.0.1:2003", "50")
		return proxy.NewLoadBalanceReverseProxy(c, rb)
	}

	log.Println("Starting httpserver at " + addr)

	sliceRouter := middleware.NewSliceRouter()
	sliceRouter.Group("/").Use(middleware.RateLimiter())

	routerHandler := middleware.NewSliceRouterHandler(coreFunc, sliceRouter)
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}
