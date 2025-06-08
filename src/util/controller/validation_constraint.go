package controller_uitl

import (
	"reflect"
	"users_orchestrator/vo/resp/common"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindValidation[T any](ctx *gin.Context) (*T, *common.ValidationError) {
	var bind T
	err := ctx.ShouldBind(&bind)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			field, msg := getValidationMsg(validationErrors, &bind)
			validationErr := common.ValidationError{Field: field, Msg: msg}
			return nil, &validationErr
		}
		return nil, &common.ValidationError{Msg: "missing data"}
	}
	return &bind, nil
}

func getValidationMsg(validationErrors validator.ValidationErrors, bind any) (string, string) {
	obj := reflect.TypeOf(bind)
	for _, validationError := range validationErrors {
		if field, ok := obj.Elem().FieldByName(validationError.Field()); ok {
			return field.Tag.Get("form"), field.Tag.Get("msg")
		}
	}
	return "", ""
}
