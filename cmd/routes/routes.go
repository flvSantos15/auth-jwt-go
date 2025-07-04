package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	// Setup routes
	authRoutes := router.Group("/api/auth")

	authRoutes.POST("/register", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	authRoutes.POST("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	authRoutes.POST("/refresh", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	authRoutes.POST("/logout", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})
}
