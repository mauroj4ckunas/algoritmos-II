package main

import (
	red "algogram/redsocial"
	usuarios "algogram/usuarios"
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
	Algogram, err := red.CrearAlgoGram[string](archivoUsuarios[0])

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
			err := usuarios.Login(usuario)
			fmt.Fprintf(os.Stdout, "%s\n", err)

		case COMANDO2:
			err := usuarios.Logout()
			fmt.Fprintf(os.Stdout, "%s\n", err)

		case COMANDO3:
			post := comandos[1:]
			err := usuarios.Publicar(post)
			fmt.Fprintf(os.Stdout, "%s\n", err)
		default:
			/* code */
			return
		}
	}
}
