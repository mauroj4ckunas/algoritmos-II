package lista

type IteradorLista[T any] interface {

	// Estos metodos no deben mezclarse con los metodos de la interfaz Lista.
	// En caso contrario puede producir errores o cambios no queridos.

	// VerActual obtiene el valor del elemento en el que este parado el iterador y lo devuelve.
	VerActual() T

	// HaySiguiente devuelve verdadero si en el siguiente elemento no es vacio, caso contrario
	// sera falso.
	HaySiguiente() bool

	// Siguiente hace que el iterador pase al siguiente elemento, devolviendo el valor del elemento
	// actual antes de pasar al siguiente.
	Siguiente() T

	// Insertar agrega en la posicion anterior al actual un nuevo elemento en la lista original.
	Insertar(T)

	// Borrar saca el elemento de la lista original. Devuelve su valor. Y nos posiciona en el siguiente elemento.
	Borrar() T
}

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, y falso si no tiene.
	EstaVacia() bool

	// InsertarPrimero agrega un elemento en la primera posicion de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un elemento en la ultima posicion de la lista.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista, devolviendo su valor. En caso en que
	// estuviera vacia al momento de usarla entra en panico con el mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero devuelve el valor del primer elemento de la lista. En caso en que
	// estuviera vacia al momento de usarla entra en panico con el mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo devuelve el valor del ultimo elemento de la lista. En caso en que
	// estuviera vacia al momento de usarla entra en panico con el mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo develve el valor del largo de la lista.
	Largo() int

	// Iterar nos permite aplicarle una funcion que queramos uno por uno a cada elemento de la lista.
	Iterar(visitar func(*T) bool)

	// Iterador devuelve un iterador que nos permite recorrer todos los elementos de la lista.
	Iterador() IteradorLista[T]
}
