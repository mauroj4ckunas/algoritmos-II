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
	return nil
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {

}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	return ""
}
