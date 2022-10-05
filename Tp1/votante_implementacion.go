package votos

import(
    TDAPila "main/Pila"
    Err "main/errores"
)

type votanteImplementacion struct {
	dni 		int
	voto 		*Voto
	decisiones	TDAPila.Pila[[CANT_VOTACION]int]
	FinDeVoto	bool
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.decisiones = TDAPila.CrearPilaDinamica[[CANT_VOTACION]int]()
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if votante.FinDeVoto == true{
		votante.voto.Impugnado = true
		var error Err.ErrorVotanteFraudulento = Err.ErrorVotanteFraudulento{Dni:votante.dni}
		return error
	}
	
	votante.decisiones.Apilar(votante.voto.VotoPorTipo)
	votante.voto.VotoPorTipo[tipo] = alternativa
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {

	if votante.FinDeVoto == true {
		votante.voto.Impugnado = true
		var error1 Err.ErrorVotanteFraudulento = Err.ErrorVotanteFraudulento{Dni:votante.dni}
		return error1
	} else if votante.decisiones.EstaVacia(){
		error2 := new(Err.ErrorNoHayVotosAnteriores)
		return error2
	}

	votante.voto.VotoPorTipo = votante.decisiones.Desapilar()
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.FinDeVoto == true {
		votante.voto.Impugnado = true
		var error Err.ErrorVotanteFraudulento = Err.ErrorVotanteFraudulento{Dni:votante.dni}
		return *votante.voto, error
	}
	votante.FinDeVoto = true
	return *votante.voto, nil
}
