package convert_test

import (
	"reflect"
	"testing"

	"github.com/JustinSo1/TVShowFinder/pkg/convert"
)

// TestFileToNumberedList calls conversions.FileToNumberedList with a file, checking
// for a valid return value.
func TestFileToNumberedList(t *testing.T) {
	file := []byte("Hello")
	expected := []string{
		"Image Links",
		"1. Hello",
	}
	actual := convert.FileToNumberedList(file)
	if !reflect.DeepEqual(&expected, &actual) {
		t.Fatalf(`FileToNumberedList(file) = %q, want match for %#q`, actual, expected)
	}
}

// TestFileToList calls conversions.FileToList with a file, checking
// for a valid return value.
func TestFileToList(t *testing.T) {
	file := []byte("Hello\nWorld\nTest")
	expected := []string{
		"Hello",
		"World",
		"Test",
	}
	actual := convert.FileToList(file)
	if !reflect.DeepEqual(&expected, &actual) {
		t.Fatalf(`FileToList(file) = %q, want match for %#q`, actual, expected)
	}
}
