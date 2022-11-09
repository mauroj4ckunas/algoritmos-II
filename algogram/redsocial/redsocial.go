package redsocial

type AlgoGram[T comparable] interface {
	Login(usuario string) (string, error)

	Logout() string

	Publicar(posteo T)

	VerSiguientePost() string
}
