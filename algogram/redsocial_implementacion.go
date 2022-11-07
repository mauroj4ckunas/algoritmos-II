package main

import (
	diccionario "algogram/Hash"
	err "algogram/errores"
	usuarios "algogram/usuarios"
	"fmt"
	"os"
)

func sacarPrioridad[V any](usuario1 V,usuario2 V) V {
	if usuario1 < usuario2{
		return usuario2 - usuario1
	}
	return usuario1 - usuario2
}


type redSocial[T comparable,V any] struct {
	actual            *usuarios.Usuario[T,V]
	registroUsuarios  diccionario.Diccionario[string, usuarios.Usuario[T,V]]
	idPosteos         int
}

func crearAlgoGram[T comparable,V any](nombreArchivo string,compararPosteos func(T,T)int) AlgoGram[T,V] {
	archivoListas, err := os.Open(ruta)
	defer archivoListas.Close()
}

func (red *redSocial[T, V]) Login(usuario string) string {
	if red.actual == nil {
		if red.registroUsuarios.Pertenece(usuario) {
			*red.actual = red.registroUsuarios.Obtener(usuario)
			return fmt.Sprintf("Hola %s", usuario)
		} else {
			return err.Error()
		}
	} else {
		return err.Error()
	}
}

func (red *redSocial[T, V]) Logout() string {
	if red.actual != nil {
		red.actual = nil
		return "Adios"
	}
	return new(err.ErrorLogout).Error()
}

func (red *redSocial[T, V]) Publicar(posteo []string) string {
	if red.actual != nil {
		losUsuarios := red.registroUsuarios.Iterador()
		for losUsuarios.HaySiguiente() {
			_, usuario := losUsuarios.VerActual()
			if usuario != (*red.actual) {
				usuarioActual := *red.actual
				aPublicar := usuarios.CrearPosteo[V](sacarPrioridad(usuarioPublicando.Prioridad(), usuario.Prioridad()), posteo, red.idPosteos)
				usuario.Publicar(aPublicar)
			}
		}
		red.idPosteos++
		return "Post publicado"
	}
	return new(err.TALERROR).Error()
}
