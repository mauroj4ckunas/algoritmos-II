package diccionario

import (
	FuncHash "FuncHash"
)

var CAPACIDAD = uint64(87)

const (
	POSICIONESHABILES          uint64 = 3
	MULTIPLICADOR_DE_CAPACIDAD int    = 3
)

type elementos[K comparable, V any] struct {
	clave     K
	valor     V
	ubicacion uint64
}

type diccionario_implementacion[K comparable, V any] struct {
	array []*elementos[K, V]
	largo int
}

func crearElemento[K comparable, V any](clave K, valor V, ubicacionHash uint64) *elementos[K, V] {
	dato := new(elementos[K, V])
	(*dato).clave = clave
	(*dato).valor = valor
	(*dato).ubicacion = ubicacionHash
	return dato
}

func (dicc *diccionario_implementacion[K, V]) redimensionar(nuevoTam int) {
	CAPACIDAD = uint64(nuevoTam)
	nuevoArray := crearArrayHash[K, V](CAPACIDAD)
	arrayViejo := dicc.array
	dicc.array = nuevoArray
	dicc.largo = 0
	for i := 0; i < len(arrayViejo); i++ {
		if arrayViejo[i] != nil {
			dicc.Guardar((*arrayViejo[i]).clave, (*arrayViejo[i]).valor)
		}
	}
}

func (dicc *diccionario_implementacion[K, V]) hacerEspacio(indice uint64, lugarNecesario uint64) (uint64, bool) {
	if (CAPACIDAD+indice)%CAPACIDAD < (CAPACIDAD+lugarNecesario+POSICIONESHABILES)%CAPACIDAD {
		return indice % CAPACIDAD, false
	}

	if (*dicc.array[(CAPACIDAD+indice-2)%CAPACIDAD]).ubicacion+1 == indice || (*dicc.array[(CAPACIDAD+indice-2)%CAPACIDAD]).ubicacion+2 == indice {

		dicc.array[(CAPACIDAD+indice-2)%CAPACIDAD], dicc.array[indice] = dicc.array[indice], dicc.array[(CAPACIDAD+indice-2)%CAPACIDAD]

		return dicc.hacerEspacio(((CAPACIDAD + indice - 2) % CAPACIDAD), lugarNecesario)

	} else if (*dicc.array[(CAPACIDAD+indice-1)%CAPACIDAD]).ubicacion+1 == indice || (*dicc.array[(CAPACIDAD+indice-1)%CAPACIDAD]).ubicacion+2 == indice {

		dicc.array[(CAPACIDAD+indice-1)%CAPACIDAD], dicc.array[indice] = dicc.array[indice], dicc.array[(CAPACIDAD+indice-1)%CAPACIDAD]

		return dicc.hacerEspacio(((CAPACIDAD + indice - 1) % CAPACIDAD), lugarNecesario)

	}

	dicc.redimensionar(cap(dicc.array) * MULTIPLICADOR_DE_CAPACIDAD)
	return 0, true
}

func (dicc *diccionario_implementacion[K, V]) Guardar(clave K, dato V) {

	indiceHash := FuncHash.Hashear[K](clave) % CAPACIDAD
	var posicion uint64

	if dicc.Pertenece(clave) {
		for posicion = indiceHash; posicion < (indiceHash + POSICIONESHABILES); posicion++ {

			if (*dicc.array[posicion%CAPACIDAD]).clave == clave {

				dicc.array[posicion%CAPACIDAD].valor = dato
				return

			}
		}
	}

	for posicion = indiceHash; posicion < (indiceHash + POSICIONESHABILES); posicion++ {

		if dicc.array[posicion%CAPACIDAD] == nil {

			paraGuardar := crearElemento[K, V](clave, dato, indiceHash)
			dicc.array[posicion%CAPACIDAD] = paraGuardar
			dicc.largo++
			return

		}
	}

	if posicion == (indiceHash + POSICIONESHABILES) {

		for true {
			if dicc.array[posicion%CAPACIDAD] == nil {

				break

			}
			posicion = (posicion + 1) % CAPACIDAD
		}

		posicion, redimension := dicc.hacerEspacio(posicion%CAPACIDAD, indiceHash)

		if redimension {

			dicc.Guardar(clave, dato)

		} else {

			paraGuardar := crearElemento[K, V](clave, dato, indiceHash)
			dicc.array[posicion] = paraGuardar
			dicc.largo++

		}
	}

}

func (dicc *diccionario_implementacion[K, V]) Pertenece(clave K) bool {

	ubicacion := FuncHash.Hashear[K](clave) % CAPACIDAD

	for i := ubicacion; i < (ubicacion + POSICIONESHABILES); i++ {
		if dicc.array[i%CAPACIDAD] == nil {
			continue
		} else if (*dicc.array[i%CAPACIDAD]).clave == clave {
			return true
		}
	}
	return false
}

func (dicc *diccionario_implementacion[K, V]) Obtener(clave K) V {

	ubicacion := FuncHash.Hashear[K](clave) % CAPACIDAD

	for i := ubicacion; i < (ubicacion + POSICIONESHABILES); i++ {
		if dicc.array[i%CAPACIDAD] == nil {
			continue
		} else if (*dicc.array[i%CAPACIDAD]).clave == clave {
			return (*dicc.array[i%CAPACIDAD]).valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *diccionario_implementacion[K, V]) Borrar(clave K) V {

	ubicacion := FuncHash.Hashear[K](clave) % CAPACIDAD

	if dicc.Cantidad() < int(CAPACIDAD)/2 && dicc.Cantidad() > int(CAPACIDAD)/4 {
		dicc.redimensionar(int(CAPACIDAD) / 2)
	}

	for i := ubicacion; i < (ubicacion + POSICIONESHABILES); i++ {

		if dicc.array[i%CAPACIDAD] == nil {

			continue

		} else if (*dicc.array[i%CAPACIDAD]).clave == clave {

			devolver := (*dicc.array[i%CAPACIDAD]).valor
			dicc.array[i%CAPACIDAD] = nil
			dicc.largo--
			return devolver
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *diccionario_implementacion[K, V]) Cantidad() int {
	return dicc.largo
}

func crearArrayHash[K comparable, V any](tam uint64) []*elementos[K, V] {
	nuevoArray := make([]*elementos[K, V], tam)
	return nuevoArray
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	diccio := new(diccionario_implementacion[K, V])
	(*diccio).array = crearArrayHash[K, V](CAPACIDAD)
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

	for i := 0; i < len(dicc.array); i++ {
		if dicc.array[i] == nil {
			continue
		}
		if !f(dicc.array[i].clave, dicc.array[i].valor) {
			break
		}
	}

}
