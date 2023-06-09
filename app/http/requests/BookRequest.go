package requests

import (
	"github.com/gin-gonic/gin"
	"gorm-test/app/models"
	"gorm-test/utils"
	utilsValidator "gorm-test/utils/validator"
)

type bookCreateRequest struct {
	Title       string `json:"title" validate:"required,min=4,max=255" `
	Description string `json:"description" validate:"required" gpc:"required=Email Is Require"`
}

func BookCreateRequestHandler(ctx *gin.Context, bookModel *models.Book) bool {
	var bookCreateRequest bookCreateRequest
	err := ctx.ShouldBindJSON(&bookCreateRequest)
	if err != nil {
		return false
	}
	res, _ := utilsValidator.Validator(bookCreateRequest)
	if res != nil {
		utils.FailedResponse(ctx, "validation Error", 422, res)
		return false
	}
	response := utils.Responses{}
	response.MakeResponse(bookCreateRequest, &bookModel)

	return true
}
