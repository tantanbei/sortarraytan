package sortarraytan

import (
	"testing"
)

func TestInt(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	aR := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	b := IntArrayReverse(a)
	if b[1] != aR[1] {
		t.Error("int error")
	}
}

func TestByte(t *testing.T) {
	a := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	aR := []byte{'9', '8', '7', '6', '5', '4', '3', '2', '1'}
	b := ByteArrayReverse(a)
	if b[1] != aR[1] {
		t.Error("int error")
	}
}

func TestInterface(t *testing.T) {
	a := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	aR := []byte{'9', '8', '7', '6', '5', '4', '3', '2', '1'}
	b, _ := InterfaceArrayReverse(a)
	c := b.([]byte)
	if c[1] != aR[1] {
		t.Error("int error")
	}

	a2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	aR2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	b2, _ := InterfaceArrayReverse(a2)
	c2 := b2.([]int)
	if c2[1] != aR2[1] {
		t.Error("int error")
	}

	a3 := []string{"aa", "bb", "cc", "dd"}
	aR3 := []string{"dd", "cc", "bb", "aa"}
	b3, _ := InterfaceArrayReverse(a3)
	c3 := b3.([]string)
	if c3[1] != aR3[1] {
		t.Error("string error")
	}
}

func TestSearchMax(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	num, index := SearchMax(a)
	if num != 9 || index != 8 {
		t.Error("search max error")
	}
}

func TestSearchMin(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	num, index := SearchMin(a)
	if num != 1 || index != 0 {
		t.Error("search min error")
	}
}

func TestSearchNum(t *testing.T) {
	a := []int{1, 1, 2, 3, 4, 5, 6, 1, 5, 6, 7, 4, 2, 6, 74, 1}
	index := SearchNumber(a, 1)
	if len(index) != 4 || index[2] != 7 {
		t.Error("search num error")
	}
}

func TestIsExist(t *testing.T) {
	a := []int{1, 1, 2, 3, 4, 5, 6, 1, 5, 6, 7, 4, 2, 6, 74, 1}
	f := IsExist(a, 5)
	if f != true {
		t.Error("isExist error")
	}
	f = IsExist(a, 9)
	if f != false {
		t.Error("isExist error")
	}
}

func BenchmarkIntReInterface(b *testing.B) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		b, _ := InterfaceArrayReverse(a)
		c := b.([]int)
		_ = c
	}
}

func BenchmarkIntRe(b *testing.B) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		c := IntArrayReverse(a)
		_ = c
	}
}
