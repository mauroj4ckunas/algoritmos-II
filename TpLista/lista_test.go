package lista_test

import (
	TDALista "lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	validarListaVacia(lista, t)
}

/* Test de TDA Lista (sin iteradores) */

func TestInsertarPrimero(t *testing.T) {
	listaBool := TDALista.CrearListaEnlazada[bool]()
	listaBool.InsertarPrimero(true)
	require.EqualValues(t, 1, listaBool.Largo())
	require.EqualValues(t, true, listaBool.VerPrimero())
	require.EqualValues(t, true, listaBool.VerUltimo())
	require.EqualValues(t, true, listaBool.BorrarPrimero())
}

func TestInsertarUltimo(t *testing.T) {
	listaBool := TDALista.CrearListaEnlazada[bool]()
	listaBool.InsertarUltimo(false)
	require.EqualValues(t, false, listaBool.VerPrimero())
	require.EqualValues(t, false, listaBool.VerUltimo())
	require.EqualValues(t, false, listaBool.BorrarPrimero())
	require.EqualValues(t, 0, listaBool.Largo())
}

func TestChicoDeInserciones(t *testing.T) {
	listaStr := TDALista.CrearListaEnlazada[string]()
	var (
		a string = "Primero"
		b string = "Segundo"
		c string = "Tercero"
		d string = "Cuarto"
	)

	listaStr.InsertarPrimero(b)
	require.EqualValues(t, b, listaStr.VerPrimero())
	require.EqualValues(t, 1, listaStr.Largo())
	listaStr.InsertarPrimero(a)
	require.EqualValues(t, b, listaStr.VerUltimo())
	require.EqualValues(t, a, listaStr.VerPrimero())
	require.EqualValues(t, 2, listaStr.Largo())
	listaStr.InsertarUltimo(c)
	require.EqualValues(t, c, listaStr.VerUltimo())
	require.EqualValues(t, 3, listaStr.Largo())
	listaStr.InsertarUltimo(d)
	require.EqualValues(t, d, listaStr.VerUltimo())
	require.False(t, listaStr.EstaVacia())
	require.EqualValues(t, 4, listaStr.Largo())
}

func TestOrdenDeListado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("Hola")
	require.EqualValues(t, 1, lista.Largo())
	lista.InsertarPrimero("Todo bien")
	require.EqualValues(t, 2, lista.Largo())
	lista.InsertarPrimero("Esto")
	require.EqualValues(t, 3, lista.Largo())
	lista.InsertarPrimero("Es")
	require.EqualValues(t, 4, lista.Largo())
	lista.InsertarPrimero("Una")
	require.EqualValues(t, 5, lista.Largo())
	lista.InsertarPrimero("Prueba")
	require.EqualValues(t, 6, lista.Largo())
	lista.InsertarPrimero("De")
	lista.InsertarPrimero("Funcionamiento")
	require.EqualValues(t, "Funcionamiento", lista.VerPrimero())
	require.EqualValues(t, "Hola", lista.VerUltimo())
	lista.InsertarUltimo("Buenas tardes")
	lista.InsertarUltimo("Este")
	lista.InsertarUltimo("Es")
	require.EqualValues(t, 11, lista.Largo())
	lista.InsertarUltimo("El")
	require.EqualValues(t, 12, lista.Largo())
	lista.InsertarUltimo("Ultimo")
	require.EqualValues(t, 13, lista.Largo())
	lista.InsertarUltimo("Yey")
	require.EqualValues(t, 14, lista.Largo())
	require.EqualValues(t, "Funcionamiento", lista.VerPrimero())
	require.EqualValues(t, "Yey", lista.VerUltimo())
	require.EqualValues(t, "Funcionamiento", lista.BorrarPrimero())
	require.EqualValues(t, "De", lista.VerPrimero())
	require.EqualValues(t, "Yey", lista.VerUltimo())
	require.EqualValues(t, "De", lista.BorrarPrimero())
	require.EqualValues(t, "Prueba", lista.VerPrimero())
	require.EqualValues(t, "Yey", lista.VerUltimo())
	require.EqualValues(t, "Prueba", lista.BorrarPrimero())
	require.EqualValues(t, "Una", lista.VerPrimero())
	require.EqualValues(t, "Yey", lista.VerUltimo())
	require.EqualValues(t, "Una", lista.BorrarPrimero())
	require.EqualValues(t, "Es", lista.VerPrimero())
	require.EqualValues(t, "Es", lista.BorrarPrimero())
	require.EqualValues(t, "Esto", lista.VerPrimero())
	require.EqualValues(t, "Esto", lista.BorrarPrimero())
	require.EqualValues(t, "Todo bien", lista.VerPrimero())
	require.EqualValues(t, "Todo bien", lista.BorrarPrimero())
	require.EqualValues(t, "Hola", lista.VerPrimero())
	require.EqualValues(t, "Hola", lista.BorrarPrimero())
	require.EqualValues(t, "Buenas tardes", lista.VerPrimero())
	require.EqualValues(t, "Buenas tardes", lista.BorrarPrimero())
	require.EqualValues(t, "Este", lista.VerPrimero())
	require.EqualValues(t, "Este", lista.BorrarPrimero())
	require.EqualValues(t, "Es", lista.VerPrimero())
	require.EqualValues(t, "Es", lista.BorrarPrimero())
	require.EqualValues(t, "El", lista.VerPrimero())
	require.EqualValues(t, "El", lista.BorrarPrimero())
	require.EqualValues(t, "Ultimo", lista.VerPrimero())
	require.EqualValues(t, "Ultimo", lista.BorrarPrimero())
	require.EqualValues(t, "Yey", lista.VerPrimero())
	require.EqualValues(t, "Yey", lista.BorrarPrimero())
	validarListaVacia(lista, t)
}

