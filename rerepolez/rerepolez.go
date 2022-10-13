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

func impresionFinalDeLaVotacion(lista []votos.Partido) {

	for candidato := votos.PRESIDENTE; candidato <= votos.INTENDENTE; candidato++ {

		switch candidato {
		case 0:
			fmt.Fprintf(os.Stdout, "%s\n", "Presidente:")
		case 1:
			fmt.Fprintf(os.Stdout, "%s\n", "Gobernador:")
		case 2:
			fmt.Fprintf(os.Stdout, "%s\n", "Intendente:")
		}

		for p := 0; p < len(lista); p++ {

			fmt.Fprintf(os.Stdout, "%s\n", (lista)[p].ObtenerResultado(candidato))

		}
		fmt.Println()
	}

	if votos.LISTA_IMPUGNA == 1 {

		fmt.Fprintf(os.Stdout, "Votos Impugnados: %d voto\n", votos.LISTA_IMPUGNA)

	} else {

		fmt.Fprintf(os.Stdout, "Votos Impugnados: %d votos\n", votos.LISTA_IMPUGNA)

	}

}

func main() {

	parametros := os.Args[1:]

	if len(parametros) == 0 { //si no hay parametros tira error

		err := new(Err.ErrorParametros)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return

	}

	rutaListas := parametros[0]

	listaDeLosPartidos, err := PrepararListaPartidos(rutaListas)

	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	if len(parametros) != 2 { //si los parametros no terminan de ser suficientes

		err := new(Err.ErrorParametros)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return

	}

	rutaPadrones := parametros[1]

	Votantes, err := PrepararListaVotantes(rutaPadrones)

	if err != nil {

		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return

	}

	defer impresionFinalDeLaVotacion(listaDeLosPartidos)

	//implementacion de elecciones

	filaVotacion := TDACola.CrearColaEnlazada[votos.Votante]()

	entradaUsuario := bufio.NewScanner(os.Stdin)
	for entradaUsuario.Scan() {

		comandos := strings.Split(entradaUsuario.Text(), " ")

		if comandos[0] == "ingresar" {

			dni, err := strconv.Atoi(comandos[1])

			if err != nil || dni <= 0 {

				err = new(Err.DNIError)
				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				continue

			}

			posicionEnLaLista := BusquedaVotante(Votantes, dni, 0, len(Votantes)-1)

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

		case "votar":

			numeroDeBoleta, err := strconv.Atoi(comandos[2])

			if err != nil || numeroDeBoleta > len(listaDeLosPartidos)-1 || numeroDeBoleta < 0 {

				err = new(Err.ErrorAlternativaInvalida)
				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				continue

			}

			switch comandos[1] {

			case "Presidente":
				err = filaVotacion.VerPrimero().Votar(votos.PRESIDENTE, numeroDeBoleta)

			case "Gobernador":
				err = filaVotacion.VerPrimero().Votar(votos.GOBERNADOR, numeroDeBoleta)

			case "Intendente":
				err = filaVotacion.VerPrimero().Votar(votos.INTENDENTE, numeroDeBoleta)

			default:
				err = filaVotacion.VerPrimero().Votar(votos.CUALQUIERCOSA, numeroDeBoleta)
			}

			if err != nil {

				if err.Error() == fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", filaVotacion.VerPrimero().LeerDNI()) { //si el votante es fraudulento

					filaVotacion.Desencolar()

				}

				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				continue

			}

		case "deshacer":

			err = filaVotacion.VerPrimero().Deshacer()

			if err != nil {

				if err.Error() == fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", filaVotacion.VerPrimero().LeerDNI()) { //si el votante es fraudulento

					filaVotacion.Desencolar()

				}

				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				continue

			}

		case "fin-votar":

			VotoTerminado, err := filaVotacion.VerPrimero().FinVoto()
			filaVotacion.Desencolar()

			if err != nil {

				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				continue

			} else if VotoTerminado.Impugnado != true {

				for puesto := votos.PRESIDENTE; puesto <= votos.INTENDENTE; puesto++ {
					listaDeLosPartidos[VotoTerminado.VotoPorTipo[puesto]].VotadoPara(puesto)
				}

			}

		}

		fmt.Fprintf(os.Stdout, "%s\n", "OK")

	}
	if !filaVotacion.EstaVacia() { //si los votantes no terminaron de votar

		err = new(Err.ErrorCiudadanosSinVotar)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
	}
}
