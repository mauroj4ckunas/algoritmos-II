package main

import (
	TDAPila "Pila"
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
		var error ErrorParametros = ErrorParametros
		return error.Error()
	}

	rutaListas := parametros[0]    //el primer parametro era el nombre del archivo de las listas
	rutaPadrones := parametros[1]  //el segundo el de los padrones

	


	//implementacion array de partidos

	archivoListas, err := os.Open(rutaListas)
	defer archivoListas.Close()
	if err != nil { //si la ruta no se puede leer o algo, error
		var error ErrorLeerArchivo = ErrorLeerArchivo
		return error.Error()
	}



	//implementacion array de votantes

	pila := TDAPila.CrearPilaDinamica[string]() // esto es para usarlo para crear los array

	archivoVotantes, err := os.Open(rutaPadrones)
	defer archivoVotantes.Close()
	cantidad_votantes := 0

	if err != nil { //si la ruta no se puede leer o algo, error
		var error ErrorLeerArchivo = ErrorLeerArchivo
		fmt.Println(error.Error())
		return
	}

	padron := bufio.NewScanner(archivoVotantes)
  	for padron.Scan() {
  		pila.Apilar(padron.Text())
     	cantidad_votantes += 1
  	}

  	error = padron.Err()
  	if err != nil {
     	fmt.Println(err)
  	}

  	votantes := make([cantidad_votantes]Voto.Votante,cantidad_votantes)
  	for i:= 0 ; i < cantidad_votantes ; i++ {
  		votantes[i] = Voto.CrearVotante(strconv.Atoi(pila.Desapilar()))
  	}



}
