package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTaxInBatch(t *testing.T) {
	cases := []struct {
		amount   float64
		expected float64
	}{
		{0, 0},
		{500, 5},
		{1000, 10},
		{1500, 10},
	}

	for _, c := range cases {
		result := CalculateTax(c.amount)

		if result != c.expected {
			t.Errorf("Expected %f, got %f", c.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500)
	}
}

// teste de mutação - testar valores que não previamos
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 0, 500, 1000, 1500, 1501.2}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount < 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}

		if amount > 20000 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}
	})
}
