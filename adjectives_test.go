package randomtext

import (
	"testing"
)

func TestAdjectiveGenerator(t *testing.T) {
	randomAdjective := Adjective();
	if(randomAdjective == ""){
		t.Fatalf("Random adjaective cannot be empty")
	}
}