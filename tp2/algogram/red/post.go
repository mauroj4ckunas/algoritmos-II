package red

type Post interface {
	PrioridadDelAutor() int

	ImprimirMensaje() string

	AgregarLike(string)

	VerTodosLosLikes()

	VerIDPost() int
}
