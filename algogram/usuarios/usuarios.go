package usuarios

type Usuario[T comparable] interface {
	Prioridad() int

	PublicarPosteo()
}
