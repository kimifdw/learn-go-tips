package chapter18_reflect

import (
	"reflect"
	"testing"
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
