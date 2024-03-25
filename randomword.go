package randomtext

import "errors"

func RandomWord(length uint, separator string) (string, error) {
	if length == 0 {
		return "", errors.New("length cannot be 0")
	}
	var randomWords = ""
	for i := 0; i < int(length); i++ {
		if i%2 == 0 {
			if i == int(length)-1 {
				randomWords = randomWords + Noun()
			} else {
				randomWords = randomWords + Noun() + separator
			}

		} else {
			if i == int(length)-1 {
				randomWords = randomWords + Adjactive()
			} else {
				randomWords = randomWords + Adjactive() + separator
			}
		}
	}
	return randomWords, nil

}
