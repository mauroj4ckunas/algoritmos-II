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
	heap := TDAHeap.CrearHeap[int](comparar)
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

func TestGuardarYActualizaElMax(t *testing.T) {
	comparar := func(clave1 string, clave2 string) int {
		if clave1[0] < clave2[0] {
			return -1
		} else if clave1[0] > clave2[0] {
			return 1
		}
		return 0
	}

	OrdenDeEntrada := []string{"Teodoro", "Sofia", "Rodrigo", "Pablo", "Mauro", "Leandro", "Veronica", "Zoe"}

	heap := TDAHeap.CrearHeap[string](comparar)

	for i := 0; i < len(OrdenDeEntrada)-2; i++ {
		heap.Encolar(OrdenDeEntrada[i])
		require.EqualValues(t, OrdenDeEntrada[0], heap.VerMax())
	}

	heap.Encolar(OrdenDeEntrada[6])
	require.EqualValues(t, OrdenDeEntrada[6], heap.VerMax())

	heap.Encolar(OrdenDeEntrada[7])
	require.EqualValues(t, OrdenDeEntrada[7], heap.VerMax())

	require.EqualValues(t, 8, heap.Cantidad())
}

func TestDesencolarEnOrdenCorrecto(t *testing.T) {
	comparar := func(clave1 string, clave2 string) int {
		if clave1[0] < clave2[0] {
			return -1
		} else if clave1[0] > clave2[0] {
			return 1
		}
		return 0
	}
	ordenDeEntrada := []string{"Teodoro", "Sofia", "Rodrigo", "Pablo", "Mauro", "Leandro", "Veronica", "Zoe"}
	ordenDeSalida := []string{"Zoe", "Veronica", "Teodoro", "Sofia", "Rodrigo", "Pablo", "Mauro", "Leandro"}
	heap := TDAHeap.CrearHeap[string](comparar)

	for i := 0; i < len(ordenDeEntrada); i++ {
		heap.Encolar(ordenDeEntrada[i])
	}

	require.EqualValues(t, ordenDeSalida[0], heap.Desencolar())
	require.EqualValues(t, 7, heap.Cantidad())
	require.EqualValues(t, ordenDeSalida[1], heap.Desencolar())
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, ordenDeSalida[2], heap.Desencolar())
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, ordenDeSalida[3], heap.Desencolar())
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, ordenDeSalida[4], heap.Desencolar())
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, ordenDeSalida[5], heap.Desencolar())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, ordenDeSalida[6], heap.Desencolar())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, ordenDeSalida[7], heap.Desencolar())
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestUnElemento(t *testing.T) {
	comparar := func(clave1 int, clave2 int) int {
		if clave1 < clave2 {
			return -1
		} else if clave1 > clave2 {
			return 1
		}
		return 0
	}

	heap := TDAHeap.CrearHeap[int](comparar)
	heap.Encolar(0)
	heap.Desencolar()
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestUnElementoArr(t *testing.T) {
	comparar := func(clave1 string, clave2 string) int {
		if clave1[0] < clave2[0] {
			return -1
		} else if clave1[0] > clave2[0] {
			return 1
		}
		return 0
	}

	arr := []string{"Hola"}

	heap := TDAHeap.CrearHeapArr[string](arr, comparar)
	require.EqualValues(t, "Hola", heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestFuncionalidadHeapArray(t *testing.T) {
	comparar := func(clave1 int, clave2 int) int {
		if clave1 < clave2 {
			return -1
		} else if clave1 > clave2 {
			return 1
		}
		return 0
	}
	array := []int{7, 2, 9, 1, 8, 3, 5, 4, 6, 10}

	heap := TDAHeap.CrearHeapArr[int](array, comparar)

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
	compararString := func(clave1 string, clave2 string) int {
		if clave1[0] < clave2[0] {
			return -1
		} else if clave1[0] > clave2[0] {
			return 1
		}
		return 0
	}

	listaAlumnos := []string{"Maidana", "DeCarmen", "Podesta", "Alfano", "Fort"}
	TDAHeap.HeapSort(listaAlumnos, compararString)
	listaOrdenada := []string{"Alfano", "DeCarmen", "Fort", "Maidana", "Podesta"}

	for indice, alumno := range listaAlumnos {
		require.EqualValues(t, listaOrdenada[indice], alumno)
	}

	compararInt := func(clave1 int, clave2 int) int {
		if clave1 < clave2 {
			return -1
		} else if clave1 > clave2 {
			return 1
		}
		return 0
	}

	numerosDesordenados := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	TDAHeap.HeapSort(numerosDesordenados, compararInt)

	for k := 0; k <= 9; k++ {
		require.EqualValues(t, k+1, numerosDesordenados[k])
	}

	numerosDesordenados2 := []int{45, 32124, 12, 78, 2, 0, -123, -3}
	TDAHeap.HeapSort(numerosDesordenados2, compararInt)

	require.EqualValues(t, -123, numerosDesordenados2[0])
	require.EqualValues(t, -3, numerosDesordenados2[1])
	require.EqualValues(t, 0, numerosDesordenados2[2])
	require.EqualValues(t, 2, numerosDesordenados2[3])
	require.EqualValues(t, 12, numerosDesordenados2[4])
	require.EqualValues(t, 45, numerosDesordenados2[5])
	require.EqualValues(t, 78, numerosDesordenados2[6])
	require.EqualValues(t, 32124, numerosDesordenados2[7])
}

func TestVolumen(t *testing.T) {
	comparar := func(clave1 int, clave2 int) int {
		if clave1 < clave2 {
			return -1
		} else if clave1 > clave2 {
			return 1
		}
		return 0
	}

	volumen := TDAHeap.CrearHeap[int](comparar)

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
