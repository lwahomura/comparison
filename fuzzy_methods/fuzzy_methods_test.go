package fuzzy_methods

import (
	"testing"
	"fmt"
	"lwahomura/comparison/pkg"
)

func Test_getCodes1(t *testing.T) {
	w := Word{Text: ""}
	w.getCodes()
	if len(w.Codes) != 0 {
		t.Error("Expected empty arr, got ", w.Codes)
	}
}

func Test_getCodes2(t *testing.T) {
	w := Word{Text: "AS"}
	w.getCodes()
	if len(w.Codes) != 0 {
		t.Error("Expected empty arr, got ", w.Codes)
	}
}

func Test_getCodes3(t *testing.T) {
	w := Word{Text: "", Language: "Test"}
	w.getCodes()
	if len(w.Codes) != 0 {
		t.Error("Expected empty arr, got ", w.Codes)
	}
}

func Test_getCodes4(t *testing.T) {
	w := Word{Text: "абв", Language: "Rus"}
	w.getCodes()
	if !pkg.IntArrEqual(w.Codes, []int{70, 188, 68}) {
		t.Error("Expected [70, 188, 68], got ", w.Codes)
	}
}

func Test_getCodes5(t *testing.T) {
	w := Word{Text: "yui", Language: "Eng"}
	w.getCodes()
	if !pkg.IntArrEqual(w.Codes, []int{89, 85, 73}) {
		t.Error("Expected [89, 85, 73], got ", w.Codes)
	}
}

func Test_getCodes6(t *testing.T) {
	w := Word{Text: "yui", Language: "Rus"}
	w.getCodes()
	if !pkg.IntArrEqual(w.Codes, []int{0, 0, 0}) {
		t.Error("Expected [0, 0, 0], got ", w.Codes)
	}
}

func Test_getCodes7(t *testing.T) {
	w := Word{Text: "абв", Language: "Rus", Codes: []int{1, 1, 1}}
	w.getCodes()
	if !pkg.IntArrEqual(w.Codes, []int{70, 188, 68}) {
		t.Error("Expected [70, 188, 68], got ", w.Codes)
	}
}

func Test_costDistanceSymbol1(t *testing.T) {
	first := "test"
	firstW := Word{Text: first, Language: "Eng"}
	second := "try"
	secondW := Word{Text: second, Language: "Eng"}
	firstW.getCodes()
	secondW.getCodes()
	cost := costDistanceSymbol(firstW, secondW, 0, 0)
	if cost != 0 {
		t.Error("Expected 0, got ", cost)
	}
}

func Test_costDistanceSymbol2(t *testing.T) {
	first := "test"
	firstW := Word{Text: first, Language: "Eng"}
	second := "try"
	secondW := Word{Text: second, Language: "Eng"}
	firstW.getCodes()
	secondW.getCodes()
	cost := costDistanceSymbol(firstW, secondW, 3, 0)
	if cost != 0 {
		t.Error("Expected 0, got", cost)
	}
}

func Test_costDistanceSymbol3(t *testing.T) {
	first := "test"
	firstW := Word{Text: first, Language: "Eng"}
	second := "try"
	secondW := Word{Text: second, Language: "Eng"}
	firstW.getCodes()
	secondW.getCodes()
	cost := costDistanceSymbol(firstW, secondW, 0, 1)
	if cost != 1 {
		t.Error("Expected 1, got ", cost)
	}
}

func Test_costDistanceSymbol4(t *testing.T) {
	first := "test"
	firstW := Word{Text: first, Language: "Eng"}
	second := "try"
	secondW := Word{Text: second, Language: "Eng"}
	firstW.getCodes()
	secondW.getCodes()
	cost := costDistanceSymbol(firstW, secondW, 1, 0)
	if cost != 2 {
		t.Error("Expected 2, got ", cost)
	}
}

func Test_Levenstein1(t *testing.T) {
	first := ""
	second := ""
	ratio := Levenshtein(first, second, "Rus")
	if ratio != 1 {
		t.Error("Expected 1, got ", ratio)
	}
}

