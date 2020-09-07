package main

import (
	"fmt"
	"github.com/howeyc/crc16"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	data := []byte("test")
	for i := 0; i < b.N; i++ {
		checksum := crc16.ChecksumCCITTFalse(data)
		fmt.Sprintf("%d", checksum)
	}
}
