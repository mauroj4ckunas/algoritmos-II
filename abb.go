package diccionario

import (
	TDApila "diccionario/Pila"
)

type hojas[K comparable, V any] struct {
	clave   K
	valor   V
	hijoIzq *hojas[K, V]
	hijoDer *hojas[K, V]
}

type arbolBinario[K comparable, V any] struct {
	raiz       *hojas[K, V]
	comparador func(K, K) int
	cantidad   int
}

func crearHoja[K comparable, V any](clave K, dato V) *hojas[K, V] {

	hojaNueva := new(hojas[K, V])
	hojaNueva.clave = clave
	hojaNueva.valor = dato
	return hojaNueva

}

func (hoja *hojas[K, V]) encontrarClave(compara func(K, K) int, clave K) (**hojas[K, V], string) {

	if compara(clave, hoja.clave) < 0 {
		if hoja.hijoIzq == nil {

			return &hoja.hijoIzq, "La clave no pertenece al diccionario"

		} else if compara(clave, hoja.hijoIzq.clave) == 0 {

			return &hoja.hijoIzq, ""

		}

		return hoja.hijoIzq.encontrarClave(compara, clave)

	} else {

		if hoja.hijoDer == nil {

			return &hoja.hijoDer, "La clave no pertenece al diccionario"

		} else if compara(clave, hoja.hijoDer.clave) == 0 {

			return &hoja.hijoDer, ""

		}

		return hoja.hijoDer.encontrarClave(compara, clave)
	}

}

