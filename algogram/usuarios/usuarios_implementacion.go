package usuarios

import (
	Heap "algogram/Heap"
)

type post struct {
	prioridadPosteo int
	posteado        []string
	id              int
}

type usuarioImplementacion[T comparable] struct {
	nivel int
	feed  Heap.ColaPrioridad[T]
}

func CrearPosteo(prioridadPost int, posteo []string, id int) *post {
	post := new(post)
	post.prioridadPosteo = prioridadPost
	post.posteado = posteo
	post.id = id
	return post
}

func CrearUsuario[T comparable](prioridadUsuario int, comparadorPosteos func(T, T) int) Usuario[T] {
	usuario := new(usuarioImplementacion[T])
	usuario.nivel = prioridadUsuario
	usuario.feed = Heap.CrearHeap[T](comparadorPosteos)
	return usuario
}

func (usu *usuarioImplementacion[T]) Prioridad() int {
	return usu.nivel
}

func (usu *usuarioImplementacion[T]) Publicar(posteo T) {
	usu.feed.Encolar(posteo)
}
