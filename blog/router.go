package main

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()

	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.Handle)
	}

	return router
}
