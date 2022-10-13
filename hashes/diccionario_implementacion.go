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


type diccionario_implementacion[K comparable, V any] struct {
	array		[]elementos[K,V]
	largo		int
}