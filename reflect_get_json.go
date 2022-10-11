package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func test() {
	m1 := M1{
		name:      "123",
		createdAt: 1,
		image:     "sxxxx",
	}
	fields := []string{"name", "image", "createdAt"}
	result, err := JsonByFields(m1, fields)
	fmt.Println(result, err)
}

type Model interface {
	Get()
}

type M1 struct {
	name      string
	createdAt int
	image     string
}

func (m M1) Get() {

}

type returnJson struct {
}

func JsonByFields(model Model, fields []string) (string, error) {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fieldNum := t.NumField()
	fieldsMap := make(map[string]interface{}, len(fields))
	if t.Kind() != reflect.Struct {
		return "", errors.New("类型错误")
	}
	for _, v := range fields {
		fieldsMap[v] = nil
	}
	values := reflect.ValueOf(model)
	var i int
	for i = 0; i < fieldNum; i += 1 {
		if _, ok := fieldsMap[t.Field(i).Name]; ok {
			value := values.FieldByName(t.Field(i).Name)
			var trueValue interface{}
			switch value.Kind() {
			case reflect.Bool:
				trueValue = value.Bool()
			case reflect.String:
				trueValue = value.String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				trueValue = value.Int()
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				trueValue = value.Uint()
			case reflect.Float32, reflect.Float64:
				trueValue = value.Float()
			case reflect.Complex64, reflect.Complex128:
				trueValue = value.Complex()
			case reflect.Pointer, reflect.UnsafePointer:
				trueValue = value.Pointer()
			}
			fieldsMap[t.Field(i).Name] = trueValue
		}
	}
	fieldsJson, err := json.Marshal(fieldsMap)
	if err != nil {
		return "", err
	}
	return string(fieldsJson), nil
}
