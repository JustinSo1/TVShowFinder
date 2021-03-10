package search_test

import (
	"testing"

	"github.com/JustinSo1/TVShowFinder/pkg/search"
)

func TestValidIsUrl(t *testing.T) {
	URL := "https://www.google.ca/"
	expected := true
	actual := search.IsURL(URL)
	if expected != actual {
		t.Fatalf(`IsURL(URL) = %t, want match for %t`, actual, expected)
	}
}

func TestInvalidIsUrl(t *testing.T) {
	URL := "www3.goo//gle..com"
	expected := false
	actual := search.IsURL(URL)
	if expected != actual {
		t.Fatalf(`IsURL(URL) = %t, want match for %t`, actual, expected)
	}
}
