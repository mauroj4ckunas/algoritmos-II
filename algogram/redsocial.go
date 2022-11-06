package main

type AlgoGram[T comparable,V comparable] interface {

	Login(usuario string) string

	Logout() string

	Publicar(posteo []T)
}
