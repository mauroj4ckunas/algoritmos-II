package main

import (
	diccionario "algogram/Hash"
	err "algogram/errores"
	usuarios "algogram/usuarios"
	"fmt"
	"os"
)

func sacarPrioridad(usuario1 int, usuario2 int) int {
	if usuario1 < usuario2 {
		return usuario2 - usuario1
	}
	return usuario1 - usuario2
}

type redSocial[T comparable] struct {
	actual           *usuarios.Usuario[T]
	registroUsuarios diccionario.Diccionario[string, usuarios.Usuario[T]]
	idPosteos        int
}

func crearAlgoGram[T comparable](nombreArchivo string, compararPosteos func(T, T) int) AlgoGram[T] {
	archivoListas, err := os.Open(ruta)
	defer archivoListas.Close()
}

func (red *redSocial[T]) Login(usuario string) string {
	if red.actual == nil {
		if red.registroUsuarios.Pertenece(usuario) {
			*red.actual = red.registroUsuarios.Obtener(usuario)
			return fmt.Sprintf("Hola %s", usuario)
		} else {
			return err.Error()
		}
	} else {
		return err.Error()
	}
}

func (red *redSocial[T]) Logout() string {
	if red.actual != nil {
		red.actual = nil
		return "Adios"
	}
	return new(err.ErrorLogout).Error()
}

func (red *redSocial[T]) Publicar(posteo []string) string {
	if red.actual != nil {
		losUsuarios := red.registroUsuarios.Iterador()
		for losUsuarios.HaySiguiente() {
			_, usuario := losUsuarios.VerActual()
			if usuario != (*red.actual) {
				usuarioActual := *red.actual
				aPublicar := usuarios.CrearPosteo(sacarPrioridad(usuarioPublicando.Prioridad(), usuario.Prioridad()), posteo, red.idPosteos)
				usuario.Publicar(aPublicar)
			}
		}
		red.idPosteos++
		return "Post publicado"
	}
	return new(err.TALERROR).Error()
}
