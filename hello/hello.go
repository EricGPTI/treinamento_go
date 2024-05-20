package main

import "fmt"
//import "rsc.io/quote"
import "example.com/greetings"
import "log"


//Criando primeira função.
func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	//Request a greeting message.
	message, err := greetings.Hellos(names)

	//If an error was returned, print it to the console and
	//exit the program.
	if err != nil {
		log.Fatal(err)
	}
	//If no error was returned, print the returned message
	//to the console.
	for _, msg := range message {
		fmt.Println(msg)
	}
}