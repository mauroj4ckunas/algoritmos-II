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
	COMANDO4 = "ver_siguiente_feed"
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
			mensaje := algogram.Login(usuario)
			fmt.Fprintf(os.Stdout, "%s\n", mensaje)

		case COMANDO2:
			mensaje := algogram.Logout()
			fmt.Fprintf(os.Stdout, "%s\n", mensaje)

		case COMANDO3:
			post := strings.Join(comandos[1:], " ")
			mensaje := algogram.Publicar(post)
			fmt.Fprintf(os.Stdout, "%s\n", mensaje)
		case COMANDO4:
			mensaje := algogram.VerSiguientePost()
			fmt.Fprintf(os.Stdout, "%s\n", mensaje)
		default:
			/* code */
			return
		}
	}
}
