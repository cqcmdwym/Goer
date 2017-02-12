package main

import "./greeting"

func main() {

	//message := "Hello Go World!"
	//var greeting *string = &message

	//var message Salutation = "Hello"

	//fmt.Println(message, greeting)
	//fmt.Println(message, *greeting)
	var s = greeting.Salutation{"Cookie", "Hello"}

	greeting.Greet(s, greeting.CreatePrintFunction("!!!"), true)

	greeting.TypeSwitchTest("Hello")
}
