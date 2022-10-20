package diccionario_test

import (
	TDADiccionario "diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiccVacio(t *testing.T) {

	funcionComparable := func(clave1 string,clave2 string) int {

		if clave1[0] < clave2[0]{

			return -1

		} else if clave1[0] > clave2[0] {

			return 1

		}

		return 0
	}

	diccio := TDADiccionario.CrearABB[string, int](funcionComparable)
	require.EqualValues(t, 0, diccio.Cantidad())
	require.False(t, diccio.Pertenece("Obama"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Obtener("RicardoFord") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Borrar("Buchwald") })
}

func TestConUnElemento(t *testing.T) {
	funcionComparable := func(clave1 int,clave2 int) int {

		if clave1 < clave2{

			return -1

		} else if clave1 > clave2 {

			return 1

		}

		return 0
	}
	diccio := TDADiccionario.CrearABB[int, bool](funcionComparable)
	diccio.Guardar(22, true)
	require.EqualValues(t, 1, diccio.Cantidad())
	require.True(t, diccio.Pertenece(22))
	require.False(t, diccio.Pertenece(5))
	require.EqualValues(t, true, diccio.Obtener(22))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Obtener(5) })
	require.EqualValues(t, true, diccio.Borrar(22))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Borrar(22) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Obtener(5) })
}