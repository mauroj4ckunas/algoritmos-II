package diccionario_test

import (
	TDADiccionario "diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func funcionComparableString(clave1 string, clave2 string) int {
	if clave1[0] < clave2[0] {
		return -1
	} else if clave1[0] > clave2[0] {
		return 1
	}
	return 0
}

func funcionComparableInt(num1 int, num2 int) int {
	if num1 < num2 {
		return -1
	} else if num1 > num2 {
		return 1
	}
	return 0
}

//Test ABB: primitivas.

func TestDiccVacio(t *testing.T) {
	diccio := TDADiccionario.CrearABB[string, int](funcionComparableString)
	require.EqualValues(t, 0, diccio.Cantidad())
	require.False(t, diccio.Pertenece("Obama"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Obtener("RicardoFord") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { diccio.Borrar("Buchwald") })
}

func TestConUnElemento(t *testing.T) {
	diccio := TDADiccionario.CrearABB[int, bool](funcionComparableInt)
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
	diccio := TDADiccionario.CrearABB[string, string](funcionComparableString)
	require.False(t, diccio.Pertenece("papa"))
	diccio.Guardar("papa", "tierra")
	require.True(t, diccio.Pertenece("papa"))
	require.EqualValues(t, 1, diccio.Cantidad())
	require.EqualValues(t, "tierra", diccio.Obtener("papa"))

	require.False(t, diccio.Pertenece("galletitas"))
	diccio.Guardar("galletitas", "industria")
	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, 2, diccio.Cantidad())
	require.EqualValues(t, "industria", diccio.Obtener("galletitas"))

	diccio.Guardar("galletitas", "desayuno")
	require.True(t, diccio.Pertenece("galletitas"))
	require.EqualValues(t, 2, diccio.Cantidad())
	require.EqualValues(t, "desayuno", diccio.Obtener("galletitas"))

	require.False(t, diccio.Pertenece("labrador"))
	diccio.Guardar("labrador", "nocomestible")
	require.True(t, diccio.Pertenece("labrador"))
	require.EqualValues(t, 3, diccio.Cantidad())
	require.EqualValues(t, "nocomestible", diccio.Obtener("labrador"))

	require.False(t, diccio.Pertenece("michi"))
	diccio.Guardar("michi", "ungatito")
	require.True(t, diccio.Pertenece("michi"))
	require.EqualValues(t, 4, diccio.Cantidad())
	require.EqualValues(t, "ungatito", diccio.Obtener("michi"))

	diccio.Guardar("michi", "elVerdaderoGatito")
	require.EqualValues(t, 4, diccio.Cantidad())
	require.EqualValues(t, "elVerdaderoGatito", diccio.Obtener("michi"))

	require.False(t, diccio.Pertenece("bastaDeTesteos"))
	diccio.Guardar("bastaDeTesteos", "plottwist")
	require.True(t, diccio.Pertenece("bastaDeTesteos"))
	require.EqualValues(t, 5, diccio.Cantidad())
	require.EqualValues(t, "plottwist", diccio.Obtener("bastaDeTesteos"))

}

func TestConCeroHijos(t *testing.T) {
	diccio := TDADiccionario.CrearABB[string, string](funcionComparableString)
	diccio.Guardar("labrador", "nocomestible")
	diccio.Guardar("papa", "tierra")
	diccio.Guardar("bastaDeTesteos", "plottwist")
	diccio.Guardar("galletitas", "desayuno")
	diccio.Guardar("michi", "elVerdaderoGatito")

	cantidad := 5

	pruebaConClaveyValor[string, string](diccio, t, "michi", "elVerdaderoGatito", cantidad)
	cantidad--

	pruebaConClaveyValor[string, string](diccio, t, "papa", "tierra", cantidad)
	cantidad--

	pruebaConClaveyValor[string, string](diccio, t, "galletitas", "desayuno", cantidad)
	cantidad--

	pruebaConClaveyValor[string, string](diccio, t, "bastaDeTesteos", "plottwist", cantidad)
	cantidad--

	pruebaConClaveyValor[string, string](diccio, t, "labrador", "nocomestible", cantidad)
	cantidad--

	require.EqualValues(t, 0, diccio.Cantidad())
}

func TestConUnoYDosHijos(t *testing.T) {
	diccio := TDADiccionario.CrearABB[string, int](funcionComparableString)
	diccio.Guardar("Messi", 10)
	diccio.Guardar("Gonzales", 3)
	diccio.Guardar("Pezzella", 0)
	diccio.Guardar("Correa", 4)
	diccio.Guardar("Otamendi", 1)
	diccio.Guardar("Rulli", 0)
	diccio.Guardar("Acuna", 3)
	diccio.Guardar("Dybala", 5)

	cantidad := 8

	//Eliminar uno con DOS hijos

	pruebaConClaveyValor[string, int](diccio, t, "Pezzella", 0, cantidad)
	cantidad--

	pruebaConClaveyValor[string, int](diccio, t, "Correa", 4, cantidad)
	cantidad--

	pruebaConClaveyValor[string, int](diccio, t, "Messi", 10, cantidad)
	cantidad--

	//Eliminar uno con UN hijo

	pruebaConClaveyValor[string, int](diccio, t, "Gonzales", 3, cantidad)
	cantidad--

	pruebaConClaveyValor[string, int](diccio, t, "Otamendi", 1, cantidad)
	cantidad--

	pruebaConClaveyValor[string, int](diccio, t, "Dybala", 5, cantidad)
	cantidad--

	pruebaConClaveyValor[string, int](diccio, t, "Rulli", 0, cantidad)
	cantidad--

	pruebaConClaveyValor[string, int](diccio, t, "Acuna", 3, cantidad)
	cantidad--

	require.EqualValues(t, 0, diccio.Cantidad())
}

