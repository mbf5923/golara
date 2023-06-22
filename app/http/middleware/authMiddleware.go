package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm-test/app/models"
	"gorm-test/utils"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")
		if authorizationHeader == "" {
			utils.FailedResponse(ctx, "Authorization is required for this endpoint", http.StatusForbidden, nil)
			defer ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		var token = strings.Split(authorizationHeader, " ")[1]
		var userModel models.User
		UserModel := models.NewUserModel()

		err := UserModel.GetUserByApiKey(&userModel, token)
		if err != nil || userModel.ID == 0 {
			utils.FailedResponse(ctx, "accessToken invalid or expired", http.StatusUnauthorized, nil)
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {
			ctx.Set("user", &userModel)
			ctx.Next()
		}
	}
}
