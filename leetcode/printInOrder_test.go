package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestByteToH(t *testing.T) {
	var m int32 = 0x12345678
	var n = int8(m)
	str := strconv.FormatInt(int64(n), 16)
	bytes := Hextob(str)
	h := ByteToH(bytes)
	fmt.Println(h)
}
