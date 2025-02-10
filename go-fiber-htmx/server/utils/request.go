package utils

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

var snakeCaseRegexp = regexp.MustCompile(`([a-z0-9])([A-Z])`)

func ConvertToFiberMap(data interface{}) *fiber.Map {
	result := make(fiber.Map)

	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			result[ToSnakeCase(fmt.Sprintf("%d", i))] = val.Index(i).Interface()
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			strKey := fmt.Sprintf("%v", key.Interface()) // Convert key to string
			result[ToSnakeCase(strKey)] = val.MapIndex(key).Interface()
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			result[ToSnakeCase(field.Name)] = val.Field(i).Interface()
		}
	}

	return &result
}
