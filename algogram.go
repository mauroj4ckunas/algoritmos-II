package main


func crearRedSocial(nombreArchivo string) diccionario{
	archivoListas, err := os.Open(ruta)
	defer archivoListas.Close()
}

const(
	COMANDO1 = "login"
	COMANDO2 = "logout"
	COMANDO3 = "publicar"
)


func main(){
	archivoUsuarios := os.Args[1:]
	TDAAlgogram := crearRedSocial(archivoUsuarios)

	logeados := TDAcola.CrearCola()
	entradaUsuario := bufio.NewScanner(os.Stdin)
	for entradaUsuario.Scan() {
		comandos := strings.Split(entradaUsuario.Text(), " ")
		switch comandos[0] {
		case COMANDO1:
			usuario := comandos[1]
			err, elUsuario:= TDAAlgogram.Login(usuario)
			if err != nil{
				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
			}else{
				logeados.Encolar(elUsuario)
			}
			
		case COMANDO2:
			if !logeados.EstaVacia(){
				logeados.VerTope().Logout()
				logeados.Desencolar()
			} else {
				fmt.Fprintf(os.Stdout, "%s\n", new(err.TALERROR).Error())
			}

		case COMANDO3:
			post := comandos[1:]
			if !logeados.EstaVacia(){
				TDAAlgogram.Publicar(logeados.VerTope(),post)
				fmt.Fprintf(os.Stdout, "%s\n", "Post publicado")
			} else {
				fmt.Fprintf(os.Stdout, "%s\n", new(err.TALERROR).Error())
			}
		default:
			/* code */
			return
		}
	}
}