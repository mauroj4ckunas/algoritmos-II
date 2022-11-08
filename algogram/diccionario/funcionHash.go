package diccionario

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

func Hashear[K comparable](clave K) uint64 {
	elementoHasheable := []byte(fmt.Sprintf("%v", clave))
	hasheado := sha1.Sum(elementoHasheable)
	var arrayUint64 []byte = hasheado[:]
	return binary.BigEndian.Uint64(arrayUint64)
}
