package usuarios

type Usuario[T comparable] interface{

	Prioridad() int

	PrioridadEntre(otroUsuario int) int

	Publicar(posteo []T, prioridadPost int)
}