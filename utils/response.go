package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Responses struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

type Meta struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
}

func (responses *Responses) SuccessResponse(ctx *gin.Context, Data interface{}, Message string, StatusCode int) {

	jsonResponse := Responses{
		Data: Data,
		Meta: Meta{
			Status:  true,
			Message: Message,
			Error:   nil,
		},
	}

	ctx.JSON(StatusCode, jsonResponse)
	defer ctx.AbortWithStatus(StatusCode)

}

func FailedResponse(ctx *gin.Context, message string, statusCode int, Error interface{}) {

	response := Responses{
		Data: nil,
		Meta: Meta{
			Status:  false,
			Message: message,
			Error:   Error,
		},
	}

	ctx.JSON(statusCode, response)
	defer ctx.AbortWithStatus(statusCode)
}

func (responses *Responses) MakeResponse(sourceData interface{}, targetData any) *Responses {
	js, marshalErr := json.Marshal(sourceData)
	if marshalErr != nil {
		return nil
	}
	unmarshalErr := json.Unmarshal(js, &targetData)
	if unmarshalErr != nil {
		return nil
	}
	return responses
}