func Test_Levenstein2(t *testing.T) {
	first := "test"
	second := ""
	ratio := Levenshtein(first, second, "Rus")
	if ratio != 0 {
		t.Error("Expected 0, got ", ratio)
	}
}

func Test_Levenstein3(t *testing.T) {
	first := "test"
	second := "test"
	ratio := Levenshtein(first, second, "Rus")
	if ratio != 1 {
		t.Error("Expected 0, got ", ratio)
	}
}

func Test_Levenstein4(t *testing.T) {
	first := "тес"
	second := "тест"
	ratio := Levenshtein(first, second, "Rus")
	if ratio == 0 {
		t.Error("Expected > 0, got ", ratio)
	}
}

func Test_Levenstein5(t *testing.T) {
	first := "test"
	second := "tes"
	ratio := Levenshtein(first, second, "Eng")
	if ratio == 0 {
		t.Error("Expected > 0, got ", ratio)
	}
}

func Test_Levenstein6(t *testing.T) {
	first := "test"
	second := "tes"
	ratio := Levenshtein(first, second, "Test")
	if ratio != 0 {
		t.Error("Expected 0, got ", ratio)
	}
}

func Test_Levenstein7(t *testing.T) {
	first := "test"
	second := "testшщ"
	ratio := Levenshtein(first, second, "Rus")
	fmt.Println(ratio)
	if ratio == 0 {
		t.Error("Expected > 0, got ", ratio)
	}
}

func Test_CompareWords1(t *testing.T) {
	comp := Comparator{SubtokenLength: 2}
	first := ""
	second := ""
	ratio := comp.CompareWords(first, second)
	if ratio != 1 {
		t.Error("Expected 1, got ", ratio)
	}
}

func Test_CompareWords2(t *testing.T) {
	comp := Comparator{SubtokenLength: 2}
	first := "test"
	second := "test"
	ratio := comp.CompareWords(first, second)
	if ratio != 1 {
		t.Error("Expected 1, got ", ratio)
	}
}

func Test_CompareWords3(t *testing.T) {
	comp := Comparator{SubtokenLength: 2}
	first := "test"
	second := ""
	ratio := comp.CompareWords(first, second)
	if ratio != 0 {
		t.Error("Expected 0, got ", ratio)
	}
}

func Test_CompareWords4(t *testing.T) {
	comp := Comparator{SubtokenLength: 2}
	first := "test"
	second := "тест"
	ratio := comp.CompareWords(first, second)
	if ratio != 0 {
		t.Error("Expected 0, got ", ratio)
	}
}

func Test_CompareWords5(t *testing.T) {
	comp := Comparator{SubtokenLength: 2}
	first := "test"
	second := "tes"
	ratio := comp.CompareWords(first, second)
	if ratio == 0 {
		t.Error("Expected > 0, got ", ratio)
	}
}

