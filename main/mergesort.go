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

func Merge(izquierda,derecha []int)[]int{
	array := make([]int, len(izquierda)+len(derecha))
	k:=0
	i:=0
	j:=0
	for i < len(izquierda) && j < len(derecha){
		if izquierda[i] <= derecha[j]{
			array[k] = izquierda[i]
			i++
		} else if izquierda[i] > derecha[j] {
			array[k] = derecha[j]
			j++
		}
		k++
	}
	for i < len(izquierda){
		array[k] = izquierda[i]
		i++
		k++
	}
	for j < len(derecha){
		array[k] = derecha[j]
		j++
		k++
	}
	return array

}


func Mergesort(arreglo []int) []int{
	if len(arreglo) == 1{
		return arreglo
	}
	mitad := len(arreglo)/2
	izquierda := Mergesort(arreglo[:mitad])
	derecha := Mergesort(arreglo[mitad:])
	return Merge(izquierda, derecha)
}