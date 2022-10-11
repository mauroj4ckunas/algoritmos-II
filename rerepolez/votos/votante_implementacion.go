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

		var error Err.ErrorVotanteFraudulento = Err.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
		return error

	}

	if alternativa == 0 {

		votante.votosAnteriores.Apilar(*votante.voto)

		if votante.voto.Impugnado == false { //se impugna su voto por primera vez

			votante.voto.Impugnado = true
			LISTA_IMPUGNA += 1

		}

		return nil

	} else if tipo == CUALQUIERCOSA {

		error := new(Err.ErrorTipoVoto)
		return error

	}

	votante.votosAnteriores.Apilar(*votante.voto)
	votante.voto.VotoPorTipo[tipo] = alternativa
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {

	if votante.finalizoSuVoto == true {

		var error1 Err.ErrorVotanteFraudulento = Err.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
		return error1

	} else if votante.votosAnteriores.EstaVacia() {

		error2 := new(Err.ErrorNoHayVotosAnteriores)
		return error2
	}

	if votante.voto.Impugnado == true && votante.votosAnteriores.VerTope().Impugnado == false {

		LISTA_IMPUGNA -= 1

	}

	*votante.voto = votante.votosAnteriores.Desapilar()
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {

	if votante.finalizoSuVoto == true {

		var error Err.ErrorVotanteFraudulento = Err.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
		return Voto{}, error

	}

	votante.finalizoSuVoto = true
	return *votante.voto, nil
}
