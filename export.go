package comparison

import (
	"lwahomura/comparison/fuzzy_methods"
	"sort"
	"lwahomura/comparison/dictionary"
	"lwahomura/comparison/pkg"
	"fmt"
	"lwahomura/comparison/transliteration"
)

// compMethods returns the highest returned similarity ratio (gotten from comparison methods).
// Argument first should be a string in Russian or English.
// Argument second should be a string in Russian or English.
// Argument lang should be a string determining the language - "Rus" or "Eng".
func compMethods(first, second, lang string) float32 {
	var results []float32
	comp := fuzzy_methods.Comparator{SubtokenLength: 2}
	results = append(results, fuzzy_methods.LevDistance(first, second), comp.CompareWords(first, second))
	if lang == "Rus" {
		res := fuzzy_methods.RusMetaphoneCompare(first, second)
		if res >= 0.85 {
			results = append(results, res)
		} else {
			results = append(results, 0)
		}
	}
	if lang == "Eng" {
		res := fuzzy_methods.MetaphoneCompare(first, second)
		if res >= 0.85 {
			results = append(results, res)
		} else {
			results = append(results, 0)
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i] > results[j]
	})
	return results[0]
}

// Compare returns true if strings are similar.
// Argument first should be a string in Russian or English.
// Argument second should be a string in Russian or English.
func Compare(first, second string) bool {
	var results []float32
	db := dictionary.GetDatabases()
	if pkg.FindInArrays(db.GetEnToRusTranslation(first), db.GetEnToRusTranslation(second), first, second) {
		fmt.Println("etr")
		results = append(results, 1)
	}
	if pkg.FindInArrays(db.GetRusToEnTranslation(first), db.GetRusToEnTranslation(second), first, second) {
		fmt.Println("rte")
		results = append(results, 1)
	}
	firstEnT := transliteration.RusToEn(first)
	secondEnT := transliteration.RusToEn(second)
	firstRusT := transliteration.EnToRus(first)
	secondRusT := transliteration.EnToRus(second)
	results = append(results, compMethods(first, second, "Rus"), compMethods(first, second, "En"),
		compMethods(firstEnT, secondEnT, "Eng"), compMethods(firstRusT, secondRusT, "Rus"))
	sort.Slice(results, func(i, j int) bool {
		return results[i] > results[j]
	})
	return results[0] > pkg.Threshold
}
