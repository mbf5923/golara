package requests

import (
	"github.com/gin-gonic/gin"
	"gorm-test/app/models"
	"gorm-test/utils"
	utilsValidator "gorm-test/utils/validator"
)

type bookCreateRequest struct {
	Title       string  `json:"title" validate:"required,min=4,max=255" `
	Description string  `json:"description" validate:"required" gpc:"required=Description Is Require"`
	Price       *uint64 `json:"price" validate:"required"`
}

type bookGetByIdRequest struct {
	ID int `uri:"id" validate:"required"`
}

type bookUpdateRequest struct {
	Title       string  `json:"title,omitempty" validate:"omitempty,min=4,max=255"`
	Description string  `json:"description,omitempty" validate:"omitempty" `
	Price       *uint64 `json:"price,omitempty" validate:"omitempty"`
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

func BookGetByIdRequestHandler(ctx *gin.Context, bookModel *models.Book) bool {
	var bookGetByIdRequest bookGetByIdRequest
	err := ctx.ShouldBindUri(&bookGetByIdRequest)
	if err != nil {
		return false
	}
	res, _ := utilsValidator.Validator(bookGetByIdRequest)
	if res != nil {
		utils.FailedResponse(ctx, "validation Error", 422, res)
		return false
	}
	response := utils.Responses{}
	response.MakeResponse(bookGetByIdRequest, &bookModel)
	return true
}

func BookUpdateRequestHandler(ctx *gin.Context, bookModel *models.Book) bool {
	var bookUpdateRequest bookUpdateRequest
	err := ctx.ShouldBindJSON(&bookUpdateRequest)
	if err != nil {
		return false
	}
	res, _ := utilsValidator.Validator(bookUpdateRequest)
	if res != nil {
		utils.FailedResponse(ctx, "validation Error", 422, res)
		return false
	}
	response := utils.Responses{}
	response.MakeResponse(bookUpdateRequest, &bookModel)
	return true
}
