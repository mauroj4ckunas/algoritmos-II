package main

import (
	diccionario "algogram/Hash"
	err "algogram/errores"
	usuarios "algogram/usuarios"
	"fmt"
	"os"
)

type redSocial struct {
	actual           *string
	registroUsuarios diccionario.Diccionario[string, usuarios.Usuario]
	idPosteos        int
}

func crearAlgoGram(nombreArchivo string) AlgoGram {
	archivoListas, err := os.Open(ruta)
	defer archivoListas.Close()
}

func (red *redSocial) Login(usuario string) string {
	if red.actual == nil {
		if red.registroUsuarios.Pertenece(usuario) {
			red.actual = &usuario
			return fmt.Sprintf("Hola %s", usuario)
		} else {
			return err.Error()
		}
	} else {
		return err.Error()
	}
}

func (red *redSocial) Logout() string {
	if red.actual != nil {
		red.actual = nil
		return "Adios"
	}
	return new(err.ErrorLogout).Error()
}

func (red *redSocial) Publicar(posteo []string) string {
	if red.actual != nil {
		losUsuarios := red.registroUsuarios.Iterador()
		for losUsuarios.HaySiguiente() {
			_, usuario := losUsuarios.VerActual()
			if usuario != red.registroUsuarios.Obtener(*red.actual) {
				usuarioActual := red.registroUsuarios.Obtener(*red.actual)
				aPublicar := usuarios.CrearPosteo(usuarioPublicando.afinidadCon(usuario), posteo, red.idPosteos)
				usuario.Publicar(aPublicar)
			}
		}
		red.idPosteos++
		return "Post publicado"
	}
	return new(err.TALERROR).Error()
}
