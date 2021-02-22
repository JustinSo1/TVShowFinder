package search

import (
	"net/url"
	"strconv"
	"strings"
)

// IntArrToString converts int array to string
func IntArrToString(arr []int) string {
	b := make([]string, len(arr))
	for i, v := range arr {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, ",")
}

// IsURL Checks if string is url
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
