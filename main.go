package main

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Initialize Logs
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	// initialize gin
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Add our api routes
	LoadAPIRoutes(router)

	// Enable CORS
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origins, Authorizations, Content-Type, X-Auth-Token",
		MaxAge:         60 * time.Second,
	}))

	router.Run(":4000")
}
