package utilsValidator

import (
	"reflect"
	"regexp"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func FormatError(err error, trans ut.Translator, customMessage interface{}) (interface{}, error) {
	errRes := make(map[string][]map[string]interface{})
	var tags []string
	rt := reflect.TypeOf(customMessage)
	if rt.Kind() != reflect.Struct {
		rt = nil
	}
	for i, e := range err.(validator.ValidationErrors) {

		errResult := make(map[string]interface{})

		if rt != nil {
			f, _ := rt.FieldByName(e.StructField())
			v := strings.Split(f.Tag.Get("json"), ",")[0]
			errResult["param"] = v
		} else {
			errResult["param"] = e.StructField()
		}

		if _, ok := reflect.TypeOf(customMessage).Field(i).Tag.Lookup("gpc"); !ok {
			errResult["msg"] = e.Translate(trans)
		} else {
			structField, _ := reflect.TypeOf(customMessage).FieldByName(e.StructField())
			structTags := structField.Tag.Get("gpc")

			regexTag := regexp.MustCompile(`=+\w.*`)
			regexVal := regexp.MustCompile(`\w+=`)
			strArr := strings.Split(structTags, ",")
			tags = append(tags, MergeSlice(strArr)...)

			for j, v := range tags {
				if ok := regexTag.ReplaceAllString(tags[j], ""); ok == e.ActualTag() {
					errResult["msg"] = regexVal.ReplaceAllString(v, "")
					tags = append(tags, "")
				}
			}
		}

		errResult["tag"] = e.ActualTag()
		errRes["errors"] = append(errRes["errors"], errResult)
	}

	return errRes["errors"], nil
}
