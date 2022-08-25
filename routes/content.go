package routes

import (
	"firestore/controllers"

	"github.com/gin-gonic/gin"
)

func ContentRoutes(c *gin.RouterGroup) {
	c.POST("/email/", controllers.SendMessageToUser)
	c.GET("/email/", controllers.GetAllEmails)
}
