package lista_test

import (
	TDALista "lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrarListaEnlazada[int]()
	validarListaVacia(lista, t)
}

/* Test de TDA Lista (sin iteradores) */

func TestConUnElemento(t *testing.T) {
	listaBool := TDALista.CrarListaEnlazada[bool]()
	listaBool.InsertarPrimero(true)
	require.EqualValues(t, 1, listaBool.Largo())
	require.EqualValues(t, true, listaBool.VerPrimero())
	require.EqualValues(t, true, listaBool.VerUltimo())
	require.EqualValues(t, true, listaBool.BorrarPrimero())
	validarListaVacia[bool](listaBool, t)

	listaBool.InsertarUltimo(false)
	require.EqualValues(t, false, listaBool.VerPrimero())
	require.EqualValues(t, false, listaBool.VerUltimo())
	require.EqualValues(t, false, listaBool.BorrarPrimero())
	require.EqualValues(t, 0, listaBool.Largo())
	validarListaVacia[bool](listaBool, t)
}

func TestConVariosElementos(t *testing.T) {
	listaStr := TDALista.CrarListaEnlazada[string]()
	var (
		a string = "Primero"
		b string = "Segundo"
		c string = "Tercero"
		d string = "Cuarto"
	)

	listaStr.InsertarPrimero(b)
	listaStr.InsertarPrimero(a)
	listaStr.InsertarUltimo(c)
	listaStr.InsertarUltimo(d)
	require.False(t, listaStr.EstaVacia())
	require.EqualValues(t, a, listaStr.VerPrimero())
	require.EqualValues(t, d, listaStr.VerUltimo())
	require.EqualValues(t, 4, listaStr.Largo())
}

func TestVolumen(t *testing.T) {
	listaInt := TDALista.CrarListaEnlazada[int]()

	for i := 5000; i >= 1; i-- {
		listaInt.InsertarPrimero(i)
		require.EqualValues(t, i, listaInt.VerPrimero())
	}
	require.EqualValues(t, 5000, listaInt.VerUltimo())
	for j := 5001; j <= 10000; j++ {
		listaInt.InsertarUltimo(j)
		require.EqualValues(t, j, listaInt.VerUltimo())
	}
	require.EqualValues(t, 1, listaInt.VerPrimero())

	for k := 1; k <= 10000; k++ {
		require.EqualValues(t, k, listaInt.BorrarPrimero())
		if k < 10000 {
			require.EqualValues(t, k+1, listaInt.VerPrimero())
		}
	}
	validarListaVacia[int](listaInt, t)
}

/* Test de TDA Lista con Iteradores */

func TestIterarVacio(t *testing.T) {
	lista := TDALista.CrarListaEnlazada[bool]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIterarUnElemento(t *testing.T) {
	listaStr := TDALista.CrarListaEnlazada[string]()
	listaStr.InsertarPrimero("Segundo")
	iter := listaStr.Iterador()
	require.EqualValues(t, "Segundo", iter.VerActual())
	iter.Insertar("Primero")
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, "Primero", listaStr.VerPrimero())
}

func validarListaVacia[T any](lista TDALista.Lista[T], test *testing.T) {
	require.True(test, lista.EstaVacia())
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.BorrarPrimero() })
}
