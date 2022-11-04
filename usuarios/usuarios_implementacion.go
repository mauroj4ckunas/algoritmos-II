package usuarios

type usuarioImplementacion struct {
	nivel		int 
	feed		ColaPrioridad[post]
	logueado	bool
}