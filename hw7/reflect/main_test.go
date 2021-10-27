package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Contacts struct {
	Email string
	Phone string
}

type Person struct {
	Name     string
	Age      int64
	Contacts Contacts
}

func buildData() map[string]interface{} {
	contacts := make(map[string]interface{})
	contacts["Phone"] = "79988007766"
	contacts["Email"] = "genre@gmail.com"

	data := make(map[string]interface{})
	data["Name"] = "Genre"
	data["Age"] = int64(50)
	data["Contacts"] = contacts

	return data
}

func TestFillStruct(t *testing.T) {
	data := buildData()
	valid := &Person{
		Name: "Genre",
		Age:  50,
		Contacts: Contacts{
			Phone: "79988007766",
			Email: "genre@gmail.com",
		},
	}
	result := &Person{}
	err := FillStruct(result, data)
	if err != nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, valid)
	}
	if !reflect.DeepEqual(result, valid) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", result, valid)
	}
}

func TestFieldNotExists(t *testing.T) {
	data := buildData()
	result := &Person{}

	data["Invalid"] = "Invalid"
	validError := fmt.Errorf("no such field: %s in o", "Invalid")
	err := FillStruct(result, data)
	if err == nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
	if !reflect.DeepEqual(err, validError) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
}
func TestInvalidType(t *testing.T) {
	data := buildData()
	result := &Person{}

	validError := fmt.Errorf("provided value type didn't match o field type")
	data["Age"] = float64(3.1415926)
	err := FillStruct(result, data)
	if err == nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
	if !reflect.DeepEqual(err, validError) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
}
