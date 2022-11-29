package errores

type ErrorLecturaArchivo struct{}

func (e ErrorLecturaArchivo) Error() string {
	return "Error: Lectura de archivos"
}

type ErrorUsuarioLoggeado struct{}

func (e ErrorUsuarioLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioNoExiste struct{}

func (e ErrorUsuarioNoExiste) Error() string {
	return "Error: usuario no existente"
}

type ErrorLogout struct{}

func (e ErrorLogout) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorNoMasPost struct{}

func (e ErrorNoMasPost) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorLikeoPostInexistente struct{}

func (e ErrorLikeoPostInexistente) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorVerLikes struct{}

func (e ErrorVerLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}