package randomtext

import (
	"errors"
	"math/rand"
	"strings"
)
func getRandomNumber() int {
	return rand.Intn(3)+1
}
func RandomWord(length uint, separator string) (string, error) {
	if length == 0 {
		return "", errors.New("length cannot be 0")
	}

	var words []string
	for i := 0; i < int(length); i++ {
		randomIndex := getRandomNumber();
		if randomIndex == 1 {
			words = append(words, Noun())
		} else if randomIndex == 2{
			words = append(words, Adjective())
		} else {
			words = append(words, Animal())
		}
	}

	randomWords :=  strings.Join(words, separator)
	return randomWords, nil
}

