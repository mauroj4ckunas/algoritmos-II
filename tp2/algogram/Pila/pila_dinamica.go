package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	if pila.cantidad == cap(pila.datos) {
		redimensionar(pila, pila.cantidad*2)
	}
	pila.datos[pila.cantidad] = elem
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	if pila.cantidad*4 <= cap(pila.datos) {
		redimensionar(pila, cap(pila.datos)/2)
	}
	tope := pila.VerTope()
	pila.cantidad--
	return tope
}

func CrearPilaDinamica[T any]() Pila[T] {
	const tamInicial int = 4
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, tamInicial)
	return pila
}

func redimensionar[T any](pil *pilaDinamica[T], nuevoTam int) {
	copiaPila := make([]T, nuevoTam)
	copy(copiaPila, pil.datos)
	pil.datos = copiaPila
}
