package red

type RedSocial interface {
	Login(string) string
	Logout() string
	Publicar(string) string
	VerSiguientePost() string
	LikearPost(int) string
	ImprimirLikesPost(int)
}
