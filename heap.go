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

func (cola heap[T]) maximoEntreHijos(indPadre int) int {
	hijoIzq := (indPadre * 2) + 1
	hijoDer := (indPadre * 2) + 2
	var indice int
	if cola.comparar(cola.datos[hijoIzq], cola.datos[hijoDer]) >= 0 {
		indice = hijoIzq
	} else {
		indice = hijoDer
	}
	return indice
}

func (cola *heap[T]) downheap(padre int) {
	if padre == cola.cantidad-1 {
		return
	}
	mayor := cola.maximoEntreHijos(padre)
	if cola.comparar(cola.datos[padre], cola.datos[mayor]) < 0 {
		cola.swap(padre, mayor)
		padre = mayor
		cola.downheap(padre)
	}
}

func (cola *heap[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	if cola.cantidad*4 <= cap(cola.datos) {
		cola.redimensionar(cap(cola.datos) / 2)
	}

	pos_ultimo := cola.cantidad - 1
	cola.swap(0, pos_ultimo)
	devolver := cola.datos[pos_ultimo]
	cola.downheap(0)
	cola.cantidad--
	return devolver
}

func (cola *heap[T]) redimensionar(nuevoTam int) {
	nuevaCola := make([]T, nuevoTam)
	copy(nuevaCola, cola.datos)
	cola.datos = nuevaCola
}

func CrearHeap[T comparable](f_comparar func(T, T) int) ColaPrioridad[T] {
	nuevoHeap := new(heap[T])
	const tamFijo int = 20
	array := make([]T, tamFijo)
	nuevoHeap.datos = array
	nuevoHeap.comparar = f_comparar

	return nuevoHeap
}
