package controllers

import "github.com/gin-gonic/gin"

func Routes() {
	router := gin.Default()
	router.POST("/", PostUrl)
	router.GET("/:hash", RedirectUrl)

	router.Run("localhost:8080")

	return
}
