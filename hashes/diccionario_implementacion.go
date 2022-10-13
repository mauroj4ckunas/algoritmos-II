package diccionario

import (
	"fmt"
	"crypto/md5"
	"encoding/binary"
)

func hashear[K comparable](clave K) uint64 {
	elementoHasheable := []byte(fmt.Sprintf("%v", clave))
	hasheado := md5.Sum(elementoHasheable)
	var convertirAEntero []byte = hasheado[:]
	return binary.BigEndian.Uint64(convertirAEntero)
}


type elementos[K comparable, V any] struct {
	clave 		K
	valor 		V
	ubicacion 	int 
}

func crearElemento [K comparable, V any] (clave K, valor V,ubicacionHash int) elementos[K,V] {
	dato := new(elementos[K,V])
	dato.clave = clave
	dato.valor = valor
	dato.ubicacion = ubicacionHash
	return dato
}


type diccionario_implementacion[K comparable, V any] struct {
	array		[]elementos[K,V]
	largo		int
}

func (dicc *diccionario_implementacion) Guardar(clave K, dato V) {
	ubicacion := hashear[K](clave) % len(dicc.array)
	paraGuardar := crearElemento[K,V](clave,dato,ubicacion)
	var i int
	for i = ubicacion ; i < (ubicacion + 3) ; i++ {
		if dicc.array[i] == nil {
			dic.array[i] = paraGuardar
		}
	}

	if i == (ubicacion + 3){
		// Buscar el mas cercano espacio libre
		// buscar 2 arriba del espacio libre si alguno puede moverse
		// si se puede se hace un swap y luego se fija si se puede poner el nuevo elemento
		//en caso de q no se pueda se vuelve a hacer lo mismo
		// y en caso de no poder mover nada se redimensiona
	}
}

func (dicc *diccionario_implementacion) Pertenece(clave K) bool {
	ubicacion := hashear[K](clave) % len(dicc.array)
	for i = ubicacion ; i < (ubicacion + 3) ; i++ {
		if dicc.array[i].clave == clave {
			return true
		}
	}
	return false
}


func (dicc *diccionario_implementacion) Obtener(clave K) V {
	ubicacion := hashear[K](clave) % len(dicc.array)
	for i = ubicacion ; i < (ubicacion + 3) ; i++ {
		if dicc.array[i].clave == clave {
			return dicc.array[i].valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *diccionario_implementacion) Borrar(clave K) V {
	ubicacion := hashear[K](clave) % len(dicc.array)
	for i = ubicacion ; i < (ubicacion + 3) ; i++ {
		if dicc.array[i].clave == clave {
			devolver := dicc.array[i].valor
			dicc.array[i] = nil
			dicc.largo--
			return devolver
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *diccionario_implementacion) Cantidad() int {
	return dicc.largo
}