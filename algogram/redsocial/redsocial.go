package redsocial

type AlgoGram[T comparable] interface {
	Login(usuario string) string

	Logout() string

	Publicar(posteo T)
}
