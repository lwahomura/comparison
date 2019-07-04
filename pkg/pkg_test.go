package pkg

import (
	"testing"
)

func Test_BuildString1(t *testing.T) {
	var arr []string
	ans := BuildString(arr)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_BuildStrin2(t *testing.T) {
	arr := []string{"first"}
	ans := BuildString(arr)
	if ans != "first" {
		t.Error("Expected \"first\", got ", ans)
	}
}

func Test_BuildString3(t *testing.T) {
	arr := []string{"first", "второй"}
	ans := BuildString(arr)
	if ans != "first второй" {
		t.Error("Expected \"first второй\", got ", ans)
	}
}

func Test_DivideSentence1(t *testing.T) {
	sentence := ""
	ans := DivideSentence(sentence)
	if len(ans) != 0 {
		t.Error("Expected empty array, got ", ans)
	}
}

func Test_DivideSentence2(t *testing.T) {
	sentence := "first1 Втор$ой"
	ans := DivideSentence(sentence)
	if ans[0] != "first1" && ans[1] != "второй" {
		t.Error("Expected [first1 второй], got ", ans)
	}
}

func Test_NormalizeWord1(t *testing.T) {
	sentence := ""
	ans := NormalizeWord(sentence)
	if ans != "" {
		t.Error("Expected empty string, got ", ans)
	}
}

func Test_NormalizeWord2(t *testing.T) {
	sentence := "АЁ~f 2"
	ans := NormalizeWord(sentence)
	if ans != "аёf2" {
		t.Error("Expected \"аёf2\", got ", ans)
	}
}

func Test_MaxRatio1(t *testing.T) {
	a := 0
	b := 1
	c := 0
	ans := MaxRatio(a, b, c)
	if ans != 0 {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_MaxRatio2(t *testing.T) {
	a := 0
	b := 1
	c := 3
	ans := MaxRatio(a, b, c)
	if ans != 3 {
		t.Error("Expected 3, got ", ans)
	}
}

func Test_MaxRatio3(t *testing.T) {
	a := 0
	b := 0
	c := 3
	ans := MaxRatio(a, b, c)
	if ans != 0 {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_MinRatio1(t *testing.T) {
	a := 0
	b := 1
	c := 0
	ans := MinRatio(a, b, c)
	if ans != 0 {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_MinRatio2(t *testing.T) {
	a := 0
	b := 1
	c := 3
	ans := MinRatio(a, b, c)
	if ans != 0 {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_MinRatio3(t *testing.T) {
	a := 1
	b := 2
	c := 3
	ans := MinRatio(a, b, c)
	if ans != 3 {
		t.Error("Expected 3, got ", ans)
	}
}

func Test_XorIfEmpty1(t *testing.T) {
	a := 0
	b := 0
	ans := XorIfEmpty(a, b)
	if ans != false {
		t.Error("Expected false, got ", ans)
	}
}

func Test_XorIfEmpty2(t *testing.T) {
	a := 0
	b := 3
	ans := XorIfEmpty(a, b)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}
func Test_XorIfEmpty3(t *testing.T) {
	a := 4
	b := 0
	ans := XorIfEmpty(a, b)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}

func Test_XorIfEmpty4(t *testing.T) {
	a := 3
	b := 1
	ans := XorIfEmpty(a, b)
	if ans != false {
		t.Error("Expected false, got ", ans)
	}
}

func Test_MaxFloat1(t *testing.T) {
	a := float32(0)
	b := float32(0)
	ans := MaxFloat(a, b)
	if ans != b {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_MaxFloat2(t *testing.T) {
	a := float32(0)
	b := float32(0.1)
	ans := MaxFloat(a, b)
	if ans != b {
		t.Error("Expected 0.1, got ", ans)
	}
}

func Test_MaxFloat3(t *testing.T) {
	a := float32(0.3)
	b := float32(0)
	ans := MaxFloat(a, b)
	if ans != a {
		t.Error("Expected 0.3, got ", ans)
	}
}

func Test_MinInt1(t *testing.T) {
	a := 0
	b := 0
	ans := MinInt(a, b)
	if ans != b {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_MinInt2(t *testing.T) {
	a := 1
	b := 0
	ans := MinInt(a, b)
	if ans != b {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_MinInt3(t *testing.T) {
	a := 0
	b := 1
	ans := MinInt(a, b)
	if ans != a {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_MaxInt1(t *testing.T) {
	a := 0
	b := 0
	ans := MaxInt(a, b)
	if ans != b {
		t.Error("Expected 0, got ", ans)
	}
}

func Test_MaxInt2(t *testing.T) {
	a := 1
	b := 0
	ans := MaxInt(a, b)
	if ans != a {
		t.Error("Expected 1, got ", ans)
	}
}

func Test_MaxInt3(t *testing.T) {
	a := 0
	b := 1
	ans := MaxInt(a, b)
	if ans != b {
		t.Error("Expected 1, got ", ans)
	}
}

func Test_ArrContains1(t *testing.T) {
	array := make([]int, 0)
	elem := 0
	ans := ArrContains(array, elem)
	if ans != false {
		t.Error("Expected false, got ", ans)
	}
}

func Test_ArrContains2(t *testing.T) {
	array := []int{0}
	elem := 0
	ans := ArrContains(array, elem)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}

func Test_ArrContains3(t *testing.T) {
	array := []int{0}
	elem := 1
	ans := ArrContains(array, elem)
	if ans != false {
		t.Error("Expected false, got ", ans)
	}
}

func Test_IntArrEqual1(t *testing.T) {
	a := make([]int, 0)
	b := make([]int, 0)
	ans := IntArrEqual(a, b)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}

func Test_IntArrEqual2(t *testing.T) {
	a := []int{0}
	b := []int{0}
	ans := IntArrEqual(a, b)
	if ans != true {
		t.Error("Expected true, got ", ans)
	}
}

func Test_IntArrEqual3(t *testing.T) {
	a := []int{1}
	b := []int{0}
	ans := IntArrEqual(a, b)
	if ans != false {
		t.Error("Expected false, got ", ans)
	}
}
