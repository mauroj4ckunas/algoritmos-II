package redsocial

import (
	"algogram/diccionario"
	"algogram/errores"
	"algogram/usuarios"
	"bufio"
	"fmt"
	"os"
)

const (
	TAMAÑOINICIAL = 5
)

type redSocial struct {
	actual           *string
	registroUsuarios diccionario.Diccionario[string, usuarios.Usuario]
	publicaciones    []*usuarios.Post
}

func CrearAlgoGram(nombreArchivo string) (AlgoGram, error) {
	archivoListas, err := os.Open(nombreArchivo)

	if err != nil {
		return nil, new(errores.ErrorLecturaArchivo)
	}

	defer archivoListas.Close()

	nuevosUsuarios := bufio.NewScanner(archivoListas)
	nuevaRed := new(redSocial)
	nuevaRed.publicaciones = make([]*usuarios.Post, 0, TAMAÑOINICIAL)
	nuevaRed.registroUsuarios = diccionario.CrearHash[string, usuarios.Usuario]()
	for i := 1; nuevosUsuarios.Scan(); i++ {
		nuevo := usuarios.CrearUsuario(i)
		nuevaRed.registroUsuarios.Guardar(nuevosUsuarios.Text(), nuevo)
	}

	return nuevaRed, nil
}

func (red *redSocial) Login(usuario string) string {
	if red.actual == nil {
		if red.registroUsuarios.Pertenece(usuario) {
			red.actual = &usuario
			return fmt.Sprintf("Hola %s", usuario)
		} else {
			return new(errores.ErrorUsuarioNoExiste).Error()
		}
	} else {
		return new(errores.ErrorUsuarioLoggeado).Error()
	}
}

func (red *redSocial) Logout() string {
	if red.actual != nil {
		red.actual = nil
		return "Adios"
	}
	return new(errores.ErrorLogout).Error()
}

func (red *redSocial) Publicar(posteo string) string {
	if red.actual != nil {
		usuarioPublicando := red.registroUsuarios.Obtener(*red.actual)
		red.publicaciones = append(red.publicaciones, usuarios.CrearPosteo(usuarioPublicando.Prioridad(), posteo, len(red.publicaciones), *red.actual))
		losUsuarios := red.registroUsuarios.Iterador()

		for losUsuarios.HaySiguiente() {
			_, usuario := losUsuarios.VerActual()
			if usuario != usuarioPublicando {
				usuario.PublicarPosteo(red.publicaciones[len(red.publicaciones)-1])
			}
			losUsuarios.Siguiente()
		}
		return "Post publicado"
	}
	return new(errores.ErrorLogout).Error()
}

func (red *redSocial) VerSiguientePost() string {
	if red.actual != nil {
		usuarioActual := red.registroUsuarios.Obtener(*red.actual)
		mensaje := usuarioActual.PrimerPostDelFeed()
		return mensaje
	}
	return new(errores.ErrorNoMasPost).Error()
}

func (red *redSocial) LikearPost(iD int) string {
	if red.actual != nil {
		if iD < len(red.publicaciones) {
			red.publicaciones[iD].Likes.Guardar(*red.actual, true)
			return "Post likeado"
		}
	}
	return new(errores.ErrorLikeoPostInexistente).Error()
}

func (red *redSocial) ImprimirLikesPost(iD int) {
	if iD < len(red.publicaciones) {
		if red.publicaciones[iD].Likes.Cantidad() > 0 {
			fmt.Printf("El post tiene %d likes:\n", red.publicaciones[iD].Likes.Cantidad())
			losLikes := red.publicaciones[iD].Likes.Iterador()
			for losLikes.HaySiguiente() {
				usuario, _ := losLikes.VerActual()
				fmt.Printf("\t%s\n", usuario)
				losLikes.Siguiente()
			}
			return
		}
	}
	fmt.Printf("%s\n", new(errores.ErrorVerLikes).Error())
}
