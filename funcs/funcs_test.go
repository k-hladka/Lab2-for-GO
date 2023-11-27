package funcs

import "testing"

func TestMinValues(t *testing.T) {
	x := MinValues(-10, 21, 23)
	res := -10.0
	if x != res {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f", x, res)
	}
}
func TestMinValues1(t *testing.T) {
	x := MinValues(10, -21, 23)
	res := -21.0
	if x != res {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f", x, res)
	}
}
func TestMinValues2(t *testing.T) {
	x := MinValues(10, 21, -23)
	res := -23.0
	if x != res {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f", x, res)
	}
}

func TestAvgValues(t *testing.T) {
	x := AvgValues(1, 2, 3)
	res := 2.0
	if x != res {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f", x, res)
	}
}

func TestAvgValues1(t *testing.T) {
	x := AvgValues(2, -10, 11)
	res := 1.0
	if x != res {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f", x, res)
	}
}

func TestDiffEq(t *testing.T) {
	x := DiffEq(10, 21.3)
	res := 3.13
	if x != res {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f", x, res)
	}
}
