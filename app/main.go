package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine


func main(){


	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	// router.GET("/", func(c*gin.Context){

	// 	c.HTML(
	// 		http.StatusOK,
	// 		"index.html",
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)

	// })
	intializeRoutes()
	router.Run()
}