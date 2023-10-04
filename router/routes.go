package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func printMe() {
	print("me me me")
}

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		//mostrar
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening texto texto",
			})
		})
		//POSTAR OPENING
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening texto texto",
			})
		})
		//DELETAR
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening texto texto",
			})
		})
		//UPDATE
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening texto texto",
			})
		})
		//listar todo
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening texto texto",
			})
		})
	}

	// router.Group("/api/")
}
