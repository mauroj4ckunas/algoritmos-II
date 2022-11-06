package main

import (
	"os"
)
const(
	funcionCompararPost = func (prioridad1,prioridad2 int) int {
		if prioridad1 < prioridad2 {
			return 1
		}
		return -1
	}
)
func crearAlgoGram[T comparable,V comparable](nombreArchivo string,calcularPrioridad func(V, V) int) AlgoGram[T] {
	archivoListas, err := os.Open(ruta)
	defer archivoListas.Close()
}

type redSocial[T comparable,V comparable] struct {
	actual 				*Usuario[T]
	registroUsuarios	Diccionario[string, Usuario[T]]
	idPosteos			int
	calcularPrioridad	func(V, V) int
}

func (red *redSocial[T]) Login(usuario string) string {
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

func (red *redSocial[T]) Logout() string {
	if red.actual != nil {
		red.actual = nil
		return "Adios"
	}
	return new(err.TALERROR).Error()
}

func (red *redSocial[T]) Publicar(posteo []T) string {
	if red.actual != nil {
		losUsuarios := red.registroUsuarios.Iterador()
		CrearPosteo[T](prioridadPost int,posteo []T,id)
		for losUsuarios.HaySiguiente() {
			_ , usuario := losUsuarios.VerActual()
			if usuario != (*red.actual) {
				usuario.Publicar(CrearPosteo[T](red.calcularPrioridad(red.actual.Prioridad(),usuario.Prioridad()),post,red.idPosteos))
			}
		}
		red.idPosteos++
		return "Post publicado"
	}
	return new(err.TALERROR).Error()
}