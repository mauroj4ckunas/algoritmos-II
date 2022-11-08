package usuarios

import (
	Heap "algogram/Heap"
)

type post struct {
	prioridadPosteo int
	posteado        []string
	id              int
}

type usuario[T comparable] struct {
	nivel int
	feed  Heap.ColaPrioridad[int]
}

var compararId = func(comp1 int, comp2 int) int {
	if comp1 < comp2 {
		return 1
	}
	return -1
}

func CrearUsuario[T comparable](prioridadUsuario int) Usuario[T] {
	usuario := new(usuario[T])
	usuario.nivel = prioridadUsuario
	usuario.feed = Heap.CrearHeap[int](compararId) //El heap sera de las posiciones de los posteos
	usuario.feed.Encolar(prioridadUsuario)
	return usuario
}

func CrearPosteo(prioridadPost int, posteo []string, id int) *post {
	post := new(post)
	post.prioridadPosteo = prioridadPost
	post.posteado = posteo
	post.id = id
	return post
}

func (usu *usuario[T]) Prioridad() int {
	return usu.nivel
}

func (usu *usuario[T]) PublicarPosteo(posteo string) {
	usu.feed.Encolar(posteo.posteado[0])
}
