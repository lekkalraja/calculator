package calc

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 3)

	if result != 4 {
		t.Errorf("Add Failed :: Expected : %v But got %v", 4, result)
	}

	result = Add(3, 2)

	if result == 4 {
		t.Errorf("Add Failed :: Expected : %v But got %v", result, 4)
	}
}

func TestSub(t *testing.T) {
	result := Sub(3, 2)

	if result != 1 {
		t.Errorf("Sub Failed :: Expected : %v But got %v", 1, result)
	}

}

func TestMul(t *testing.T) {
	result := Mul(3, 2)

	if result != 6 {
		t.Errorf("Mul Failed :: Expected : %v But got %v", 6, result)
	}

}

func TestDiv(t *testing.T) {
	result := Div(6, 2)

	if result != 3 {
		t.Errorf("Div Failed :: Expected : %v But got %v", 0, result)
	}

}
