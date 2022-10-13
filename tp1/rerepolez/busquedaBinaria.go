package main

import (
	votos "rerepolez/votos"
)

func BusquedaVotante(lista []votos.Votante, dni int, inicio int, fin int) int {

	if inicio > fin { //se busco en la lista y no se encontro el votante
		return -1
	}

	mitad := (inicio + fin) / 2

	if lista[mitad].LeerDNI() == dni { //se encontro

		return mitad

	} else if lista[mitad].LeerDNI() < dni { //se busca a la derecha de la lista

		return BusquedaVotante(lista, dni, mitad+1, fin)

	} else { //se busca a la izquierda de la lista

		return BusquedaVotante(lista, dni, inicio, mitad-1)

	}
}
