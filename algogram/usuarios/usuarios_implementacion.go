package usuarios

import (
	Heap "algogram/Heap"
	errores "algogram/errores"
)

type post struct {
	prioridadPosteo int
	id              int
	likes           int
	posteado        string
	quienPosteo     string
}

type usuario[T comparable] struct {
	nivel int
	feed  Heap.ColaPrioridad[*post]
}

var (
	compararPosteos = func(comp1, comp2 *post) int {
		switch {
		case (*comp1).prioridadPosteo != (*comp2).prioridadPosteo:
			if (*comp1).prioridadPosteo < (*comp2).prioridadPosteo {
				return 1
			}

		case (*comp1).prioridadPosteo == (*comp2).prioridadPosteo:
			if (*comp1).id < (*comp2).id {
				return 1
			}
		}
		return -1
	}
)

func CrearUsuario[T comparable](prioridadUsuario int) Usuario[T] {
	usuario := new(usuario[T])
	usuario.nivel = prioridadUsuario
	usuario.feed = Heap.CrearHeap(compararPosteos)
	return usuario
}

func CrearPosteo(prioridadPost int, posteo string, id int) post {
	post := new(post)
	post.prioridadPosteo = prioridadPost
	post.posteado = posteo
	post.id = id
	return *post
}

func (usu *usuario[T]) Prioridad() int {
	return usu.nivel
}

func (usu *usuario[T]) PublicarPosteo(nuevoPost post) {
	usu.feed.Encolar(&nuevoPost)
}

func (usu *usuario[T]) PrimerPostDelFeed() string {
	if !usu.feed.EstaVacia() {
		posteo := usu.feed.Desencolar()
		mensaje := (*posteo).posteado

		return mensaje
	}
	err := new(errores.ErrorNoMasPost)
	return err.Error()
}
