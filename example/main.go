package main

import (
	"fmt"

	"github.com/akp111/randomtext"
)

func main(){
	// Generate random adjectives
	randomAdjactive := randomtext.Adjective()
	fmt.Printf("Random adjactive: %v\n", randomAdjactive)

	// generate random first names
	randomFirstName := randomtext.FirstName()
	fmt.Printf("Random first name: %v\n", randomFirstName)

	// random full name
	randomFullName := randomtext.RandomName()
	fmt.Printf("Random first name: %v\n", randomFullName)

	// random animal name
	randomAnimal := randomtext.Animal()
	fmt.Printf("Random animal: %v\n", randomAnimal)
	
	// random words
	randomWord, _ := randomtext.RandomWord(5,"-")
	fmt.Printf("Random words: %v\n", randomWord)
}