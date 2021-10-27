// Написать функцию, которая принимает на вход структуру in (struct или кастомную struct) и
// values map[string]interface{} (key - название поля структуры, которому нужно присвоить value этой мапы).
// Необходимо по значениям из мапы изменить входящую структуру in с помощью пакета reflect.
// Функция может возвращать только ошибку error. Написать к данной функции тесты (чем больше,
// тем лучше - зачтется в плюс).
package main

import (
	"fmt"
	"reflect"
)

func FillStruct(in interface{}, data map[string]interface{}) error {
	for key, value := range data {
		err := setField(in, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(o interface{}, name string, value interface{}) error {
	if o == nil {
		return fmt.Errorf("empty input object")
	}

	rVal := reflect.ValueOf(o)
	if rVal.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
	}

	if rVal.Kind() != reflect.Struct {
		return fmt.Errorf("object is not struct")
	}

	fieldVal := rVal.FieldByName(name)

	if !fieldVal.IsValid() {
		return fmt.Errorf("no such field: %s in o", name)
	}

	if !fieldVal.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}

	val := reflect.ValueOf(value)

	if fieldVal.Type() == val.Type() {
		fieldVal.Set(val)
		return nil
	}

	if m, ok := value.(map[string]interface{}); ok {
		if fieldVal.Kind() == reflect.Struct {
			return FillStruct(fieldVal.Addr().Interface(), m)
		}
		if fieldVal.Kind() == reflect.Ptr && fieldVal.Type().Elem().Kind() == reflect.Struct {
			if fieldVal.IsNil() {
				fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
			}
			return FillStruct(fieldVal.Interface(), m)
		}

	}
	return fmt.Errorf("provided value type didn't match o field type")
}
