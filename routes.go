package main

import (
	"authtoken/controllers"

	"github.com/gin-gonic/gin"
)

func LoadAPIRoutes(r *gin.Engine) {
	// User
	userRoutes := r.Group("/user")
	{
		// Router doesn't need authentication
		userRoutes.GET("/generateToken", controllers.GenerateToken)
		userRoutes.POST("/login", controllers.LoginPOST)
	}
}
