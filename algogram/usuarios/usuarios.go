package usuarios

type Usuario[T comparable] interface{

	Prioridad() int

	Publicar(post[T])
}