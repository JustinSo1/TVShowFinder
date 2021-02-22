package pkg

import (
	"regexp"
	"strconv"
	"strings"
)

// ReformatFile removes carriage returns
func ReformatFile(s []byte) []string {
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
