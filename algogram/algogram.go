package main

import (
	red "algogram/redsocial"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	COMANDO1 = "login"
	COMANDO2 = "logout"
	COMANDO3 = "publicar"
)

func main() {
	archivoUsuarios := os.Args[1:]
	algogram, err := red.CrearAlgoGram[string](archivoUsuarios[0])

	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	entradaUsuario := bufio.NewScanner(os.Stdin)
	for entradaUsuario.Scan() {
		comandos := strings.Split(entradaUsuario.Text(), " ")
		switch comandos[0] {
		case COMANDO1:
			usuario := comandos[1]
			err := algogram.Login(usuario)
			fmt.Fprintf(os.Stdout, "%s\n", err)

		case COMANDO2:
			err := algogram.Logout()
			fmt.Fprintf(os.Stdout, "%s\n", err)

		case COMANDO3:
			post := comandos[1:]
			err := algogram.Publicar(post)
			fmt.Fprintf(os.Stdout, "%s\n", err)
		default:
			/* code */
			return
		}
	}
}
