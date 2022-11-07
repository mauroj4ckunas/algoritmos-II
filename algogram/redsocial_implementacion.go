package main

import (
	diccionario "algogram/Hash"
	errores "algogram/errores"
	usuarios "algogram/usuarios"
	"fmt"
	"os"
)

type redSocial[T comparable] struct {
	actual           *usuarios.Usuario[T]
	registroUsuarios diccionario.Diccionario[string, usuarios.Usuario[T]]
	idPosteos        int
}

func sacarPrioridad(usuario1 int, usuario2 int) int {
	if usuario1 < usuario2 {
		return usuario2 - usuario1
	}
	return usuario1 - usuario2
}

func crearAlgoGram[T comparable](nombreArchivo string, compararPosteos func(T, T) int) AlgoGram[T] {
	archivoListas, err := os.Open(nombreArchivo)
	defer archivoListas.Close()
}

func (red *redSocial[T]) Login(usuario string) string {
	if red.actual == nil {
		if red.registroUsuarios.Pertenece(usuario) {
			*red.actual = red.registroUsuarios.Obtener(usuario)
			return fmt.Sprintf("Hola %s", usuario)
		} else {
			err := new(errores.ErrorUsuarioNoExiste)
			return err.Error()
		}
	} else {
		err := new(errores.ErrorUsuarioLoggeado)
		return err.Error()
	}
}

func (red *redSocial[T]) Logout() string {
	if red.actual != nil {
		red.actual = nil
		return "Adios"
	}
	return new(errores.ErrorLogout).Error()
}

func (red *redSocial[T]) Publicar(posteo []string) string {
	if red.actual != nil {
		losUsuarios := red.registroUsuarios.Iterador()
		for losUsuarios.HaySiguiente() {
			_, usuario := losUsuarios.VerActual()
			if usuario != (*red.actual) {
				usuarioActual := *red.actual
				aPublicar := usuarios.CrearPosteo(sacarPrioridad(usuarioActual.Prioridad(), usuario.Prioridad()), posteo, red.idPosteos)
				usuario.Publicar(aPublicar)
			}
		}
		red.idPosteos++
		return "Post publicado"
	}
	return new(errores.ErrorUsuarioLoggeado).Error()
}
