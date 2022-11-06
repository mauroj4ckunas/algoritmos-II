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
	TDAAlgogram := crearRedSocial(archivoUsuarios[0])

	logeado := TDAcola.CrearColaEnlazada[]()
	entradaUsuario := bufio.NewScanner(os.Stdin)
	for entradaUsuario.Scan() {
		comandos := strings.Split(entradaUsuario.Text(), " ")
		switch comandos[0] {
		case COMANDO1:
			usuario := comandos[1]
			if logeado.EstaVacia() {
				err, elUsuario := TDAAlgogram.Login(usuario)
				if err != nil {
					fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				}
				logeado.Encolar(elUsuario)
			} else {
				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
			}

		case COMANDO2:
			if !logeado.EstaVacia() {
				logeado.VerTope().Logout()
				logeado.Desencolar()
			} else {
				fmt.Fprintf(os.Stdout, "%s\n", new(err.TALERROR).Error())
			}

		case COMANDO3:
			post := comandos[1:]
			if !logeado.EstaVacia() {
				losUsuarios := TDAAlgogram.Iterador()
				for losUsuarios.HaySiguiente() {
					_, usuario := losUsuarios.VerActual()
					usuario.Publicar(post)
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
