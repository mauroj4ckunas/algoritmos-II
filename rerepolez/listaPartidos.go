package main

import (
	"bufio"
	"os"
	Err "rerepolez/errores"
	votos "rerepolez/votos"
	"strings"
)

func PrepararListaPartidos(ruta string) ([]votos.Partido, error) {

	archivoListas, err := os.Open(ruta)
	defer archivoListas.Close()

	if err != nil { //si la ruta no se puede leer o algo, error

		ErrorLectura := new(Err.ErrorLeerArchivo)
		return []votos.Partido{}, ErrorLectura

	}

	listaPartidos := bufio.NewScanner(archivoListas)

	Partidos := make([]votos.Partido, 1)
	candVacio := [3]string{"", "", ""}
	partidoEnBlanco := votos.CrearPartido("Votos en Blanco", candVacio)
	Partidos[0] = partidoEnBlanco

	for listaPartidos.Scan() {

		boleta := strings.Split(listaPartidos.Text(), ",")

		if len(boleta) != 4 { //si el archivo no esta escrito como se debe o es el archivo de votantes

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
