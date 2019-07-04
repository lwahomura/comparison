package transliteration

import "testing"

func Test_processSignes1(t *testing.T) {
	word := ""
	ans := processSignes(word)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_processSignes2(t *testing.T) {
	word := "а"
	ans := processSignes(word)
	if ans != "а" {
		t.Error("Expected \"а\", got ", ans)
	}
}

func Test_processSignes3(t *testing.T) {
	word := "ъ"
	ans := processSignes(word)
	if ans != "" {
		t.Error("Expected \"\", got ", ans)
	}
}

func Test_processSignes4(t *testing.T) {
	word := "ье"
	ans := processSignes(word)
	if ans != "iе" {
		t.Error("Expected \"iе\", got ", ans)
	}
}

func Test_RusToEn1(t *testing.T) {
	target := ""
	ans := RusToEn(target)
	if ans != "" {
		t.Error("Expected empty strings, got ", ans)
	}
}

func Test_RusToEn2(t *testing.T) {
	target := "fi#rst"
	ans := RusToEn(target)
	if ans != "first" {
		t.Error("Expected \"first\", got ", ans)
	}
}

func Test_RusToEn3(t *testing.T) {
	target := "1съесть"
	ans := RusToEn(target)
	if ans != "1siest" {
		t.Error("Expected \"1siest\", got ", ans)
	}
}

func Test_RusToEn4(t *testing.T) {
	target := "first съесть"
	ans := RusToEn(target)
	if ans != "first siest" {
		t.Error("Expected \"first siest\", got ", ans)
	}
}

func Test_EnToRus1(t *testing.T) {
	target := ""
	ans := EnToRus(target)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_EnToRus2(t *testing.T) {
	target := "word"
	ans := EnToRus(target)
	if ans != "вод" {
		t.Error("Expected \"вод\", got ", ans)
	}
}

func Test_EnToRus3(t *testing.T) {
	target := "слово"
	ans := EnToRus(target)
	if ans != "слово" {
		t.Error("Expected \"слово\", got ", ans)
	}
}

func Test_EnToRus4(t *testing.T) {
	target := "wake"
	ans := EnToRus(target)
	if ans != "вак" {
		t.Error("Expected \"вак\", got ", ans)
	}
}

func Test_EnToRus5(t *testing.T) {
	target := "check"
	ans := EnToRus(target)
	if ans != "чэк" {
		t.Error("Expected \"чэк\", got ", ans)
	}
}

func Test_EnToRus6(t *testing.T) {
	target := "1make"
	ans := EnToRus(target)
	if ans != "1мак" {
		t.Error("Expected \"1мак\", got ", ans)
	}
}