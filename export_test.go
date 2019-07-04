package comparison

import "testing"

func Test_compMethods1(t *testing.T) {
	first := ""
	second := ""
	lang := ""
	ans := compMethods(first, second, lang)
	if ans != 1 {
		t.Error("Expected 1, got ", ans)
	}
}

func Test_compMethods2(t *testing.T) {
	first := "тест"
	second := ""
	lang := ""
	ans := compMethods(first, second, lang)
	if ans != 0 {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_compMethods3(t *testing.T) {
	first := "test"
	second := "тест"
	lang := "Rus"
	ans := compMethods(first, second, lang)
	if ans != 0 {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_compMethods4(t *testing.T) {
	first := "test"
	second := "тест"
	lang := "Eng"
	ans := compMethods(first, second, lang)
	if ans != 0 {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_compMethods5(t *testing.T) {
	first := "test"
	second := "tst"
	lang := "Eng"
	ans := compMethods(first, second, lang)
	if ans == 0 {
		t.Error("Expected > 0, got ", ans)
	}
}
func Test_compMethods6(t *testing.T) {
	first := "тесд"
	second := "тест"
	lang := "Rus"
	ans := compMethods(first, second, lang)
	if ans == 0 {
		t.Error("Expected > 0, got ", ans)
	}
}

func Test_Compare1(t *testing.T) {
	first := ""
	second := ""
	ans := Compare(first, second)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}

func Test_Compare2(t *testing.T) {
	first := "test"
	second := ""
	ans := Compare(first, second)
	if ans != false {
		t.Error("Expected false, got ", ans)
	}
}

func Test_Compare3(t *testing.T) {
	first := "test"
	second := "тест"
	ans := Compare(first, second)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}

func Test_Compare4(t *testing.T) {
	first := "test"
	second := "golang"
	ans := Compare(first, second)
	if ans != false {
		t.Error("Expected false, got ", ans)
	}
}

func Test_Compare5(t *testing.T) {
	first := "москвин"
	second := "moskvin"
	ans := Compare(first, second)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}

func Test_Compare6(t *testing.T) {
	first := "example.com"
	second := "john@example.com"
	ans := Compare(first, second)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}

func Test_Compare7(t *testing.T) {
	first := "example.com"
	second := "bobmarley@example.com"
	ans := Compare(first, second)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}

func Test_Compare8(t *testing.T) {
	first := "example.com"
	second := "johnsmith@example.com"
	ans := Compare(first, second)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}