package middleware

import (
	"context"
	"math"
	"net/http"
	"strings"
)

const abortIndex int8 = math.MaxInt8 / 2 //最多63个中间件

type HandlerFunc func(*SliceRouterContext)

//router结构体
type SliceRouter struct {
	groups []*SliceGroup
}

//group结构体
type SliceGroup struct {
	*SliceRouter
	path     string
	handlers []HandlerFunc
}

//router上下文
type SliceRouterContext struct {
	Rw  http.ResponseWriter
	Req *http.Request
	Ctx context.Context
	*SliceGroup
	index int8
}

type SliceRouterHandler struct {
	coreFunc func(*SliceRouterContext) http.Handler
	router   *SliceRouter
}

func (g *SliceRouter) Group(path string) *SliceGroup {
	return &SliceGroup{
		SliceRouter: g,
		path:        path,
	}
}

func (g *SliceGroup) Use(middlewares ...HandlerFunc) *SliceGroup {
	g.handlers = append(g.handlers, middlewares...)
	existsFlag := false
	for _, oldGroup := range g.SliceRouter.groups {
		if oldGroup == g {
			existsFlag = true
		}
	}

	if !existsFlag {
		g.SliceRouter.groups = append(g.SliceRouter.groups, g)
	}

	return g
}

func (w *SliceRouterHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := newSliceRouterContext(rw, req, w.router)
	if w.coreFunc != nil {
		c.handlers = append(c.handlers, func(routerContext *SliceRouterContext) {
			w.coreFunc(c).ServeHTTP(rw, req)
		})
	}
	c.Reset()
	c.Next()
}

//重制回调
func (c *SliceRouterContext) Reset() {
	c.index = -1
}

//从最先加入的中间件开始回调
func (c *SliceRouterContext) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *SliceRouterContext) Abort() {
	c.index = abortIndex
}

func (c *SliceRouterContext) IsAbort() bool {
	return c.index >= abortIndex
}

func (c *SliceRouterContext) Get(key interface{}) interface{} {
	return c.Ctx.Value(key)
}

func (c *SliceRouterContext) Set(key, val interface{}) {
	c.Ctx = context.WithValue(c.Ctx, key, val)
}

func newSliceRouterContext(rw http.ResponseWriter, req *http.Request, r *SliceRouter) *SliceRouterContext {
	newSliceGroup := &SliceGroup{}

	//最长url前缀匹配
	matchUrlLen := 0
	for _, group := range r.groups {
		if strings.HasPrefix(req.RequestURI, group.path) {
			pathLen := len(group.path)
			if pathLen > matchUrlLen {
				matchUrlLen = pathLen
				*newSliceGroup = *group
			}
		}
	}

	c := &SliceRouterContext{
		Rw:         rw,
		Req:        req,
		Ctx:        req.Context(),
		SliceGroup: newSliceGroup,
	}
	c.Reset()
	return c
}

func NewSliceRouterHandler(coreFunc func(*SliceRouterContext) http.Handler, router *SliceRouter) *SliceRouterHandler {
	return &SliceRouterHandler{
		coreFunc: coreFunc,
		router:   router,
	}
}

func NewSliceRouter() *SliceRouter {
	return &SliceRouter{}
}
