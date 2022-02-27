package main

import (
	"fmt"
	"regexp"
	"strconv"
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
	fmt.Println("is StringSpec valid :", testValidate(input))
	fmt.Println("whole story: " + wholeStory(input))
	fmt.Println("average of numbers:", averageNumber(input))
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

/* func wholeStory
   param: input
   return: wholeStory as string
wholeStory string consist of all words in given input string spec.
Example:
inputStingSpec: "48-subodh-30-kumar-22-cisco"
outputstring : subodh kumar cisco
*/
func wholeStory(input string) string {
	// to recover from panic if any while creating whole story.
	// if given string does not match string spec pattern ,there is highly chance of panic.
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("%v while creating whole story for given string: %s\n", r, input)
		}
	}()
	var wholeStory strings.Builder
	for _, str := range getStrSlice(input) {
		strings.Split(str, delimiter)
		wholeStory.WriteString(strings.Split(str, delimiter)[1] + " ")
	}
	return strings.TrimSpace(wholeStory.String())
}

/*
 func averageNumber
param: input
return average of given no in given input string spec.
Example:
inputStingSpec: "48-subodh-30-kumar-22-cisco"
averageOfno : (48+30+22)/3 => 33.3333
*/
func averageNumber(input string) float64 {
	// to recover from panic if any while calculating average of no in given string input  .
	// if given string input does not match string spec pattern ,there is highly chance of panic.
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("%v while finding avg number for given string: %s\n", r, input)
		}
	}()
	sum := 0
	index := -1
	str := ""
	for index, str = range getStrSlice(input) {
		numStr := strings.Split(str, delimiter)[0]
		num, _ := strconv.Atoi(numStr)
		sum = sum + num
	}
	return float64(sum) / float64(index+1)

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
