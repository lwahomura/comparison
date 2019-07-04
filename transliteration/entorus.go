package transliteration

import (
	"strings"
	"regexp"
	"github.com/lwahomura/comparison/pkg"
)

func EnToRus(target string) string {
	words := pkg.DivideSentence(target)
	var translitWords []string
	for _, word := range words {
		var translitWord string
		for sub := range pkg.EnToRusSR {
			//fmt.Println("found", sub)
			word = strings.Replace(word, sub, string(pkg.EnToRusSR[sub]), -1)
		}
		for _, elem := range word {
			trmain := pkg.EnToRusRR[elem]
			switch trmain {
			case 0:
				trex := pkg.EnToRusRS[elem]
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
		if regexp.MustCompile(`^.*э$`).MatchString(translitWord) {
			translitWord = translitWord[:len(translitWord) - 2]
		}
		translitWords = append(translitWords, translitWord)
	}
	return pkg.BuildString(translitWords)
}