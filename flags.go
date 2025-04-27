package main

import (
	"flag"
	"fmt"
	"reflect"
)

func BindFlags(cfg any) {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := range v.NumField() {
		field := v.Field(i)
		fieldType := t.Field(i)

		flagName := fieldType.Tag.Get("flag")
		help := fieldType.Tag.Get("help")
		defaultValue := fieldType.Tag.Get("default")

		if flagName == "" {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			flag.StringVar(field.Addr().Interface().(*string), flagName, defaultValue, help)
		case reflect.Int:
			defaultInt := 0
			if defaultValue != "" {
				fmt.Sscanf(defaultValue, "%d", &defaultInt)
			}
			flag.IntVar(field.Addr().Interface().(*int), flagName, defaultInt, help)
		case reflect.Bool:
			defaultBool := false
			if defaultValue == "true" {
				defaultBool = true
			}
			flag.BoolVar(field.Addr().Interface().(*bool), flagName, defaultBool, help)
		default:
			fmt.Printf("Unsupported field type: %s\n", field.Kind())
		}
	}
}
