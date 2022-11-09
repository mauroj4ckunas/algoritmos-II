package main

import (
	usuarios "algogram/usuarios"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	COMANDO1               = "login"
	COMANDO2               = "logout"
	COMANDO3               = "publicar"
)


func main() {
	archivoUsuarios := os.Args[1:]
	Algogram , err:= crearAlgoGram(archivoUsuarios[0])
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
			mensaje := Algogram.Login(usuario)
			fmt.Fprintf(os.Stdout, "%s\n", mensaje)

		case COMANDO2:
			mensaje := Algogram.Logout()
			fmt.Fprintf(os.Stdout, "%s\n", mensaje)

		case COMANDO3:
			post := strings.Join(comandos[1:], " ")
			mensaje := Algogram.Publicar(post)
			fmt.Fprintf(os.Stdout, "%s\n", mensaje)
		default:
			/* code */
			return
		}
	}
}
