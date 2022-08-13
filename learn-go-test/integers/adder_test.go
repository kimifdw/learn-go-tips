package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	excepted := 4

	if sum != excepted {
		t.Errorf("expected '%d' but got '%d'", excepted, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}

// Add takes two integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}
