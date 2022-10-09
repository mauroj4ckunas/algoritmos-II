package main

import (
	Voto "votos"
)

func BusquedaVotante(lista []Voto.Votante, dni int, inicio int, fin int) int {
	if inicio >= fin {
		return -1
	}
	mitad := (inicio + fin) / 2
	if lista[mitad].LeerDNI() == dni {
		return mitad
	} else if lista[mitad].LeerDNI() < dni {
		return BusquedaVotante(lista, dni, mitad, fin)
	} else {
		return BusquedaVotante(lista, dni, inicio, mitad)
	}
}
