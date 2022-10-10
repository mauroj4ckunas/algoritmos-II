package main

import (
	"bufio"
	Err "errores"
	"os"
	"strings"
	"votos"
)

func PrepararListaPartidos(ruta string) ([]votos.Partido, Err.Errores) {
	archivoListas, err := os.Open(ruta)
	defer archivoListas.Close()
	if err != nil { //si la ruta no se puede leer o algo, error
		ErrorLectura := new(Err.ErrorLeerArchivo)
		return []votos.Partido{}, ErrorLectura
	}

	listaPartidos := bufio.NewScanner(archivoListas)

	partido := make([]votos.Partido, 1)
	candVacio := [3]string{"", "", ""}
	partidoEnBlanco := votos.CrearPartido("Votos en Blanco", candVacio)
	partido[0] = partidoEnBlanco

	for listaPartidos.Scan() {
		grupo := strings.Split(listaPartidos.Text(), ",")
		nombrePartido := grupo[0]
		candidatosPartido := [votos.CANT_VOTACION]string{grupo[1], grupo[2], grupo[3]}
		nuevoPartido := votos.CrearPartido(nombrePartido, candidatosPartido)
		partido = append(partido, nuevoPartido)
	}

	return partido, nil
}
