package reflects

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got '%s', want '%s'", got[0], expected)
	}
}

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
	//fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
}
