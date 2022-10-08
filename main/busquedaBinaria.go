package main

import (
	TDAPila "Pila"
	"bufio"
	Err "errores"
	"fmt"
	"os"
	"strconv"
	Voto "votos"
)

func BusquedaVotante(lista []Voto.Votante, dni int) int {
	if len(lista) == 0 {
		return -1
	}
	mitad = len(lista) / 2
	if lista[mitad].LeerDNI() == dni {
		return mitad
	} else if lista[mitad].LeerDNI() < dni {
		return mitad + 1 + BusquedaVotante(lista[mitad+1:], dni)
	} else {
		return BusquedaVotante(lista[:mitad], dni)
	}
}
