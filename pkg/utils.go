package pkg

import (
	"strings"
	"regexp"
)

// BuildString returns a string built from string array.
// Argument arr should be an array of string to be built together.
func BuildString(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, str := range arr {
		sb.WriteString(str)
		sb.WriteString(" ")
	}
	return sb.String()[:len(sb.String())-1]
}

// DivideSentence returns a string array of words (divided with space) from a string.
// Argument sentence should be a string inRussian or English.
func DivideSentence(sentence string) []string {
	if sentence == "" {
		return []string{}
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		words[i] = NormalizeWord(word)
	}
	return words
}

// NormalizeWord returns the word in lowercase cleared from special symbols.
// Argument word should be a string inRussian or English.
func NormalizeWord(word string) string {
	word = strings.ToLower(word)
	re := regexp.MustCompile(`[а-яa-z0-9ё]`)
	wordArr := re.FindAllStringSubmatch(word, -1)
	var normWord string
	for _, arr := range wordArr {
		normWord += arr[0]
	}
	return normWord
}

func MaxRatio(a, b, c int) float32 {
	if a == 0 && b == 0 {
		return 0
	}
	return float32(c) / float32(MaxInt(a, b))
}

func MinRatio(a, b, c int) float32 {
	if a == 0 || b == 0 {
		return 0
	}
	return float32(c) / float32(MinInt(a, b))
}

func XorIfEmpty(a, b int) bool {
	return (a != 0 && b == 0) || (a == 0 && b != 0)
}

func MaxFloat(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ArrContains(array []int, elem int) bool {
	for _, a := range array {
		if a == elem {
			return true
		}
	}
	return false
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntArrEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func FindInArrays(firstArr, secondArr []string, firstWord, secondWord string) bool {
	for _, item := range firstArr {
		if item == secondWord {
			return true
		}
	}
	for _, item := range secondArr {
		if item == firstWord {
			return true
		}
	}
	return false
}
