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

func TestDeGuardado(t *testing.T) {
	funcionComparable := func(clave1 string,clave2 string) int {

		if clave1[0] < clave2[0]{

			return -1

		} else if clave1[0] > clave2[0] {

			return 1

		}

		return 0
	}
	diccio := TDADiccionario.CrearABB[string, string](funcionComparable)
	require.False(t, diccio.Pertenece("papa"))
	diccio.Guardar("papa","tierra")
	require.True(t, diccio.Pertenece("papa"))
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, 1, diccio.Cantidad())
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))
	
	require.False(t, diccio.Pertenece("galletitas"))
	diccio.Guardar("galletitas","industria")
	require.True(t, diccio.Pertenece("galletitas"))
	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, 2, diccio.Cantidad())
	require.EqualValues(t, "industria", diccio.Obtener("galletitas"))
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))

	require.True(t, diccio.Pertenece("galletitas"))
	diccio.Guardar("galletitas","desayuno")
	require.True(t, diccio.Pertenece("galletitas"))
	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, 2, diccio.Cantidad())
	require.EqualValues(t, "desayuno", diccio.Obtener("galletitas"))
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))

	require.False(t, diccio.Pertenece("labrador"))
	diccio.Guardar("labrador","nocomestible")
	require.True(t, diccio.Pertenece("labrador"))
	require.True(t, diccio.Pertenece("labrador"))
	require.EqualValues(t, 3, diccio.Cantidad())
	require.EqualValues(t, "nocomestible", diccio.Obtener("labrador"))
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))
	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, "desayuno", diccio.Obtener("galletitas"))

	require.False(t, diccio.Pertenece("michi"))
	diccio.Guardar("michi","ungatito")
	require.True(t, diccio.Pertenece("michi"))
	require.True(t, diccio.Pertenece("michi"))
	require.EqualValues(t, 4, diccio.Cantidad())
	require.EqualValues(t, "ungatito", diccio.Obtener("michi"))
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))
	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, "desayuno", diccio.Obtener("galletitas"))
	require.True(t, diccio.Pertenece("labrador"))
	require.EqualValues(t, "nocomestible", diccio.Obtener("labrador"))

	require.True(t, diccio.Pertenece("michi"))
	diccio.Guardar("michi","elVerdaderoGatito")
	require.True(t, diccio.Pertenece("michi"))
	require.True(t, diccio.Pertenece("michi"))
	require.EqualValues(t, 4, diccio.Cantidad())
	require.EqualValues(t, "elVerdaderoGatito", diccio.Obtener("michi"))
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))
	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, "desayuno", diccio.Obtener("galletitas"))
	require.True(t, diccio.Pertenece("labrador"))
	require.EqualValues(t, "nocomestible", diccio.Obtener("labrador"))

	require.False(t, diccio.Pertenece("bastaDeTesteos"))
	diccio.Guardar("bastaDeTesteos","plottwist")
	require.True(t, diccio.Pertenece("bastaDeTesteos"))
	require.True(t, diccio.Pertenece("bastaDeTesteos"))
	require.EqualValues(t, 5, diccio.Cantidad())
	require.EqualValues(t, "plottwist", diccio.Obtener("bastaDeTesteos"))
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))
	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, "desayuno", diccio.Obtener("galletitas"))
	require.True(t, diccio.Pertenece("labrador"))
	require.EqualValues(t, "nocomestible", diccio.Obtener("labrador"))
	require.True(t, diccio.Pertenece("michi"))
	require.EqualValues(t, "elVerdaderoGatito", diccio.Obtener("michi"))

	
}

func TestDeBorrados(t *testing.T) {
	funcionComparable := func(clave1 string,clave2 string) int {

		if clave1[0] < clave2[0]{

			return -1

		} else if clave1[0] > clave2[0] {

			return 1

		}

		return 0
	}
	diccio := TDADiccionario.CrearABB[string, string](funcionComparable)
	diccio.Guardar("labrador","nocomestible")
	diccio.Guardar("papa","tierra")
	diccio.Guardar("bastaDeTesteos","plottwist")
	diccio.Guardar("galletitas","desayuno")
	diccio.Guardar("michi","elVerdaderoGatito")
	
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))
	require.EqualValues(t, 5, diccio.Cantidad())
	require.EqualValues(t, "tierra", diccio.Borrar("papa"))
	require.EqualValues(t, 4, diccio.Cantidad())
	require.False(t, diccio.Pertenece("papa"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Borrar("papa") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Obtener("papa") })

	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, "desayuno", diccio.Obtener("galletitas"))
	require.EqualValues(t, 4, diccio.Cantidad())
	require.EqualValues(t, "desayuno", diccio.Borrar("galletitas"))
	require.EqualValues(t, 3, diccio.Cantidad())
	require.False(t, diccio.Pertenece("galletitas"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Borrar("galletitas") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Obtener("galletitas") })

	require.True(t, diccio.Pertenece("labrador"))
	require.EqualValues(t, "nocomestible", diccio.Obtener("labrador"))
	require.EqualValues(t, 3, diccio.Cantidad())
	require.EqualValues(t, "nocomestible", diccio.Borrar("labrador"))
	require.EqualValues(t, 2, diccio.Cantidad())
	require.False(t, diccio.Pertenece("labrador"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Borrar("labrador") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Obtener("labrador") })

	require.True(t, diccio.Pertenece("michi"))
	require.EqualValues(t, "elVerdaderoGatito", diccio.Obtener("michi"))
	require.EqualValues(t, 2, diccio.Cantidad())
	require.EqualValues(t, "elVerdaderoGatito", diccio.Borrar("michi"))
	require.EqualValues(t, 1, diccio.Cantidad())
	require.False(t, diccio.Pertenece("michi"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Borrar("michi") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Obtener("michi") })

	require.True(t, diccio.Pertenece("bastaDeTesteos"))
	require.EqualValues(t, "plottwist", diccio.Obtener("bastaDeTesteos"))
	require.EqualValues(t, 1, diccio.Cantidad())
	require.EqualValues(t, "plottwist", diccio.Borrar("bastaDeTesteos"))
	require.EqualValues(t, 0, diccio.Cantidad())
	require.False(t, diccio.Pertenece("bastaDeTesteos"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Borrar("bastaDeTesteos") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Obtener("bastaDeTesteos") })
}

/*
func TestConClavesStructs(t *testing.T) {
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	dic := TDADiccionario.CrearHash[avanzado, int]()

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))

}*/