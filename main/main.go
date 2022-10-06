package main

import (
	TDAPila "Pila"
	Err "errores"
	"bufio"
	"fmt"
	"os"
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

	pila := TDAPila.CrearPilaDinamica[string]() // esto es para usarlo para crear los array de los partidos y los participantes


	//implementacion array de partidos

	archivoListas, err := os.Open(rutaListas)
	defer archivoListas.Close()
	if err != nil { //si la ruta no se puede leer o algo, error
		var error ErrorLeerArchivo = ErrorLeerArchivo
		return error.Error()
	}



	//implementacion array de votantes

	archivoVotantes, err := os.Open(rutaPadrones)
	defer archivoVotantes.Close()
	if err != nil { //si la ruta no se puede leer o algo, error
		var error ErrorLeerArchivo = ErrorLeerArchivo
		return error.Error()
	}

}
