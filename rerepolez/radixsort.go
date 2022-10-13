package main

const (
	DIGITOS_DE_UN_NUMERO int = 10
)

func cifrasDeUnNumero(numero int, cifra int) int {
	var parteEntera int = numero
	var resto int
	for i := 0; i < cifra; i++ {
		resto, parteEntera = parteEntera%DIGITOS_DE_UN_NUMERO, parteEntera/DIGITOS_DE_UN_NUMERO
	}
	return resto
}

func CountingSort(array []int, cifra int) []int {

	frecuencias := make([]int, DIGITOS_DE_UN_NUMERO)
	cifrasEnOrden := make([]int, len(array))

	for i := 0; i < len(array); i++ {

		cifrasEnOrden[i] = cifrasDeUnNumero(array[i], cifra)
		frecuencias[cifrasEnOrden[i]] += 1

	}

	sumasAcumuladas := make([]int, DIGITOS_DE_UN_NUMERO)

	for i := 0; i < len(sumasAcumuladas)-1; i++ {

		sumasAcumuladas[i+1] = frecuencias[i] + sumasAcumuladas[i]

	}

	arrayNuevo := make([]int, len(array))

	for i := 0; i < len(array); i++ {

		arrayNuevo[sumasAcumuladas[cifrasEnOrden[i]]] = array[i]
		sumasAcumuladas[cifrasEnOrden[i]] += 1

	}
	return arrayNuevo

}

func RadixSort(array []int, cifrasMaximas int) []int {
	for cifra := 1; cifra <= cifrasMaximas; cifra++ {
		array = CountingSort(array, cifra)
	}
	return array
}