func Test_dedup1(t *testing.T) {
	token := ""
	ans := dedup(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_dedup2(t *testing.T) {
	token := "f"
	ans := dedup(token)
	if ans != "f" {
		t.Error("Expected \"f\", got ", ans)
	}
}

func Test_dedup3(t *testing.T) {
	token := "bitter"
	ans := dedup(token)
	if ans != "biter" {
		t.Error("Expected \"biter\", got ", ans)
	}
}

func Test_dedup4(t *testing.T) {
	token := "soccer"
	ans := dedup(token)
	if ans != "soccer" {
		t.Error("Expected \"soccer\", got ", ans)
	}
}

func Test_dropInitialLetters1(t *testing.T) {
	token := ""
	ans := dropInitialLetters(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_dropInitialLetters2(t *testing.T) {
	token := "knit"
	ans := dropInitialLetters(token)
	if ans != "nit" {
		t.Error("Expected \"nit\", got ", ans)
	}
}

func Test_dropBafterMAtEnd1(t *testing.T) {
	token := ""
	ans := dropBafterMAtEnd(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_dropBafterMAtEnd2(t *testing.T) {
	token := "bomb"
	ans := dropBafterMAtEnd(token)
	if ans != "bom" {
		t.Error("Expected \"bom\", got ", ans)
	}
}

func Test_cTransform1(t *testing.T) {
	token := ""
	ans := cTransform(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_cTransform2(t *testing.T) {
	token := "c"
	ans := cTransform(token)
	if ans != "k" {
		t.Error("Expected \"c\", got ", ans)
	}
}

func Test_cTransform3(t *testing.T) {
	token := "ci"
	ans := cTransform(token)
	if ans != "si" {
		t.Error("Expected \"si\", got ", ans)
	}
}

func Test_cTransform4(t *testing.T) {
	token := "cia"
	ans := cTransform(token)
	if ans != "xia" {
		t.Error("Expected \"xia\", got ", ans)
	}
}

func Test_dTransform1(t *testing.T) {
	token := ""
	ans := dTransform(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_dTransform2(t *testing.T) {
	token := "dge"
	ans := dTransform(token)
	if ans != "jge" {
		t.Error("Expected \"jge\", got ", ans)
	}
}

func Test_dTransform3(t *testing.T) {
	token := "d"
	ans := dTransform(token)
	if ans != "t" {
		t.Error("Expected \"t\", got ", ans)
	}
}

func Test_dropG1(t *testing.T) {
	token := ""
	ans := dropG(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_dropG2(t *testing.T) {
	token := "ghe"
	ans := dropG(token)
	if ans != "ghe" {
		t.Error("Expected \"ghe\", got ", ans)
	}
}

func Test_dropG3(t *testing.T) {
	token := "tegh"
	ans := dropG(token)
	if ans != "tegh" {
		t.Error("Expected \"tegh\", got ", ans)
	}
}

func Test_dropG4(t *testing.T) {
	token := "agned"
	ans := dropG(token)
	if ans != "aned" {
		t.Error("Expected \"aned\", got ", ans)
	}
}

func Test_transformG1(t *testing.T) {
	token := ""
	ans := transformG(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformG2(t *testing.T) {
	token := "gi"
	ans := transformG(token)
	if ans != "ji" {
		t.Error("Expected \"ji\", got ", ans)
	}
}

func Test_transformG3(t *testing.T) {
	token := "g"
	ans := transformG(token)
	if ans != "k" {
		t.Error("Expected \"k\", got ", ans)
	}
}

func Test_dropH1(t *testing.T) {
	token := ""
	ans := dropH(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_dropH2(t *testing.T) {
	token := "wahr"
	ans := dropH(token)
	if ans != "war" {
		t.Error("Expected \"war\", got ", ans)
	}
}

func Test_dropH3(t *testing.T) {
	token := "ha"
	ans := dropH(token)
	if ans != "ha" {
		t.Error("Expected \"ha\", got ", ans)
	}
}

func Test_transformCK1(t *testing.T) {
	token := ""
	ans := transformCK(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformCK2(t *testing.T) {
	token := "ck"
	ans := transformCK(token)
	if ans != "k" {
		t.Error("Expected \"k\", got ", ans)
	}
}

func Test_transformPH1(t *testing.T) {
	token := ""
	ans := transformPH(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformPH2(t *testing.T) {
	token := "ph"
	ans := transformPH(token)
	if ans != "f" {
		t.Error("Expected \"f\", got ", ans)
	}
}

func Test_transformQ1(t *testing.T) {
	token := ""
	ans := transformQ(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformQ2(t *testing.T) {
	token := "q"
	ans := transformQ(token)
	if ans != "k" {
		t.Error("Expected \"k\", got ", ans)
	}
}

func Test_transformV1(t *testing.T) {
	token := ""
	ans := transformV(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformV2(t *testing.T) {
	token := "v"
	ans := transformV(token)
	if ans != "f" {
		t.Error("Expected \"f\", got ", ans)
	}
}

func Test_transformZ1(t *testing.T) {
	token := ""
	ans := transformZ(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformZ2(t *testing.T) {
	token := "z"
	ans := transformZ(token)
	if ans != "s" {
		t.Error("Expected \"s\", got ", ans)
	}
}

func Test_transformS1(t *testing.T) {
	token := ""
	ans := transformS(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformS2(t *testing.T) {
	token := "sia"
	ans := transformS(token)
	if ans != "xia" {
		t.Error("Expected \"xia\", got ", ans)
	}
}

func Test_transformT1(t *testing.T) {
	token := ""
	ans := transformT(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformT2(t *testing.T) {
	token := "tia"
	ans := transformT(token)
	if ans != "xia" {
		t.Error("Expected \"xia\", got ", ans)
	}
}

func Test_transformT3(t *testing.T) {
	token := "th"
	ans := transformT(token)
	if ans != "0" {
		t.Error("Expected \"0\", got ", ans)
	}
}

func Test_dropT1(t *testing.T) {
	token := ""
	ans := dropT(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_dropT2(t *testing.T) {
	token := "tch"
	ans := dropT(token)
	if ans != "ch" {
		t.Error("Expected \"ch\", got ", ans)
	}
}

func Test_transformWH1(t *testing.T) {
	token := ""
	ans := transformWH(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformWH2(t *testing.T) {
	token := "whale"
	ans := transformWH(token)
	if ans != "wale" {
		t.Error("Expected \"wale\", got ", ans)
	}
}

func Test_transformWH3(t *testing.T) {
	token := "mowh"
	ans := transformWH(token)
	if ans != "mowh" {
		t.Error("Expected \"0\", got ", ans)
	}
}

func Test_dropW1(t *testing.T) {
	token := ""
	ans := dropW(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_dropW2(t *testing.T) {
	token := "wow"
	ans := dropW(token)
	if ans != "wo" {
		t.Error("Expected \"wo\", got ", ans)
	}
}

func Test_transformX1(t *testing.T) {
	token := ""
	ans := transformX(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformX2(t *testing.T) {
	token := "xin"
	ans := transformX(token)
	if ans != "sin" {
		t.Error("Expected \"sin\", got ", ans)
	}
}

func Test_transformX3(t *testing.T) {
	token := "box"
	ans := transformX(token)
	if ans != "boks" {
		t.Error("Expected \"boks\", got ", ans)
	}
}

func Test_dropVowels1(t *testing.T) {
	token := ""
	ans := dropVowels(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_dropVowels2(t *testing.T) {
	token := "ara"
	ans := dropVowels(token)
	if ans != "ar" {
		t.Error("Expected \"ar\", got ", ans)
	}
}

func Test_Metaphone1(t *testing.T) {
	token := ""
	ans := Metaphone(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_Metaphone2(t *testing.T) {
	token := "Serafim"
	ans := Metaphone(token)
	if ans != "SRFM" {
		t.Error("Expected \"SRFM\", got ", ans)
	}
}

func Test_Metaphone3(t *testing.T) {
	token := "тест"
	ans := Metaphone(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_MetaphoneCompare1(t *testing.T) {
	first := ""
	second := ""
	ans := MetaphoneCompare(first, second)
	if ans != 1 {
		t.Error("Expected 1, got ", ans)
	}
}

func Test_MetaphoneCompare2(t *testing.T) {
	first := "test"
	second := ""
	ans := MetaphoneCompare(first, second)
	if ans != 0 {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_MetaphoneCompare3(t *testing.T) {
	first := "test"
	second := "tezt"
	ans := MetaphoneCompare(first, second)
	if ans == 0 {
		t.Error("Expected > 0, got ", ans)
	}
}
func Test_MetaphoneCompare4(t *testing.T) {
	first := "test"
	second := "тест"
	ans := MetaphoneCompare(first, second)
	if ans != 0 {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_dropDoubles1(t *testing.T) {
	token := ""
	ans := dropDoubles(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_dropDoubles2(t *testing.T) {
	token := "c"
	ans := dropDoubles(token)
	if ans != "c" {
		t.Error("Expected \"c\", got ", ans)
	}
}

func Test_dropDoubles3(t *testing.T) {
	token := "cc"
	ans := dropDoubles(token)
	if ans != "c" {
		t.Error("Expected \"c\", got ", ans)
	}
}

func Test_dropDoubles4(t *testing.T) {
	token := "ccc"
	ans := dropDoubles(token)
	if ans != "c" {
		t.Error("Expected \"c\", got ", ans)
	}
}

func Test_transformVowels1(t *testing.T) {
	token := ""
	ans := transformVowels(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformVowels2(t *testing.T) {
	token := "йо"
	ans := transformVowels(token)
	if ans != "и" {
		t.Error("Expected \"и\", got ", ans)
	}
}

func Test_transformVowels3(t *testing.T) {
	token := "о"
	ans := transformVowels(token)
	if ans != "а" {
		t.Error("Expected \"а\", got ", ans)
	}
}

func Test_transformVowels4(t *testing.T) {
	token := "е"
	ans := transformVowels(token)
	if ans != "и" {
		t.Error("Expected \"и\", got ", ans)
	}
}

func Test_transformVowels5(t *testing.T) {
	token := "ю"
	ans := transformVowels(token)
	if ans != "у" {
		t.Error("Expected \"у\", got ", ans)
	}
}

func Test_transformConsonants1(t *testing.T) {
	token := ""
	ans := transformConsonants(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformConsonants2(t *testing.T) {
	token := "б"
	ans := transformConsonants(token)
	if ans != "п" {
		t.Error("Expected \"п\", got ", ans)
	}
}

func Test_transformConsonants3(t *testing.T) {
	token := "з"
	ans := transformConsonants(token)
	if ans != "с" {
		t.Error("Expected \"с\", got ", ans)
	}
}

func Test_transformConsonants4(t *testing.T) {
	token := "д"
	ans := transformConsonants(token)
	if ans != "т" {
		t.Error("Expected \"т\", got ", ans)
	}
}

func Test_transformConsonants5(t *testing.T) {
	token := "вл"
	ans := transformConsonants(token)
	if ans != "вл" {
		t.Error("Expected \"вл\", got ", ans)
	}
}

func Test_transformConsonants6(t *testing.T) {
	token := "гх"
	ans := transformConsonants(token)
	if ans != "кх" {
		t.Error("Expected \"кх\", got ", ans)
	}
}

func Test_transformRusDoubles1(t *testing.T) {
	token := ""
	ans := transformRusDoubles(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_transformRusDoubles2(t *testing.T) {
	token := "тс"
	ans := transformRusDoubles(token)
	if ans != "ц" {
		t.Error("Expected \"ц\", got ", ans)
	}
}

func Test_dropExtra1(t *testing.T) {
	token := ""
	ans := dropExtra(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_dropExtra2(t *testing.T) {
	token := "ъ"
	ans := dropExtra(token)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_isConsonantExcept1(t *testing.T) {
	token := ""
	ans := isConsonantExcept(token)
	if ans != false {
		t.Error("Expected false, got ", ans)
	}
}

func Test_isConsonantExcept2(t *testing.T) {
	token := "л"
	ans := isConsonantExcept(token)
	if ans != false {
		t.Error("Expected false, got ", ans)
	}
}

func Test_isConsonantExcept3(t *testing.T) {
	token := "б"
	ans := isConsonantExcept(token)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}

func Test_isSonorous1(t *testing.T) {
	token := ""
	ans := isSonorous(token)
	if ans != false {
		t.Error("Expected false, got ", ans)
	}
}

func Test_isSonorous2(t *testing.T) {
	token := "л"
	ans := isSonorous(token)
	if ans != false {
		t.Error("Expected false, got ", ans)
	}
}

func Test_isSonorous3(t *testing.T) {
	token := "б"
	ans := isSonorous(token)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}

func Test_RusMetaphoneCompare1(t *testing.T) {
	first := ""
	second := ""
	ans := RusMetaphoneCompare(first, second)
	if ans != 1 {
		t.Error("Expected 1, got ", ans)
	}
}

func Test_RusMetaphoneCompare2(t *testing.T) {
	first := "тест"
	second := ""
	ans := RusMetaphoneCompare(first, second)
	if ans != 0 {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_RusMetaphoneCompare3(t *testing.T) {
	first := "тест"
	second := "тесд"
	ans := RusMetaphoneCompare(first, second)
	if ans != 1 {
		t.Error("Expected 1, got ", ans)
	}
}
func Test_RusMetaphoneCompare4(t *testing.T) {
	first := "test"
	second := "тест"
	ans := RusMetaphoneCompare(first, second)
	if ans != 0 {
		t.Error("Expected 0, got ", ans)
	}
}
