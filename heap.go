package cola_prioridad

type heap[T comparable] struct {
	datos    []T
	cantidad int
	comparar func(T, T) int
}

const TAM_INICIAL int = 20

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

func swap[T comparable](ind1 int, ind2 int, arr []T) {
	reemplazado := arr[ind1]
	arr[ind1] = arr[ind2]
	arr[ind2] = reemplazado
}

func upheap[T comparable](hijo int, array []T, comparar func(T, T) int) {
	if hijo == 0 {
		return
	}
	padre := (hijo - 1) / 2
	if comparar(array[padre], array[hijo]) < 0 {
		swap(padre, hijo, array)
		upheap(padre, array, comparar)
	}
}

func (cola *heap[T]) Encolar(elem T) {

	if cola.cantidad == cap(cola.datos) {
		cola.redimensionar(cap(cola.datos) * 2)
	}

	nuevaPosicion := cola.cantidad
	cola.datos[nuevaPosicion] = elem
	upheap(nuevaPosicion, cola.datos, cola.comparar)
	cola.cantidad++
}

func maximoEntreHijos[T comparable](arreglo []T, indPadre int, f func(T, T) int, cant int) int {
	hijoIzq := (indPadre * 2) + 1
	hijoDer := (indPadre * 2) + 2
	if hijoDer >= cant || f(arreglo[hijoIzq], arreglo[hijoDer]) > 0 {
		return hijoIzq
	}
	return hijoDer
}

func downheap[T comparable](padre int, array []T, cantidad int, comparar func(T, T) int) {
	if padre == cantidad-1 {
		return
	}
	mayor := maximoEntreHijos(array, padre, comparar, cantidad)
	if mayor >= cantidad {
		return
	}
	if comparar(array[padre], array[mayor]) < 0 {
		swap(padre, mayor, array)
		padre = mayor
		downheap(padre, array, cantidad, comparar)
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
	swap(0, pos_ultimo, cola.datos)
	devolver := cola.datos[pos_ultimo]
	cola.cantidad--
	downheap(0, cola.datos, cola.cantidad, cola.comparar)
	return devolver
}

func (cola *heap[T]) redimensionar(nuevoTam int) {
	nuevaCola := make([]T, nuevoTam)
	copy(nuevaCola, cola.datos)
	cola.datos = nuevaCola
}

func CrearHeap[T comparable](f_comparar func(T, T) int) ColaPrioridad[T] {
	nuevoHeap := new(heap[T])
	array := make([]T, TAM_INICIAL)
	nuevoHeap.datos = array
	nuevoHeap.comparar = f_comparar

	return nuevoHeap
}

func heapify[T comparable](arr []T, cantidad int, cmp func(T, T) int) {
	for i := len(arr) - 1; i >= 1; i-- {
		padre := (i - 1) / 2
		downheap(padre, arr, cantidad, cmp)
	}
}

func CrearHeapArr[T comparable](arr []T, f_comparar func(T, T) int) ColaPrioridad[T] {
	if len(arr) == 0 {
		return CrearHeap(f_comparar)
	}
	arrHeap := new(heap[T])
	arrHeap.comparar = f_comparar
	arrHeap.cantidad = len(arr)
	arrHeap.datos = arr
	arrHeap.redimensionar(arrHeap.cantidad)
	if arrHeap.cantidad > 0 {
		heapify(arrHeap.datos, arrHeap.cantidad, arrHeap.comparar)
	}
	return arrHeap
}

func HeapSort[T comparable](elementos []T, f_comparar func(T, T) int) {
	heapify(elementos, len(elementos), f_comparar)
	heapsort(elementos, len(elementos), f_comparar)
}

func heapsort[T comparable](elem []T, cant int, cmp func(T, T) int) {
	ultimoRelativo := cant - 1 //indice del ultimo relativo
	if ultimoRelativo == 0 {
		return
	}
	swap(0, ultimoRelativo, elem)
	downheap(0, elem, ultimoRelativo, cmp)
	heapsort(elem, ultimoRelativo, cmp)
}