func pruebaConClaveyValor[K comparable, V any](dic TDADiccionario.DiccionarioOrdenado[K, V], t *testing.T, clave K, valor V, conteo int) {
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, conteo, dic.Cantidad())
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.EqualValues(t, conteo-1, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(clave) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(clave) })
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

	funcionComparableStruct := func(clave1 avanzado, clave2 avanzado) int {
		if clave1.w < clave2.w {

			return -1

		} else if clave1.w > clave2.w {

			return 1

		}
		return 0
	}

	dic := TDADiccionario.CrearABB[avanzado, int](funcionComparableStruct)

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

func TestMismosNodosDiferentesPuestos(t *testing.T) {
	var (
		num1 int = 1
		num2 int = 2
		num3 int = 3
		num4 int = 4
		num5 int = 5
		num6 int = 6
		num7 int = 7
	)
	array := []int{num1, num2, num3, num4, num5, num6, num7}
	diccionario := TDADiccionario.CrearABB[int, int](funcionComparableInt)

	diccionario.Guardar(num1, 0)
	diccionario.Guardar(num2, 0)
	diccionario.Guardar(num3, 0)
	diccionario.Guardar(num5, 0)
	diccionario.Guardar(num4, 0)
	diccionario.Guardar(num6, 0)
	diccionario.Guardar(num7, 0)

	ptrNum2 := &num2
	ptrNum5 := &num5
	var i int = 1
	ptrI := &i

	diccionario.IterarRango(ptrNum2, ptrNum5, func(clave, dato int) bool {
		require.EqualValues(t, array[*ptrI], clave)
		*ptrI++
		return true
	})

	for _, valor := range array {
		diccionario.Borrar(valor)
	}

	diccionario.Guardar(num7, 0)
	diccionario.Guardar(num6, 0)
	diccionario.Guardar(num5, 0)
	diccionario.Guardar(num4, 0)
	diccionario.Guardar(num3, 0)
	diccionario.Guardar(num2, 0)
	diccionario.Guardar(num1, 0)

	i = 1
	diccionario.IterarRango(ptrNum2, ptrNum5, func(clave, dato int) bool {
		require.EqualValues(t, array[*ptrI], clave)
		*ptrI++
		return true
	})

	for _, valor := range array {
		diccionario.Borrar(valor)
	}

	diccionario.Guardar(num5, 0)
	diccionario.Guardar(num6, 0)
	diccionario.Guardar(num7, 0)
	diccionario.Guardar(num3, 0)
	diccionario.Guardar(num1, 0)
	diccionario.Guardar(num2, 0)
	diccionario.Guardar(num4, 0)

	i = 1
	diccionario.IterarRango(ptrNum2, ptrNum5, func(clave, dato int) bool {
		require.EqualValues(t, array[*ptrI], clave)
		*ptrI++
		return true
	})
}

func TestVolumen(t *testing.T) {

	volumen := TDADiccionario.CrearABB[int, int](funcionComparableInt)

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

//Test Iterarador Interno: Sin y Con Rangos

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

	funcionComparableOtroTipo := func(clave1 letras, clave2 letras) int {
		if clave1[0] < clave2[0] {
			return -1
		} else if clave1[0] > clave2[0] {
			return 1
		}
		return 0
	}

	abc := TDADiccionario.CrearABB[letras, palabra](funcionComparableOtroTipo)

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

	arrayAconOrdenadoPorClave := []string{acon6, acon8, acon5, acon9, acon1, acon3, acon10, acon4, acon7, acon2}

	lineaDeTiempo := TDADiccionario.CrearABB[int, string](funcionComparableInt)

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

	desdeFueraDelAbb := 1900
	ptrdesdeFueraDelAbb := &desdeFueraDelAbb

	hastaFueraDelAbb := 2000
	ptrhastaFueraDelAbb := &hastaFueraDelAbb

	var i int = 4
	ptrI := &i

	lineaDeTiempo.IterarRango(ptrdesdeFueraDelAbb, ptrhastaFueraDelAbb, func(year int, acon string) bool {
		require.EqualValues(t, acon, arrayAconOrdenadoPorClave[*ptrI])
		*ptrI++
		return true
	})

	ptrFecha8 := &fecha8
	var j int = 1
	ptrj := &j

	lineaDeTiempo.IterarRango(ptrFecha8, ptrFecha7, func(year int, acon string) bool {
		require.EqualValues(t, acon, arrayAconOrdenadoPorClave[*ptrj])
		*ptrj++
		return true
	})

}

