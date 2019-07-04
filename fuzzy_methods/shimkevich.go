package fuzzy_methods

import (
	"lwahomura/comparison/pkg"
)

type Comparator struct {
	SubtokenLength int
}

// CompareWords returns similarity ratio of two words using shingles algorithm and Shimkevich-Simpson ratio.
// Argument first should be a string in Russian or English.
// Argument second should be a string in Russian or English.
func (comp *Comparator) CompareWords(first, second string) float32 {
	first = pkg.NormalizeWord(first)
	second = pkg.NormalizeWord(second)
	firstRunic := []rune(first)
	secondRunic := []rune(second)
	firstLen := len([]rune(first))
	secondLen := len([]rune(second))
	if pkg.XorIfEmpty(firstLen, secondLen) {
		return 0
	}
	if first == second {
		return 1
	}
	var subtokenFirst, subtokenSecond string
	equalPartsCount := 0
	usedTokens := make([]bool, secondLen-comp.SubtokenLength+1)
	for i := 0; i < (firstLen - comp.SubtokenLength + 1); i++ {
		subtokenFirst = string(firstRunic[i : comp.SubtokenLength+i])
		for j := 0; j < (secondLen - comp.SubtokenLength + 1); j++ {
			subtokenSecond = string(secondRunic[j : comp.SubtokenLength+j])
			if subtokenFirst == subtokenSecond {
				equalPartsCount += 1
				usedTokens[j] = true
				break
			}
		}
	}
	subtokenFirstCount := firstLen - comp.SubtokenLength + 1
	subtokenSecondCount := secondLen - comp.SubtokenLength + 1
	compRatio := pkg.MaxRatio(subtokenFirstCount, subtokenSecondCount, equalPartsCount)
	return compRatio
}
