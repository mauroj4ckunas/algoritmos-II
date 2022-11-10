package redsocial

type AlgoGram interface {
	Login(usuario string) string

	Logout() string

	Publicar(posteo string) string

	VerSiguientePost() string
}
