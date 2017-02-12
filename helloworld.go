package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting)

	fmt.Println(message)
	fmt.Println(alternate)
}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY:" + name
	return
}

func main() {

	//message := "Hello Go World!"
	//var greeting *string = &message

	//var message Salutation = "Hello"

	//fmt.Println(message, greeting)
	//fmt.Println(message, *greeting)
	var s = Salutation{"Cookie", "Hello"}
	Greet(s)
}
