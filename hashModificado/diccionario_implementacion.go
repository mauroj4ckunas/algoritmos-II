package diccionario

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

type elementos[K comparable, V any] struct {
	clave     K
	valor     V
	ubicacion uint64
}

type diccionario_implementacion[K comparable, V any] struct {
	array       []*elementos[K, V]
	largo       int
	factorCarga float64
}

func hashear[K comparable](clave K) uint64 {
	elementoHasheable := []byte(fmt.Sprintf("%v", clave))
	hasheado := md5.Sum(elementoHasheable)
	var arrayUint64 []byte = hasheado[:]
	return binary.BigEndian.Uint64(arrayUint64)
}

func crearElemento[K comparable, V any](clave K, valor V, ubicacionHash uint64) *elementos[K, V] {
	dato := new(elementos[K, V])
	(*dato).clave = clave
	(*dato).valor = valor
	(*dato).ubicacion = ubicacionHash
	return dato
}



func (dicc *diccionario_implementacion[K, V]) redimensionar(nuevoTam int) {
	nuevoArray := make([]*elementos[K, V], nuevoTam)
	arrayViejo := dicc.array
	dicc.array = nuevoArray
	dicc.largo = 0
	for i := 0; i < len(arrayViejo); i++ {
		if arrayViejo[i] != nil {
			dicc.Guardar((*arrayViejo[i]).clave, (*arrayViejo[i]).valor)
		}
	}
}

func (dicc *diccionario_implementacion[K, V]) hacerEspacio(indice uint64, lugarNecesario uint64) (uint64,bool) {
	if indice < lugarNecesario+3 {
		return indice, false
	}

	if (*dicc.array[indice-2]).ubicacion+1 == indice || (*dicc.array[indice-2]).ubicacion+2 == indice {

		dicc.array[indice-2], dicc.array[indice] = dicc.array[indice], dicc.array[indice-2]

		return dicc.hacerEspacio(indice-2, lugarNecesario)

	} else if (*dicc.array[indice-1]).ubicacion+1 == indice || (*dicc.array[indice-1]).ubicacion+2 == indice {

		dicc.array[indice-1], dicc.array[indice] = dicc.array[indice], dicc.array[indice-1]

		return dicc.hacerEspacio(indice-1, lugarNecesario)

	}

	dicc.redimensionar(cap(dicc.array) * 2)
	return 0 , true
}

func (dicc *diccionario_implementacion[K, V]) Guardar(clave K, dato V) {
	dicc.factorCarga = float64(dicc.Cantidad()) / float64(cap(dicc.array))

	if dicc.factorCarga >= 0.7 {

		dicc.redimensionar(cap(dicc.array) * 2)

	}

	indiceHash := hashear[K](clave) % uint64(len(dicc.array))
	var posicion uint64

	if dicc.Pertenece(clave) {
		for posicion = indiceHash; posicion < (indiceHash + 3); posicion++ {

			if (*dicc.array[posicion % uint64(len(dicc.array))]).clave == clave {

				dicc.array[posicion % uint64(len(dicc.array))].valor = dato
				return

			}
		}
	}


	for posicion = indiceHash; posicion < (indiceHash + 3); posicion++ {

		if dicc.array[posicion % uint64(len(dicc.array))] == nil {

			paraGuardar := crearElemento[K, V](clave, dato, indiceHash)
			dicc.array[posicion % uint64(len(dicc.array))] = paraGuardar
			dicc.largo++
			return

		}
	}

	if posicion == (indiceHash + 3) {

		for true {

			if dicc.array[posicion] == nil {

				break

			}
			posicion = (posicion + 1) % uint64(len(dicc.array))
		}

		posicionNueva, redimension := dicc.hacerEspacio(posicion, indiceHash)

		if redimension == true {

			dicc.Guardar(clave, dato)

		} else {

			paraGuardar := crearElemento[K, V](clave, dato, indiceHash)
			dicc.array[posicionNueva] = paraGuardar
			dicc.largo++

		}
	}
	

}

func (dicc *diccionario_implementacion[K, V]) Pertenece(clave K) bool {

	ubicacion := hashear[K](clave) % uint64(len(dicc.array))

	for i := ubicacion; i < (ubicacion + 3); i++ {
		if dicc.array[i % uint64(len(dicc.array))] == nil {
			continue
		} else if (*dicc.array[i % uint64(len(dicc.array))]).clave == clave {
			return true
		}
	}
	return false
}

func (dicc *diccionario_implementacion[K, V]) Obtener(clave K) V {

	ubicacion := hashear[K](clave) % uint64(len(dicc.array))

	for i := ubicacion; i < (ubicacion + 3); i++ {

		if dicc.array[i % uint64(len(dicc.array))] == nil {
			continue
		} else if (*dicc.array[i % uint64(len(dicc.array))]).clave == clave {
			return (*dicc.array[i % uint64(len(dicc.array))]).valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *diccionario_implementacion[K, V]) Borrar(clave K) V {

	dicc.factorCarga = float64(dicc.Cantidad()) / float64(cap(dicc.array))

	if dicc.factorCarga < 0.1 {

		dicc.redimensionar(cap(dicc.array) / 2)

	}

	ubicacion := hashear[K](clave) % uint64(len(dicc.array))

	for i := ubicacion; i < (ubicacion + 3); i++ {

		if dicc.array[i % uint64(len(dicc.array))] == nil {

			continue

		} else if (*dicc.array[i % uint64(len(dicc.array))]).clave == clave {

			devolver := (*dicc.array[i % uint64(len(dicc.array))]).valor
			dicc.array[i % uint64(len(dicc.array))] = nil
			dicc.largo--
			return devolver
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *diccionario_implementacion[K, V]) Cantidad() int {
	return dicc.largo
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	diccio := new(diccionario_implementacion[K, V])
	(*diccio).array = make([]*elementos[K, V], 100)
	return diccio
}

// iterador externo del diccionario
type iterador_externo[K comparable, V any] struct {
	actual int
	dicc   []*elementos[K, V]
}

// creador de iterador externo

func (dicc *diccionario_implementacion[K, V]) Iterador() IterDiccionario[K, V] {
	iterr := new(iterador_externo[K, V])
	iterr.dicc = dicc.array
	for iterr.actual < len(iterr.dicc) {
		if iterr.dicc[iterr.actual] != nil {
			break
		}
		iterr.actual++
	}
	return iterr
}

func (iterr *iterador_externo[K, V]) HaySiguiente() bool {
	return iterr.actual < len(iterr.dicc)
}

func (iterr *iterador_externo[K, V]) VerActual() (K, V) {
	if !iterr.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	//fmt.Println(iterr.dicc[iterr.actual].clave)
	return (*iterr.dicc[iterr.actual]).clave, (*iterr.dicc[iterr.actual]).valor
}

func (iterr *iterador_externo[K, V]) Siguiente() K {
	if !iterr.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	devolver := (*iterr.dicc[iterr.actual]).clave
	iterr.actual++
	for iterr.HaySiguiente() {
		if iterr.dicc[iterr.actual] != nil {
			break
		}
		iterr.actual++
	}
	return devolver
}

// Iterador Interno

func (dicc *diccionario_implementacion[K, V]) Iterar(f func(clave K, valor V) bool) {

	for i := 0 ; i < len(dicc.array) ; i++ {
		if dicc.array[i] == nil {
			continue
		}
		if !f(dicc.array[i].clave,dicc.array[i].valor){
			break
		}
	}

}
