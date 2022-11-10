package usuarios

type Usuario interface {
	Prioridad() int

	PublicarPosteo(nuevoPost *Post)

	PrimerPostDelFeed() string
}
