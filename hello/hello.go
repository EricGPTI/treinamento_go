package main

import "fmt"
import "rsc.io/quote"
import "example.com/greetings"

//Criando primeira função.
func main() {
	fmt.Println(quote.Go())
	fmt.Println(GetGreetings())
}

//Chamando message em greetings.
func GetGreetings() string {
	message := greetings.Hello("Eric")
	return message
	//fmt.Println(message)
}


