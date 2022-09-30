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

func TestConUnElemento(t *testing.T) {
	listaBool := TDALista.CrearListaEnlazada[bool]()
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
	listaStr := TDALista.CrearListaEnlazada[string]()
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

func TestOrdenDeListado(t *testing.T) {
	lista := Lista.CrearListaEnlazada[string]()
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
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestVolumen(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()

	for i := 5000; i >= 1; i-- {
		listaInt.InsertarPrimero(i)
		require.EqualValues(t, i, listaInt.VerPrimero())
	}
	require.False(t, listaInt.EstaVacia())
	require.EqualValues(t, 1, listaInt.VerPrimero())
	require.EqualValues(t, 5000, listaInt.VerUltimo())
	require.EqualValues(t, 5000, lista.Largo())
	for j := 5001; j <= 10000; j++ {
		listaInt.InsertarUltimo(j)
		require.EqualValues(t, j, listaInt.VerUltimo())
		require.EqualValues(t, i, lista.Largo())
	}
	require.False(t, listaInt.EstaVacia())
	require.EqualValues(t, 1, listaInt.VerPrimero())
	require.EqualValues(t, 10000, listaInt.VerUltimo())
	require.EqualValues(t, 10000, lista.Largo())

	for k := 1; k <= 10000; k++ {
		require.EqualValues(t, k, listaInt.VerPrimero())
		require.EqualValues(t, k, listaInt.BorrarPrimero())
	}
	validarListaVacia[int](listaInt, t)
}

/* Test de TDA Lista con Iteradores */

func TestIterarVacio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIterarUnElemento(t *testing.T) {
	listaStr := TDALista.CrearListaEnlazada[string]()
	listaStr.InsertarPrimero("Segundo")
	iter := listaStr.Iterador()
	require.EqualValues(t, "Segundo", iter.VerActual())
	iter.Insertar("Primero")
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, "Primero", listaStr.VerPrimero())
	require.EqualValues(t, "Segundo", listaStr.VerUltimo())
	require.EqualValues(t, "Primero", iter.Borrar())
	require.EqualValues(t, "Segundo", iter.Borrar())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
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
	lista := Lista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(1)
	require.EqualValues(t, 1, iterador.VerActual())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, 1, iterador.Siguiente())
	require.False(t, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
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
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
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
	
	/*for i := 4; i <= 20; i += 4 {
		lista.InsertarUltimo(i)
	}
	for i := 24; i <= 40; i += 2 {
		lista.InsertarUltimo(i)
	}

	suma:=0

	lista.Iterar(func(elem int) bool {
		suma +=elem
		return elem % 2 == 0
	})
	require.EqualValues(t, x, suma)
	*/

}

func validarListaVacia[T any](lista TDALista.Lista[T], test *testing.T) {
	require.True(test, lista.EstaVacia())
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(test, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.EqualValues(t, 0, lista.Largo())
}
