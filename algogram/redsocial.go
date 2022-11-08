package main

//TDAheap "algogram/Heap"

type AlgoGram interface {
	Login(usuario string) string
	Logout() string
	Publicar(posteo []string)
}
