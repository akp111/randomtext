package randomtext

import (
	"testing"
)

func TestNameGenerator(t *testing.T) {
	randomAdjective := RandomName();
	if(randomAdjective == ""){
		t.Fatalf("Random adjaective cannot be empty")
	}
}