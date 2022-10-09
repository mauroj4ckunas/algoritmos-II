package main

import (
	"bufio"
	Err "errores"
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
	array := make([]int, 0, 10)
	mayor_numero_digitos := 0 //el numero que tiene mayor cantidad de "cifras/digitos"
	for padron.Scan() {
		if mayor_numero_digitos < len(padron.Text()) {
			mayor_numero_digitos = len(padron.Text())
		}
		dni, _ := strconv.Atoi(padron.Text())
		array = append(array, dni)
	}

	/*
	  	err = padron.Err()
	  	if err != nil {
	     	fmt.Println(err)
	  	}*/

	array = RadixSort(array, mayor_numero_digitos)

	Votantes := make([]Voto.Votante, len(array))
	for j := 0; j < len(array); j++ {
		Votantes[j] = Voto.CrearVotante(array[j])
	}

	return Votantes, nil
}
