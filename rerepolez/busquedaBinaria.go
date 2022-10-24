package main

func BusquedaVotante[T any](lista []T, dni int, inicio int, fin int, f func(T, int) int) int {

	if inicio > fin { //se busco en la lista y no se encontro el votante
		return -1
	}

	mitad := (inicio + fin) / 2
	comparacion := f(lista[mitad], dni)

	switch {
	case comparacion == 0:
		return mitad

	case comparacion < 0:
		return BusquedaVotante[T](lista, dni, mitad+1, fin, f)

	default:
		return BusquedaVotante[T](lista, dni, inicio, mitad-1, f)

	}
}
