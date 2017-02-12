package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation, do func(string)) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")

	do(message)
	do(alternate)
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	//fmt.Println(len(greeting))
	message = greeting[1] + " " + name
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
	Greet(s, Print)
	Greet(s, PrintLine)
}
