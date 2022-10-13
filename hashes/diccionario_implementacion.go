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


type elementos[K comparable, V any] struct{
	clave 	K
	valor 	V
}

func crearElemento [K comparable, V any] (clave K, valor V) elementos[K,V] {
	dato := new(elementos[K,V])
	dato.clave = clave
	dato.valor = valor
	return dato
}


type diccionario_implementacion[K comparable, V any] struct {
	array		[]elementos[K,V]
	largo		int
}

func (dicc *diccionario_implementacion) Guardar(clave K, dato V) {
	paraGuardar := crearElemento[K,V](clave,dato)
	ubicacion := hashear[K](clave) % len(dicc.array)
	var i int
	for i = ubicacion ; i < (ubicacion + 3) ; i++ {
		if dicc.array[i] == nil {
			dic.array[i] = paraGuardar
		}
	}

	if i == (ubicacion + 3){
		// implementar el hacer espacio y el redimensionar
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