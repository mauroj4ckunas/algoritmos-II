package diccionario_test

import (
	TDADiccionario "diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiccVacio(t *testing.T) {

	funcionComparable := func(clave1 string, clave2 string) int {

		if clave1[0] < clave2[0] {

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
	funcionComparable := func(clave1 int, clave2 int) int {

		if clave1 < clave2 {

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
	funcionComparable := func(clave1 string, clave2 string) int {

		if clave1[0] < clave2[0] {

			return -1

		} else if clave1[0] > clave2[0] {

			return 1

		}

		return 0
	}
	diccio := TDADiccionario.CrearABB[string, string](funcionComparable)
	require.False(t, diccio.Pertenece("papa"))
	diccio.Guardar("papa", "tierra")
	require.True(t, diccio.Pertenece("papa"))
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, 1, diccio.Cantidad())
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))

	require.False(t, diccio.Pertenece("galletitas"))
	diccio.Guardar("galletitas", "industria")
	require.True(t, diccio.Pertenece("galletitas"))
	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, 2, diccio.Cantidad())
	require.EqualValues(t, "industria", diccio.Obtener("galletitas"))
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))

	require.True(t, diccio.Pertenece("galletitas"))
	diccio.Guardar("galletitas", "desayuno")
	require.True(t, diccio.Pertenece("galletitas"))
	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, 2, diccio.Cantidad())
	require.EqualValues(t, "desayuno", diccio.Obtener("galletitas"))
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))

	require.False(t, diccio.Pertenece("labrador"))
	diccio.Guardar("labrador", "nocomestible")
	require.True(t, diccio.Pertenece("labrador"))
	require.True(t, diccio.Pertenece("labrador"))
	require.EqualValues(t, 3, diccio.Cantidad())
	require.EqualValues(t, "nocomestible", diccio.Obtener("labrador"))
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))
	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, "desayuno", diccio.Obtener("galletitas"))

	require.False(t, diccio.Pertenece("michi"))
	diccio.Guardar("michi", "ungatito")
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
	diccio.Guardar("michi", "elVerdaderoGatito")
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
	diccio.Guardar("bastaDeTesteos", "plottwist")
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
	funcionComparable := func(clave1 string, clave2 string) int {

		if clave1[0] < clave2[0] {

			return -1

		} else if clave1[0] > clave2[0] {

			return 1

		}

		return 0
	}
	diccio := TDADiccionario.CrearABB[string, string](funcionComparable)
	diccio.Guardar("labrador", "nocomestible")
	diccio.Guardar("papa", "tierra")
	diccio.Guardar("bastaDeTesteos", "plottwist")
	diccio.Guardar("galletitas", "desayuno")
	diccio.Guardar("michi", "elVerdaderoGatito")

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

func TestIterarRango(t *testing.T) {

	var (
		fecha1  int = 1910
		fecha2  int = 2022
		fecha3  int = 1916
		fecha4  int = 1986
		fecha5  int = 1810
		fecha6  int = 1806
		fecha7  int = 2010
		fecha8  int = 1807
		fecha9  int = 1816
		fecha10 int = 1978
	)

	const (
		acon1  string = "Centenario"
		acon2  string = "Actualidad"
		acon3  string = "Presidencia Yrigoyen"
		acon4  string = "Mundial"
		acon5  string = "Revolucion de Mayo"
		acon6  string = "Primera Invasion Inglesa"
		acon7  string = "Bicentenario"
		acon8  string = "Segunda Invasion Inglesa"
		acon9  string = "Independencia"
		acon10 string = "Mundial"
	)

	funcionComparable := func(clave1 int, clave2 int) int {
		if clave1 < clave2 {
			return -1
		} else if clave1 > clave2 {
			return 1
		}
		return 0
	}

	lineaDeTiempo := TDADiccionario.CrearABB[int, string](funcionComparable)

	lineaDeTiempo.Guardar(fecha1, acon1)
	lineaDeTiempo.Guardar(fecha2, acon2)
	lineaDeTiempo.Guardar(fecha3, acon3)
	lineaDeTiempo.Guardar(fecha4, acon4)
	lineaDeTiempo.Guardar(fecha5, acon5)
	lineaDeTiempo.Guardar(fecha6, acon6)
	lineaDeTiempo.Guardar(fecha7, acon7)
	lineaDeTiempo.Guardar(fecha8, acon8)
	lineaDeTiempo.Guardar(fecha9, acon9)
	lineaDeTiempo.Guardar(fecha10, acon10)

	ptrFecha6 := &fecha6
	ptrFecha1 := &fecha1
	var contadorMundiales int
	lineaDeTiempo.IterarRango(ptrFecha6, ptrFecha1, func(year int, acon string) bool {
		if acon == "Mundial" {
			contadorMundiales++
		}
		return true
	})
	require.EqualValues(t, 0, contadorMundiales)

	ptrFecha4 := &fecha4
	ptrFecha7 := &fecha7
	contadorMundiales = 0
	lineaDeTiempo.IterarRango(ptrFecha4, ptrFecha7, func(year int, acon string) bool {
		if acon == "Mundial" {
			contadorMundiales++
		}
		return true
	})
	require.EqualValues(t, 1, contadorMundiales)

	contadorMundiales = 0
	lineaDeTiempo.IterarRango(nil, nil, func(year int, acon string) bool {
		if acon == "Mundial" {
			contadorMundiales++
		}
		return true
	})
	require.EqualValues(t, 2, contadorMundiales)
}

