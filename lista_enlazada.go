package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iteradorExterno[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	proximo  *nodoLista[T]
}

/* Primitivas de Lista */

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevo := crearNodo[T](elem)
	if lista.EstaVacia() {
		iniciarLista(lista, nuevo)
	} else {
		nuevo.siguiente = lista.primero
		lista.primero = nuevo
	}
	lista.largo++
}

func (lista listaEnlazada[T]) InsertarUltimo(elem T) {
	nuevo := crearNodo[T](elem)
	if lista.EstaVacia() {
		iniciarLista(lista, nuevo)
	} else {
		lista.ultimo.siguiente = nuevo
		lista.ultimo = nuevo
	}
	lista.largo++
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
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

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

// func CrarListaEnlazada[T any]() Lista[T] {
// 	lista := new(listaEnlazada[T])
// 	return lista
// }

/* Primitivas de los iteradores */

func (iter iteradorExterno[T]) VerActual() T {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter iteradorExterno[T]) HaySiguiente() bool {
	return iter.proximo != nil
}

func (iter iteradorExterno[T]) Siguiente() T {
	datoActual := iter.VerActual()
	iter.anterior = iter.actual
	iter.actual = iter.proximo
	iter.proximo = iter.proximo.siguiente
	return datoActual
}

// El iterador va bien pero no esta modificando la lista original

func (iter iteradorExterno[T]) Insertar(elem T) {
	nuevo := crearNodo[T](elem)
	nuevo.siguiente = iter.actual
	iter.anterior.siguiente = nuevo
	iter.proximo = iter.actual
	iter.actual = nuevo
}

func (iter iteradorExterno[T]) Borrar() T {
	datoFuera := iter.VerActual()
	iter.anterior.siguiente = iter.proximo
	iter.actual = iter.proximo
	iter.proximo = iter.proximo.siguiente
	return datoFuera
}

func (lista listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := new(iteradorExterno[T])
	iterador.actual = lista.primero
	iterador.proximo = lista.primero.siguiente
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {

}

// Funciones auxiliares

func iniciarLista[T any](listaVacia listaEnlazada[T], nodo *nodoLista[T]) {
	listaVacia.primero = nodo
	listaVacia.ultimo = nodo
}

func crearNodo[T any](dato T) *nodoLista[T] {
	nuevoNodo := new(nodoLista[T])
	nuevoNodo.dato = dato
	return nuevoNodo
}
