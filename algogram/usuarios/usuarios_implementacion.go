package usuarios

import (
	Heap "algogram/Heap"
)

type Post struct {
	prioridadPosteo int
	posteado        []string
	id              int
}

type usuarioImplementacion struct {
	nivel int
	feed  Heap.ColaPrioridad[Post]
}

func CrearPosteo(prioridadPost int, posteo []string, id int) *Post {
	post := new(Post)
	post.prioridadPosteo = prioridadPost
	post.posteado = posteo
	post.id = id
	return post
}

func CrearUsuario(prioridadUsuario int) Usuario {
	usuario := new(usuarioImplementacion)
	usuario.nivel = prioridadUsuario

	funcionCompararPosteos := func (comp1, comp2 Post) int {
		if comp1.prioridadPosteo < comp2.prioridadPosteo {
			return 1
		}else if comp1.prioridadPosteo == comp2.prioridadPosteo{
			if comp1.id < comp2.id {
				return 1
			}
		}
		return -1
	}

	usuario.feed = Heap.CrearHeap[Post](funcionCompararPosteos)
	return usuario
}

func (usu *usuarioImplementacion) Prioridad() int {
	return usu.nivel
}

func (usu *usuarioImplementacion) Publicar(posteo Post) {
	usu.feed.Encolar(posteo)
}

func (usu *usuarioImplementacion) afinidadCon(otro Usuario) int{
	if usu.Prioridad() < otro.Prioridad() {
		return otro.Prioridad() - usu.Prioridad()
	}
	return usu.Prioridad() - otro.Prioridad()
}