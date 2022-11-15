package cola_prioridad_test

import (
	TDAHeap "cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func compararInt(clave1 int, clave2 int) int {
	if clave1 < clave2 {
		return -1
	} else if clave1 > clave2 {
		return 1
	}
	return 0
}

func compararString(clave1 string, clave2 string) int {
	if clave1[0] < clave2[0] {
		return -1
	} else if clave1[0] > clave2[0] {
		return 1
	}
	return 0
}

func TestHeapsVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](compararInt)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestHeapConArrayVacio(t *testing.T) {
	arrayVacio := make([]int, 0)
	heapArray := TDAHeap.CrearHeapArr(arrayVacio, compararInt)
	require.True(t, heapArray.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapArray.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapArray.Desencolar() })
	require.EqualValues(t, 0, heapArray.Cantidad())
}

func TestGuardarYActualizaElMax(t *testing.T) {

	OrdenDeEntrada := []int{87, 51, 47, 38, 60, -10, 98, 101}

	heap := TDAHeap.CrearHeap(compararInt)

	for i := 0; i < len(OrdenDeEntrada); i++ {
		heap.Encolar(OrdenDeEntrada[i])
		if i < 6 {
			require.EqualValues(t, OrdenDeEntrada[0], heap.VerMax())
		} else {
			require.EqualValues(t, OrdenDeEntrada[i], heap.VerMax())
		}
	}

	require.EqualValues(t, 8, heap.Cantidad())
}

func TestDesencolarEnOrdenCorrecto(t *testing.T) {
	ordenDeEntrada := []string{"Teodoro", "Sofia", "Rodrigo", "Pablo", "Mauro", "Leandro", "Veronica", "Zoe"}
	ordenDeSalida := []string{"Zoe", "Veronica", "Teodoro", "Sofia", "Rodrigo", "Pablo", "Mauro", "Leandro"}
	heap := TDAHeap.CrearHeap[string](compararString)

	for i := 0; i < len(ordenDeEntrada); i++ {
		heap.Encolar(ordenDeEntrada[i])
	}
	j := 7
	for k := 0; k < len(ordenDeSalida); k++ {
		require.EqualValues(t, ordenDeSalida[k], heap.Desencolar())
		require.EqualValues(t, j, heap.Cantidad())
		j--
	}
}

func TestUnElemento(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](compararInt)
	heap.Encolar(0)
	heap.Desencolar()
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestUnElementoArr(t *testing.T) {
	arr := []string{"Unico Valor del Array"}
	heap := TDAHeap.CrearHeapArr[string](arr, compararString)
	require.EqualValues(t, "Unico Valor del Array", heap.VerMax())
	require.EqualValues(t, "Unico Valor del Array", heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestFuncionalidadHeapArray(t *testing.T) {
	array := []int{7, 2, 9, 1, 8, 3, 5, 4, 6, 10}

	heap := TDAHeap.CrearHeapArr[int](array, compararInt)

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 10, heap.Cantidad())
	require.EqualValues(t, 10, heap.VerMax())
	for i := 10; i >= 5; i-- {
		require.EqualValues(t, i, heap.Desencolar())
		require.EqualValues(t, i-1, heap.VerMax())
	}
	require.EqualValues(t, 4, heap.Cantidad())
	heap.Encolar(11)
	require.EqualValues(t, 11, heap.VerMax())
	heap.Encolar(12)
	require.EqualValues(t, 12, heap.VerMax())
	heap.Encolar(13)
	require.EqualValues(t, 13, heap.VerMax())
	require.EqualValues(t, 7, heap.Cantidad())
}

func TestHeapSort(t *testing.T) {
	listaAlumnos := []string{"Maidana", "DeCarmen", "Podesta", "Alfano", "Fort"}
	TDAHeap.HeapSort(listaAlumnos, compararString)
	listaOrdenada := []string{"Alfano", "DeCarmen", "Fort", "Maidana", "Podesta"}

	for indice, alumno := range listaAlumnos {
		require.EqualValues(t, listaOrdenada[indice], alumno)
	}

	numerosDesordenados := []int{45, 32124, 12, 78, 2, 0, -123, -3}
	numerosOrdenados := []int{-123, -3, 0, 2, 12, 45, 78, 32124}
	TDAHeap.HeapSort(numerosDesordenados, compararInt)

	for i := 0; i < len(numerosOrdenados); i++ {
		require.EqualValues(t, numerosOrdenados[i], numerosDesordenados[i])
	}
}

func TestVolumen(t *testing.T) {
	volumen := TDAHeap.CrearHeap[int](compararInt)

	require.True(t, volumen.EstaVacia())

	for i := 0; i <= 1000; i++ {
		volumen.Encolar(i)
		require.EqualValues(t, i, volumen.VerMax())
	}

	for j := 5000; j > 1000; j-- {
		volumen.Encolar(j)
		require.EqualValues(t, 5000, volumen.VerMax())
	}

	for k := 5001; k <= 10000; k++ {
		volumen.Encolar(k)
		require.EqualValues(t, k, volumen.VerMax())
	}

	require.False(t, volumen.EstaVacia())

	for p := 10000; p >= 0; p-- {
		require.EqualValues(t, p, volumen.Desencolar())
		if p > 0 {
			require.EqualValues(t, p-1, volumen.VerMax())
		}
	}
	require.True(t, volumen.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { volumen.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { volumen.Desencolar() })
	require.EqualValues(t, 0, volumen.Cantidad())
}

func TestModificarArrOriginal(t *testing.T) {
	original := []int{4, 6, 7, 1, 90}
	heap := TDAHeap.CrearHeapArr[int](original, compararInt)
	ordenSalida := []int{90, 7, 6, 4, 1}

	for i := 0; i < len(ordenSalida); i++ {
		require.EqualValues(t, ordenSalida[i], heap.Desencolar())
		require.NotEqualValues(t, ordenSalida[i], original[i])
	}

	require.EqualValues(t, 4, original[0])
	require.EqualValues(t, 6, original[1])
	require.EqualValues(t, 7, original[2])
	require.EqualValues(t, 1, original[3])
	require.EqualValues(t, 90, original[4])

}
