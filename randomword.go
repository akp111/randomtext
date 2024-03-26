package randomtext

import (
	"errors"
	"strings"
)

func RandomWord(length uint, separator string) (string, error) {
	if length == 0 {
		return "", errors.New("length cannot be 0")
	}

	var words []string
	for i := 0; i < int(length); i++ {
		if i%2 == 0 {
			words = append(words, Noun())
		} else {
			words = append(words, Adjective())
		}
	}

	randomWords := strings.Join(words, separator)
	return randomWords, nil
}

