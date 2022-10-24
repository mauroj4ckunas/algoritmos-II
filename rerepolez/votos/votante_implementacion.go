package votos

import (
	TDAPila "rerepolez/Pila"
	Err "rerepolez/errores"
)

type votanteImplementacion struct {
	dni             int
	voto            *Voto
	votosAnteriores TDAPila.Pila[Voto]
	finalizoSuVoto  bool
}

func CrearVotante(dni int) Votante {

	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.voto = new(Voto)
	votante.votosAnteriores = TDAPila.CrearPilaDinamica[Voto]()
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {

	if votante.finalizoSuVoto == true {
		return Err.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
	}
	if alternativa == 0 {
		votante.votosAnteriores.Apilar(*votante.voto)

		if votante.voto.Impugnado == false { //se impugna su voto por primera vez

			votante.voto.Impugnado = true

		}
		return nil

	} else if tipo == CANDIDATOERRONEO {

		return new(Err.ErrorTipoVoto)

	}
	votante.votosAnteriores.Apilar(*votante.voto)
	votante.voto.VotoPorTipo[tipo] = alternativa
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {

	if votante.finalizoSuVoto == true {
		return Err.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}

	} else if votante.votosAnteriores.EstaVacia() {
		return new(Err.ErrorNoHayVotosAnteriores)
	}

	*votante.voto = votante.votosAnteriores.Desapilar()
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {

	if votante.finalizoSuVoto == true {
		return Voto{}, Err.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
	}
	votante.finalizoSuVoto = true
	return *votante.voto, nil
}