func TestIteradorInterno(t *testing.T) {
	type letras string
	const (
		a letras = "A"
		b letras = "B"
		c letras = "C"
		d letras = "D"
		e letras = "E"
	)

	type palabra string
	const (
		primera palabra = "Debería"
		segunda palabra = " salir"
		tercera palabra = " este"
		cuarta  palabra = " mensaje"
		quinta  palabra = " correctamente."
	)

	funcionComparable := func(clave1 letras, clave2 letras) int {
		if clave1[0] < clave2[0] {
			return -1
		} else if clave1[0] > clave2[0] {
			return 1
		}
		return 0
	}

	abc := TDADiccionario.CrearABB[letras, palabra](funcionComparable)

	abc.Guardar(a, primera)
	require.EqualValues(t, primera, abc.Obtener(a))
	abc.Guardar(b, segunda)
	require.EqualValues(t, segunda, abc.Obtener(b))
	abc.Guardar(c, tercera)
	require.EqualValues(t, tercera, abc.Obtener(c))
	abc.Guardar(d, cuarta)
	require.EqualValues(t, cuarta, abc.Obtener(d))
	abc.Guardar(e, quinta)
	require.EqualValues(t, quinta, abc.Obtener(e))

	var abecedario letras
	ptrAbecedario := &abecedario
	var frase palabra
	ptrFrase := &frase

	abc.Iterar(func(clave letras, valor palabra) bool {
		*ptrAbecedario += clave
		*ptrFrase += valor
		return true
	})

	require.EqualValues(t, "ABCDE", abecedario)
	require.EqualValues(t, "Debería salir este mensaje correctamente.", frase)
}

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

	funcionComparable := func(clave1 avanzado, clave2 avanzado) int {
		if clave1.w < clave2.w {

			return -1

		} else if clave1.w > clave2.w {

			return 1

		}
		return 0
	}

	dic := TDADiccionario.CrearABB[avanzado, int](funcionComparable)

	a1 := avanzado{w: 13, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 15, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
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

}

func TestIteradorExternoSinElementos(t *testing.T) {
	funcionComparable := func(clave1 int, clave2 int) int {
		if clave1 < clave2 {
			return -1
		} else if clave1 > clave2 {
			return 1
		}
		return 0
	}
	probarIter := TDADiccionario.CrearABB[int, string](funcionComparable)

	iterVacio := probarIter.Iterador()
	verificarIterVacio(iterVacio, t)

	probarIter.Guardar(1, "")
	probarIter.Guardar(2, "")
	probarIter.Guardar(3, "")

	require.EqualValues(t, "", probarIter.Borrar(1))
	require.EqualValues(t, "", probarIter.Borrar(2))
	require.EqualValues(t, "", probarIter.Borrar(3))

	iterVacio2 := probarIter.Iterador()
	verificarIterVacio(iterVacio2, t)
}

