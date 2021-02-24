package convert

import (
	"regexp"
	"strconv"
	"strings"
)

// FileToNumberedList returns a numbered list
func FileToNumberedList(s []byte) []string {
	var listElements []string
	lines := removeCarriageReturns(s)
	listElements = strings.Split(lines, "\n")
	for i := 1; i < len(listElements); i++ {
		listElements[i] = strconv.Itoa(i) + ". " + listElements[i]
	}
	return listElements
}

func removeCarriageReturns(s []byte) string {
	re := regexp.MustCompile(`\r?`)
	lines := re.ReplaceAllString(string(s), "")
	return lines
}

// IntArrToString converts int array to string
func IntArrToString(arr []int) string {
	b := make([]string, len(arr))
	for i, v := range arr {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, ",")
}
