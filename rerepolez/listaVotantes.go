package main

import (
	"bufio"
	Err "rerepolez/errores"
	"os"
	"strconv"
	Voto "rerepolez/votos"
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
		dni, err := strconv.Atoi(padron.Text())

		if err != nil {
			ErrorLectura := new(Err.ErrorLeerArchivo)
			return []Voto.Votante{}, ErrorLectura
		}

		array = append(array, dni)
	}

	

	array = RadixSort(array, mayor_numero_digitos)

	Votantes := make([]Voto.Votante, len(array))
	for j := 0; j < len(array); j++ {
		Votantes[j] = Voto.CrearVotante(array[j])
	}

	return Votantes, nil
}
