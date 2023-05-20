package unitarytest

import "testing"

/* func TestSum(t *testing.T) {
	total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Error en  Suma, resultado: %d, esperado: %d", total, 10)
	}
} */

func TestSum(t *testing.T) {
	table := []struct{ a, b, c int }{
		{1, 2, 3},
		{3, 5, 8},
		{8, 8, 16},
	}

	for _, v := range table {
		result := v.a + v.b
		if result != v.c {
			t.Errorf("Error en  Suma, resultado: %d, esperado: %d", result, v.c)
		}
	}

}
