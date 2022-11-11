package usuarios

import (
	Heap "algogram/Heap"
	diccionario "algogram/diccionario"
	errores "algogram/errores"
	"fmt"
)

type Post struct {
	prioridadPosteo int
	id              int
	posteado        string
	likes           diccionario.DiccionarioOrdenado[string, bool]
	publicador      string
}

type usuario struct {
	nivel int
	feed  Heap.ColaPrioridad[*Post]
}

func CrearUsuario(prioridadUsuario int) Usuario {
	usuario := new(usuario)
	usuario.nivel = prioridadUsuario

	compararPosteos := func(comp1, comp2 *Post) int {
		prioridad := &usuario.nivel
		prioridadPost1 := comp1.prioridadPosteo - *prioridad
		prioridadPost2 := comp2.prioridadPosteo - *prioridad

		if prioridadPost1 < 0 {
			prioridadPost1 *= -1
		}
		if prioridadPost2 < 0 {
			prioridadPost2 *= -1
		}

		if prioridadPost1 < prioridadPost2 {
			return 1
		} else if prioridadPost1 == prioridadPost2 {
			if comp1.id < comp2.id {
				return 1
			}
		}
		return -1
	}

	usuario.feed = Heap.CrearHeap(compararPosteos)
	return usuario
}

func CrearPosteo(prioridadPost int, posteo string, id int, usuario string) *Post {
	post := new(Post)
	post.prioridadPosteo = prioridadPost
	post.publicador = usuario
	post.posteado = posteo
	post.id = id
	ordenarLikes := func(nombre1, nombre2 string) int {
		if nombre1 < nombre2 {
			return -1
		} else if nombre1 == nombre2 {
			return 0
		}
		return 1
	}
	post.likes = diccionario.CrearABB[string, bool](ordenarLikes)
	return post
}

func (usu *usuario) Prioridad() int {
	return usu.nivel
}

func (usu *usuario) PublicarPosteo(nuevoPost *Post) {
	usu.feed.Encolar(nuevoPost)
}

func (usu *usuario) PrimerPostDelFeed() string {
	if !usu.feed.EstaVacia() {
		posteo := usu.feed.Desencolar()
		mensaje := fmt.Sprintf("Post ID %d\n%s dijo: %s\nLikes: %d", posteo.id, posteo.publicador, posteo.posteado, posteo.likes.Cantidad())

		return mensaje
	}
	return new(errores.ErrorNoMasPost).Error()
}
