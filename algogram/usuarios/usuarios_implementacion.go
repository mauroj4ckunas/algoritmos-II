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
	return *usuario
}

func CrearPosteo(prioridadPost int, posteo []string, id int) *post {
	post := new(post)
	post.prioridadPosteo = prioridadPost
	post.posteado = posteo
	post.id = id
	return post
}

// func (usu *usuario[T]) Prioridad() int {
// 	return usu.nivel
// }

// func (usu *usuario[T]) Publicar(posteo post) {
// 	usu.feed.Encolar(posteo.posteado[0])
// }

// func (red *redSocial[T]) Login(usuario string) string {
// 	if red.actual == nil {
// 		if red.registroUsuarios.Pertenece(usuario) {
// 			*red.actual = red.registroUsuarios.Obtener(usuario)
// 			return fmt.Sprintf("Hola %s", usuario)
// 		} else {
// 			err := new(errores.ErrorUsuarioNoExiste)
// 			return err.Error()
// 		}
// 	} else {
// 		err := new(errores.ErrorUsuarioLoggeado)
// 		return err.Error()
// 	}
// }

// func (red *redSocial[T]) Logout() string {
// 	if red.actual != nil {
// 		red.actual = nil
// 		return "Adios"
// 	}
// 	return new(errores.ErrorLogout).Error()
// }

// func (red *redSocial[T]) Publicar(posteo []string) string {
// 	if red.actual != nil {
// 		losUsuarios := red.registroUsuarios.Iterador()
// 		for losUsuarios.HaySiguiente() {
// 			_, usuario := losUsuarios.VerActual()
// 			if usuario != (*red.actual) {
// 				usuarioActual := *red.actual
// 				aPublicar := usuarios.CrearPosteo(sacarPrioridad(usuarioActual.Prioridad(), usuario.Prioridad()), posteo, red.idPosteos)
// 				usuario.Publicar(aPublicar.id)
// 			}
// 		}
// 		red.idPosteos++
// 		return "Post publicado"
// 	}
// 	return new(errores.ErrorUsuarioLoggeado).Error()
// }