func (hoja *hojas[K, V]) iterar(comparador func(K, K) int, f func(clave K, dato V) bool, desde *K, hasta *K) bool {
	if hoja == nil {
		return true
	}

	resultado := hoja.hijoIzq.iterar(comparador, f, desde, hasta)
	if resultado == false {
		return resultado
	}

	if (hasta != nil && comparador(hoja.clave, *hasta) <= 0) || hasta == nil {
		if (desde != nil && comparador(hoja.clave, *desde) >= 0) || desde == nil {
			if f(hoja.clave, hoja.valor) == false {
				return false
			}
		}
	} else {
		return false
	}

	return hoja.hijoDer.iterar(comparador, f, desde, hasta)
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {

	arbol := new(arbolBinario[K, V])
	arbol.comparador = funcion_cmp
	return arbol

}

func (arbol *arbolBinario[K, V]) encontrarClave(clave K) (**hojas[K, V], string) {

	if arbol.raiz == nil {

		return &arbol.raiz, "La clave no pertenece al diccionario"

	} else if arbol.raiz.clave == clave {

		return &arbol.raiz, ""

	}
	resultado, err := arbol.raiz.encontrarClave(arbol.comparador, clave)
	return resultado, err
}

func (arbol *arbolBinario[K, V]) Guardar(clave K, dato V) {

	AGuardar, _ := arbol.encontrarClave(clave)
	if *AGuardar == nil {
		*AGuardar = crearHoja[K, V](clave, dato)
		arbol.cantidad++
	} else {
		(*AGuardar).valor = dato
	}

}

func (arbol *arbolBinario[K, V]) Pertenece(clave K) bool {

	_, error := arbol.encontrarClave(clave)
	return error == ""

}

func (arbol *arbolBinario[K, V]) Obtener(clave K) V {

	resultado, err := arbol.encontrarClave(clave)

	if err != "" {

		panic(err)

	}

	return (*resultado).valor
}

func (arbol *arbolBinario[K, V]) Cantidad() int {
	return arbol.cantidad
}

func (arbol *arbolBinario[K, V]) Borrar(clave K) V {

	borrar, err := arbol.encontrarClave(clave)

	if err != "" {

		panic(err)

	}

	arbol.cantidad--

	var devolver V
	devolver = (*borrar).valor
	borrarSinHijos[K, V](borrar)
	return devolver
}
func borrarSinHijos[K comparable, V any](borrar **hojas[K, V]) {
	if (*(*borrar)).hijoDer == nil && (*(*borrar)).hijoIzq == nil {
		*borrar = nil
	} else {
		tieneUnHijo[K, V](borrar)
	}
}
func tieneUnHijo[K comparable, V any](borrar **hojas[K, V]) {
	if (*(*borrar)).hijoDer == nil && (*(*borrar)).hijoIzq != nil {
		*borrar = (*borrar).hijoIzq
	} else if (*(*borrar)).hijoDer != nil && (*(*borrar)).hijoIzq == nil {
		*borrar = (*borrar).hijoDer
	} else {
		tieneDosHijos[K, V](borrar)
	}
}

func tieneDosHijos[K comparable, V any](borrar **hojas[K, V]) {
	if (*(*borrar)).hijoDer != nil && (*(*borrar)).hijoIzq != nil {

		reemplazante := &(*borrar).hijoDer

		for (*reemplazante).hijoIzq != nil {

			reemplazante = &(*reemplazante).hijoIzq

		}

		(*borrar).clave, (*reemplazante).clave = (*reemplazante).clave, (*borrar).clave
		(*borrar).valor, (*reemplazante).valor = (*reemplazante).valor, (*borrar).valor

		borrarSinHijos(reemplazante)
	}
}

func (arbol *arbolBinario[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	arbol.raiz.iterar(arbol.comparador, visitar, desde, hasta)
}

func (arbol *arbolBinario[K, V]) Iterar(f func(clave K, dato V) bool) {

	arbol.IterarRango(nil, nil, f)

}

type iterExterno[K comparable, V any] struct {
	pilaRecursiva TDApila.Pila[*hojas[K, V]]
	hasta         *K
	comparador    func(K, K) int
}

func (arbol *arbolBinario[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

	iterr := new(iterExterno[K, V])
	iterr.pilaRecursiva = TDApila.CrearPilaDinamica[*hojas[K, V]]()
	iterr.hasta = hasta
	iterr.comparador = arbol.comparador

	if desde != nil && hasta != nil && arbol.comparador(*desde, *hasta) > 0 {

		return iterr

	}
	todoIzquierda := arbol.raiz
	for todoIzquierda != nil {
		if (iterr.hasta != nil && iterr.comparador(todoIzquierda.clave, *iterr.hasta) <= 0) || iterr.hasta == nil {
			iterr.pilaRecursiva.Apilar(todoIzquierda)
		}
		todoIzquierda = todoIzquierda.hijoIzq
	}
	if desde != nil {
		for !iterr.pilaRecursiva.EstaVacia() && iterr.comparador(iterr.pilaRecursiva.VerTope().clave, *desde) < 0 {
			iterr.Siguiente()
		}
	}

	return iterr
}

func (arbol *arbolBinario[K, V]) Iterador() IterDiccionario[K, V] {

	return arbol.IteradorRango(nil, nil)

}

func (iterr *iterExterno[K, V]) HaySiguiente() bool {

	return !iterr.pilaRecursiva.EstaVacia()

}

func (iterr *iterExterno[K, V]) VerActual() (K, V) {

	if !iterr.HaySiguiente() {

		panic("El iterador termino de iterar")

	}

	return iterr.pilaRecursiva.VerTope().clave, iterr.pilaRecursiva.VerTope().valor
}

func (iterr *iterExterno[K, V]) Siguiente() K {

	if !iterr.HaySiguiente() {

		panic("El iterador termino de iterar")

	}

	devolver := iterr.pilaRecursiva.Desapilar()
	if iterr.hasta != nil && iterr.comparador(devolver.clave, *iterr.hasta) == 0 {
		for !iterr.pilaRecursiva.EstaVacia() {
			iterr.pilaRecursiva.Desapilar()
		}
	} else if devolver.hijoDer != nil {
		todoIzquierda := devolver.hijoDer
		for todoIzquierda != nil {
			if (iterr.hasta != nil && iterr.comparador(todoIzquierda.clave, *iterr.hasta) <= 0) || iterr.hasta == nil {
				iterr.pilaRecursiva.Apilar(todoIzquierda)
			}
			todoIzquierda = todoIzquierda.hijoIzq
		}
	}

	return devolver.clave

}
