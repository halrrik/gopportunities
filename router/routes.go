package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/halrrik/gopportunities/docs"
	"github.com/halrrik/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// func printMe() {
// 	print("me me me")
// }

func initializeRoutes(router *gin.Engine) {

	//initialize handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
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
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// router.Group("/api/")
}
