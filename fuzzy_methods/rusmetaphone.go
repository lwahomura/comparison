package fuzzy_methods

import (
	"strings"
	"regexp"
	"fmt"
	"lwahomura/comparison/pkg"
)

func dropDoubles(token string) string {
	runes := []rune(token)
	if len(runes) == 1 || token == "" {
		return token
	}
	result := fmt.Sprintf("%c", runes[0])
	for i := 1; i < len(runes); i ++ {
		if runes[i] != runes[i-1] {
			result += fmt.Sprintf("%c", runes[i])
		}
	}
	return result
}

func transformVowels(token string) string {
	token = regexp.MustCompile("(йо|ио|йе|ие)").ReplaceAllLiteralString(token, "и")
	token = regexp.MustCompile("[оыя]").ReplaceAllLiteralString(token, "а")
	token = regexp.MustCompile("[еёэ]").ReplaceAllLiteralString(token, "и")
	token = regexp.MustCompile("ю").ReplaceAllLiteralString(token, "у")
	return token
}

func transformConsonants(token string) string {
	runes := []rune(token)
	result := ""
	for i := 0; i < len(runes); i ++ {
		if (i != len(runes)-1 && isSonorous(fmt.Sprintf("%c", runes[i])) && isConsonantExcept(fmt.Sprintf("%c", runes[i+1]))) || (i == len(runes)-1 && isSonorous(fmt.Sprintf("%c", runes[i]))) {
			switch runes[i] {
			case 'б':
				result += "п"
			case 'з':
				result += "с"
			case 'д':
				result += "т"
			case 'в':
				result += "ф"
			case 'г':
				result += "к"
			}
		} else {
			result += fmt.Sprintf("%c", runes[i])
		}
	}
	return result
}

func transformRusDoubles(token string) string {
	return regexp.MustCompile("(тс|дс)").ReplaceAllLiteralString(token, "ц")
}

func dropExtra(token string) string {
	return regexp.MustCompile("[ъь]").ReplaceAllLiteralString(token, "")
}

// RusMetaphone returns the target string drive through Russian Metaphone.
// Argument token should be a string in Russian or English.
func RusMetaphone(token string) string {
	if len([]rune(token)) == 0 {
		return ""
	}

	token = pkg.NormalizeWord(token)
	token = dropDoubles(token)
	token = transformVowels(token)
	token = transformConsonants(token)
	token = transformRusDoubles(token)
	token = dropExtra(token)
	token = dropDoubles(token)
	token = strings.ToUpper(token)

	return token
}

func isConsonantExcept(elem string) bool {
	return regexp.MustCompile("[бвгджзкпстфхцчшщ]").MatchString(elem)
}

func isSonorous(elem string) bool {
	return regexp.MustCompile("[бздвг]").MatchString(elem)
}

// RusMetaphoneCompare returns similarity ratio of two words using Russian Metaphone Algorithm.
// Argument first should be a string in Russian or English.
// Argument second should be a string in Russian or English.
func RusMetaphoneCompare(first, second string) float32 {
	first = RusMetaphone(first)
	second = RusMetaphone(second)
	comp := Comparator{SubtokenLength: 2}
	return pkg.MaxFloat(LevDistance(first, second), comp.CompareWords(first, second))
}
