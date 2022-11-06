package main

import (
	TDAcola "algogram/Cola"
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
	usuariosRegistrados := crearRedSocial(archivoUsuarios[0])

	var logeado *usuarios.Usuario
	entradaUsuario := bufio.NewScanner(os.Stdin)
	for entradaUsuario.Scan() {
		comandos := strings.Split(entradaUsuario.Text(), " ")
		switch comandos[0] {
		case COMANDO1:
			usuario := comandos[1]
			if logeado == nil {
				if usuariosRegistrados.Pertenece(usuario) {
					logeado = usuariosRegistrados.Obtener(usuario)
					fmt.Fprintf(os.Stdout, "Hola %s\n", usuario)
				} else {
					fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				}
			} else {
				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
			}

		case COMANDO2:
			if logeado != nil {
				logeado = nil
				fmt.Fprintf(os.Stdout, "%s\n", "Adios")
			} else {
				fmt.Fprintf(os.Stdout, "%s\n", new(err.TALERROR).Error())
			}

		case COMANDO3:
			post := comandos[1:]
			if logeado != nil {
				losUsuarios := usuariosRegistrados.Iterador()
				for losUsuarios.HaySiguiente() {
					_ , usuario := losUsuarios.VerActual()
					if usuario != *logeado {
						usuario.Publicar(post,logeado.PrioridadEntre(usuario.Prioridad()))
					}
				}
				fmt.Fprintf(os.Stdout, "%s\n", "Post publicado")
			} else {
				fmt.Fprintf(os.Stdout, "%s\n", new(err.TALERROR).Error())
			}
		default:
			/* code */
			return
		}
	}
}
