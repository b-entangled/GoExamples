package main

import (
	"fmt"
	"github.com/b-entangled/GoExamples/CustomErrors/errors"
)

func main(){
	var customError error = errors.CError("Custom Error", 101)
	fmt.Printf("Error Type: %T\n", customError)
	fmt.Println(customError.Error())
}