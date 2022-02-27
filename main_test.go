package main

import (
	"gotest.tools/assert"
	"strings"
	"testing"
)

func TestStringSpecPattern_testValidateFucn(t *testing.T) {
	str := generateStrSpec(true)
	t.Log("valid string spec:" + str)
	assert.Assert(t, testValidate(str))
	str = generateStrSpec(false)
	t.Log("Invalid string spec:" + str)
	assert.Assert(t, !testValidate(str))
	if testValidate(str) {
		t.Failed()
	}
}
func TestStringSpecPattern_wholeStoryFunc(t *testing.T) {
	str := generateStrSpec(true)
	assert.Assert(t, testValidate(str))
	assert.Assert(t, len(strings.Split(wholeStory(str), " ")) == 3)
	str = generateStrSpec(false)
	assert.Assert(t, !testValidate(str))
	assert.Assert(t, len(strings.Split(wholeStory(str), " ")) != 3)
}

func TestStringSpecPattern_averageNumberFunc(t *testing.T) {
	str := generateStrSpec(true)
	assert.Assert(t, testValidate(str))
	assert.Assert(t, averageNumber(str) != 0)
	str = generateStrSpec(false)
	assert.Assert(t, !testValidate(str))
	assert.Assert(t, averageNumber(str) == 0)
}

func TestStringSpecPattern_storyState(t *testing.T) {
	str := generateStrSpec(true)
	assert.Assert(t, testValidate(str))
	shortestWord, longestWord, avg, wordList := storyStats(str)
	assert.Assert(t, len(shortestWord) != 0)
	assert.Assert(t, len(longestWord) != 0)
	assert.Assert(t, avg != 0)
	assert.Assert(t, len(wordList) != 0)
}
