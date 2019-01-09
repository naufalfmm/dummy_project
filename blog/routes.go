package main

import "github.com/gin-gonic/gin"

type Route struct {
	Method string
	Path   string
	Handle gin.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		Index,
	},
	Route{
		"GET",
		"/posts",
		PostIndex,
	},
	Route{
		"GET",
		"/posts/:id",
		PostShow,
	},
	Route{
		"POST",
		"/posts",
		PostCreate,
	},
}
