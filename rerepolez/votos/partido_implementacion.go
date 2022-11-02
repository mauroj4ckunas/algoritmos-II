package votos

import "fmt"

type partidoImplementacion struct {
	nombrePart string
	candidatos [CANT_VOTACION]candidatosParaEleccion
}

type candidatosParaEleccion struct {
	nombre    string
	cantVotos int
}

func CrearPartido(nombre string, candidatos []string) Partido {

	nuevoPartido := new(partidoImplementacion)
	nuevoPartido.nombrePart = nombre
	for i := TipoVoto(0); i < CANT_VOTACION; i++ {
		nuevoPartido.candidatos[i].nombre = candidatos[i]
	}
	return nuevoPartido

}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.candidatos[tipo].cantVotos += 1
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	var palabra string

	if partido.candidatos[tipo].cantVotos == 1 {
		palabra = "voto"
	} else {
		palabra = "votos"
	}

	if partido.nombrePart == "Votos en Blanco" {
		return fmt.Sprintf("%s: %d %s", partido.nombrePart, partido.candidatos[tipo].cantVotos, palabra)
	}
	return fmt.Sprintf("%s - %s: %d %s", partido.nombrePart, partido.candidatos[tipo].nombre, partido.candidatos[tipo].cantVotos, palabra)
}
