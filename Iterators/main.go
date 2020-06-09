package main

import (
	"fmt"
	funciterator "github.com/b-entangled/GoExamples/Iterators/func_iterator"
	chaniterator "github.com/b-entangled/GoExamples/Iterators/channel_iterator"
)

func main(){
	var n int = 5
	// == Iterator using Callback Func ==
	data := funciterator.GetData(n)

	// callback function, To be called for each data.
	cprint := func(val int){
		fmt.Println("Callback :", val)
	}

	funciterator.Iterator(data, cprint)
	// ===================================

	// ===== Iterator using Channels =====
	// Receive data from iterator one by one until end of data
	for val := range chaniterator.Iterate(n){
		fmt.Println("Channel :", val)
	}
	// ===================================
}