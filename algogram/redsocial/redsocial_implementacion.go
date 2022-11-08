package redsocial

import (
	diccionario "algogram/Hash"
	errores "algogram/errores"
	usuarios "algogram/usuarios"
	"bufio"
	"fmt"
	"os"
)

type redSocial[T comparable] struct {
	actual           *usuarios.Usuario[T]
	registroUsuarios diccionario.Diccionario[string, usuarios.Usuario[T]]
	idPosteos        int
}

func CrearAlgoGram[T comparable](nombreArchivo string) (*redSocial[T], error) {
	archivoListas, err := os.Open(nombreArchivo)
	defer archivoListas.Close()

	if err != nil {
		err = new(errores.ErrorLecturaArchivo)
		return nil, err
	}

	nuevos := bufio.NewScanner(archivoListas)
	nuevaRed := new(redSocial[T])
	nuevaRed.agregarRegistroUsuarios(nuevos)

	return nuevaRed, nil
}

func (red *redSocial[T]) agregarRegistroUsuarios(nuevosUsuarios *bufio.Scanner) {
	registro := diccionario.CrearHash[string, usuarios.Usuario[T]]()
	pos := 1
	for nuevosUsuarios.Scan() {
		newUser := usuarios.CrearUsuario[T](pos)
		registro.Guardar(nuevosUsuarios.Text(), newUser)
		pos++
	}
	red.registroUsuarios = registro
}

func (red *redSocial[T]) Login(usuario string) string {
	if red.actual == nil {
		if red.registroUsuarios.Pertenece(usuario) {
			*red.actual = red.registroUsuarios.Obtener(usuario)
			return fmt.Sprintf("Hola %s", usuario)
		} else {
			return new(errores.ErrorUsuarioNoExiste).Error()
		}
	} else {
		return new(errores.ErrorUsuarioLoggeado).Error()
	}
}

func (red *redSocial[T]) Logout() string {
	if red.actual != nil {
		red.actual = nil
		return "Adios"
	}
	return new(errores.ErrorLogout).Error()
}

func (red *redSocial[T]) Publicar(posteo []string) string {
	if red.actual != nil {
		losUsuarios := red.registroUsuarios.Iterador()
		for losUsuarios.HaySiguiente() {
			_, usuario := losUsuarios.VerActual()
			if usuario != (*red.actual) {
				usuarioActual := *red.actual
				aPublicar := usuarios.CrearPosteo(sacarPrioridad(usuarioActual.Prioridad(), usuario.Prioridad()), posteo, red.idPosteos)
				usuario.PublicarPosteo()
			}
		}
		red.idPosteos++
		return "Post publicado"
	}
	return new(errores.ErrorUsuarioLoggeado).Error()
}

func sacarPrioridad(usuario1 int, usuario2 int) int {
	if usuario1 < usuario2 {
		return usuario2 - usuario1
	}
	return usuario1 - usuario2
}

// type redSocial[T comparable] struct {
// 	actual           Usuario[T]
// 	registroUsuarios diccionario.Diccionario[string, Usuario[T]]
// 	idPosteos        int
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
