package main

import (
	"encoding/json"
	"reflect"
)

type Test struct {
	String string `json:"this-is-string-field"`
	Number int    `json:"this-is-number-field"`
}

func main() {
}

// magical code by bxcodec -> https://gist.github.com/bxcodec/c2a25cfc75f6b21a0492951706bc80b8
func structToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = structToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}

func jsonToMap(item interface{}) map[string]interface{} {
	m, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	var s map[string]interface{}
	err = json.Unmarshal(m, &s)
	if err != nil {
		panic(err)
	}
	return s
}
