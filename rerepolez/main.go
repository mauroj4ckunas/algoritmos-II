package main

import (
	TDACola "rerepolez/Cola"
	"bufio"
	Err "rerepolez/errores"
	"fmt"
	"os"
	"strconv"
	"strings"
	votos "rerepolez/votos"
)

func finDeEjecucion(listaPartidos []votos.Partido) {

	for candidato := votos.PRESIDENTE; candidato <= votos.INTENDENTE; candidato++ {

		switch candidato {
		case 0:
			fmt.Println("Presidente:")
		case 1:
			fmt.Println("Gobernador:")
		case 2:
			fmt.Println("Intendente:")
		}

		for p := 0; p < len(listaPartidos); p++ {

			fmt.Println((listaPartidos)[p].ObtenerResultado(candidato))

		}
		fmt.Println()
	}

}

func main() {

	parametros := os.Args[1:] //recibe los nombres de los archivos pasados por parametro en un array
	// el [1:] es para sacar el nombre del archivo (main)

	if len(parametros) == 0 {
		ErrorInicial := new(Err.ErrorParametros)
		fmt.Println(ErrorInicial.Error())
		return
	}

	//implementacion array de partidos

	rutaListas := parametros[0] //el primer parametro era el nombre del archivo de las listas

	partido, errPartido := PrepararListaPartidos(rutaListas)

	if errPartido != nil {
		fmt.Println(errPartido.Error())
		return
	}

	if len(parametros) != 2 { //si los parametros son mas chicos que dos, error
		ErrorInicial := new(Err.ErrorParametros)
		fmt.Println(ErrorInicial.Error())
		return
	}

	//implementacion array de votantes

	rutaPadrones := parametros[1] //el segundo el de los padrones

	Votantes, err := PrepararListaVotantes(rutaPadrones)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//implementacion de final de la votacion

	//PODRIAMOS MEJORARLO Y REHACERLO EN UNA FUNCION
	defer func() {
		if votos.LISTA_IMPUGNA == 1 {
			fmt.Printf("Votos Impugnados: %d voto\n", votos.LISTA_IMPUGNA)
		} else {
			fmt.Printf("Votos Impugnados: %d votos\n", votos.LISTA_IMPUGNA)
		}

	}()
	defer finDeEjecucion(partido)

	//implementacion de elecciones

	filaVotacion := TDACola.CrearColaEnlazada[votos.Votante]()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {

		comandos := strings.Split(s.Text(), " ")

		if comandos[0] == "ingresar" {

			dni, _ := strconv.Atoi(comandos[1])

			if dni <= 0 {

				err = new(Err.DNIError)
				fmt.Println(err.Error())
				continue

			}

			posicion := BusquedaVotante(Votantes, dni, 0, len(Votantes)-1)

			if posicion == -1 {

				err = new(Err.DNIFueraPadron)
				fmt.Println(err.Error())
				continue

			}

			filaVotacion.Encolar(Votantes[posicion])
			fmt.Println("OK")
			continue
		}

		if filaVotacion.EstaVacia() {

			err := new(Err.FilaVacia)
			fmt.Println(err.Error())
			continue

		}

		switch comandos[0] {

		case "votar":

			comand2, err := strconv.Atoi(comandos[2])

			if err != nil || comand2 > len(partido)-1 || comand2 < 0 {

				err = new(Err.ErrorAlternativaInvalida)
				fmt.Println(err.Error())
				continue

			}

			switch comandos[1] {

			case "Presidente":
				err = filaVotacion.VerPrimero().Votar(votos.PRESIDENTE, comand2)

			case "Gobernador":
				err = filaVotacion.VerPrimero().Votar(votos.GOBERNADOR, comand2)

			case "Intendente":
				err = filaVotacion.VerPrimero().Votar(votos.INTENDENTE, comand2)

			default:
				err = filaVotacion.VerPrimero().Votar(4, comand2)
			}

			if err != nil {
				if err.Error() == fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", filaVotacion.VerPrimero().LeerDNI()) {

					filaVotacion.Desencolar()

				}

				fmt.Println(err.Error())
				continue

			}

		case "deshacer":

			err = filaVotacion.VerPrimero().Deshacer()
			if err != nil {

				if err.Error() == fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", filaVotacion.VerPrimero().LeerDNI()) {

					filaVotacion.Desencolar()

				}

				fmt.Println(err.Error())
				continue

			}

		case "fin-votar":

			VotoTerminado, err := filaVotacion.VerPrimero().FinVoto()
			filaVotacion.Desencolar()

			if err != nil {

				fmt.Println(err.Error())
				continue

			} else if VotoTerminado.Impugnado == true {

				fmt.Println("OK")
				continue

			}
			for puesto := votos.PRESIDENTE; puesto <= votos.INTENDENTE; puesto++ {
				partido[VotoTerminado.VotoPorTipo[puesto]].VotadoPara(puesto)
			}

		}
		fmt.Println("OK")

	}
	if !filaVotacion.EstaVacia() {

		err = new(Err.ErrorCiudadanosSinVotar)
		fmt.Println(err.Error())

	}
}
