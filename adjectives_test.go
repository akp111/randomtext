package randomtext

import (
	"testing"
)

func TestAdjectiveGenerator(t *testing.T) {
	randomAdjective := Adjactive();
	if(randomAdjective == ""){
		t.Fatalf("Random adjaective cannot be empty")
	}
}