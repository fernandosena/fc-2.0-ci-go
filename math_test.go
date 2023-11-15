package main
import "testing"

func TestSoma(t *testing.T){
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resiltado %d. Esperdado: %d", total, 30)
	}
}