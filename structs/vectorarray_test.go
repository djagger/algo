package structs

import (
	"reflect"
	"testing"
)

func removeEmpty[T any](arr []T) []T {
	var res []T
	for _, v := range arr {
		if !reflect.ValueOf(v).IsZero() {
			res = append(res, v)
		}
	}
	return res
}

func TestVectorArray_Add(t *testing.T) {
	ar := NewVectorArray[string](2)
	ar.Add("a")
	ar.Add("b")
	ar.Add("c")

	want := []string{"a", "b", "c"}

	resArray := removeEmpty[string](ar.array)

	ok := reflect.DeepEqual(resArray, want)
	if !ok {
		t.Errorf("got %v, want %v", ar.array, want)
	}
}

func TestVectorArray_Remove(t *testing.T) {
	ar := NewVectorArray[string](2)
	ar.Add("x")
	ar.Add("y")
	ar.Add("z")
	removedElement := ar.Remove(1)

	want := []string{"x", "z"}

	resArray := removeEmpty[string](ar.array)

	ok := reflect.DeepEqual(resArray, want)
	if !ok {
		t.Errorf("got %v, want %v", ar.array, want)
	}

	wantRemoved := "y"
	if removedElement != wantRemoved {
		t.Errorf("got %s, want %s", removedElement, wantRemoved)
	}
}
