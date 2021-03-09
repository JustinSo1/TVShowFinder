package convert_test

import (
	"reflect"
	"testing"

	"github.com/JustinSo1/TVShowFinder/pkg/convert"
)

// TestFileToNumberedList calls conversions.FileToNumberedList with a file with mulitple mock links, checking
// for a valid return value.
func TestValidFileToNumberedList(t *testing.T) {
	file := []byte("Hello\nWorld\nTest")
	expected := []string{
		"Image Links",
		"1. Hello",
		"2. World",
		"3. Test",
	}
	actual, err := convert.FileToNumberedList(file)
	if !reflect.DeepEqual(&expected, &actual) || err != nil {
		t.Fatalf(`FileToNumberedList(file) = %q, %v, want match for %#q`, actual, err, expected)
	}
}

// TestFileToNumberedList calls conversions.FileToNumberedList with an empty file, checking
// for a valid return value.
func TestEmptyFileToNumberedList(t *testing.T) {
	file := []byte{}

	expected := []string{
		"Image Links",
	}

	actual, err := convert.FileToNumberedList(file)
	if !reflect.DeepEqual(&expected, &actual) || err == nil {
		t.Fatalf(`FileToNumberedList(file) = %q wants match for %#q, %v`, actual, expected, "Cannot give an empty file")
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