func TestIterarRangoIncluyeUno(t *testing.T) {

	var (
		primer  int = 45
		segundo int = 90
		tercero int = 92
		cuarto  int = 140
		quinto  int = 237
	)

	dic1 := TDADiccionario.CrearABB[int, int](funcionComparableInt)

	dic1.Guardar(cuarto, 0)
	dic1.Guardar(quinto, 0)
	dic1.Guardar(primer, 0)
	dic1.Guardar(tercero, 0)
	dic1.Guardar(segundo, 0)

	var (
		a string = "C"
		b string = "F"
		c string = "M"
		d string = "T"
		e string = "V"
	)

	dic2 := TDADiccionario.CrearABB[string, string](funcionComparableString)

	dic2.Guardar(c, "")
	dic2.Guardar(a, "")
	dic2.Guardar(e, "")
	dic2.Guardar(b, "")
	dic2.Guardar(d, "")

	num1 := 130
	ptrNum1 := &num1
	num2 := 210
	ptrNum2 := &num2

	dic1.IterarRango(ptrNum1, ptrNum2, func(clave, dato int) bool {
		require.EqualValues(t, clave, cuarto)
		return true
	})

	num3 := 230
	ptrNum3 := &num3
	num4 := 10000
	ptrNum4 := &num4

	dic1.IterarRango(ptrNum3, ptrNum4, func(clave, dato int) bool {
		if clave == quinto {
			require.EqualValues(t, clave, quinto)
		} else {
			require.NotEqualValues(t, clave, quinto)
		}

		return true
	})

	letra1 := "P"
	ptrLetra1 := &letra1
	letra2 := "U"
	ptrLetra2 := &letra2

	dic2.IterarRango(ptrLetra1, ptrLetra2, func(clave, dato string) bool {
		require.EqualValues(t, clave, d)
		return true
	})

}

func TestIterarConNilyCortes(t *testing.T) {
	var (
		par1   int = 2
		par2   int = 4
		par3   int = 6
		par4   int = 8
		par5   int = 10
		impar1 int = 3
		impar2 int = 7
	)

	buscarImpar := TDADiccionario.CrearABB[int, bool](funcionComparableInt)
	buscarImpar.Guardar(par3, false)
	buscarImpar.Guardar(impar2, true)
	buscarImpar.Guardar(par4, false)
	buscarImpar.Guardar(par2, false)
	buscarImpar.Guardar(par1, false)
	buscarImpar.Guardar(impar1, true)
	buscarImpar.Guardar(par5, false)

	var contador int
	ptrContador := &contador

	ptrPar1 := &par1
	ptrPar3 := &par3

	buscarImpar.IterarRango(ptrPar1, ptrPar3, func(clave int, dato bool) bool {
		*ptrContador++
		return clave%2 == 0
	})
	require.EqualValues(t, 2, contador)

	ptrPar2 := &par2
	ptrPar5 := &par5

	contador = 0
	buscarImpar.IterarRango(ptrPar2, ptrPar5, func(clave int, dato bool) bool {
		*ptrContador++
		return clave%2 == 0
	})
	require.EqualValues(t, 3, contador)
}

//Test Iterador Externo: Sin y Con Rangos

func TestIteradorExternoEInternoSinElementos(t *testing.T) {

	probarIter := TDADiccionario.CrearABB[int, string](funcionComparableInt)

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

	probarIter.Iterar(func(clave int, dato string) bool {
		return true
	})

}

func TestIteradorRangos(t *testing.T) {
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
	dic := TDADiccionario.CrearABB[string, int](funcionComparableString)
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

	animales := TDADiccionario.CrearABB[string, string](funcionComparableString)

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
	verificarIterVacio(iterSinRangos, t)
}

func TestIteradoresSalidaInOrder(t *testing.T) {

	const (
		a string = "A"
		b string = "B"
		c string = "C"
		d string = "D"
		e string = "E"
	)

	ordenDeIngreso := []string{e, d, a, b, c}

	dic := TDADiccionario.CrearABB[string, int](funcionComparableString)

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

func TestIteradorConNil(t *testing.T) {

	alumnos := TDADiccionario.CrearABB[string, string](funcionComparableString)

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

	require.False(t, iteradorDesdeNil.HaySiguiente())

	ptrDesde := &listado[5]

	iteradorHastaNil := alumnos.IteradorRango(ptrDesde, nil)

	presentesDia2 := []string{"Mora", "Nicolas", "Pablo", "Roberto", "Tamara"}

	var j int = 0
	for iteradorHastaNil.HaySiguiente() {
		clave, valor := iteradorHastaNil.VerActual()
		require.EqualValues(t, presentesDia2[j], clave)
		require.EqualValues(t, "Presente", valor)
		require.EqualValues(t, presentesDia2[j], iteradorHastaNil.Siguiente())
		j++
	}
}

func verificarIterVacio[K comparable, V any](iter TDADiccionario.IterDiccionario[K, V], t *testing.T) {
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
