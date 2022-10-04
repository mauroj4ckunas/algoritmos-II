package cola

type Nodo[T any] struct {
	dato T
	prox *Nodo[T]
}

func crearNodo[T any](elem T) *Nodo[T] {
	n := new(Nodo[T])
	n.dato = elem
	return n
}
