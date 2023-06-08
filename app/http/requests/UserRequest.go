package requests

import (
	"github.com/gin-gonic/gin"
	"gorm-test/app/models"
	"gorm-test/utils"
	utilsValidator "gorm-test/utils/validator"
)

type userRegisterRequest struct {
	Name     string `json:"name" validate:"required,min=4,max=15" `
	Email    string `json:"email" validate:"required,email" gpc:"required=Email is require"`
	Password string `json:"password" validate:"required,min=4,max=15" gpc:"required=Password Is Require,min=minimum char for password is 4"`
}

type userGetByIdRequest struct {
	ID int `uri:"id" validate:"required"`
}

type userUpdateRequest struct {
	Name     string `json:"name" validate:"min=4,max=15" `
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=4,max=15" gpc:"min=minimum char for password is 4"`
}

func UserRegisterRequestHandler(ctx *gin.Context, userModel *models.User) bool {
	var userRegisterRequest userRegisterRequest
	err := ctx.ShouldBindJSON(&userRegisterRequest)
	if err != nil {
		return false
	}
	res, _ := utilsValidator.Validator(userRegisterRequest)
	if res != nil {
		utils.FailedResponse(ctx, "validation Error", 422, res)
		return false
	}
	response := utils.Responses{}
	_ = response.MakeResponse(userRegisterRequest, &userModel)

	return true
}

func UserGetByIdRequestHandler(ctx *gin.Context, userModel *models.User) bool {
	var userGetByIdRequest userGetByIdRequest
	err := ctx.ShouldBindUri(&userGetByIdRequest)
	if err != nil {
		return false
	}
	res, _ := utilsValidator.Validator(userGetByIdRequest)
	if res != nil {
		utils.FailedResponse(ctx, "validation Error", 422, res)
		return false
	}
	response := utils.Responses{}
	_ = response.MakeResponse(userGetByIdRequest, &userModel)

	return true
}

func UserUpdateRequestHandler(ctx *gin.Context, userModel *models.User) bool {
	var userUpdateRequest userUpdateRequest
	err := ctx.ShouldBind(&userUpdateRequest)
	if err != nil {
		return false
	}
	res, _ := utilsValidator.Validator(userUpdateRequest)
	if res != nil {
		utils.FailedResponse(ctx, "validation Error", 422, res)
		return false
	}
	response := utils.Responses{}
	_ = response.MakeResponse(userUpdateRequest, &userModel)

	return true
}