func verificarIterVacio[K comparable, V any](iter TDADiccionario.IterDiccionario[K, V], t *testing.T) {
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangos(t *testing.T) {
	funcionComparable := func(clave1 string, clave2 string) int {
		if clave1[0] < clave2[0] {
			return -1
		} else if clave1[0] > clave2[0] {
			return 1
		}
		return 0
	}
	clave1 := "Rawson"
	clave2 := "Jackunas"
	clave3 := "Gerez"
	clave4 := "Navarro"
	clave5 := "Alberti"
	clave6 := "Mendez"
	valor1 := 10
	valor2 := 2
	valor3 := 7
	valor4 := 4
	valor5 := 6
	valor6 := 1
	claves := []string{clave1, clave2, clave3, clave4, clave5, clave6}
	valores := []int{valor1, valor2, valor3, valor4, valor5, valor6}
	dic := TDADiccionario.CrearABB[string, int](funcionComparable)
	dic.Guardar(claves[0], valores[0])

	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	dic.Guardar(claves[3], valores[3])
	dic.Guardar(claves[4], valores[4])
	dic.Guardar(claves[5], valores[5])
	ptrClave5 := &clave5
	ptrClave4 := &clave4
	iter := dic.IteradorRango(ptrClave5, ptrClave4)

	require.True(t, iter.HaySiguiente())
	resClave5, resValor5 := iter.VerActual()
	require.EqualValues(t, resClave5, clave5)
	require.EqualValues(t, resValor5, valor5)
	require.EqualValues(t, resClave5, iter.Siguiente())

	require.True(t, iter.HaySiguiente())
	resClave3, resValor3 := iter.VerActual()
	require.EqualValues(t, resClave3, clave3)
	require.EqualValues(t, resValor3, valor3)
	require.EqualValues(t, resClave3, iter.Siguiente())

	require.True(t, iter.HaySiguiente())
	resClave2, resValor2 := iter.VerActual()
	require.EqualValues(t, resClave2, clave2)
	require.EqualValues(t, resValor2, valor2)
	require.EqualValues(t, resClave2, iter.Siguiente())
}

func TestComparacionIteradores(t *testing.T) {
	const (
		animal1 string = "Gato"
		animal2 string = "Perro"
		animal3 string = "Vaca"
	)
	arrayAnimal := []string{animal1, animal2, animal3}

	const (
		ruido1 string = "Miau"
		ruido2 string = "Guau"
		ruido3 string = "Muu"
	)
	arrayRuido := []string{ruido1, ruido2, ruido3}

	funcionComparable := func(clave1 string, clave2 string) int {
		if clave1[0] < clave2[0] {
			return -1
		} else if clave1[0] > clave2[0] {
			return 1
		}
		return 0
	}

	animales := TDADiccionario.CrearABB[string, string](funcionComparable)

	animales.Guardar(animal1, ruido1)
	animales.Guardar(animal2, ruido2)
	animales.Guardar(animal3, ruido3)

	iterSinRangos := animales.Iterador()
	iterConRangos := animales.IteradorRango(nil, nil)

	for i := 0; i < 3; i++ {
		resClaveCon, resValorCon := iterConRangos.VerActual()
		resClaveSin, resValorSin := iterSinRangos.VerActual()

		require.EqualValues(t, resClaveCon, resClaveSin)
		require.EqualValues(t, resValorCon, resValorSin)
		require.EqualValues(t, resValorCon, arrayRuido[i])
		require.EqualValues(t, resValorSin, arrayRuido[i])

		claveCon := iterConRangos.Siguiente()
		claveSin := iterSinRangos.Siguiente()

		require.EqualValues(t, claveCon, arrayAnimal[i])
		require.EqualValues(t, claveSin, arrayAnimal[i])
	}

	verificarIterVacio(iterConRangos, t)
	verificarIterVacio(iterConRangos, t)
}

func TestIteradoresSalidaInOrder(t *testing.T) {

	const (
		a string = "A"
		b string = "B"
		c string = "C"
		d string = "D"
		e string = "E"
	)

	funcionComparable := func(clave1 string, clave2 string) int {
		if clave1 < clave2 {
			return -1
		} else if clave1 > clave2 {
			return 1
		}
		return 0
	}

	ordenDeIngreso := []string{e, d, a, b, c}

	dic := TDADiccionario.CrearABB[string, int](funcionComparable)

	for _, valor := range ordenDeIngreso {
		dic.Guardar(valor, 0)
	}

	iter := dic.Iterador()

	clave1, _ := iter.VerActual()
	require.NotEqualValues(t, clave1, ordenDeIngreso[0])
	require.EqualValues(t, clave1, a)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	clave2, _ := iter.VerActual()
	require.NotEqualValues(t, clave2, ordenDeIngreso[1])
	require.EqualValues(t, clave2, b)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	clave3, _ := iter.VerActual()
	require.NotEqualValues(t, clave3, ordenDeIngreso[2])
	require.EqualValues(t, clave3, c)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	clave4, _ := iter.VerActual()
	require.NotEqualValues(t, clave4, ordenDeIngreso[3])
	require.EqualValues(t, clave4, d)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	clave5, _ := iter.VerActual()
	require.NotEqualValues(t, clave5, ordenDeIngreso[4])
	require.EqualValues(t, clave5, e)
	iter.Siguiente()
	verificarIterVacio(iter, t)
}

func TestVolumen(t *testing.T) {
	funcionComparable := func(clave1 int, clave2 int) int {
		if clave1 < clave2 {
			return -1
		} else if clave1 > clave2 {
			return 1
		}
		return 0
	}

	volumen := TDADiccionario.CrearABB[int, int](funcionComparable)

	for j := 5000; j <= 10000; j++ {
		volumen.Guardar(j, j*2)
	}
	for i := 1; i <= 5000; i++ {
		volumen.Guardar(i, i*2)
	}

	iter := volumen.IteradorRango(nil, nil)
	var num int = 1
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		require.EqualValues(t, clave, num)
		require.EqualValues(t, valor, clave*2)
		num++
		iter.Siguiente()
	}
	require.EqualValues(t, 10000, volumen.Cantidad())
	for k := 1; k <= 10000; k++ {
		require.True(t, volumen.Pertenece(k))
		require.EqualValues(t, k*2, volumen.Obtener(k))
		require.EqualValues(t, k*2, volumen.Borrar(k))
	}
	for m := 10000; m >= 1; m-- {
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { volumen.Obtener(m) })
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { volumen.Borrar(m) })
	}
	require.EqualValues(t, 0, volumen.Cantidad())
}

