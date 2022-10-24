package votos

import "fmt"

type partidoImplementacion struct {
	nombre_Part string
	candidatos  [CANT_VOTACION]candidatosParaEleccion
}

type candidatosParaEleccion struct {
	nombre     string
	cant_votos int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {

	nuevoPartido := new(partidoImplementacion)
	nuevoPartido.nombre_Part = nombre
	for i := PRESIDENTE; i < CANT_VOTACION; i++ {
		nuevoPartido.candidatos[i].nombre = candidatos[i]
	}
	return nuevoPartido

}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.candidatos[tipo].cant_votos += 1
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	var palabra string

	if partido.candidatos[tipo].cant_votos == 1 {
		palabra = "voto"
	} else {
		palabra = "votos"
	}

	if partido.nombre_Part == "Votos en Blanco" {
		return fmt.Sprintf("%s: %d %s", partido.nombre_Part, partido.candidatos[tipo].cant_votos, palabra)
	}
	return fmt.Sprintf("%s - %s: %d %s", partido.nombre_Part, partido.candidatos[tipo].nombre, partido.candidatos[tipo].cant_votos, palabra)
}
