package randomtext

import (
	"fmt"
	"testing"
)

func TestRandomWordGeneratorFailingCase(t *testing.T) {
	_, err := RandomWord(0,"");
	if(err == nil){
		t.Fatalf("Shoudl throw error for length to be 0")
	}
}

func TestRandomWordGeneratorPassingCase(t *testing.T) {
	word, err := RandomWord(5," ");
	if(err != nil){
		t.Fatalf("Shoudl throw error for length to be 0")
	}
	fmt.Println(word)
}