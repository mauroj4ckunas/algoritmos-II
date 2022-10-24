package main

import (
	"bufio"
	"fmt"
	"os"
	TDACola "rerepolez/Cola"
	Err "rerepolez/errores"
	votos "rerepolez/votos"
	"strconv"
	"strings"
)

const (
	CANDIDATO1 = "Presidente"
	CANDIDATO2 = "Gobernador"
	CANDIDATO3 = "Intendente"
	COMAND1    = "ingresar"
	COMAND2    = "votar"
	COMAND3    = "deshacer"
	COMAND4    = "fin-votar"
)

func impresionFinalDeLaVotacion(lista []votos.Partido, impugnados int) {

	for candidato := votos.PRESIDENTE; candidato <= votos.INTENDENTE; candidato++ {

		switch candidato {
		case 0:
			fmt.Fprintf(os.Stdout, "%s:\n", CANDIDATO1)
		case 1:
			fmt.Fprintf(os.Stdout, "%s:\n", CANDIDATO2)
		case 2:
			fmt.Fprintf(os.Stdout, "%s:\n", CANDIDATO3)
		}

		for p := 0; p < len(lista); p++ {

			fmt.Fprintf(os.Stdout, "%s\n", (lista)[p].ObtenerResultado(candidato))

		}
		fmt.Println()
	}

	if impugnados == 1 {

		fmt.Fprintf(os.Stdout, "Votos Impugnados: %d voto\n", impugnados)

	} else {

		fmt.Fprintf(os.Stdout, "Votos Impugnados: %d votos\n", impugnados)

	}

}

func main() {

	parametros := os.Args[1:]

	listaDeLosPartidos, Votantes, err := PrepararMesa(parametros)

	if err != nil {

		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	var contadorImpugnados int

	//implementacion de elecciones

	filaVotacion := TDACola.CrearColaEnlazada[votos.Votante]()

	entradaUsuario := bufio.NewScanner(os.Stdin)
	for entradaUsuario.Scan() {

		comandos := strings.Split(entradaUsuario.Text(), " ")

		if comandos[0] == COMAND1 {

			dni, err := strconv.Atoi(comandos[1])

			if err != nil || dni <= 0 {

				err = new(Err.DNIError)
				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				continue

			}
			posicionEnLaLista := BusquedaVotante[votos.Votante](Votantes, dni, 0, len(Votantes)-1, func(votante votos.Votante, buscado int) int {
				if votante.LeerDNI() == buscado {
					return 0
				} else if votante.LeerDNI() < buscado {
					return -1
				}
				return 1
			})

			if posicionEnLaLista == -1 { //-1 es q no esta en la lista

				err = new(Err.DNIFueraPadron)
				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				continue

			}

			filaVotacion.Encolar(Votantes[posicionEnLaLista])
			fmt.Fprintf(os.Stdout, "%s\n", "OK")
			continue
		}

		if filaVotacion.EstaVacia() { //para cualquier otro comando si la fila esta vacia tira error

			err := new(Err.FilaVacia)
			fmt.Fprintf(os.Stdout, "%s\n", err.Error())
			continue

		}

		switch comandos[0] {

		case COMAND2:

			numeroDeBoleta, err := strconv.Atoi(comandos[2])

			if err != nil || numeroDeBoleta > len(listaDeLosPartidos)-1 || numeroDeBoleta < 0 {

				err = new(Err.ErrorAlternativaInvalida)
				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				continue

			}

			switch comandos[1] {

			case CANDIDATO1:
				err = filaVotacion.VerPrimero().Votar(votos.PRESIDENTE, numeroDeBoleta)

			case CANDIDATO2:
				err = filaVotacion.VerPrimero().Votar(votos.GOBERNADOR, numeroDeBoleta)

			case CANDIDATO3:
				err = filaVotacion.VerPrimero().Votar(votos.INTENDENTE, numeroDeBoleta)

			default:
				err = filaVotacion.VerPrimero().Votar(votos.CANDIDATOERRONEO, numeroDeBoleta)
			}

			if err != nil {

				if err.Error() == fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", filaVotacion.VerPrimero().LeerDNI()) { //si el votante es fraudulento

					filaVotacion.Desencolar()

				}

				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				continue

			}

		case COMAND3:

			err = filaVotacion.VerPrimero().Deshacer()

			if err != nil {

				if err.Error() == fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", filaVotacion.VerPrimero().LeerDNI()) { //si el votante es fraudulento

					filaVotacion.Desencolar()

				}

				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				continue

			}

		case COMAND4:

			VotoTerminado, err := filaVotacion.VerPrimero().FinVoto()
			filaVotacion.Desencolar()

			if err != nil {

				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				continue

			} else if VotoTerminado.Impugnado != true {

				for puesto := votos.PRESIDENTE; puesto < votos.CANT_VOTACION; puesto++ {
					listaDeLosPartidos[VotoTerminado.VotoPorTipo[puesto]].VotadoPara(puesto)
				}

			} else {
				contadorImpugnados += 1
			}

		}

		fmt.Fprintf(os.Stdout, "%s\n", "OK")

	}
	if !filaVotacion.EstaVacia() { //si los votantes no terminaron de votar

		err = new(Err.ErrorCiudadanosSinVotar)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
	}

	impresionFinalDeLaVotacion(listaDeLosPartidos, contadorImpugnados)
}
