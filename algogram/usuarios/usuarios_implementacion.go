package usuarios

type post[T comparable] struct {
	prioridadPosteo	int 
	posteado		[]T
	id				int
}

type usuarioImplementacion[T comparable] struct {
	nivel		int 
	feed		ColaPrioridad[post[T]]
}

func CrearPosteo[T comparable](prioridadPost int,posteo []T,id){
	post:= new(post[T])
	post.prioridadPosteo = prioridadPost
	post.posteado = posteo
	post.id = id
}

func CrearUsuario[T comparable](prioridadUsuario int, comparadorPosteos func(post[T],post[T])int) Usuario[T] {
	usuario := new(usuarioImplementacion[T])
	usuario.nivel = prioridadUsuario
	usuario.feed = CrearHeap[post[T]](comparadorPosteos)
	return usuario
}

func (usu *usuarioImplementacion[T]) Prioridad() int {
	return usu.nivel
}

func (usu *usuarioImplementacion[T]) Publicar(posteo post[T]) {
	usu.feed.Encolar(posteo)
}