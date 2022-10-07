package main

import (
	TDAPila "Pila"
	TDACola "Cola"
	Err "errores"
	Voto "votos"
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {

	parametros := os.Args[1:] //recibe los nombres de los archivos pasados por parametro en un array
							  // el [1:] es para sacar el nombre del archivo (main)
	
	if len(parametros) != 2 { //si los parametros son mas chicos que dos, error
		ErrorInicial := new(Err.ErrorParametros)
		fmt.Println(ErrorInicial.Error())
		return
	}

	rutaListas := parametros[0]    //el primer parametro era el nombre del archivo de las listas
	rutaPadrones := parametros[1]  //el segundo el de los padrones

	


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
	partido := make([]Votos.Partido, cantidad_partidos)

	//Creo el partido que recibira los votos en blanco
	candVacio := [3]string{"", "", ""}
	partidoEnBlanco := Votos.CrearPartido("Votos en Blanco", candVacio)
	partido[0] = partidoEnBlanco

	i := 1
	for listaPartidos.Scan() {
		grupo := strings.Split(listaPartidos.Text(), ",")
		nombrePartido := grupo[0]
		candidatosPartido := [3]string{grupo[1], grupo[2], grupo[3]}

		nuevoPartido := Votos.CrearPartido(nombrePartido, candidatosPartido)
		partido[i] = nuevoPartido
		i++
	}

	//implementacion array de votantes
  	Votantes, err := PrepararListaVotantes(rutaPadrones)

  	if err != nil{
  		fmt.Println(err.Error())
  		return
  	}


  	//implementacion de elecciones
  	filaVotacion := TDACola.CrearColaEnlazada[*Voto.Votante]()
  	s := bufio.NewScanner(os.Stdin)
  	for s.Scan() {
  		comandos := strings.Split(s.Text())
     	switch comandos[0]{
     	case "ingresar"{
     		dni := strconv.Atoi(comandos[1])
     		if dni <= 0{
     			err := new(DNIError)
     			fmt.Println(err.Error())
     			break
     		}
     		posicion := BusquedaVotante(Votantes,dni)
     		if posicion== -1{
     			err := new(DNIFueraPadron)
     			fmt.Println(err.Error())
     			break
     		}
     		filaVotacion.Encolar(&Votantes[posicion])
     	}
     	case "votar"{
		if comandos[1] != "Presidente" &&
		   comandos[1] != "Gobernador" &&
		   comandos[1] != "Intendente" {
			ErrorMsj:= new(Err.ErrorTipoVoto)
			fmt.Println(ErrorMsj.Error())
		} else if comandos[2] > len(partido) - 1 || comandos[2] < 0 {
			ErrorMsj:= new(Err.ErrorAlternativaInvalida)
			fmt.Println(ErrorMsj.Error())
			} else {
				if comandos[1] == "Presidente" {
					partido[comandos[2]].VotadoPara(0)
				} else if comandos[1] == "Gobernador" {
					partido[comandos[2]].VotadoPara(1)
				} else if comandos[1] == "Intendente" {
					partido[comandos[2]].VotadoPara(2)
				}
			}
		}
	}
     	case "deshacer"{}
     	case "fin-votar"{}
     	default{}
     	}
    }


}
