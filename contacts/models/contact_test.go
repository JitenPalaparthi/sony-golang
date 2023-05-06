package models_test

import (
	"contacts/models"
	"fmt"
	"testing"
)

func TestContactValidateSuccess(t *testing.T) {
	c1 := new(models.Contact)
	c1.Name = "Sony"
	c1.Email = "Sony@Sony.Com"

	err := c1.Validate()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestContactValidateFirstCase(t *testing.T) {
	c1 := new(models.Contact)
	c1.Name = "Sony"
	c1.Email = "Sony@Sony.Com"

	err := c1.Validate()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestContactValidateFailure(t *testing.T) {
	c1 := new(models.Contact)
	c1.Name = "Sony"
	//c1.Email = "Sony@Sony.Com"

	err := c1.Validate()
	if err == nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestContactToJson(t *testing.T) {
	c1 := new(models.Contact)
	c1.Name = "Sony"
	c1.Email = "Sony@Sony.Com"

	expected := `{"id":0,"name":"Sony","email":"Sony@Sony.Com","status":"","lastmodified":0}`
	str, err := c1.ToJson()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if str != expected {
		fmt.Println("Expected result is ", expected, ";Actual result is ", str)
		t.Fail()
	}
}
