package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	delimiter = "-"
	// randomNoMaxRange is used to specify the range of a no used in string spec.
	randomNoMaxRange = 65536
	//strPattern is regex for a string that is a sequence of numbers followed by dash followed by text
	//Example: "48-subodh-30-kumar-22-cisco"
	strSpecPattern = "^([0-9]+)-([a-zA-Z]+)$"
	// charSet this is combination of upper and lower case which help creating random alphabet string only.
	charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// subStringSize is used to specify the size of substring in string spec
	// Example of string spec with subStringSize=3 is "48-subodh-30-kumar-22-cisco"
	subStringSize = 3
	// strLengthMaxRange used to specify the substring length range.
	strLengthMaxRange = 10
)

var (
	identifierValidater = regexp.MustCompile(strSpecPattern)
	charSetRune         = []rune(charSet)
)

func main() {
	input := "48-subodh-30-kumar-22-cisco"
	fmt.Println("inputStr:" + input)
	fmt.Println("is StringSpec valid :", testValidate(input))
	fmt.Println("whole story: " + wholeStory(input))
	fmt.Println("average of numbers:", averageNumber(input))
	fmt.Printf("storySate:")
	fmt.Println(storyStats(input))
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

/*-
storyStats  function takes string spec as input and returns
- the shortest word
-the longest word
-the average word length
-the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
in given input string spec if given string match with string spec pattern.
Example: "48-subodh-30-kumar-22-cisco" is correct string spec for strSpecPattern where kumar is shortestword,
subodh is longest word and average of word length is (6+5+5)/3 =5.333 rounded values is 5 hence wordlist is
[kumar cisco]
*/
func storyStats(input string) (shortestWord, longestWord string, avgWordLen float64, wordList []string) {
	//length of all words.
	totalWordLenth := 0
	// count of word in given input string spec
	var wordCount = 0
	words := strings.Split(wholeStory(input), " ")
	for _, word := range words {
		wordLenth := len(word)
		// to find out longest word
		if wordLenth > len(longestWord) {
			longestWord = word
		}
		// to find out shortest word
		if wordLenth < len(shortestWord) || len(shortestWord) == 0 {
			shortestWord = word
		}
		wordCount++
		totalWordLenth = totalWordLenth + wordLenth
	}
	// to find out ave of words length
	avgWordLen = float64(totalWordLenth) / float64(wordCount)
	roundedAvgWordLen := getRoundedValue(avgWordLen)
	// to filter all words which length is eq to roundedAvgWordLen
	for _, word := range words {
		if float64(len(word)) == roundedAvgWordLen {
			wordList = append(wordList, word)
		}
	}
	return
}

/*
func getRoundedValue
param: n
return rounded value
if n is 33.3 then rounded value is 33
if n is 33.5 then rounded value is 34.
*/
func getRoundedValue(n float64) float64 {
	fractionNO := n - float64(int64(n))
	if fractionNO >= 0.5 {
		n = math.Ceil(n)
	} else {
		n = math.Floor(n)
	}
	return n
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

/*
func generateStrSpec
input: flag
output: string
if flag is true then return valid string spec for given string spec pattern strSpecPattern otherwise invalid string spec.
*/
func generateStrSpec(flag bool) string {
	strBuilder := strings.Builder{}
	for i := 0; i < subStringSize; i++ {
		intNo := getIntRandomNo(randomNoMaxRange)
		strLen := 0
		for strLen == 0 {
			// make sure that str length is not zero.
			strLen = getIntRandomNo(strLengthMaxRange)
		}
		str := getRandomString(strLen)
		if flag {
			if i == 0 {
				strBuilder.WriteString(strconv.Itoa(intNo) + "-" + str)
			} else {
				strBuilder.WriteString("-" + strconv.Itoa(intNo) + "-" + str)
			}
		} else {
			// generating wrong string pattern
			switch strLen % 2 {
			case 0:
				strBuilder.WriteString(str + "#")
			case 1:
				strBuilder.WriteString("_" + strconv.Itoa(intNo))
			}
		}
	}

	return strBuilder.String()
}
func getIntRandomNo(randomNoRange int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(randomNoRange)
}

func getRandomString(strLength int) string {
	rand.Seed(time.Now().UnixNano())
	randString := make([]rune, strLength)
	for i := 0; i < strLength; i++ {
		randString[i] = charSetRune[rand.Intn(len(charSet))]
	}
	return string(randString)
}

func getStrSlice(input string) []string {
	var delimiterCounter = 0
	return strings.FieldsFunc(input, func(r rune) bool {
		if strings.ContainsRune(delimiter, r) {
			delimiterCounter = delimiterCounter + 1
		}
		return delimiterCounter%2 == 0 && strings.ContainsRune(delimiter, r)
	})
}
