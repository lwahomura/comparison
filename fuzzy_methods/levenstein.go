package fuzzy_methods

import (
	"lwahomura/comparison/pkg"
)

type Word struct {
	Text     string
	Codes    []int
	Language string
}

func (word *Word) getCodes() {
	word.Codes = nil
	switch word.Language {
	case "Rus":
		for _, char := range word.Text {
			word.Codes = append(word.Codes, pkg.RusCodeKeys[char])
		}
	case "Eng":
		for _, char := range word.Text {
			word.Codes = append(word.Codes, pkg.EnCodeKeys[char])
		}
	}
}

// costDistanceSymbol returns the cost of transformation for symbols.
// Argument first should have type Word.
// Argument second should have type Word.
// Argument firstPostion should be an integer defining which symbol we want to transform.
// Argument secondPostion should be an integer defining to which symbol we want to transform.
func costDistanceSymbol(first, second Word, firstPosition, secondPosition int) int {
	if first.Text[firstPosition] == second.Text[secondPosition] {
		return 0
	}
	if first.Codes[firstPosition] != 0 && first.Codes[firstPosition] == second.Codes[secondPosition] {
		return 0
	}
	resultWeight := 0
	nearKeys := pkg.ClosestKeys[first.Codes[firstPosition]]
	if len(nearKeys) == 0 {
		resultWeight = 2
	} else {
		if pkg.ArrContains(nearKeys, second.Codes[secondPosition]) {
			resultWeight = 1
		} else {
			resultWeight = 2
		}
	}
	return resultWeight
}

func LevDistance(s1, s2 string) float32 {
	min := func(values ...int) int {
		m := values[0]
		for _, v := range values {
			if v < m {
				m = v
			}
		}
		return m
	}
	r1, r2 := []rune(s1), []rune(s2)
	n, m := len(r1), len(r2)
	if n > m {
		r1, r2 = r2, r1
		n, m = m, n
	}
	currentRow := make([]int, n+1)
	previousRow := make([]int, n+1)
	for i := range currentRow {
		currentRow[i] = i
	}
	for i := 1; i <= m; i++ {
		for j := range currentRow {
			previousRow[j] = currentRow[j]
			if j == 0 {
				currentRow[j] = i
				continue
			} else {
				currentRow[j] = 0
			}
			add, del, change := previousRow[j]+1, currentRow[j-1]+1, previousRow[j-1]
			if r1[j-1] != r2[i-1] {
				change++
			}
			currentRow[j] = min(add, del, change)
		}
	}
	return 1 - pkg.MaxRatio(n, m, currentRow[n])
}

// Levenstein returns similarity ratio of two words using Levenstein distance.
// Argument first should be a string in Russian or English.
// Argument second should be a string in Russian or English.
// Argument lang should be a string determining the language - "Rus" or "Eng".
func Levenshtein(first, second, lang string) float32 {
	first = pkg.NormalizeWord(first)
	second = pkg.NormalizeWord(second)
	firstWord := Word{Text: first, Language: lang}
	firstWord.getCodes()
	secondWord := Word{Text: second, Language: lang}
	secondWord.getCodes()
	firstLen := len([]rune(firstWord.Text))
	secondLen := len([]rune(secondWord.Text))

	if pkg.XorIfEmpty(firstLen, secondLen) {
		return 0
	}
	if first == second {
		return 1
	}
	if len(firstWord.Codes) == 0 || len(secondWord.Codes) == 0 {
		return 0
	}

	firstLenEx := firstLen + 1
	secondLenEx := secondLen + 1
	distance := make([][]int, 3)

	for i := 0; i < 3; i++ {
		distance[i] = make([]int, secondLenEx+1)
	}
	for j := 1; j < secondLenEx; j++ {
		distance[0][j] = j * 2
	}
	currentRow := 0
	for i := 1; i < firstLenEx; i++ {
		currentRow = i % 3
		distance[currentRow][0] = i * 2
		previousRow := (i - 1) % 3
		for j := 1; j < secondLenEx; j++ {
			distance[currentRow][j] = pkg.MinInt(pkg.MinInt(distance[previousRow][j]+2, distance[currentRow][j-1]+2), distance[previousRow][j-1]+costDistanceSymbol(firstWord, secondWord, i-1, j-1))
			if i > 1 && j > 1 && firstWord.Text[i-1] == secondWord.Text[j-2] && firstWord.Text[i-2] == secondWord.Text[j-1] {
				distance[currentRow][j] = pkg.MinInt(distance[currentRow][j], distance[(i-2)%3][j-2]+2)
			}
		}

	}
	levCost := distance[currentRow][secondLenEx-1]
	return 1 - pkg.MaxRatio(firstLen, secondLen, levCost)/2
}
