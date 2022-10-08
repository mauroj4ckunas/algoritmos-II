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

func PrepararListaVotantes(rutaPadrones string) ([]Voto.Votante, Err.Errores) {

	archivoVotantes, err := os.Open(rutaPadrones)
	defer archivoVotantes.Close()
	if err != nil { //si la ruta no se puede leer o algo, error
		ErrorLectura := new(Err.ErrorLeerArchivo)
		return []Voto.Votante{}, ErrorLectura
	}

	padron := bufio.NewScanner(archivoVotantes)
	cantidad_votantes := 0
	mayor_numero_digitos := 0 //el numero que tiene mayor cantidad de "cifras/digitos"
	for padron.Scan() {
		if mayor_numero_digitos < len(padron.Text()) {
			mayor_numero_digitos = len(padron.Text())
		}
		cantidad_votantes += 1
	}

	/*
	  	err = padron.Err()
	  	if err != nil {
	     	fmt.Println(err)
	  	}*/

	padron = bufio.NewScanner(archivoVotantes)
	array := make([]int, cantidad_votantes)
	for i := 0; padron.Scan(); i++ {
		array[i] = strconv.Atoi(padron.Text())
	}
	array = RadixSort(array, mayor_numero_digitos)

	Votantes := make([]Voto.Votante, cantidad_votantes)
	for j := 0; j < len(array); j++ {
		Votantes[j] = Voto.CrearVotante(array[j])
	}

	return Votantes, nil
}
