package votos

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

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	nuevoPartido := new(partidoImplementacion)
	nuevoPartido.nombre_part = nombre
	nuevoPartido.inscribirCandidatos(candidatos)
	return nuevoPartido
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	if tipo == 0 {
		partido.presidente.cant_votos++
	} else if tipo == 1 {
		partido.gobernador.cant_votos++
	} else if tipo == 2 {
		partido.intendente.cant_votos++
	}
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	var nombreCandidato string
	var cantidadVotos int
	if tipo == 0 {
		nombreCandidato = partido.presidente.nombre
		cantidadVotos = partido.presidente.cant_votos
	} else if tipo == 1 {
		nombreCandidato = partido.gobernador.nombre
		cantidadVotos = partido.gobernador.cant_votos
	} else if tipo == 2 {
		nombreCandidato = partido.intendente.nombre
		cantidadVotos = partido.intendente.cant_votos
	}
	return fmt.Sprintf("%s: %d votos.", nombreCandidato, cantidadVotos)
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
