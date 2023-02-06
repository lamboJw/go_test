package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func Test() {
	m1 := &M1{
		BaseModel: BaseModel{},
		Name:      "123",
		CreatedAt: 1,
		Image:     "sxxxx",
	}
	fields := []string{"Name", "Image"}
	result, err := m1.JsonByFields(m1, fields)
	fmt.Println(result, err)
}

type Model interface {
	Get()
	JsonByFields(model Model, fields []string) (string, error)
}

type M1 struct {
	BaseModel
	Name      string
	CreatedAt int    `json:"created_at"`
	Image     string `json:"image"`
}

func (m M1) Get() {

}

type BaseModel struct {
}

func (baseModel *BaseModel) JsonByFields(model Model, fields []string) (string, error) {
	t := reflect.TypeOf(model).Elem()
	fieldNum := t.NumField()
	fieldsMap := make(map[string]interface{}, len(fields))
	if t.Kind() != reflect.Struct {
		return "", errors.New("类型错误")
	}
	for _, v := range fields {
		fieldsMap[v] = nil
	}
	jsonData := make(map[string]interface{}, len(fields))
	values := reflect.ValueOf(model).Elem()
	var i int
	for i = 0; i < fieldNum; i += 1 {
		fieldName := t.Field(i).Name
		tagName := ""
		if tag, ok := t.Field(i).Tag.Lookup("json"); ok {
			tagName = tag
		} else {
			tagName = fieldName
		}
		if _, ok := fieldsMap[fieldName]; ok {
			value := values.Field(i)
			if value.CanSet() { //已导出的属性，直接用Interface()方法获取值
				jsonData[tagName] = value.Interface()
			} else {
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
				}
				jsonData[tagName] = trueValue
			}
		}
	}
	fieldsJson, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}
	return string(fieldsJson), nil
}
