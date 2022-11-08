package usuarios

type Usuario interface {
	Prioridad() int

	Publicar(posteo Post)

	afinidadCon(otro Usuario) int
}
