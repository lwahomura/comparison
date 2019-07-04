package transliteration

import (
	"regexp"
	"lwahomura/comparison/pkg"
)

// RusToEn returns the target string transliterated in English.
// Argument target must be a string.
func RusToEn(target string) string {
	words := pkg.DivideSentence(target)
	var translitWords []string
	for _, word := range words {
		word = processSignes(word)
		var translitWord string
		for _, elem := range word {
			trmain := pkg.RusToEnRR[elem]
			switch trmain {
			case 0:
				trex := pkg.RusToEnRS[elem]
				switch trex {
				case "":
					translitWord += string(elem)
				default:
					translitWord += trex
				}
			default:
				translitWord += string(trmain)
			}
		}
		translitWords = append(translitWords,translitWord)
	}
	return pkg.BuildString(translitWords)
}

func processSignes(word string) string{
	var result string
	var previous bool
	for _, elem := range word {
		if regexp.MustCompile("(ь|ъ)").MatchString(string(elem)) {
			previous = true
		} else {
			if previous {
				if regexp.MustCompile("(е|ё|и)").MatchString(string(elem)) {
					result += "i"
					result += string(elem)
					previous = false
				} else {
					result += string(elem)
					previous = false
				}
			} else {
				result += string(elem)
			}
		}
	}
	return result
}