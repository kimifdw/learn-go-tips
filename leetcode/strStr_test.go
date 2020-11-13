package main

import (
	"fmt"
	"testing"
)

func TestStrStr1(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	fmt.Println(StrStr1(haystack, needle))
}
