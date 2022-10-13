package diccionario

import (
	"fmt"
	"crypto/md5"
	"encoding/binary"
)
//funcion de hash
func hashear[K comparable](clave K) uint64 {
	elementoHasheable := []byte(fmt.Sprintf("%v", clave))
	hasheado := md5.Sum(elementoHasheable)
	var convertirAEntero []byte = hasheado[:]
	return binary.BigEndian.Uint64(convertirAEntero)
}

//celdas del diccionario
type elementos[K comparable, V any] struct {
	clave 		K
	valor 		V
	ubicacion 	int 
}
//creador de celdas
func crearElemento [K comparable, V any] (clave K, valor V,ubicacionHash int) elementos[K,V] {
	dato := new(elementos[K,V])
	dato.clave = clave
	dato.valor = valor
	dato.ubicacion = ubicacionHash
	return dato
}

//diccionario
type diccionario_implementacion[K comparable, V any] struct {
	array		[]elementos[K,V]
	largo		int
	factorCarga	float32
}

//creador del diccionario
func CrearDiccionario [K comparable, V any] () Diccionario[K, V] {
	diccio := new(diccionario_implementacion[K,V])
	diccio.array := make([]elementos[K,V],100)
	return diccio
}

//iterador externo del diccionario
type iterador_externo[K comparable, V any] struct {
	actual 		int
	dicc 		[]elementos[K,V]
}

//creador de iterador externo
func crearIteradorExterno [K comparable, V any] (dicc diccionario_implementacion[K,V]) IterDiccionario[K, V] {
	iterr := new(iterador_externo[K,V])
	iterr.dicc = dicc.array
	for iterr.actual < len(iterr.dicc){
		if dicc.array[iterr.actual] == nil{
			iterr.actual++
		}
		break
	}
	return iterr
}

func (dicc *diccionario_implementacion) Guardar(clave K, dato V) {
	dicc.factorCarga = dicc.Cantidad()/len(dicc.array)
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
/*
func (dicc *diccionario_implementacion) Iterador() IterDiccionario[K, V] {
	return crearIteradorExterno [K, V] (dicc)
}

func (iterr *iterador_externo) HaySiguiente() bool {
	return iterr.actual < len(iterr.dicc)
}

func (iterr *iterador_externo) VerActual() (K, V){
	if !iterr.HaySiguiente(){
		panic("El iterador termino de iterar")
	}
	return iterr.dicc[iterr.actual].clave , iterr.dicc[iterr.actual].valor
}

func (iterr *iterador_externo) Siguiente() K {
	if !iterr.HaySiguiente(){
		panic("El iterador termino de iterar")
	}
	devolver := iterr.dicc[iterr.actual].clave
	iterr.actual++
	for iterr.HaySiguiente() {
		if dicc.array[iterr.actual] == nil{
			iterr.actual++
		}
		break
	}
}*/