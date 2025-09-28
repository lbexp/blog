package main

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type ContextKey string

type route struct {
	method       string
	path         *regexp.Regexp
	innerHandler http.HandlerFunc
	paramKeys    []string
}

type router struct {
	routes []route
}

func NewRouter() *router {
	return &router{routes: []route{}}
}

func buildContext(req *http.Request, paramKeys, paramValues []string) *http.Request {
	ctx := req.Context()

	for i := 0; i < len(paramKeys); i++ {
		ctx = context.WithValue(ctx, ContextKey(paramKeys[i]), paramValues[i])
	}

	return req.WithContext(ctx)
}

func (r *route) handler(w http.ResponseWriter, req *http.Request) {
	requestString := fmt.Sprint(req.Method, " ", req.URL)
	fmt.Println("Received ", requestString)
	r.innerHandler(w, req)
}

func (r *router) addRoute(method string, path string, handler http.HandlerFunc) {
	pathPattern := regexp.MustCompile(":([a-z]+)")
	matches := pathPattern.FindAllStringSubmatch(path, -1)
	paramKeys := []string{}

	if len(matches) > 0 {
		path = pathPattern.ReplaceAllLiteralString(path, "([^/]+)")

		for i := 0; i < len(matches); i++ {
			paramKeys = append(paramKeys, matches[i][1])
		}
	}

	route := route{
		method,
		regexp.MustCompile("^" + path + "$"),
		handler,
		paramKeys,
	}

	r.routes = append(r.routes, route)
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var allow []string

	for _, route := range r.routes {
		matches := route.path.FindStringSubmatch(req.URL.Path)

		if len(matches) > 0 {
			if req.Method != route.method {
				allow = append(allow, route.method)
				continue
			}

			context := buildContext(req, route.paramKeys, matches[1:])
			route.handler(w, context)
			return
		}

	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(w, req)
}

func (r *router) GET(path string, handler http.HandlerFunc) {
	r.addRoute("GET", path, handler)
}
