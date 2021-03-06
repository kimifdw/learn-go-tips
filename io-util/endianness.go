package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	buf1 := new(bytes.Buffer)
	err := binary.Write(buf1, binary.LittleEndian, int32(-0x21524111))
	if err != nil {
		fmt.Println("binary write failed:" + err.Error())
	}
	fmt.Printf("%# s\n", buf1)
}
