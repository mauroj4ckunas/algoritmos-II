package red

import (
	"algogram/diccionario"
	errores "algogram/errores"
	"fmt"
)

type post struct {
	prioriDelAutor int
	id             int
	publicacion    string
	likes          diccionario.DiccionarioOrdenado[string, bool]
	autor          string
}

var ordenarLikes = func(nombre1, nombre2 string) int {
	if nombre1 < nombre2 {
		return -1
	} else if nombre1 == nombre2 {
		return 0
	}
	return 1
}

func CrearPosteo(prioridadPost int, posteo string, id int, usuario string) Post {
	post := new(post)
	post.prioriDelAutor = prioridadPost
	post.id = id
	post.autor = usuario
	post.publicacion = posteo
	post.likes = diccionario.CrearABB[string, bool](ordenarLikes)
	return post
}

func (pst *post) ImprimirMensaje() string {
	return fmt.Sprintf("Post ID %d\n%s dijo: %s\nLikes: %d", pst.id, pst.autor, pst.publicacion, pst.likes.Cantidad())
}

func (pst *post) PrioridadDelAutor() int {
	return pst.prioriDelAutor
}

func (pst *post) AgregarLike(usuario string) {
	pst.likes.Guardar(usuario, true)
}

func (pst *post) VerIDPost() int {
	return pst.id
}

func (pst *post) VerTodosLosLikes() {
	if pst.likes.Cantidad() > 0 {
		fmt.Printf("El post tiene %d likes:\n", pst.likes.Cantidad())
		losLikes := pst.likes.Iterador()
		for losLikes.HaySiguiente() {
			usuario, _ := losLikes.VerActual()
			fmt.Printf("\t%s\n", usuario)
			losLikes.Siguiente()
		}
		return
	}
	fmt.Printf("%s\n", new(errores.ErrorVerLikes).Error())
}
