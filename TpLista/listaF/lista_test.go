package lista_test

import (
	"github.com/stretchr/testify/require"
	Lista "lista"
	"testing"
)

func TestListaVacia(t *testing.T) {
	lista := Lista.CrearListaEnlazada[string]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
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

func TestVolumetria(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	//volumetria agregando al final
	for i := 0; i < 10000; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, i, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, i+1, lista.Largo())
	}
	for i := 0; i < 10000; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())

	//Volumetria agregando al principio
	for i := 0; i < 10000; i++ {
		lista.InsertarPrimero(i)
		require.EqualValues(t, i, lista.VerPrimero())
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, i+1, lista.Largo())
	}
	for i := 9999; i >= 0; i-- {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestIteradorInterno(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	for i := 0; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}
	lista.Iterar(func(elem *int) bool{
		*elem *= 2
		return true
	})
	for i := 0; i <= 10; i++ {
		require.EqualValues(t, i * 2, lista.VerPrimero())
		require.EqualValues(t, i * 2, lista.BorrarPrimero())
	}

	for i := 4; i <= 20; i += 4 {
		lista.InsertarUltimo(i)
	}
	for i := 24; i <= 40; i += 2 {
		lista.InsertarUltimo(i)
	}

	lista.Iterar(func(elem *int) bool{
		*elem /=2
		return *elem % 2 == 0
	})
	for i := 4; i <= 24; i += 4 {
		require.EqualValues(t, i / 2, lista.VerPrimero())
		require.EqualValues(t, i / 2, lista.BorrarPrimero())
	}

	require.EqualValues(t, 13, lista.VerPrimero())
	require.EqualValues(t, 13, lista.BorrarPrimero())

	for i := 28; i <= 40; i += 2 {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}

}


func TestIteradorExterno1(t *testing.T) {
	lista := Lista.CrearListaEnlazada[string]()
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

func TestIteradorExterno2(t *testing.T) {
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