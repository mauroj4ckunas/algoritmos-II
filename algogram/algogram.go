package main

import (
	TDAcola "algogram/Cola"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	funcionPrioridadEntreUsuarios = func (prioridad1,prioridad2 int) int {
		if prioridad1 < prioridad2 {
			return prioridad2 - prioridad1
		}
		return prioridad1 - prioridad2
	}
	COMANDO1 = "login"
	COMANDO2 = "logout"
	COMANDO3 = "publicar"
)

func main() {
	archivoUsuarios := os.Args[1:]
	Algogram := crearAlgoGram[string,int](archivoUsuarios[0],funcionPrioridadEntreUsuarios)

	entradaUsuario := bufio.NewScanner(os.Stdin)
	for entradaUsuario.Scan() {
		comandos := strings.Split(entradaUsuario.Text(), " ")
		switch comandos[0] {
		case COMANDO1:
			usuario := comandos[1]
			err := Algogram.Login(usuario)
			fmt.Fprintf(os.Stdout, "%s\n", err)

		case COMANDO2:
			err := Algogram.Logout()
			fmt.Fprintf(os.Stdout, "%s\n", err)

		case COMANDO3:
			post := comandos[1:]
			err := Algogram.Publicar(post)
			fmt.Fprintf(os.Stdout, "%s\n", err)
		default:
			/* code */
			return
		}
	}
}
