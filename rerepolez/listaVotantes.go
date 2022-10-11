package main

import (
	"bufio"
	Err "rerepolez/errores"
	"os"
	"strconv"
	Voto "rerepolez/votos"
)

func PrepararListaVotantes(rutaPadrones string) ([]Voto.Votante, error) {

	archivoVotantes, err := os.Open(rutaPadrones)
	defer archivoVotantes.Close()

	if err != nil { //si la ruta no se puede leer o algo, error
		ErrorLectura := new(Err.ErrorLeerArchivo)
		return []Voto.Votante{}, ErrorLectura
	}

	padrones := bufio.NewScanner(archivoVotantes)
	arrayDeDnis := make([]int, 0, 10)
	mayorNumeroDeCifras := 0 //el numero que tiene mayor cantidad de "cifras/digitos"

	for padrones.Scan() {

		if mayorNumeroDeCifras < len(padrones.Text()) {

			mayorNumeroDeCifras = len(padrones.Text())

		}

		dni, err := strconv.Atoi(padrones.Text())

		if err != nil { //si el archivo no esta escrito como se debe o el archivo es el de partidos

			ErrorLectura := new(Err.ErrorLeerArchivo)
			return []Voto.Votante{}, ErrorLectura

		}

		arrayDeDnis = append(arrayDeDnis, dni)
	}

	

	arrayDeDnis = RadixSort(arrayDeDnis, mayorNumeroDeCifras) //ordenamiento de Dnis

	Votantes := make([]Voto.Votante, len(arrayDeDnis))

	for j := 0; j < len(arrayDeDnis); j++ {

		Votantes[j] = Voto.CrearVotante(arrayDeDnis[j])

	}

	return Votantes, nil
}
