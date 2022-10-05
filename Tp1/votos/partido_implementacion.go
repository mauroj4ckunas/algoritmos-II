package votos

type partidoImplementacion struct {
	nombre_part		string
	presidente		*candidatos
	gobernador		*candidatos
	intendente		*candidatos
}

type candidatos struct {
	nombre		string
	cant_votos	int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	nuevoPartido := new(partidoImplementacion)
	nuevoPartido.nombre_part = nombre
	nuevoPartido.inscribirCandidatos(candidatos)
	return nuevoPartido
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {

}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	return ""
}

func (partido partidoImplementacion) inscribirCandidatos(candidatos [CANT_VOTACION]string) {
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