func TestVolumen(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()

	//volumetria agregando al final
	for i := 0; i < 10000; i++ {
		listaInt.InsertarUltimo(i)
		require.EqualValues(t, i, listaInt.VerUltimo())
		require.False(t, listaInt.EstaVacia())
		require.EqualValues(t, i+1, listaInt.Largo())
	}
	for i := 0; i < 10000; i++ {
		require.EqualValues(t, i, listaInt.VerPrimero())
		require.EqualValues(t, i, listaInt.BorrarPrimero())
	}
	validarListaVacia[int](listaInt, t)

	//Volumetria agregando al principio
	for i := 0; i < 10000; i++ {
		listaInt.InsertarPrimero(i)
		require.EqualValues(t, i, listaInt.VerPrimero())
		require.False(t, listaInt.EstaVacia())
		require.EqualValues(t, i+1, listaInt.Largo())
	}
	for i := 9999; i >= 0; i-- {
		require.EqualValues(t, i, listaInt.VerPrimero())
		require.EqualValues(t, i, listaInt.BorrarPrimero())
	}
	validarListaVacia[int](listaInt, t)

}

/* Test de TDA Lista con Iteradores */

func TestIterarVacio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()
	iter := lista.Iterador()
	validarIteradorVacio[bool](iter, t)
}

func TestIterarUnElemento(t *testing.T) {
	listaStr := TDALista.CrearListaEnlazada[string]()
	iter := listaStr.Iterador()
	iter.Insertar("Segundo")
	require.EqualValues(t, "Segundo", iter.VerActual())
	iter.Insertar("Primero")
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, "Primero", listaStr.VerPrimero())
	require.EqualValues(t, "Segundo", listaStr.VerUltimo())
	require.EqualValues(t, "Primero", iter.Borrar())
	require.EqualValues(t, "Segundo", iter.Borrar())
	validarIteradorVacio[string](iter, t)
	validarListaVacia[string](listaStr, t)
}

func TestBorrarConElIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("Elemento del Medio")
	lista.InsertarPrimero("Primero luego de eliminar el primero")
	lista.InsertarPrimero("Primero original")
	lista.InsertarUltimo("Anteultimo")
	lista.InsertarUltimo("Ultimo")

	require.EqualValues(t, "Primero original", lista.VerPrimero())
	require.EqualValues(t, "Ultimo", lista.VerUltimo())

	iter1 := lista.Iterador()
	require.EqualValues(t, "Primero original", iter1.Borrar())
	require.EqualValues(t, "Primero luego de eliminar el primero", iter1.Siguiente())
	require.EqualValues(t, "Elemento del Medio", iter1.Borrar())
	require.EqualValues(t, "Anteultimo", iter1.Siguiente())
	require.EqualValues(t, "Ultimo", iter1.Borrar())

	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, "Primero luego de eliminar el primero", lista.VerPrimero())
	require.EqualValues(t, "Anteultimo", lista.VerUltimo())

	//Confirmamos que Elemento del Medio fue eliminado
	iter2 := lista.Iterador()
	require.EqualValues(t, "Primero luego de eliminar el primero", iter2.Siguiente())
	require.EqualValues(t, "Anteultimo", iter2.Siguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.Siguiente() })
}

func TestIteradorExternoConUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(1)
	require.EqualValues(t, 1, iterador.VerActual())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, 1, iterador.Siguiente())
	require.False(t, iterador.HaySiguiente())
	validarIteradorVacio[int](iterador, t)
}

func TestIteradorExternoPasoPorPaso(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("Externo")
	lista.InsertarPrimero("Iterador")
	lista.InsertarPrimero("Del")
	lista.InsertarPrimero("Prueba")
	lista.InsertarPrimero("La")
	lista.InsertarPrimero("Es")
	lista.InsertarPrimero("Esta")
	lista.InsertarPrimero("Gusto")
	lista.InsertarPrimero("Mucho")
	lista.InsertarPrimero("Tal")
	lista.InsertarPrimero("Que")
	lista.InsertarPrimero("Hola")
	iterador := lista.Iterador()
	require.EqualValues(t, "Hola", iterador.VerActual())
	require.EqualValues(t, "Hola", iterador.Siguiente())
	require.EqualValues(t, "Que", iterador.VerActual())
	require.EqualValues(t, "Que", iterador.Siguiente())
	require.EqualValues(t, "Tal", iterador.VerActual())
	require.EqualValues(t, "Tal", iterador.Siguiente())
	require.EqualValues(t, "Mucho", iterador.VerActual())
	require.EqualValues(t, "Mucho", iterador.Siguiente())
	require.EqualValues(t, "Gusto", iterador.VerActual())
	require.EqualValues(t, "Gusto", iterador.Siguiente())
	require.EqualValues(t, "Esta", iterador.VerActual())
	require.EqualValues(t, "Esta", iterador.Siguiente())
	require.EqualValues(t, "Es", iterador.VerActual())
	require.EqualValues(t, "Es", iterador.Siguiente())
	require.EqualValues(t, "La", iterador.VerActual())
	require.EqualValues(t, "La", iterador.Siguiente())
	require.EqualValues(t, "Prueba", iterador.VerActual())
	require.EqualValues(t, "Prueba", iterador.Siguiente())
	require.EqualValues(t, "Del", iterador.VerActual())
	require.EqualValues(t, "Del", iterador.Siguiente())
	require.EqualValues(t, "Iterador", iterador.VerActual())
	require.EqualValues(t, "Iterador", iterador.Siguiente())
	require.EqualValues(t, "Externo", iterador.VerActual())
	require.EqualValues(t, "Externo", iterador.Siguiente())
}

func TestIteradorInterno(t *testing.T) {
	lista1 := TDALista.CrearListaEnlazada[int]()
	lista2 := TDALista.CrearListaEnlazada[int]()

	for i := 0; i <= 10; i++ {
		lista1.InsertarUltimo(i)
	}

	lista1.Iterar(func(elem int) bool {
		elem *= 2
		lista2.InsertarUltimo(elem)
		return true
	})

	for i := 0; i <= 10; i++ {
		require.EqualValues(t, i*2, lista2.VerPrimero())
		require.EqualValues(t, i*2, lista2.BorrarPrimero())
	}

	var (
		a string = "Mauro"
		b string = "Facundo"
		c string = "Martin"
		d string = "Nicolas"
		e string = "Wally"
	)

	dondeEstaWally := TDALista.CrearListaEnlazada[string]()
	dondeEstaWally.InsertarPrimero(a)
	dondeEstaWally.InsertarUltimo(b)
	dondeEstaWally.InsertarUltimo(e)
	dondeEstaWally.InsertarUltimo(d)
	dondeEstaWally.InsertarUltimo(c)
	encontrar := 0
	dondeEstaWally.Iterar(func(nombre string) bool {
		encontrar++
		return nombre != "Wally"
	})
	require.Equal(t, 3, encontrar)
}

func TestIteradorVacio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()
	var verdadero bool = true
	lista.Iterar(func(elem bool) bool {
		return verdadero == elem
	})
}

//validadores externos

func validarListaVacia[T any](lista TDALista.Lista[T], test *testing.T) {
	require.True(test, lista.EstaVacia())
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.EqualValues(test, 0, lista.Largo())
}

func validarIteradorVacio[T any](iter TDALista.IteradorLista[T], test *testing.T) {
	require.PanicsWithValue(test, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(test, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(test, "El iterador termino de iterar", func() { iter.Borrar() })
}
