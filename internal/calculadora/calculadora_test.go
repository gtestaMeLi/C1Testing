package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num2 := 3
	num1 := 5
	resultadoEsperado := 2
	// Se ejecuta el test
	resultado := Restar(num1, num2)
	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num2 := 0
	num1 := 5

	// Se ejecuta el test
	_, err := Dividir(num1, num2)
	// Se validan los resultados
	assert.NotNil(t, err)
}
