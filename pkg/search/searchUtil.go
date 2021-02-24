package search

import (
	"net/url"
)

// IsURL Checks if string is url
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
