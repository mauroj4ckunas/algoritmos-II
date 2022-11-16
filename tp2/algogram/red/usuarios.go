package red

type Usuario interface {
	Prioridad() int

	PublicarPosteo(nuevoPost *post)

	PrimerPostDelFeed() string
}
