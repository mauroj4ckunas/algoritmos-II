package cola_prioridad

type heap[T comparable] struct {
	datos    []T
	cantidad int
	comparar func(T, T) int
}

func (cola *heap[T]) EstaVacia() bool {
	return cola.cantidad == 0
}

func (cola *heap[T]) VerMax() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.datos[0]
}

func (cola *heap[T]) Cantidad() int {
	return cola.cantidad
}

func (cola *heap[T]) swap(ind1 int, ind2 int) {
	reemplazado := cola.datos[ind1]
	cola.datos[ind1] = cola.datos[ind2]
	cola.datos[ind2] = reemplazado
}

func (cola *heap[T]) upheap(hijo int) {
	if hijo == 0 {
		return
	}
	padre := (hijo - 1) / 2
	if cola.comparar(cola.datos[padre], cola.datos[hijo]) < 0 {
		cola.swap(padre, hijo)
		cola.upheap(padre)
	}
}

func (cola *heap[T]) Encolar(elem T) {

	if cola.cantidad == cap(cola.datos) {
		cola.redimensionar(cap(cola.datos) * 2)
	}

	nuevaPosicion := cola.cantidad
	cola.datos[nuevaPosicion] = elem
	cola.upheap(nuevaPosicion)
	cola.cantidad++
}

func (cola *heap[T]) redimensionar(nuevoTam int) {
	nuevaCola := make([]T, nuevoTam)
	copy(nuevaCola, cola.datos)
	cola.datos = nuevaCola
}

func (cola heap[T]) downheap(hijoIzq int, hijoDer int) {

}

func (cola heap[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	pos_ultimo := cola.cantidad - 1
	cola.swap(0, pos_ultimo)
	devolver := cola.datos[pos_ultimo]

	cola.downheap(1, 2)

}
