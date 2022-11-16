package red

type AlgoGram interface {
	Login(usuario string) string

	Logout() string

	Publicar(posteo string) string

	VerSiguientePost() string

	LikearPost(iD int) string

	ImprimirLikesPost(iD int)
}
