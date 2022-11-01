package main

import (
	"bufio"
	"os"
	Err "rerepolez/errores"
	votos "rerepolez/votos"
	"strconv"
	"strings"
)

func PrepararMesa(parametros []string) ([]votos.Partido, []votos.Votante, error) {
	if len(parametros) == 0 { //si no hay parametros tira error
		return []votos.Partido{}, []votos.Votante{}, new(Err.ErrorParametros)
	}

	rutaListas := parametros[0]
	listaDeLosPartidos, err := PrepararListaPartidos(rutaListas)

	if err != nil {
		return []votos.Partido{}, []votos.Votante{}, err
	}

	if len(parametros) != 2 { //si los parametros no terminan de ser suficientes
		return []votos.Partido{}, []votos.Votante{}, new(Err.ErrorParametros)
	}

	rutaPadrones := parametros[1]

	listaDeLosVotantes, err := PrepararListaVotantes(rutaPadrones)
	if err != nil {
		return []votos.Partido{}, []votos.Votante{}, err
	}
	return listaDeLosPartidos, listaDeLosVotantes, nil
}

func PrepararListaPartidos(ruta string) ([]votos.Partido, error) {

	archivoListas, err := os.Open(ruta)
	defer archivoListas.Close()

	if err != nil { //si la ruta no se puede leer o algo, error

		ErrorLectura := new(Err.ErrorLeerArchivo)
		return []votos.Partido{}, ErrorLectura

	}

	listaPartidos := bufio.NewScanner(archivoListas)

	Partidos := make([]votos.Partido, 1)
	candVacio := [votos.CANT_VOTACION]string{"", "", ""}
	partidoEnBlanco := votos.CrearPartido("Votos en Blanco", candVacio)
	Partidos[0] = partidoEnBlanco

	for listaPartidos.Scan() {

		boleta := strings.Split(listaPartidos.Text(), ",")

		if len(boleta) != int(votos.CANT_VOTACION+1) { //si el archivo no esta escrito como se debe o es el archivo de votantes

			ErrorLectura := new(Err.ErrorLeerArchivo)
			return []votos.Partido{}, ErrorLectura

		}

		nombrePartido := boleta[0]
		candidatosPartido := [votos.CANT_VOTACION]string{boleta[1], boleta[2], boleta[3]}
		nuevoPartido := votos.CrearPartido(nombrePartido, candidatosPartido)
		Partidos = append(Partidos, nuevoPartido)
	}

	return Partidos, nil
}

func PrepararListaVotantes(rutaPadrones string) ([]votos.Votante, error) {

	archivoVotantes, err := os.Open(rutaPadrones)
	defer archivoVotantes.Close()

	if err != nil { //si la ruta no se puede leer o algo, error
		ErrorLectura := new(Err.ErrorLeerArchivo)
		return []votos.Votante{}, ErrorLectura
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
			return []votos.Votante{}, ErrorLectura

		}

		arrayDeDnis = append(arrayDeDnis, dni)
	}

	arrayDeDnis = RadixSort(arrayDeDnis, mayorNumeroDeCifras) //ordenamiento de Dnis

	Votantes := make([]votos.Votante, len(arrayDeDnis))

	for j := 0; j < len(arrayDeDnis); j++ {

		Votantes[j] = votos.CrearVotante(arrayDeDnis[j])

	}

	return Votantes, nil
}
