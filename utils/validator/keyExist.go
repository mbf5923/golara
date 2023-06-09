package utilsValidator

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func KeyExist(input interface{}) (int, error) {

	if reflect.TypeOf(input).Kind().String() != "struct" {
		return -1, fmt.Errorf("validator value not supported, because %v is not struct", reflect.TypeOf(input).Kind().String())
	}

	received := make(map[string]interface{})
	var mapsArr []string

	stringify, err := json.Marshal(&input)

	if err != nil {
		return -1, err
	}

	if err := json.Unmarshal(stringify, &received); err != nil {
		return -1, err
	}

	for i := range received {
		mapsArr = append(mapsArr, i)
	}

	return len(mapsArr), nil
}
