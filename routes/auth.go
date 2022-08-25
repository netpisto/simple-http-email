package routes

import (
	"firestore/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/sign_up/", controllers.SignUp)
	r.POST("/sign_in/", controllers.SignIn)
}
