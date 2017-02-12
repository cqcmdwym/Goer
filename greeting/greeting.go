package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting, "yo")
	if prefix := "Mr "; isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	//fmt.Println(len(greeting))
	message = greeting[1] + " " + name
	alternate = "HEY:" + name
	return
}
