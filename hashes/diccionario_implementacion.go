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

func buscarPrimo(inicio uint64) uint64 {
	const rango uint64 = 25

	for r := inicio; r < inicio+rango; r++ {
		isPrimo := true
		for k := 2; k <= 19; k++ {
			if (r % uint64(k)) == 0 {
				isPrimo = false
				break
			}
		}
		if isPrimo {
			return r
		}
	}
	return inicio
}

func (dicc *diccionario_implementacion[K, V]) redimensionar(nuevoTam uint64) {
	nuevoArray := make([]elementos[K, V], nuevoTam)
	for _, elemento := range dicc.array {
		if elemento.conValor {
			nuevoIndice := hashear[K](elemento.clave) % nuevoTam
			nuevoElemento := crearElemento[K, V](elemento.clave, elemento.valor, nuevoIndice)

			nuevoArray[nuevoIndice] = *nuevoElemento
		}
	}
	dicc.array = nuevoArray
}

func (dicc *diccionario_implementacion[K, V]) hacerEspacio(indice uint64) (uint64, bool) {

	var espacio uint64
	if dicc.array[indice].ubicacion > indice {
		espacio = dicc.array[indice].ubicacion - indice
	} else if dicc.array[indice].ubicacion == indice {
		espacio = 3
	}

	if indice+espacio > uint64(len(dicc.array))-1 {
		var indiceABuscar uint64
		switch {
		case indice == uint64(len(dicc.array))-2:
			for i := 1; i <= 3; i++ {
				switch i {
				case 1:
					indiceABuscar = indice
				case 2:
					indiceABuscar = indice + 1
				case 3:
					indiceABuscar = 0
				}

				if dicc.array[indiceABuscar].conValor == false {
					dicc.array[i].ubicacion = indiceABuscar
					return indiceABuscar, false
				}
			}

		case indice == uint64(len(dicc.array))-1:

			for i := 1; i <= 3; i++ {
				switch i {
				case 1:
					indiceABuscar = indice
				case 2:
					indiceABuscar = 0
				case 3:
					indiceABuscar = 1
				}

				if dicc.array[indiceABuscar].conValor == false {
					dicc.array[i].ubicacion = indiceABuscar
					return indiceABuscar, false
				}
			}
		}
		return 0, true
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
		tamNuevo := buscarPrimo(uint64(len(dicc.array)) * 2)
		dicc.redimensionar(tamNuevo)
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
