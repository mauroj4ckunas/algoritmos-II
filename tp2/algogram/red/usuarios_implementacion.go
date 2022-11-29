package red

import (
	Heap "algogram/Heap"
	errores "algogram/errores"
)

type usuario struct {
	nombre string
	nivel  int
	feed   Heap.ColaPrioridad[*Post]
}

func crearFuncionComparativa(prioridad int) func(*Post, *Post) int {

	ptrPriori := &prioridad

	var compararPosteos = func(comp1, comp2 *Post) int {
		prioridadPost1 := (*comp1).PrioridadDelAutor() - *ptrPriori
		prioridadPost2 := (*comp2).PrioridadDelAutor() - *ptrPriori

		if prioridadPost1 < 0 {
			prioridadPost1 *= -1
		}
		if prioridadPost2 < 0 {
			prioridadPost2 *= -1
		}

		if prioridadPost1 < prioridadPost2 {
			return 1
		} else if prioridadPost1 == prioridadPost2 {
			if (*comp1).VerIDPost() < (*comp2).VerIDPost() {
				return 1
			}
		}
		return -1
	}

	return compararPosteos
}

func CrearUsuario(nombreUsuario string, prioridadUsuario int) Usuario {
	usuario := new(usuario)
	usuario.nombre = nombreUsuario
	usuario.nivel = prioridadUsuario

	usuario.feed = Heap.CrearHeap(crearFuncionComparativa(prioridadUsuario))

	return usuario
}

func (usu *usuario) PublicarPosteo(nuevoPost *Post) {
	usu.feed.Encolar(nuevoPost)
}

func (usu *usuario) Prioridad() int {
	return usu.nivel
}

func (usu *usuario) PrimerPostDelFeed() (*Post, string) {
	if !usu.feed.EstaVacia() {
		posteo := usu.feed.Desencolar()
		return posteo, ""
	}
	return nil, new(errores.ErrorNoMasPost).Error()
}
