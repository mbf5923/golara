package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm-test/app/http/requests"
	"gorm-test/app/http/resources"
	"gorm-test/app/models"
	"gorm-test/utils"
	"gorm.io/gorm"
	"net/http"
)

type UserRepo struct {
	UserModel *models.UserModel
}

func New() *UserRepo {
	return &UserRepo{
		UserModel: models.NewUserModel(),
	}
}

func (repository *UserRepo) CreateUser(ctx *gin.Context) {
	var userModel models.User
	if !requests.UserRegisterRequestHandler(ctx, &userModel) {
		utils.FailedResponse(ctx, "Validation Error", 422, nil)
		return
	}
	// make password md5 hash
	userModel.Password = utils.Md5Hash(userModel.Password)
	err := repository.UserModel.CreateUser(&userModel)
	if err != nil {
		utils.FailedResponse(ctx, "Create User Error", 500, err)
		return
	}
	var userResource resources.UserResource
	response := utils.Responses{}
	response.MakeResponse(userModel, &userResource).SuccessResponse(ctx, userResource, "ok", 200)

}

func (repository *UserRepo) GetUsers(ctx *gin.Context) {
	var user []models.User
	err := repository.UserModel.GetUsers(&user)
	if err != nil {
		utils.FailedResponse(ctx, "get users error", 500, err)
		return
	}
	var userResource []resources.UserResource
	response := utils.Responses{}
	response.MakeResponse(user, &userResource).SuccessResponse(ctx, userResource, "ok", 200)
}

func (repository *UserRepo) GetUser(ctx *gin.Context) {
	var user models.User
	requests.UserGetByIdRequestHandler(ctx, &user)
	err := repository.UserModel.GetUser(&user, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailedResponse(ctx, "Not Found", http.StatusNotFound, nil)
			return
		}
		utils.FailedResponse(ctx, "Server Error", http.StatusInternalServerError, err)
		return
	}
	var userResource resources.UserResource
	response := utils.Responses{}
	response.MakeResponse(user, &userResource).SuccessResponse(ctx, userResource, "ok", 200)
}
func (repository *UserRepo) UpdateUser(ctx *gin.Context) {
	var user models.User
	requests.UserGetByIdRequestHandler(ctx, &user)
	err := repository.UserModel.GetUser(&user, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailedResponse(ctx, "Not Found", http.StatusNotFound, nil)
			return
		}

		utils.FailedResponse(ctx, "Server Error", http.StatusInternalServerError, err)
		return
	}
	requests.UserUpdateRequestHandler(ctx, &user)
	err = repository.UserModel.UpdateUser(&user)
	if err != nil {
		utils.FailedResponse(ctx, "Server Error", http.StatusInternalServerError, err)
		return
	}
	var userResource resources.UserResource
	response := utils.Responses{}
	response.MakeResponse(user, &userResource).SuccessResponse(ctx, userResource, "ok", 200)
}

func (repository *UserRepo) DeleteUser(ctx *gin.Context) {
	var user models.User
	requests.UserGetByIdRequestHandler(ctx, &user)
	err := repository.UserModel.DeleteUser(&user, user.ID)
	if err != nil {
		utils.FailedResponse(ctx, "Server Error", http.StatusInternalServerError, err)
		return
	}
	response := utils.Responses{}
	response.SuccessResponse(ctx, nil, "ok", 200)
}

func (repository *UserRepo) Login(ctx *gin.Context) {
	var user models.User
	requests.UserLoginRequestHandler(ctx, &user)
	//generate api key
	apiKey := utils.GenerateRandomString()

	err := repository.UserModel.LoginUser(&user, user.Email, utils.Md5Hash(user.Password))
	if err != nil {
		utils.FailedResponse(ctx, "Email or Password is wrong", http.StatusUnauthorized, err)
		return
	}

	if user.ID == 0 {
		utils.FailedResponse(ctx, "Email or Password is wrong", http.StatusUnauthorized, err)
		return
	}
	user.ApiKey = &apiKey
	err = repository.UserModel.UpdateApiKey(&user)

	if err != nil {
		utils.FailedResponse(ctx, "Server Error", http.StatusInternalServerError, err)
		return
	}
	response := utils.Responses{}
	data := map[string]string{
		"api_key": apiKey,
	}
	response.SuccessResponse(ctx, data, "ok", 200)

}
