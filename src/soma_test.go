package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(5, 5)

	if resultado != 10 {
		t.Errorf("Resultado esperado: %d, Resultado obtido: %d", 10, resultado)
	}
}
