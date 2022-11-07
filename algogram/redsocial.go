package main

//TDAheap "algogram/Heap"

type AlgoGram [T comparable,V any] interface {
	Login(usuario string) string
	Logout() string
	Publicar(posteo []string)
}
