package helper

import "reflect"

func StructToMap(obj interface{}) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	if objType.Kind() != reflect.Struct {
		panic("obj is not a struct")
	}

	data := make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i).Interface()
		data[field.Name] = fieldValue
	}

	return data
}
