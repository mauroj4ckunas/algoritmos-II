package errores

type ErrorUsuarioLoggeado struct{}

func (e ErrorUsuarioLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioNoExiste struct{}

func (e ErrorUsuarioNoExiste) Error() string {
	return "ERROR: usuario no existente"
}

type ErrorLogout struct{}

func (e ErrorLogout) Error() string {
	return "ERROR: no habia usuario loggeado"
}

type ErrorNoMasPost struct{}

func (e ErrorNoMasPost) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorPostInexistente struct{}

func (e ErrorPostInexistente) Error() string {
	return "Error: Post inexistente o sin likes"
}