package main

import (
	"testing"
)

// --- Слайд 12: Кодування на основі тестування (Приклад тесту) ---
// run test | debug test
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}
