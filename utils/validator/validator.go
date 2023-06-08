package utilsValidator

import (
	"fmt"
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

func Validator(s interface{}) (interface{}, error) {
	if reflect.TypeOf(s).Kind().String() != "struct" {
		return nil, fmt.Errorf("validator value not supported, because %v is not struct", reflect.TypeOf(s).Kind().String())
	} else if res, err := KeyExist(s); err != nil || res == 0 {
		return nil, fmt.Errorf("validator value can't be empty struct %v", s)
	}

	val := validator.New()
	err := val.Struct(s)

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("fa_IR")

	if err := enTranslations.RegisterDefaultTranslations(val, trans); err != nil {
		return nil, err
	}

	if err == nil {
		return nil, err
	}

	return FormatError(err, trans, s)
}
