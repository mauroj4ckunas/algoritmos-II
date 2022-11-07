package usuarios

import(
	Heap "algogram/Heap"
)

type post[V any] struct {
	prioridadPosteo	V 
	posteado		[]string
	id				int
}

type usuarioImplementacion[T comparable,V any] struct {
	nivel		V 
	feed		Heap.ColaPrioridad[T]
}

func CrearPosteo[V any](prioridadPost V ,posteo []string, id int) *post[V] { 
	post := new(post[V])
	post.prioridadPosteo = prioridadPost
	post.posteado = posteo
	post.id = id
	return post
}

func CrearUsuario[T comparable,V any](prioridadUsuario V, comparadorPosteos func(T,T)int) Usuario[T,V] {
	usuario := new(usuarioImplementacion[T,V])
	usuario.nivel = prioridadUsuario
	usuario.feed = Heap.CrearHeap[T](comparadorPosteos)
	return usuario
}

func (usu *usuarioImplementacion[T,V]) Prioridad() V {
	return usu.nivel
}

func (usu *usuarioImplementacion[T,V]) Publicar(posteo T) {
	usu.feed.Encolar(posteo)
}