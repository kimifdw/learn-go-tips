package chapter18_reflect

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unsafe"
)

type People struct {
	Age  int
	Name string
}

func New() *People {
	return &People{
		Age:  18,
		Name: "Emily",
	}
}

func NewUseReflect() interface{} {
	var p People
	t := reflect.TypeOf(p)
	v := reflect.New(t)

	v.Elem().Field(0).Set(reflect.ValueOf(18))
	v.Elem().Field(1).Set(reflect.ValueOf("Emily"))
	return v.Interface()
}

func String2ByteSlice(str string) (bs []byte) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	sliceHdr.Data = strHdr.Data
	sliceHdr.Len = strHdr.Len
	sliceHdr.Cap = strHdr.Len
	runtime.KeepAlive(&str)
	return
}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkNewUseReflect(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewUseReflect()
	}
}

func TestString2ByteSlice(t *testing.T) {
	t.Run("String2ByteSlice", func(t *testing.T) {
		str := strings.Join([]string{"Go", "Land"}, "")
		s := String2ByteSlice(str)
		fmt.Printf("%s\n", s)
		s[5] = 'g'
		fmt.Println(str)
	})
}
