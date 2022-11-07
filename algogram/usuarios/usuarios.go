package usuarios

type Usuario[T comparable,V any] interface{

	Prioridad() V

	Publicar(posteo T)
}