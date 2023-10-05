package handler

import (
	"github.com/halrrik/gopportunities/config"
	"gorm.io/gorm"
)

// func CreateOpeningHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "POST Opening texto texto",
// 	})
// }

// func ShowOpeningHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "GET Opening texto texto",
// 	})
// }

// func DeleteOpeningHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "delete Opening texto texto",
// 	})
// }

// func UpdateOpeningHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "put Opening texto texto",
// 	})
// }

// func ListOpeningHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "GET Opening texto texto",
// 	})
// }

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
