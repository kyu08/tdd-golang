package main

import (
	"fmt"
	"reflect"
)

func main() {

	x := struct {
		name string
		ok   bool
	}{"a", false}

	val := reflect.ValueOf(x)
	fmt.Printf("numFiled: %d\n", val.NumField())
	fmt.Printf("field(0): %v\n", val.Field(0))
	fmt.Printf("field(1): %v", val.Field(1))
}
