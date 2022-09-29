package lista

type iteradorLista[T any] struct {
	anterior *nodo[T]
	actual   *nodo[T]
	lista    *listaEnlazada[T]
}

func crearIteradorExterno[T any](actual *nodo[T], lista *listaEnlazada[T]) IteradorLista[T] {
	iterador := new(iteradorLista[T])
	iterador.actual = actual
	iterador.lista = lista
	return iterador
}

func (iter *iteradorLista[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iteradorLista[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorLista[T]) Siguiente() T {
	devolver := iter.VerActual()
	iter.anterior = iter.actual
	iter.actual = iter.actual.prox
	return devolver
}

func (iter *iteradorLista[T]) Insertar(elem T) {
	nuevoNodo := crearNodo[T](elem)

	if !iter.HaySiguiente() {
		iter.lista.ulti = nuevoNodo
	}

	if iter.anterior == nil {
		iter.lista.prim = nuevoNodo
	} else {
		iter.anterior.prox = nuevoNodo
	}

	nuevoNodo.prox = iter.actual
	iter.actual = nuevoNodo
	iter.lista.cantidad++
}

func (iter *iteradorLista[T]) Borrar() T {
	devolver := iter.VerActual()

	if iter.actual.prox == nil {
		iter.lista.ulti = iter.anterior
	}

	if iter.anterior == nil {
		iter.lista.prim = iter.actual.prox
	} else {
		iter.anterior.prox = iter.actual.prox
	}

	iter.actual = iter.actual.prox
	iter.lista.cantidad--
	return devolver
}


