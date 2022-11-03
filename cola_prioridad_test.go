package cola_prioridad_test

import (
	TDAHeap "cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeapVacio(t *testing.T) {
	comparar := func(clave1 int, clave2 int) int {
		if clave1 < clave2 {
			return -1
		} else if clave1 > clave2 {
			return 1
		}
		return 0
	}
	heap := TDAHeap.CrearHeap[int](comparar)
	require.True(t, heap.EstaVacia())
}
