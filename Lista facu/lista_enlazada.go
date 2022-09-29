package lista

type listaEnlazada[T any] struct {
	cantidad int
	prim     *nodo[T]
	ulti     *nodo[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.cantidad == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nodo := crearNodo[T](elem)
	if lista.EstaVacia() {
		lista.ulti = nodo
	} else {
		nodo.prox = lista.prim
	}
	lista.prim = nodo
	lista.cantidad++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	nodo := crearNodo[T](elem)
	if lista.EstaVacia() {
		lista.prim = nodo
	} else {
		lista.ulti.prox = nodo
	}
	lista.ulti = nodo
	lista.cantidad++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	retorno := lista.VerPrimero()
	lista.prim = lista.prim.prox
	lista.cantidad--
	if lista.prim == nil {
		lista.ulti = nil
	}
	return retorno
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.prim.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ulti.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.cantidad
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.prim
	for i := 0; i < lista.cantidad && visitar(actual.dato); i++ {
		actual = actual.prox
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return crearIteradorExterno[T](lista.prim, lista)
}
