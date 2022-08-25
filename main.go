package main

import (
	"firestore/middlewares"
	"firestore/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	port := "80"
	apiV1 := router.Group("/apiV1/")
	apiV1.Use(middlewares.Authenticate)
	routes.AuthRoutes(apiV1.Group("/auth/"))
	routes.ContentRoutes(apiV1.Group("/content/"))
	router.Run("0.0.0.0:"+port)
}