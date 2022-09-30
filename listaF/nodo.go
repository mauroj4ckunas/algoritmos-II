package lista

type nodo[T any] struct {
	dato T
	prox *nodo[T]
}

func crearNodo[T any](elem T) *nodo[T] {
	nodoNuevo := new(nodo[T])
	nodoNuevo.dato = elem
	return nodoNuevo
}
