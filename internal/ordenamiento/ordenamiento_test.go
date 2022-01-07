package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	s := []int{5, 4, 1, 2, 3}
	resultadoEsperado := []int{1, 2, 3, 4, 5}
	// Se ejecuta el test
	resultado := orderSlice(s)
	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
