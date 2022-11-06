package main

import (
	"fmt"
	"os"
)

type redSocial[T comparable, V comparable] struct {
	actual            *Usuario[T]
	registroUsuarios  Diccionario[string, Usuario[T]]
	idPosteos         int
	calcularPrioridad func(V, V) int
}

var (
	funcionCompararPost = func(prioridad1, prioridad2 int) int {
		if prioridad1 < prioridad2 {
			return 1
		}
		return -1
	}

	funcionPrioridadEntreUsuarios = func(prioridad1, prioridad2 int) int {
		if prioridad1 < prioridad2 {
			return prioridad2 - prioridad1
		}
		return prioridad1 - prioridad2
	}
)

func crearAlgoGram[T comparable, V comparable](nombreArchivo string) AlgoGram[T] {
	archivoListas, err := os.Open(ruta)
	defer archivoListas.Close()
}

func (red *redSocial[T, V]) Login(usuario string) string {
	if red.actual == nil {
		if red.registroUsuarios.Pertenece(usuario) {
			red.actual = usuariosRegistrados.Obtener(usuario)
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
	return new(err.TALERROR).Error()
}

func (red *redSocial[T, v]) Publicar(posteo []T) string {
	if red.actual != nil {
		losUsuarios := red.registroUsuarios.Iterador()
		CrearPosteo[T](prioridadPost, posteo, id)
		for losUsuarios.HaySiguiente() {
			_, usuario := losUsuarios.VerActual()
			if usuario != (*red.actual) {
				usuario.Publicar(CrearPosteo[T](red.calcularPrioridad(red.actual.Prioridad(), usuario.Prioridad()), post, red.idPosteos))
			}
		}
		red.idPosteos++
		return "Post publicado"
	}
	return new(err.TALERROR).Error()
}
