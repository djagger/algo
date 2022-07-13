package structs

import (
	"reflect"
	"testing"
)

func TestSingleArray_Add(t *testing.T) {
	ar := NewSingleArray[string]()
	ar.Add("a")
	ar.Add("b")
	ar.Add("c")

	want := []string{"a", "b", "c"}

	ok := reflect.DeepEqual(ar.array, want)
	if !ok {
		t.Errorf("got %v, want %v", ar.array, want)
	}
}

func TestSingleArray_Remove(t *testing.T) {
	ar := NewSingleArray[string]()
	ar.Add("x")
	ar.Add("y")
	ar.Add("z")
	removedElement := ar.Remove(1)

	want := []string{"x", "z"}

	ok := reflect.DeepEqual(ar.array, want)
	if !ok {
		t.Errorf("got %v, want %v", ar.array, want)
	}

	wantRemoved := "y"
	if removedElement != wantRemoved {
		t.Errorf("got %s, want %s", removedElement, wantRemoved)
	}
}
