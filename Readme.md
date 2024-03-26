# randomtext

[![Go Reference](https://pkg.go.dev/badge/github.com/akp111/randomtext.svg)](https://pkg.go.dev/github.com/akp111/randomtext)

A go library that helps you generate:
1. A random adjective
2. A random noun
3. A random first name
4. A random last name
5. A random full name
6. A radom word of custom length and delimiter
7. A random animal name

## For a quickie, please see the below example

```go
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
```

## Next step:

1. ~~Add animals~~
2. Add colours
3. Add languages
4. Add anime


