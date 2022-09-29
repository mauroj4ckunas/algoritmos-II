package lista_test

import (
	Lista "lista"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListaVacia(t *testing.T){
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
	for i:= 0; i < 10000 ; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, i, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, i + 1, lista.Largo())
	}
	for i:= 0; i < 10000 ; i-- {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())

	//Volumetria agregando al principio
	for i:= 0; i < 10000 ; i++ {
		lista.InsertarPrimero(i)
		require.EqualValues(t, i, lista.VerPrimero())
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, i + 1, lista.Largo())
	}
	for i:= 9999; i >= 0 ; i-- {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
}
