// https://aprendagolang.com.br/2023/01/04/como-utilizar-tags-customizadas/
// https://medium.com/golangspec/tags-in-golang-3e5db0b8ef3e
// https://www.tutorialspoint.com/how-to-assign-default-value-for-struct-field-in-golang

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func SetDefaults(obj any) error {
	// Iterate over the fields struct using reflection
	// and set the default value for each field if the field is not provided
	// by the caller of the constructor function.
	for i := 0; i < reflect.TypeOf(obj).NumField(); i++ {
		field := reflect.TypeOf(obj).Field(i)

		if defaultVal, ok := field.Tag.Lookup("default"); ok {
			value := reflect.ValueOf(obj)
			fmt.Println(value)
			if err := setField(value.Field(i), defaultVal); err != nil {
				return err
			}
		} else {
			// Required variable
			fmt.Println(field.Name)
		}
	}
	return nil
}

func setField(field reflect.Value, defaultVal string) error {

	if !field.CanSet() {
		return fmt.Errorf("Can't set value\n")
	}
	fmt.Printf("Kind %s\n", field.Kind())
	switch field.Kind() {

	case reflect.Int:
		if val, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}
	case reflect.String:
		field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
	default:
		return fmt.Errorf("Cannot set value %s", field.Kind())
	}

	return nil
}
