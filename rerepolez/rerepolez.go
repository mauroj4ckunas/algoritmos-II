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

var CANDIDATO = [votos.CANT_VOTACION]string{"Presidente", "Gobernador", "Intendente"}

const (
	COMAND1 = "ingresar"
	COMAND2 = "votar"
	COMAND3 = "deshacer"
	COMAND4 = "fin-votar"
)

func impresionFinalDeLaVotacion(lista []votos.Partido, impugnados int) {

	for candidato := votos.TipoVoto(0); candidato < votos.CANT_VOTACION; candidato++ {
		fmt.Fprintf(os.Stdout, "%s:\n", CANDIDATO[candidato])
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
			} else {
				posicionEnLaLista := BusquedaVotante[votos.Votante](Votantes, votos.CrearVotante(dni), 0, len(Votantes)-1, func(votante votos.Votante, buscado votos.Votante) int {
					if votante.LeerDNI() == buscado.LeerDNI() {
						return 0
					} else if votante.LeerDNI() < buscado.LeerDNI() {
						return -1
					}
					return 1
				})
				if posicionEnLaLista == -1 { //-1 es q no esta en la lista
					err = new(Err.DNIFueraPadron)
					fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				} else {
					filaVotacion.Encolar(Votantes[posicionEnLaLista])
					fmt.Fprintf(os.Stdout, "%s\n", "OK")
				}
			}
		} else {
			if filaVotacion.EstaVacia() { //para cualquier otro comando si la fila esta vacia tira error
				err := new(Err.FilaVacia)
				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
			} else {
				switch comandos[0] {
				case COMAND2:
					numeroDeBoleta, err := strconv.Atoi(comandos[2])
					if err != nil || numeroDeBoleta > len(listaDeLosPartidos)-1 || numeroDeBoleta < 0 {
						err = new(Err.ErrorAlternativaInvalida)
						fmt.Fprintf(os.Stdout, "%s\n", err.Error())
					} else {
						for candidatoElegido := votos.TipoVoto(0); candidatoElegido <= votos.CANDIDATOERRONEO; candidatoElegido++ {
							if candidatoElegido == votos.CANDIDATOERRONEO {
								err = filaVotacion.VerPrimero().Votar(candidatoElegido, numeroDeBoleta)
							} else if comandos[1] == CANDIDATO[candidatoElegido] {
								err = filaVotacion.VerPrimero().Votar(candidatoElegido, numeroDeBoleta)
								break
							}
						}
						if err != nil {
							if err.Error() == fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", filaVotacion.VerPrimero().LeerDNI()) { //si el votante es fraudulento
								filaVotacion.Desencolar()
							}
							fmt.Fprintf(os.Stdout, "%s\n", err.Error())
						} else {
							fmt.Fprintf(os.Stdout, "%s\n", "OK")
						}
					}
				case COMAND3:
					err = filaVotacion.VerPrimero().Deshacer()
					if err != nil {
						if err.Error() == fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", filaVotacion.VerPrimero().LeerDNI()) { //si el votante es fraudulento
							filaVotacion.Desencolar()
						}
						fmt.Fprintf(os.Stdout, "%s\n", err.Error())
					} else {
						fmt.Fprintf(os.Stdout, "%s\n", "OK")
					}
				case COMAND4:
					VotoTerminado, err := filaVotacion.VerPrimero().FinVoto()
					filaVotacion.Desencolar()
					if err != nil {
						fmt.Fprintf(os.Stdout, "%s\n", err.Error())
					} else if VotoTerminado.Impugnado != true {
						for puesto := votos.PRESIDENTE; puesto < votos.CANT_VOTACION; puesto++ {
							listaDeLosPartidos[VotoTerminado.VotoPorTipo[puesto]].VotadoPara(puesto)
						}
						fmt.Fprintf(os.Stdout, "%s\n", "OK")
					} else {
						contadorImpugnados += 1
						fmt.Fprintf(os.Stdout, "%s\n", "OK")
					}

				}
			}

		}
	}
	if !filaVotacion.EstaVacia() { //si los votantes no terminaron de votar
		err = new(Err.ErrorCiudadanosSinVotar)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
	}
	impresionFinalDeLaVotacion(listaDeLosPartidos, contadorImpugnados)
}
