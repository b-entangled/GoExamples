package main

import (
	"fmt"
	spro "github.com/b-entangled/GoExamples/Middlewares/pkg/string_processor"
	mware "github.com/b-entangled/GoExamples/Middlewares/pkg/middlewares"
)

func main(){
	var sentence string = `
	Physicists at the National Institute of Standards and Technology have boosted their control of the 
	fundamental properties of molecules at the quantum level by linking or "entangling" 
	an electrically charged atom and an electrically charged molecule, showcasing a way to 
	build hybrid quantum information systems that could manipulate, store and transmit different forms of data.
	`
	fmt.Println(len(spro.SplitString(sentence)))
	fmt.Println(mware.StringsMiddlewares(mware.StringsMiddlewares(sentence)))
}