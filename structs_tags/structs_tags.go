//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/golang-playground
// Date: Fri Jul 11 01:20:44 PM UTC 2025
//

// {{{ package

package structs_tags

// }}}
// {{{ imports

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// }}}
// ---- types ----
// {{{ type User struct

type User struct {
	Name string `validate:"required"`
	Age  int    `validate:"min=18,max=65"`
}

// }}}
// ---- funcs ----
// {{{ func ruleCheck()

func ruleCheck(rule string, field reflect.Value, fieldName string) error {
	switch {
	case rule == "required":
		if field.String() == "" {
			return fmt.Errorf("%s is required", fieldName)
		}
	case strings.HasPrefix(rule, "min="):
		m, _ := strconv.Atoi(rule[4:])
		if field.Int() < int64(m) {
			return fmt.Errorf("%s=%d (min=%d)", fieldName, field.Int(), m)
		}
	case strings.HasPrefix(rule, "max="):
		m, _ := strconv.Atoi(rule[4:])
		if field.Int() > int64(m) {
			return fmt.Errorf("%s=%d (max=%d)", fieldName, field.Int(), m)
		}
	}
	return nil
}

// }}}
// {{{ func validate()

func validate(val any) error {
	fmt.Println("======== reflect.TypeOf(val) ========")
	t := reflect.TypeOf(val)
	fmt.Println("t.Name() ........", t.Name())
	fmt.Println("t.Kind() ........", t.Kind())
	fmt.Println("--------")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("validate")
		fmt.Println("field ...........", field)
		fmt.Println("tag .............", tag)
	}
	fmt.Println("======== reflect.ValueOf(val) ========")
	v := reflect.ValueOf(val)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get("validate")
		fmt.Println("field ...........", field)
		fmt.Println("tag .............", tag)
		if tag == "" {
			continue
		}
		rules := strings.Split(tag, ",")
		for _, rule := range rules {
			fieldName := v.Type().Field(i).Name
			err := ruleCheck(rule, field, fieldName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// }}}
// ---- Main ----
// {{{ func Main()

func Main() {
	users := []User{{"", 25}, {"Tom", 17}, {"Bob", 33}, {"Alan", 69}}
	for _, user := range users {
		fmt.Println("================", user, "================")
		err := validate(user)
		if err != nil {
			fmt.Println("******** validation failed:", err, "********")
			continue
		}
		fmt.Println("******** validation passed ********")
		fmt.Println("----------------", user, "----------------")
	}
}

// }}}
