package diccionario

import (
	TDApila "diccionario/Pila"
)

type hojas[K comparable, V any] struct{

	clave		K 
	valor		V
	hijoIzq		*hojas[K,V]
	hijoDer		*hojas[K,V]
	 
}

func crearHoja[K comparable, V any] (clave K, dato V) *hojas[K, V] {

	hojaNueva := new(hojas[K, V])
	hojaNueva.clave = clave 
	hojaNueva.valor = dato
	return hojaNueva
	
}


func (hoja *hojas[K,V]) borrar(compara func(K, K) int ,clave K ) (**hojas[K,V],string) {

	if compara(clave,hoja.clave) < 0 {
		if hoja.hijoIzq == nil {

			return nil,"La clave no pertenece al diccionario"

		} else if compara(clave,hoja.hijoIzq.clave) == 0 {
			
			return &hoja.hijoIzq,""

		}

		return hoja.hijoIzq.borrar(compara,clave)

	} else {

		if hoja.hijoDer == nil{

			return nil,"La clave no pertenece al diccionario"

		} else if compara(clave,hoja.hijoDer.clave) == 0{

			return &hoja.hijoDer,""

		}

		return hoja.hijoDer.borrar(compara,clave)
	}
		
		
}

func (arbol *arbolBinario[K,V]) Borrar(clave K) V {

	var err string
	var borrar **hojas[K,V]

	if arbol.raiz == nil {

		panic("La clave no pertenece al diccionario")

	} else if arbol.raiz.clave == clave {

		borrar = &arbol.raiz

	} else {

		borrar, err = arbol.raiz.borrar(arbol.comparador,clave)

	}
	

	if borrar == nil {

		panic(err)

	}

	arbol.cantidad--

	var devolver V

	for true {

		if (*(*borrar)).hijoDer == nil && (*(*borrar)).hijoIzq == nil {

		
			devolver = (*borrar).valor
			*borrar = nil
			return devolver

		}

		//caso 1 hijo
		if (*(*borrar)).hijoDer == nil && (*(*borrar)).hijoIzq != nil {

		
			devolver = (*borrar).valor
			*borrar = (*borrar).hijoIzq
			return devolver

		} else if (*(*borrar)).hijoDer != nil && (*(*borrar)).hijoIzq == nil {

		
			devolver = (*borrar).valor
			*borrar = (*borrar).hijoDer
			return devolver

		}

		//caso 2 hijos
		if (*(*borrar)).hijoDer != nil && (*(*borrar)).hijoIzq != nil {

		
			reemplazante := &(*borrar).hijoDer

			for (*reemplazante).hijoIzq != nil {
				
				reemplazante = &(*reemplazante).hijoIzq

			}

			(*borrar).clave , (*reemplazante).clave = (*reemplazante).clave , (*borrar).clave
			(*borrar).valor , (*reemplazante).valor = (*reemplazante).valor , (*borrar).valor
			borrar = reemplazante
		}
	}

	return devolver

}

func (hoja *hojas[K,V]) guardar(compara func(K, K) int ,hojaNueva *hojas[K,V],cantidad *int) {
	resultado := compara(hojaNueva.clave,hoja.clave)
	switch {

	case resultado < 0:

		if hoja.hijoIzq == nil {

			hoja.hijoIzq = hojaNueva
			*cantidad++

		}

		hoja.hijoIzq.guardar(compara,hojaNueva,cantidad)

	case resultado > 0:

		if hoja.hijoDer == nil {

			hoja.hijoDer = hojaNueva
			*cantidad++

		}

		hoja.hijoDer.guardar(compara,hojaNueva,cantidad)

	default:

		hoja.valor = hojaNueva.valor
		
	}
}

func (hoja *hojas[K,V]) encontrarClave(compara func(K, K) int ,clave K ) (*hojas[K,V],string) {

	if hoja == nil {

		return nil,"La clave no pertenece al diccionario"

	}

	resultado := compara(clave,hoja.clave)

	switch {

	case resultado < 0:

		if hoja.hijoIzq == nil {

			return nil,"La clave no pertenece al diccionario"

		}

		return hoja.hijoIzq.encontrarClave(compara,clave)

		

	case resultado > 0:

		if hoja.hijoDer == nil {

			return nil,"La clave no pertenece al diccionario"

		}

		return hoja.hijoDer.encontrarClave(compara,clave)
		
	default :

		return hoja , ""

	}
}


