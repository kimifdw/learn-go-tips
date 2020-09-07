package main

import "testing"

const N = 1024 * 1024

type Element = int64

var xForMake = make([]Element, N)
var xForMakeCopy = make([]Element, N)
var xForAppend = make([]Element, N)
var yForMake []Element
var yForMakeCopy []Element
var yForAppend []Element

// make copy 在go 1.15中要快于make
func Benchmark_PureMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForMake = make([]Element, N)
	}
}

func Benchmark_MakeAndCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForMakeCopy = make([]Element, N)
		copy(yForMakeCopy, xForMakeCopy)
	}
}

func Benchmark_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForAppend = append([]Element(nil), xForAppend...)
	}
}
