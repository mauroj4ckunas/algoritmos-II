package diccionario

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

type elementos[K comparable, V any] struct {
	clave     K
	valor     V
	ubicacion int
}

type diccionario_implementacion[K comparable, V any] struct {
	array []elementos[K, V]
	largo int
}

func hashear[K comparable](clave K) int {
	elementoHasheable := []byte(fmt.Sprintf("%v", clave))
	hasheado := md5.Sum(elementoHasheable)
	var arrayUint64 []byte = hasheado[:]
	var convertirAEntero int = int(binary.BigEndian.Uint64(arrayUint64))
	if convertirAEntero < 0 {
		return convertirAEntero * -1
	}
	return convertirAEntero
}

func crearElemento[K comparable, V any](clave K, valor V, ubicacionHash int) *elementos[K, V] {
	dato := new(elementos[K, V])
	dato.clave = clave
	dato.valor = valor
	dato.ubicacion = ubicacionHash
	return dato

}

func (dicc *diccionario_implementacion[K, V]) Guardar(clave K, dato V) {
	ubicacion := hashear[K](clave) % len(dicc.array)
	paraGuardar := crearElemento[K, V](clave, dato, ubicacion)
	var i int
	for i = ubicacion; i < (ubicacion + 3); i++ {
		if dicc.array[i] == nil {
			dicc.array[i] = paraGuardar
			break
		}
	}

	if i == (ubicacion + 3) {
		// Buscar el mas cercano espacio libre
		// buscar 2 arriba del espacio libre si alguno puede moverse
		// si se puede se hace un swap y luego se fija si se puede poner el nuevo elemento
		//en caso de q no se pueda se vuelve a hacer lo mismo
		// y en caso de no poder mover nada se redimensiona
		// claveAMover := dicc.array[ubicacion+1].clave
		// indice := hashear[K]()

	}
}

func (dicc *diccionario_implementacion[K, V]) Pertenece(clave K) bool {
	ubicacion := hashear[K](clave) % len(dicc.array)
	for i := ubicacion; i < (ubicacion + 3); i++ {
		if dicc.array[i].clave == clave {
			return true
		}
	}
	return false
}

func (dicc *diccionario_implementacion[K, V]) Obtener(clave K) V {
	ubicacion := hashear[K](clave) % len(dicc.array)
	for i := ubicacion; i < (ubicacion + 3); i++ {
		if dicc.array[i].clave == clave {
			return dicc.array[i].valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

// func (dicc *diccionario_implementacion[K, V]) Borrar(clave K) V {
// 	ubicacion := hashear[K](clave) % len(dicc.array)
// 	for i := ubicacion; i < (ubicacion + 3); i++ {
// 		if dicc.array[i].clave == clave {
// 			devolver := dicc.array[i].valor
// 			dicc.array[i] = nil
// 			dicc.largo--
// 			return devolver
// 		}
// 	}
// 	panic("La clave no pertenece al diccionario")
// }

func (dicc *diccionario_implementacion[K, V]) Cantidad() int {
	return dicc.largo
}
