package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	usuarios "algogram/usuarios"
)

const (
	COMANDO1 = "login"
	COMANDO2 = "logout"
	COMANDO3 = "publicar"
	funcionCompararPosteos = func (comp1, comp2 usuarios.post[V]) int {
		if comp1 < comp2 {
			return 1
		}
		return -1
	}
)

func main() {
	archivoUsuarios := os.Args[1:]
	Algogram := crearAlgoGram[string, int](archivoUsuarios[0],funcionCompararPosteos)

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
