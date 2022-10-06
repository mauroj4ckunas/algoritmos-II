package cola

/* Definición del struct cola proporcionado por la cátedra. */

type colaEnlazada[T any] struct {
	primero *Nodo[T] //precondicion: si primero es nil
	ultimo  *Nodo[T] // ultimo tambien debe ser nil

}

type Nodo[T any] struct {
	dato T
	prox *Nodo[T]
}

func crearNodo[T any](elem T) *Nodo[T] {
	n := new(Nodo[T])
	n.dato = elem
	return n
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(colaEnlazada[T])
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(elem T) {
	nodoNuevo := crearNodo[T](elem)

	if cola.primero == nil {

		cola.primero = nodoNuevo
	} else {

		cola.ultimo.prox = nodoNuevo
	}

	cola.ultimo = nodoNuevo
}

func (cola *colaEnlazada[T]) Desencolar() T {

	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	primero_encolado := (*cola.primero).dato
	cola.primero = cola.primero.prox

	if cola.primero == nil {

		cola.ultimo = nil
	}
	return primero_encolado
}