func TestIteradorConNil(t *testing.T) {
	funcionComparable := func(clave1 string, clave2 string) int {
		if clave1[0] < clave2[0] {
			return -1
		} else if clave1[0] > clave2[0] {
			return 1
		}
		return 0
	}

	alumnos := TDADiccionario.CrearABB[string, string](funcionComparable)

	listado := []string{"Nicolas", "Roberto", "Leonardo", "Cesar", "Pablo", "Mora", "Tamara"}

	alumnos.Guardar(listado[0], "Presente")
	alumnos.Guardar(listado[1], "Presente")
	alumnos.Guardar(listado[2], "Presente")
	alumnos.Guardar(listado[3], "Presente")
	alumnos.Guardar(listado[4], "Presente")
	alumnos.Guardar(listado[5], "Presente")
	alumnos.Guardar(listado[6], "Presente")

	ptrHasta := &listado[4]

	iteradorDesdeNil := alumnos.IteradorRango(nil, ptrHasta)

	presentesDia1 := []string{"Cesar", "Leonardo", "Mora", "Nicolas", "Pablo"}

	var i int = 0
	for iteradorDesdeNil.HaySiguiente() {
		clave, valor := iteradorDesdeNil.VerActual()
		require.EqualValues(t, presentesDia1[i], clave)
		require.EqualValues(t, "Presente", valor)
		require.EqualValues(t, presentesDia1[i], iteradorDesdeNil.Siguiente())
		i++
	}

	ptrDesde := &listado[5]

	iteradorHastaNil := alumnos.IteradorRango(ptrDesde, nil)

	presentesDia2 := []string{"Mora", "Nicolas", "Pablo", "Roberto", "Tamara"}

	var j int = 0
	for iteradorHastaNil.HaySiguiente() {
		clave, valor := iteradorHastaNil.VerActual()
		require.EqualValues(t, presentesDia2[j], clave)
		require.EqualValues(t, "Presente", valor)
		require.EqualValues(t, presentesDia2[j], iteradorHastaNil.Siguiente())
		iteradorHastaNil.Siguiente()
		j++
	}

}
