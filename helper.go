package kucoin

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// IntToString converts int64 to string.
func IntToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// ToJsonString converts any value to JSON string.
func ToJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func StructToMap(s interface{}) map[string]string {
	m := make(map[string]string)
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		if value.CanInterface() {
			tag := field.Tag.Get("json")
			if tag == "" {
				tag = field.Name
			}

			if !value.IsZero() {
				m[tag] = fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return m
}
