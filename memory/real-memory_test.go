package main

import (
	"sort"
	"testing"
)

func BenchmarkSortString(b *testing.B) {
	s := []string{"heart", "lungs", "brain", "kidneys", "pancreas"}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var ss sort.StringSlice = s
		var si sort.Interface = ss
		sort.Sort(si)
	}
}
