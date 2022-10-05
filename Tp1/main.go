package main

import(
    TDAPila "main/Pila"
    Err "main/errores"
	"bufio"
	"fmt"
	"os"
)


func main(){
	
	parametros := os.Args[1:]
	if len(parametros) != 2 {
		var error ErrorParametros = ErrorParametros
		return error.Error()
	}

	rutaListas := parametros[0]
	rutaPadrones := parametros[1]

	pila := TDAPila.CrearPilaDinamica[string]()

	archivoListas, err := os.Open(rutaListas) 
	defer archivoListas.Close()
	if err != nil {
		var error ErrorLeerArchivo = ErrorLeerArchivo
		return error.Error()
	}
}
