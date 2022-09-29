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

func TestConUnElemento(t *testing.T) {
	listaBool := TDALista.CrarListaEnlazada[bool]()
	listaBool.InsertarPrimero(true)
	require.EqualValues(t, true, listaBool.VerPrimero())
	require.EqualValues(t, true, listaBool.VerUltimo())
	require.EqualValues(t, true, listaBool.BorrarPrimero())
	validarListaVacia[bool](listaBool, t)

	listaBool.InsertarUltimo(false)
	require.EqualValues(t, false, listaBool.VerPrimero())
	require.EqualValues(t, false, listaBool.VerUltimo())
	require.EqualValues(t, false, listaBool.BorrarPrimero())
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

	listaStr.InsertarPrimero(a)
	listaStr.InsertarPrimero(b)
	listaStr.InsertarUltimo(c)
	listaStr.InsertarUltimo(d)

	require.EqualValues(t, a, listaStr.VerPrimero())
	require.EqualValues(t, d, listaStr.VerUltimo())

}

func validarListaVacia[T any](lista TDALista.Lista[T], test *testing.T) {
	require.True(test, lista.EstaVacia())
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.BorrarPrimero() })

}
