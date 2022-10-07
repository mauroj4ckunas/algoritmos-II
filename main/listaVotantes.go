package main 


import (
	TDAPila "Pila"
	Err "errores"
	Voto "votos"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PrepararListaVotantes(rutaPadrones string) []Voto.Votante, Err.Errores {
	pila := TDAPila.CrearPilaDinamica[int]() // esto es para usarlo para crear los array

	archivoVotantes, err := os.Open(rutaPadrones)
	defer archivoVotantes.Close()
	if err != nil { //si la ruta no se puede leer o algo, error
		ErrorLectura := new(Err.ErrorLeerArchivo)
		return []Voto.Votante{} , ErrorLectura
	}



	padron := bufio.NewScanner(archivoVotantes)
	cantidad_votantes := 0
  	for padron.Scan() {
  		pila.Apilar(strconv.Atoi(padron.Text()))
     	cantidad_votantes += 1
  	}

  	/*
  	err = padron.Err()
  	if err != nil {
     	fmt.Println(err)
  	}*/



  	array := make([]int,cantidad_votantes)
  	for i:= 0 ; i < cantidad_votantes ; i++ {
  		array[i] = pila.Desapilar()
  	}
  	array = Mergesort(array[:])



  	Votantes := make([]Voto.Votante,cantidad_votantes)
  	for j:= 0 ; j < len(array) ; j++ {
  		Votantes[j] = Voto.CrearVotante(array[j])
  	}

  	return Votantes, nil
}