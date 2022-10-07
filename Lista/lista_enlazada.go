package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	nuevoNodo := new(nodoLista[T])
	nuevoNodo.dato = dato
	return nuevoNodo
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

type iteradorExterno[T any] struct {
	actual       *nodoLista[T]
	anterior     *nodoLista[T]
	listaExterna *listaEnlazada[T]
}

/* Primitivas de Lista */

func (lista *listaEnlazada[T]) iniciarLista(nodo *nodoLista[T]) {
	lista.primero = nodo
	lista.ultimo = nodo
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevo := crearNodo[T](elem)
	if lista.EstaVacia() {
		lista.iniciarLista(nuevo)
	} else {
		nuevo.siguiente = lista.primero
		lista.primero = nuevo
	}

	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	nuevo := crearNodo[T](elem)
	if lista.EstaVacia() {
		lista.iniciarLista(nuevo)
	} else {
		lista.ultimo.siguiente = nuevo
		lista.ultimo = nuevo
	}

	lista.largo++
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	datoPrimero := lista.VerPrimero()
	if lista.primero == lista.ultimo {
		lista.primero = nil
		lista.ultimo = nil
	} else {
		lista.primero = lista.primero.siguiente
	}
	lista.largo--
	return datoPrimero
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := new(iteradorExterno[T])
	iterador.actual = lista.primero
	iterador.listaExterna = lista
	return iterador
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for i := 0; i < lista.largo && visitar(actual.dato); i++ {
		actual = actual.siguiente
	}
}

/* Primitivas de los iteradores */

func (iter *iteradorExterno[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iteradorExterno[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorExterno[T]) Siguiente() T {
	devolver := iter.VerActual()
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
	return devolver
}

func (iter *iteradorExterno[T]) Insertar(elem T) {

	if !iter.HaySiguiente() && iter.anterior != nil {

		iter.listaExterna.InsertarUltimo(elem)
		iter.actual = iter.listaExterna.ultimo

	} else if iter.anterior == nil {

		iter.listaExterna.InsertarPrimero(elem)
		iter.actual = iter.listaExterna.primero

	} else {
		nuevoNodo := crearNodo[T](elem)
		iter.anterior.siguiente = nuevoNodo
		nuevoNodo.siguiente = iter.actual
		iter.actual = nuevoNodo
		iter.listaExterna.largo++
	}
}

func (iter *iteradorExterno[T]) Borrar() T {
	devolver := iter.VerActual()
	if iter.actual.siguiente == nil {
		iter.listaExterna.ultimo = iter.anterior
	}
	if iter.anterior == nil {

		iter.listaExterna.primero = iter.actual.siguiente
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
	}
	iter.actual = iter.actual.siguiente
	iter.listaExterna.largo--
	return devolver
}
