package greeting

import "fmt"
import "os"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting, "yo")
	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func GetPrefix(name string) (prefix string) {
	// switch name {
	// case "Bob":
	// 	prefix = "Mr "
	// case "Joe", "Amy":
	// 	prefix = "Dr "
	// case "Mary":
	// 	prefix = "Mrs "
	// default:
	// 	prefix = "Dude "
	// }
	switch {
	case name == "Bob":
		prefix = "Mr "
	case name == "Joe", name == "Amy", len(name) == 10:
		prefix = "Dr "
	case name == "Mary":
		prefix = "Mrs "
	default:
		prefix = "Dude "
	}
	return
}

func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("salutation")
	default:
		fmt.Println("unknown")
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

func PrintOS(){
	var o = os.Getenv("OS")
	fmt.Println("User: "+o)
}
