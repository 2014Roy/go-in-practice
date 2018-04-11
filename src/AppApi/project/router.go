package main

import (
	. "AppApi/project/apis"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexApi)
	router.POST("/person", AddPerson)
	router.GET("/person/:id", GetPerson)
	router.DELETE("/person", DeletePerson)

	return router
}
