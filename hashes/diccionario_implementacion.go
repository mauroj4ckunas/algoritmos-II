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
	conValor  bool
}

type diccionario_implementacion[K comparable, V any] struct {
	array []elementos[K, V]
	largo int
}

func hashear[K comparable](clave K) uint64 {
	elementoHasheable := []byte(fmt.Sprintf("%v", clave))
	hasheado := md5.Sum(elementoHasheable)
	var arrayUint64 []byte = hasheado[:]
	// var convertirAEntero int = int(binary.BigEndian.Uint64(arrayUint64))
	// if convertirAEntero < 0 {
	// 	return convertirAEntero * -1
	// }
	return binary.BigEndian.Uint64(arrayUint64)
}

func crearElemento[K comparable, V any](clave K, valor V, ubicacionHash uint64) *elementos[K, V] {
	dato := new(elementos[K, V])
	dato.clave = clave
	dato.valor = valor
	dato.conValor = true
	dato.ubicacion = ubicacionHash
	return dato
}

func (dicc *diccionario_implementacion[K, V]) redimensionar() {

}

func (dicc *diccionario_implementacion[K, V]) hacerEspacio(indice uint64) (uint64, bool) {

	var espacio uint64
	if dicc.array[indice].ubicacion > indice {
		espacio = dicc.array[indice].ubicacion - indice
	} else if dicc.array[indice].ubicacion == indice {
		espacio = 3
	}

	for i := indice; i < (indice + espacio); i++ {
		if dicc.array[i].conValor == false {
			dicc.array[i].ubicacion = i
			return i, false
		}
	}

	for j := indice + 1; j < (indice + espacio); j++ {
		k, cambioDeLugar := dicc.hacerEspacio(j)
		if !cambioDeLugar {
			dicc.array[k].clave = dicc.array[j].clave
			dicc.array[k].valor = dicc.array[j].valor
			dicc.array[k].conValor = true
			dicc.array[k].ubicacion = k
			return j, false
		}
	}
	return 0, true
}

// En caso que algun espacio este vacio, entra en el primer if y termina el for con un break
// pero si tiene algun valor, no entrara al if y no terminara el ciclo, por ende en ultimo caso
// que seria indice+2 antes que cierra el for, se llamara a una funcion.
func (dicc *diccionario_implementacion[K, V]) Guardar(clave K, dato V) {
	indice := hashear[K](clave) % uint64(len(dicc.array))
	paraGuardar := crearElemento[K, V](clave, dato, indice)
	var i uint64
	var hayRedimension bool

	for ind := indice; ind < (indice + 3); ind++ {
		if dicc.array[ind].clave == clave {
			dicc.array[ind].valor = dato
			// En el caso que se borre un elemento, simplemente debe ver el conValor para saber si se puede sobreescribir otro elemento.
			// Pero que pasa si se quiere guardar el mismo elemento despues de haberlo borrado? Se debe devuelta agregar true al conValor.
			if !dicc.array[ind].conValor {
				dicc.array[ind].conValor = true
			}
			return
		}
	}

	switch {
	case dicc.largo == 0:
		i = indice
		hayRedimension = false

	case dicc.array[indice].conValor == true:
		i, hayRedimension = dicc.hacerEspacio(indice)
	}

	if hayRedimension {
		dicc.redimensionar()
	} else {
		dicc.array[i] = *paraGuardar
		dicc.largo++
	}

}

func (dicc *diccionario_implementacion[K, V]) Pertenece(clave K) bool {
	ubicacion := hashear[K](clave) % uint64(len(dicc.array))
	for i := ubicacion; i < (ubicacion + 3); i++ {
		if dicc.array[i].clave == clave {
			return true
		}
	}
	return false
}

func (dicc *diccionario_implementacion[K, V]) Obtener(clave K) V {
	ubicacion := hashear[K](clave) % uint64(len(dicc.array))
	for i := ubicacion; i < (ubicacion + 3); i++ {
		if dicc.array[i].clave == clave {
			return dicc.array[i].valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *diccionario_implementacion[K, V]) Borrar(clave K) V {
	ubicacion := hashear[K](clave) % uint64(len(dicc.array))
	for i := ubicacion; i < (ubicacion + 3); i++ {
		if dicc.array[i].clave == clave {
			devolver := dicc.array[i].valor
			dicc.array[i].conValor = false
			dicc.largo--
			return devolver
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *diccionario_implementacion[K, V]) Cantidad() int {
	return dicc.largo
}
