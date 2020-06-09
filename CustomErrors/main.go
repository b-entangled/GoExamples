package main

import (
	"fmt"
	"github.com/b-entangled/GoExamples/CustomErrors/errors"
)

func main() {
	var err error = errors.CError("Custom Error", 101)
	if cerr, ok := err.(*errors.CustomError); ok {
		fmt.Printf("Error Type: %T\n", cerr)
		fmt.Println(cerr.Error())
		fmt.Println(cerr.Code())
	}

}
