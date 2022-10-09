package votos

import "fmt"

type partidoImplementacion struct {
	nombre_part string
	presidente  *candidatosParaEleccion
	gobernador  *candidatosParaEleccion
	intendente  *candidatosParaEleccion
}

type candidatosParaEleccion struct {
	nombre     string
	cant_votos int
}

func CrearPartido(nombre string, candidatos [3]string) Partido {
	nuevoPartido := new(partidoImplementacion)
	nuevoPartido.nombre_part = nombre
	nuevoPartido.inscribirCandidatos(candidatos)
	return nuevoPartido
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	switch tipo {
	case 0:
		partido.presidente.cant_votos++
	case 1:
		partido.gobernador.cant_votos++
	case 2:
		partido.intendente.cant_votos++
	}
}

func (partido *partidoImplementacion) RestarVoto(tipo TipoVoto) {
	switch tipo {
	case 0:
		partido.presidente.cant_votos--
	case 1:
		partido.gobernador.cant_votos--
	case 2:
		partido.intendente.cant_votos--
	}
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	var nombreCandidato string
	var cantidadVotos int
	var palabra string

	switch tipo {
	case 0:
		nombreCandidato = partido.presidente.nombre
		cantidadVotos = partido.presidente.cant_votos
	case 1:
		nombreCandidato = partido.gobernador.nombre
		cantidadVotos = partido.gobernador.cant_votos
	case 2:
		nombreCandidato = partido.intendente.nombre
		cantidadVotos = partido.intendente.cant_votos
	}

	if cantidadVotos == 1 {
		palabra = "voto"
	} else {
		palabra = "votos"
	}

	if partido.nombre_part == "Votos en Blanco" {
		return fmt.Sprintf("%s: %d %s", partido.nombre_part, cantidadVotos, palabra)
	}
	return fmt.Sprintf("%s - %s: %d %s", partido.nombre_part, nombreCandidato, cantidadVotos, palabra)
}

func (partido *partidoImplementacion) inscribirCandidatos(candidatos [CANT_VOTACION]string) {
	candidatoPres := new(candidatosParaEleccion)
	candidatoPres.nombre = candidatos[0]
	partido.presidente = candidatoPres

	candidatoGob := new(candidatosParaEleccion)
	candidatoGob.nombre = candidatos[1]
	partido.presidente = candidatoGob

	candidatoInt := new(candidatosParaEleccion)
	candidatoInt.nombre = candidatos[2]
	partido.intendente = candidatoInt
}
