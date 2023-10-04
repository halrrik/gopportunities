package router

import (
	"github.com/gin-gonic/gin"
	"github.com/halrrik/gopportunities/handler"
)

func printMe() {
	print("me me me")
}

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		//mostrar
		v1.GET("/opening", handler.ShowOpeningHandler)
		//POSTAR OPENING
		v1.POST("/opening", handler.CreateOpeningHandler)
		//DELETAR
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		//UPDATE
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		//listar todo
		v1.GET("/openings", handler.ListOpeningHandler)
	}

	// router.Group("/api/")
}
