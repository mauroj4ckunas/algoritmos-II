package red

import (
	"algogram/diccionario"
	errores "algogram/errores"
	"bufio"
	"fmt"
	"os"
)

type redSocial struct {
	actual           *string
	registroUsuarios diccionario.Diccionario[string, Usuario]
	publicaciones    diccionario.Diccionario[int, Post]
}

func CrearAlgoGram(nombreArchivo string) (RedSocial, error) {
	archivoListas, err := os.Open(nombreArchivo)

	if err != nil {
		return nil, new(errores.ErrorLecturaArchivo)
	}

	defer archivoListas.Close()

	nuevosUsuarios := bufio.NewScanner(archivoListas)
	nuevaRed := new(redSocial)
	nuevaRed.publicaciones = diccionario.CrearHash[int, Post]()
	nuevaRed.registroUsuarios = diccionario.CrearHash[string, Usuario]()
	for i := 1; nuevosUsuarios.Scan(); i++ {
		nuevo := CrearUsuario(nuevosUsuarios.Text(), i)
		nuevaRed.registroUsuarios.Guardar(nuevosUsuarios.Text(), nuevo)
	}
	return nuevaRed, nil
}

func (sesion *redSocial) Login(usuario string) string {
	if sesion.actual == nil {
		if sesion.registroUsuarios.Pertenece(usuario) {
			sesion.actual = &usuario
			return fmt.Sprintf("Hola %s", usuario)
		} else {
			return new(errores.ErrorUsuarioNoExiste).Error()
		}
	} else {
		return new(errores.ErrorUsuarioLoggeado).Error()
	}
}

func (sesion *redSocial) Logout() string {
	if sesion.actual != nil {
		sesion.actual = nil
		return "Adios"
	}
	return new(errores.ErrorLogout).Error()
}

func (sesion *redSocial) Publicar(posteo string) string {
	if sesion.actual != nil {
		publicador := sesion.registroUsuarios.Obtener(*sesion.actual)
		postNuevo := CrearPosteo(publicador.Prioridad(), posteo, sesion.publicaciones.Cantidad(), *sesion.actual)
		sesion.publicaciones.Guardar(sesion.publicaciones.Cantidad(), postNuevo)

		losUsuarios := sesion.registroUsuarios.Iterador()

		for losUsuarios.HaySiguiente() {
			_, usuario := losUsuarios.VerActual()
			if usuario != publicador {
				usuario.PublicarPosteo(&postNuevo)
			}
			losUsuarios.Siguiente()
		}
		return "Post publicado"
	}
	return new(errores.ErrorLogout).Error()
}

func (sesion *redSocial) VerSiguientePost() string {
	if sesion.actual != nil {
		usuarioActual := sesion.registroUsuarios.Obtener(*sesion.actual)
		posteo, err := usuarioActual.PrimerPostDelFeed()
		if posteo != nil {
			return (*posteo).ImprimirMensaje()
		}
		return err
	}
	return new(errores.ErrorNoMasPost).Error()

}

func (sesion *redSocial) LikearPost(iD int) string {
	if sesion.actual != nil {
		if iD < sesion.publicaciones.Cantidad() {
			postLikeado := sesion.publicaciones.Obtener(iD)
			postLikeado.AgregarLike(*sesion.actual)
			return "Post likeado"
		}
	}
	return new(errores.ErrorLikeoPostInexistente).Error()
}

func (sesion *redSocial) ImprimirLikesPost(iD int) {
	if iD < sesion.publicaciones.Cantidad() {
		post := sesion.publicaciones.Obtener(iD)
		post.VerTodosLosLikes()
		return
	}
	fmt.Printf("%s\n", new(errores.ErrorVerLikes).Error())
}
