package convert

import (
	"regexp"
	"strconv"
	"strings"
)

// FileToNumberedList takes in a file and returns a numbered list
func FileToNumberedList(s []byte) []string {
	var listElements []string
	if len(s) == 0 {
		return listElements
	}
	listElements = append(listElements, "Image Links")
	listElements = append(listElements, FileToList(s)[:]...)
	for i := 1; i < len(listElements); i++ {
		listElements[i] = strconv.Itoa(i) + ". " + listElements[i]
	}
	return listElements
}

// FileToList returns file in a list separated by new lines
func FileToList(s []byte) []string {
	lines := removeCarriageReturns(s)
	return strings.Split(lines, "\n")
}

func removeCarriageReturns(s []byte) string {
	re := regexp.MustCompile(`\r?`)
	lines := re.ReplaceAllString(string(s), "")
	return lines
}
