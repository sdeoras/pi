package pi

import (
	"testing"
)

func TestGetter_Get_Start_0(t *testing.T) {
	g := NewGetter()
	digits, decimal, err := g.Get(0, 10)

	if err != nil {
		t.Fatal(err)
	}

	if decimal != 1 {
		t.Fatal("expected decimal to be present at index 1, got", decimal)
	}

	if string(digits) != string([]byte{3, 1, 4, 1, 5, 9, 2, 6, 5}) {
		t.Fatal("expected [3,1,4,1,5,9,2,6,5], got", digits)
	}
}

func TestGetter_Get_Start_1(t *testing.T) {
	g := NewGetter()
	digits, decimal, err := g.Get(1, 10)

	if err != nil {
		t.Fatal(err)
	}

	if decimal != 0 {
		t.Fatal("expected decimal to be present at index 0, got", decimal)
	}

	if string(digits) != string([]byte{1, 4, 1, 5, 9, 2, 6, 5, 3}) {
		t.Fatal("expected [1,4,1,5,9,2,6,5,3], got", digits)
	}
}

func TestGetter_Get_Start_10(t *testing.T) {
	g := NewGetter()
	digits, decimal, err := g.Get(10, 10)

	if err != nil {
		t.Fatal(err)
	}

	if decimal != -1 {
		t.Fatal("expected decimal to be present at index 0, got", decimal)
	}

	if string(digits) != string([]byte{3, 5, 8, 9, 7, 9, 3, 2, 3, 8}) {
		t.Fatal("expected [3,5,8,9,7,9,3,2,3,8], got", digits)
	}
}

func TestGetter_Get_Start_Neg1(t *testing.T) {
	g := NewGetter()
	_, _, err := g.Get(-1, 10)

	if err == nil {
		t.Fatal("expected err but got nil")
	}
}

func TestGetter_Get_Start_Neg2(t *testing.T) {
	g := NewGetter()
	_, _, err := g.Get(0, -1)

	if err == nil {
		t.Fatal("expected err but got nil")
	}
}

func TestGetter_Get_Start_Neg3(t *testing.T) {
	g := NewGetter()
	_, _, err := g.Get(-1, -1)

	if err == nil {
		t.Fatal("expected err but got nil")
	}
}

func TestGetter_Get_Start_Zero(t *testing.T) {
	g := NewGetter()
	digits, decimal, err := g.Get(0, 0)

	if err != nil {
		t.Fatal(err)
	}

	if decimal != -1 {
		t.Fatal("expected decimal to be present at index 0, got", decimal)
	}

	if string(digits) != string([]byte{}) {
		t.Fatal("expected [], got", digits)
	}
}

func TestGetter_Get_Start_Zero2(t *testing.T) {
	g := NewGetter()
	digits, decimal, err := g.Get(100, 0)

	if err != nil {
		t.Fatal(err)
	}

	if decimal != -1 {
		t.Fatal("expected decimal to be present at index 0, got", decimal)
	}

	if string(digits) != string([]byte{}) {
		t.Fatal("expected [], got", digits)
	}
}

// ---------

func TestGetter_Pkg_Get_Start_0(t *testing.T) {
	digits, decimal, err := Get(0, 10)

	if err != nil {
		t.Fatal(err)
	}

	if decimal != 1 {
		t.Fatal("expected decimal to be present at index 1, got", decimal)
	}

	if string(digits) != string([]byte{3, 1, 4, 1, 5, 9, 2, 6, 5}) {
		t.Fatal("expected [3,1,4,1,5,9,2,6,5], got", digits)
	}
}

func TestGetter_Pkg_Get_Start_1(t *testing.T) {
	digits, decimal, err := Get(1, 10)

	if err != nil {
		t.Fatal(err)
	}

	if decimal != 0 {
		t.Fatal("expected decimal to be present at index 0, got", decimal)
	}

	if string(digits) != string([]byte{1, 4, 1, 5, 9, 2, 6, 5, 3}) {
		t.Fatal("expected [1,4,1,5,9,2,6,5,3], got", digits)
	}
}

func TestGetter_Pkg_Get_Start_10(t *testing.T) {
	digits, decimal, err := Get(10, 10)

	if err != nil {
		t.Fatal(err)
	}

	if decimal != -1 {
		t.Fatal("expected decimal to be present at index 0, got", decimal)
	}

	if string(digits) != string([]byte{3, 5, 8, 9, 7, 9, 3, 2, 3, 8}) {
		t.Fatal("expected [3,5,8,9,7,9,3,2,3,8], got", digits)
	}
}

func TestGetter_Pkg_Get_Start_Neg1(t *testing.T) {
	_, _, err := Get(-1, 10)

	if err == nil {
		t.Fatal("expected err but got nil")
	}
}

func TestGetter_Pkg_Get_Start_Neg2(t *testing.T) {
	_, _, err := Get(0, -1)

	if err == nil {
		t.Fatal("expected err but got nil")
	}
}

func TestGetter_Pkg_Get_Start_Neg3(t *testing.T) {
	_, _, err := Get(-1, -1)

	if err == nil {
		t.Fatal("expected err but got nil")
	}
}

func TestGetter_Pkg_Get_Start_Zero(t *testing.T) {
	digits, decimal, err := Get(0, 0)

	if err != nil {
		t.Fatal(err)
	}

	if decimal != -1 {
		t.Fatal("expected decimal to be present at index 0, got", decimal)
	}

	if string(digits) != string([]byte{}) {
		t.Fatal("expected [], got", digits)
	}
}

func TestGetter_Pkg_Get_Start_Zero2(t *testing.T) {
	digits, decimal, err := Get(100, 0)

	if err != nil {
		t.Fatal(err)
	}

	if decimal != -1 {
		t.Fatal("expected decimal to be present at index 0, got", decimal)
	}

	if string(digits) != string([]byte{}) {
		t.Fatal("expected [], got", digits)
	}
}
