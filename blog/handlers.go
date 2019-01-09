package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to blog service</h1>")
}

func PostIndex(c *gin.Context) {
	c.Set("Content-Type", "application/json; charset=UTF-8")
	// c.Status(http.StatusCreated)

	posts := FindAll()

	c.JSON(http.StatusOK, posts)

	// if err := json.NewEncoder(w).Encode(posts); err != nil {
	// 	panic(err)
	// }
}

func PostShow(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	HandleError(err)

	post := FindPost(id)

	c.JSON(http.StatusOK, post)

	// if err := json.NewEncoder(w).Encode(post); err != nil {
	// 	panic(err)
	// }
}

func PostCreate(c *gin.Context) {
	var post Post

	err := c.Bind(&post)
	if err != nil {
		c.Set("Content-Type", "application/json; charset=UTF-8")
		c.Status(http.StatusUnprocessableEntity)
	}

	CreatePost(post)
	c.Set("Content-Type", "application/json; charset=UTF-8")
	c.Status(http.StatusCreated)

	// body, err := ioutil.ReadAll(io.LimitReader(c.Body, 1048576))
}
