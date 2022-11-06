package usuarios

type post[T comparable] struct {
	prioridadPosteo	int 
	posteado		[]T
	id				int
}
func CrearPosteo[T comparable](prioridadPost int,posteo []T,id){
	post:= new(post[T])
	post.prioridadPosteo = prioridadPost
	post.posteado = posteo
	post.id = id
}

type usuarioImplementacion[T comparable] struct {
	nivel		int 
	feed		ColaPrioridad[post[T]]
	comparable	func(int,int) int
}

func CrearUsuario[T comparable](prioridadUsuario int, comparadorPosteos func(T, T) int, comparador func(int,int)int) Usuario {
	usuario := new(usuarioImplementacion[T])
	usuario.nivel = prioridadUsuario
	usuario.feed = CrearHeap[T](comparadorPosteos)
	usuario.comparable = comparador
	return usuario
}

func (usu *usuarioImplementacion[T]) Prioridad() int {
	return usu.nivel
}

func (usu *usuarioImplementacion[T]) PrioridadEntre(otroUsuario int) int {
	return usu.comparable(usu.nivel,otroUsuario)
}

func (usu *usuarioImplementacion[T]) Publicar(posteo post[T]) {
	usu.feed.Encolar(posteo)
}