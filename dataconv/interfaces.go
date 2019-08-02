package dataconv

import "fmt"

func CheckTypes(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Int")
	default:
		fmt.Println("naoo sei")
	}
}

func Interfaces() {
	CheckTypes("teste")
	CheckTypes(1)
	CheckTypes(false)

	var i interface{}
	i = "teste"

	if val, ok := i.(string); ok {
		fmt.Println("val is", val)
	}

	if _, ok := i.(int); !ok {
		fmt.Println("nao vai aparecer")
	}
}
