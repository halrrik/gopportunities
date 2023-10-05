package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halrrik/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"msg": "POST Opening texto texto",
	// })

	// request := struct {
	// 	Role string `json:"role"`
	// }{}

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	// logger.Infof("*request recebida  %+v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation erro : %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creando opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")

		return
	}
	sendSuccess(ctx, "create-opening", opening)

}
