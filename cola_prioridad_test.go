package cola_prioridad_test

import (
	TDAHeap "cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeapsVacio(t *testing.T) {
	comparar := func(clave1 int, clave2 int) int {
		if clave1 < clave2 {
			return -1
		} else if clave1 > clave2 {
			return 1
		}
		return 0
	}
	heap := TDAHeap.CrearHeap(comparar)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	arrayVacio := make([]int, 0)
	heapArray := TDAHeap.CrearHeapArr(arrayVacio, comparar)
	require.True(t, heapArray.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapArray.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapArray.Desencolar() })
	require.EqualValues(t, 0, heapArray.Cantidad())

}
