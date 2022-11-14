package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const capacidad_inicial = 10
const cuadruple = 4

func (p *pilaDinamica[T]) redimensionar(capacidadNueva int) {
	/*duplica o divide la capacidad de un arreglo segun lo que se pase por parametro*/

	datos2 := make([]T, capacidadNueva)
	copy(datos2, p.datos)
	p.datos = datos2
}

func (p pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p pilaDinamica[T]) VerTope() T {

	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elem T) {

	if p.cantidad == cap(p.datos) {
		/*si la cantidad de elementos llega al tope del arreglo se duplica su capacidad*/
		p.redimensionar(cap(p.datos) * 2)
	}
	p.datos[p.cantidad] = elem
	p.cantidad += 1
}

func (p *pilaDinamica[T]) Desapilar() T {

	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	if p.cantidad*cuadruple <= cap(p.datos) {
		/*si la cantidad de elementos es 4 veces mas chica que la capacidad
		se divide a la mitad la capacidad para no usar memoria por demas*/
		p.redimensionar(cap(p.datos) / 2)
	}
	p.cantidad -= 1
	return p.datos[p.cantidad]
}

func CrearPilaDinamica[T any]() Pila[T] {
	p := new(pilaDinamica[T])
	p.datos = make([]T, capacidad_inicial)
	return p
}
