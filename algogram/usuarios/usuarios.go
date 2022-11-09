package usuarios

type Usuario[T comparable] interface {
	Prioridad() int

	PublicarPosteo(nuevoPost post)

	PrimerPostDelFeed() string
}
