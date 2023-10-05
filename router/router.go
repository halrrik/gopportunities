package router

//usa la ultima palavra para nomear o import
import "github.com/gin-gonic/gin"

//para exportar pacotes e func TEM QUE INICIAR COM LETRA MAISCUULA
func Initialize() {

	//todo que esta no package es accesivel no pacage entero, se fuera
	//fuera del package usaria router.printMe()
	// printMe()

	//inicializa o router utilizando las confiuracoes defauilt do gin
	// trocamos r que normalmente se usa por router para quedar mas intuitivo
	router := gin.Default()
	//define una ruta
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	initializeRoutes(router)

	//rodando
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
