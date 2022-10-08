package main

import (
	TDACola "Cola"
	"strings"

	//TDAPila "Pila"
	"bufio"
	Err "errores"
	"fmt"
	"os"
	"strconv"
	Voto "votos"
)

func main() {

	parametros := os.Args[1:] //recibe los nombres de los archivos pasados por parametro en un array
	// el [1:] es para sacar el nombre del archivo (main)

	if len(parametros) != 2 { //si los parametros son mas chicos que dos, error
		ErrorInicial := new(Err.ErrorParametros)
		fmt.Println(ErrorInicial.Error())
		return
	}

	rutaListas := parametros[0]   //el primer parametro era el nombre del archivo de las listas
	rutaPadrones := parametros[1] //el segundo el de los padrones

	//implementacion array de partidos

	archivoListas, err := os.Open(rutaListas)
	defer archivoListas.Close()
	if err != nil { //si la ruta no se puede leer o algo, error
		ErrorLectura := new(Err.ErrorLeerArchivo)
		fmt.Println(ErrorLectura.Error())
		return
	}

	listaPartidos := bufio.NewScanner(archivoListas)
	var cantidad_partidos int
	for listaPartidos.Scan() {
		cantidad_partidos++ //cuenta la cantidad de partidos que hay para hacer el arreglo
	}
	partido := make([]Voto.Partido, cantidad_partidos)

	//Creo el partido que recibira los votos en blanco
	candVacio := [3]string{"", "", ""}
	partidoEnBlanco := Voto.CrearPartido("Votos en Blanco", candVacio)
	partido[0] = partidoEnBlanco

	i := 1
	for listaPartidos.Scan() {
		grupo := strings.Split(listaPartidos.Text(), ",")
		nombrePartido := grupo[0]
		candidatosPartido := [3]string{grupo[1], grupo[2], grupo[3]}

		nuevoPartido := Voto.CrearPartido(nombrePartido, candidatosPartido)
		partido[i] = nuevoPartido
		i++
	}

	//implementacion array de votantes
	Votantes, err := PrepararListaVotantes(rutaPadrones)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {

		for k := PRESIDENTE; k <= INTENDENTE; k++ {
			switch k {
			case 0:
				fmt.Println("Presidente: ")
			case 1:
				fmt.Println("Gobernador: ")
			case 2:
				fmt.Println("Intendente: ")
			}
			for p := 0; p < cantidad_partidos; p++ {
				partido[p].ObtenerResultado(k)
			}
		}

	}()

	//implementacion de elecciones

	filaVotacion := TDACola.CrearColaEnlazada[Voto.Votante]()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {

		comandos := strings.Split(s.Text())

		if comandos[0] == "ingresar" {
			dni := strconv.Atoi(comandos[1])
			if dni <= 0 {

				err := new(Err.DNIError)
				fmt.Println(err.Error())
				continue

			}

			posicion := BusquedaVotante(Votantes, dni)

			if posicion == -1 {

				err := new(Err.DNIFueraPadron)
				fmt.Println(err.Error())
				continue

			}

			filaVotacion.Encolar(&Votantes[posicion])
			continue
		}

		if filaVotacion.EstaVacia() {

			err := new(Err.FilaVacia)
			fmt.Println(err.Error())
			continue

		} else if filaVotacion.VerPrimero().FraudulentoPorPrimeraVez() {

			err := filaVotacion.VerPrimero().Votar(0, 0)
			filaVotacion.Desencolar()
			fmt.Println(err.Error())

			//IMPLEMENTACION DE RESTAR VOTO A LOS PARTIDOS
			continue
		}

		switch comandos[0] {

		case "votar":

			if comandos[1] != "Presidente" && comandos[1] != "Gobernador" && comandos[1] != "Intendente" {

				err := new(Err.ErrorTipoVoto)
				fmt.Println(err.Error())
				continue

			} else if comandos[2] > len(partido)-1 || comandos[2] < 0 {

				err := new(Err.ErrorAlternativaInvalida)
				fmt.Println(err.Error())
				continue

			}

			switch comandos[1] {

			case "Presidente":
				err := filaVotacion.VerPrimero().Votar(PRESIDENTE, comandos[2])
				if err != nil {

					filaVotacion.Desencolar()
					fmt.Println(err.Error())
					continue

				}

			case "Gobernador":
				err := filaVotacion.VerPrimero().Votar(GOBERNADOR, comandos[2])
				if err != nil {

					filaVotacion.Desencolar()
					fmt.Println(err.Error())
					continue

				}

			case "Intendente":
				err := filaVotacion.VerPrimero().Votar(GOBERNADOR, comandos[2])
				if err != nil {

					filaVotacion.Desencolar()
					fmt.Println(err.Error())
					continue

				}

			}

		case "deshacer":

			err := filaVotacion.VerPrimero().Deshacer()
			if err != nil {

				filaVotacion.Desencolar()
				fmt.Println(err.Error())
				continue

			}

		case "fin-votar":

			VotoTerminado, err := filaVotacion.VerPrimero().FinVoto()
			if err != nil {

				filaVotacion.Desencolar()
				fmt.Println(err.Error())
				continue

			}

			for puesto := PRESIDENTE; puesto < INTENDENTE+1; puesto++ {
				partido[VotoTerminado[puesto]].VotadoPara(puesto)
			}

		}
		fmt.Println("OK")
	}

}