func (hoja *hojas[K,V]) iterar(f func(clave K, dato V) bool, hasta *hojas[K,V]) {

	if hasta != nil{
		if hoja.clave == hasta.clave {

			return

		}
	}
	

	if hoja.hijoIzq != nil {

		hoja.hijoIzq.iterar(f,hasta)

	}

	switch f(hoja.clave,hoja.valor) {

	case true:

		if hoja.hijoDer != nil {

			hoja.hijoDer.iterar(f,hasta)

		}

	default:

		return 
		
	}
}


type arbolBinario[K comparable, V any] struct {

	raiz		*hojas[K,V]
	comparador	func(K, K) int
	cantidad	int

}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V]{

	arbol := new(arbolBinario[K,V])
	arbol.comparador = funcion_cmp
	return arbol

}


func (arbol *arbolBinario[K,V]) Guardar(clave K, dato V){

	hoja := crearHoja[K,V](clave,dato)
	

	if arbol.raiz == nil {

		arbol.raiz = hoja
		arbol.cantidad++

	} else {

		arbol.raiz.guardar(arbol.comparador,hoja,&arbol.cantidad)

	}


}

func (arbol *arbolBinario[K,V]) Pertenece(clave K) bool{

	resultado, _ := arbol.raiz.encontrarClave(arbol.comparador,clave)
	return resultado != nil

}

func (arbol *arbolBinario[K,V]) Obtener(clave K) V {

	resultado, err := arbol.raiz.encontrarClave(arbol.comparador,clave)

	if resultado == nil {

		panic(err)

	}

	return resultado.valor
}


func (arbol *arbolBinario[K,V]) Cantidad() int{
	return arbol.cantidad
}


func (arbol *arbolBinario[K,V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

	var err string
	var empieza *hojas[K,V]
	var termina *hojas[K,V]

	if desde != nil {

		empieza, err = arbol.raiz.encontrarClave(arbol.comparador,*desde)

	} else {

		empieza = arbol.raiz

	}

	if err != "" {

		panic(err)

	}

	if hasta != nil {

		termina, err = arbol.raiz.encontrarClave(arbol.comparador,*hasta)

	}

	if err != "" {

		panic(err)

	}

	if empieza != nil {

		empieza.iterar(visitar,termina)

	}
	
}

func (arbol *arbolBinario[K,V]) Iterar(f func(clave K, dato V) bool){
	
	arbol.IterarRango(nil,nil,f)

}

type iterExterno[K comparable, V any] struct {

	pilaRecursiva	TDApila.Pila[*hojas[K,V]]
	hasta			K

}

func (arbol *arbolBinario[K,V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

	iterr := new(iterExterno[K,V])
	iterr.pilaRecursiva = TDApila.CrearPilaDinamica[*hojas[K,V]]()
	var err string
	var todoIzquierda *hojas[K,V]
	var termina *hojas[K,V]

	if desde != nil {

		todoIzquierda, err = arbol.raiz.encontrarClave(arbol.comparador,*desde)

	} else {

		todoIzquierda = arbol.raiz

	}

	if err != "" {

		panic(err)
		
	}

	if hasta != nil {

		termina, err = arbol.raiz.encontrarClave(arbol.comparador,*hasta)
		if err != "" {

			panic(err)
		
		}
		iterr.hasta = termina.clave
		
	}

	

	for todoIzquierda != nil && todoIzquierda.clave != *hasta {

		iterr.pilaRecursiva.Apilar(todoIzquierda)
		todoIzquierda = todoIzquierda.hijoIzq

	}

	return iterr
}

func (arbol *arbolBinario[K,V]) Iterador() IterDiccionario[K, V]{

	return arbol.IteradorRango(nil,nil)

}

func (iterr *iterExterno[K,V]) HaySiguiente() bool {

	return !iterr.pilaRecursiva.EstaVacia()

}

func (iterr *iterExterno[K,V]) VerActual() (K, V) {

	if !iterr.HaySiguiente(){

		panic("El iterador termino de iterar")

	}

	return iterr.pilaRecursiva.VerTope().clave , iterr.pilaRecursiva.VerTope().valor
}

func (iterr *iterExterno[K,V]) Siguiente() K {

	if !iterr.HaySiguiente(){

		panic("El iterador termino de iterar")

	}

	devolver := iterr.pilaRecursiva.Desapilar()
	if devolver.hijoDer != nil {

		todoIzquierda := devolver.hijoDer

		for todoIzquierda != nil && todoIzquierda.clave != iterr.hasta {

			iterr.pilaRecursiva.Apilar(todoIzquierda)
			todoIzquierda = todoIzquierda.hijoIzq

		}

	}

	return devolver.clave

}