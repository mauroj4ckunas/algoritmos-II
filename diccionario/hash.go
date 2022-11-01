package diccionario

const (
	POSICIONES_HABILES         uint64 = 3
	MULTIPLICADOR_DE_CAPACIDAD int    = 3
)

type elementos[K comparable, V any] struct {
	clave     K
	valor     V
	ubicacion uint64
}

type tablaDeHash[K comparable, V any] struct {
	array           []*elementos[K, V]
	largo           int
	cap_diccionario uint64
}

func crearElemento[K comparable, V any](clave K, valor V, ubicacionHash uint64) *elementos[K, V] {
	dato := new(elementos[K, V])
	(*dato).clave = clave
	(*dato).valor = valor
	(*dato).ubicacion = ubicacionHash
	return dato
}

func (dicc *tablaDeHash[K, V]) redimensionar(nuevoTam int) {
	dicc.cap_diccionario = uint64(nuevoTam)
	arrayViejo := dicc.array
	dicc.array = crearArrayHash[K, V](dicc.cap_diccionario)
	dicc.largo = 0
	for i := 0; i < len(arrayViejo); i++ {
		if arrayViejo[i] != nil {
			dicc.Guardar((*arrayViejo[i]).clave, (*arrayViejo[i]).valor)
		}
	}
}

func (dicc *tablaDeHash[K, V]) hacerEspacio(indice uint64, lugarNecesario uint64) (uint64, bool) {

	capacidad := dicc.cap_diccionario
	if (capacidad+indice)%capacidad < (capacidad+lugarNecesario+POSICIONES_HABILES)%capacidad {
		return indice % capacidad, false
	}

	if (*dicc.array[(capacidad+indice-2)%capacidad]).ubicacion+1 == indice ||
		(*dicc.array[(capacidad+indice-2)%capacidad]).ubicacion+2 == indice {

		dicc.array[(capacidad+indice-2)%capacidad], dicc.array[indice] = dicc.array[indice], dicc.array[(capacidad+indice-2)%capacidad]

		return dicc.hacerEspacio(((capacidad + indice - 2) % capacidad), lugarNecesario)

	} else if (*dicc.array[(capacidad+indice-1)%capacidad]).ubicacion+1 == indice ||
		(*dicc.array[(capacidad+indice-1)%capacidad]).ubicacion+2 == indice {

		dicc.array[(capacidad+indice-1)%capacidad], dicc.array[indice] = dicc.array[indice], dicc.array[(capacidad+indice-1)%capacidad]

		return dicc.hacerEspacio(((capacidad + indice - 1) % capacidad), lugarNecesario)
	}

	dicc.redimensionar(cap(dicc.array) * MULTIPLICADOR_DE_CAPACIDAD)
	return 0, true
}

func (dicc *tablaDeHash[K, V]) Guardar(clave K, dato V) {

	capacidad := dicc.cap_diccionario
	indiceHash := Hashear(clave) % capacidad
	var posicion uint64

	pertenece, indice := dicc.perteneceElemento(clave)

	if pertenece {
		dicc.array[indice].valor = dato
		return
	} else {
		for posicion = indiceHash; posicion < (indiceHash + POSICIONES_HABILES); posicion++ {
			if dicc.array[posicion%capacidad] == nil {
				paraGuardar := crearElemento(clave, dato, indiceHash)
				dicc.array[posicion%capacidad] = paraGuardar
				dicc.largo++
				return
			}
		}
	}

	for true {
		if dicc.array[posicion%capacidad] == nil {
			break
		}
		posicion = (posicion + 1) % capacidad
	}

	posicion, redimension := dicc.hacerEspacio(posicion%capacidad, indiceHash)

	if redimension {
		dicc.Guardar(clave, dato)
	} else {
		paraGuardar := crearElemento(clave, dato, indiceHash)
		dicc.array[posicion] = paraGuardar
		dicc.largo++
	}

}

func (dicc *tablaDeHash[K, V]) perteneceElemento(clave K) (bool, int) {

	capacidad := dicc.cap_diccionario
	ubicacion := Hashear(clave) % capacidad

	for i := ubicacion; i < (ubicacion + POSICIONES_HABILES); i++ {
		if dicc.array[i%capacidad] == nil {
			continue
		} else if (*dicc.array[i%capacidad]).clave == clave {
			return true, int(i % capacidad)
		}
	}
	return false, -1
}

func (dicc *tablaDeHash[K, V]) Pertenece(clave K) bool {
	pertenece, _ := dicc.perteneceElemento(clave)
	return pertenece
}

func (dicc *tablaDeHash[K, V]) Obtener(clave K) V {
	pertenece, indice := dicc.perteneceElemento(clave)
	if pertenece {
		return (*dicc.array[indice]).valor
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *tablaDeHash[K, V]) Borrar(clave K) V {

	if dicc.Cantidad() < int(dicc.cap_diccionario)/2 && dicc.Cantidad() > int(dicc.cap_diccionario)/4 {
		dicc.redimensionar(int(dicc.cap_diccionario) / 2)
	}

	pertenece, indice := dicc.perteneceElemento(clave)
	if pertenece {
		devolver := (*dicc.array[indice]).valor
		dicc.array[indice] = nil
		dicc.largo--
		return devolver
	}
	panic("La clave no pertenece al diccionario")
}

func (dicc *tablaDeHash[K, V]) Cantidad() int {
	return dicc.largo
}

func crearArrayHash[K comparable, V any](tam uint64) []*elementos[K, V] {
	nuevoArray := make([]*elementos[K, V], tam)
	return nuevoArray
}

const CAPACIDAD_INICIAL uint64 = 87

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	diccio := new(tablaDeHash[K, V])
	diccio.cap_diccionario = CAPACIDAD_INICIAL
	(*diccio).array = crearArrayHash[K, V](CAPACIDAD_INICIAL)
	return diccio
}

// Iterador Externo
type iterTablaDeHash[K comparable, V any] struct {
	actual   int
	iter_arr []*elementos[K, V]
}

func (dicc *tablaDeHash[K, V]) Iterador() IterDiccionario[K, V] {
	iterr := new(iterTablaDeHash[K, V])
	iterr.iter_arr = dicc.array
	for iterr.actual < len(iterr.iter_arr) {
		if iterr.iter_arr[iterr.actual] != nil {
			break
		}
		iterr.actual++
	}
	return iterr
}

func (iterr *iterTablaDeHash[K, V]) HaySiguiente() bool {
	return iterr.actual < len(iterr.iter_arr)
}

func (iterr *iterTablaDeHash[K, V]) VerActual() (K, V) {
	if !iterr.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return (*iterr.iter_arr[iterr.actual]).clave, (*iterr.iter_arr[iterr.actual]).valor
}

func (iterr *iterTablaDeHash[K, V]) Siguiente() K {
	if !iterr.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	devolver := (*iterr.iter_arr[iterr.actual]).clave
	iterr.actual++
	for iterr.HaySiguiente() {
		if iterr.iter_arr[iterr.actual] != nil {
			break
		}
		iterr.actual++
	}
	return devolver
}

// Iterador Interno

func (dicc *tablaDeHash[K, V]) Iterar(f func(clave K, valor V) bool) {

	for i := 0; i < len(dicc.array); i++ {
		if dicc.array[i] == nil {
			continue
		}
		if !f(dicc.array[i].clave, dicc.array[i].valor) {
			break
		}
	}

}
