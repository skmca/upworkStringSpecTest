package main

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	delimiter = "-"
	//strPattern is regex for a string that is a sequence of numbers followed by dash followed by text
	//Example: "48-subodh-30-kumar-22-cisco"
	strSpecPattern = "^([0-9]+)-([a-zA-Z]+)$"
)

var (
	identifierValidater = regexp.MustCompile(strSpecPattern)
)

func main() {
	input := "48-subodh-30-kumar-22-cisco"
	fmt.Println("inputStr:" + input)
	fmt.Println(" is StringSpec valid :", testValidate(input))
}

// testValidate func is used to check that whether given input string match with strSpecPattern
//Example: "48-subodh-30-kumar-22-cisco" is correct string spec
func testValidate(input string) bool {
	if len(input) == 0 {
		return false
	}

	for _, str := range getStrSlice(input) {
		if !identifierValidater.MatchString(str) {
			return false
		}
	}
	return true
}

/*
func getStrSlice
param: input
output: [] string
it returns all sub string in given string spec.
Example:
inputStingSpec: "48-subodh-30-kumar-22-cisco"
output: [48-subodh 30-kumar 22-cisco]
*/
func getStrSlice(input string) []string {
	var delimiterCounter = 0
	return strings.FieldsFunc(input, func(r rune) bool {
		if strings.ContainsRune(delimiter, r) {
			delimiterCounter = delimiterCounter + 1
		}
		return delimiterCounter%2 == 0 && strings.ContainsRune(delimiter, r)
	})
}
