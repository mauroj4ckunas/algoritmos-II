package red

type Usuario interface {
	PublicarPosteo(*Post)

	Prioridad() int

	PrimerPostDelFeed() (*Post, string)
}
