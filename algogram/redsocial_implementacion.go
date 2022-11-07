package main

import (
	diccionario "algogram/Hash"
	err "algogram/errores"
	usuarios "algogram/usuarios"
	"fmt"
	"os"
)

type redSocial[T comparable, V comparable] struct {
	actual            *usuarios.Usuario[T]
	registroUsuarios  diccionario.Diccionario[string, usuarios.Usuario[T]]
	idPosteos         int
	calcularPrioridad func(V, V) int
}

func compararPost[T int | string](comp1, comp2 T) int {
	if comp1 < comp2 {
		return 1
	}
	return -1
}

// funcionCompararUsuarios := func (prioridad1, prioridad2 int) int {
// 	if prioridad1 < prioridad2 {
// 		return prioridad2 - prioridad1
// 	}
// 	return prioridad1 - prioridad2
// }

func crearAlgoGram[T comparable, V comparable](nombreArchivo string) AlgoGram[T] {
	archivoListas, err := os.Open(ruta)
	defer archivoListas.Close()
}

func (red *redSocial[T, V]) Login(usuario string) string {
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

func (red *redSocial[T, V]) Logout() string {
	if red.actual != nil {
		red.actual = nil
		return "Adios"
	}
	return new(err.ErrorLogout).Error()
}

func (red *redSocial[T, V]) Publicar(posteo []T) string {
	if red.actual != nil {
		losUsuarios := red.registroUsuarios.Iterador()
		usuarios.CrearPosteo[T](prioridadPost, posteo, id)
		for losUsuarios.HaySiguiente() {
			_, usuario := losUsuarios.VerActual()
			if usuario != (*red.actual) {
				usuarioActual := *red.actual
				aPublicar := usuarios.CrearPosteo[T](red.calcularPrioridad(usuarioActual.Prioridad(), usuario.Prioridad()), posteo, red.idPosteos)
				usuario.Publicar(aPublicar)
			}
		}
		red.idPosteos++
		return "Post publicado"
	}
	return new(err.TALERROR).Error()
}
