package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil && cola.ultimo == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func crearNodo[T any](elem T) *nodoCola[T] {
	nuevoNodo := new(nodoCola[T])
	nuevoNodo.dato = elem
	return nuevoNodo
}

func (cola *colaEnlazada[T]) Encolar(elem T) {
	nodo := crearNodo(elem)
	if cola.EstaVacia() {
		cola.primero = nodo
	} else {
		cola.ultimo.prox = nodo
	}
	cola.ultimo = nodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	datoDevuelto := cola.VerPrimero()
	if cola.primero == cola.ultimo {
		cola.primero = nil
		cola.ultimo = nil
	} else {
		cola.primero = cola.primero.prox
	}
	return datoDevuelto
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}
