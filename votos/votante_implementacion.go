package votos

import (
	TDAPila "tp1/Pila"
	Err "tp1/errores"
)

type votanteImplementacion struct {
	dni        int
	voto       *Voto
	decisiones TDAPila.Pila[Voto]
	FinDeVoto  bool
}

func CrearVotante(dni int) Votante {

	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.voto = new(Voto)
	votante.decisiones = TDAPila.CrearPilaDinamica[Voto]()
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) Err.Errores {

	if votante.FinDeVoto == true {

		var error Err.ErrorVotanteFraudulento = Err.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
		return error

	}

	if alternativa == 0 {

		if votante.voto.Impugnado == false {

			LISTA_IMPUGNA += 1

		}
		votante.decisiones.Apilar(*votante.voto)
		votante.voto.Impugnado = true
		return nil

	} else if tipo != PRESIDENTE && tipo != GOBERNADOR && tipo != INTENDENTE {

		error := new(Err.ErrorTipoVoto)
		return error

	}

	votante.decisiones.Apilar(*votante.voto)
	votante.voto.VotoPorTipo[tipo] = alternativa
	return nil
}

func (votante *votanteImplementacion) Deshacer() Err.Errores {

	if votante.FinDeVoto == true {

		var error1 Err.ErrorVotanteFraudulento = Err.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
		return error1

	} else if votante.decisiones.EstaVacia() {

		error2 := new(Err.ErrorNoHayVotosAnteriores)
		return error2
	}

	if votante.voto.Impugnado == true && votante.decisiones.VerTope().Impugnado == false {

		LISTA_IMPUGNA -= 1

	}

	*votante.voto = votante.decisiones.Desapilar()
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, Err.Errores) {

	if votante.FinDeVoto == true {

		var error Err.ErrorVotanteFraudulento = Err.ErrorVotanteFraudulento{Dni: votante.dni}
		return Voto{}, error

	}

	votante.FinDeVoto = true
	return *votante.voto, nil
}
