Fallos/Errores: 

En algunos casos al ingresar el DNI dice que esta fuera del Padron, cuando esta. Por ejemplo en el test 4, al ingresar 76567790 me salta error. En cambio el 12345678 lo toma.


Test 1:
	Devuelve Error: Faltan Parametros
	Deberia devolver ERROR: Lectura de archivos
	
	
Test 2:
	Todo bien

Test 3:
	Todo bien

Test 4:
	Todo bien
	
Test 5:
	Error con la asignacion de votos (Hay gente q recibio votos que era para otra persona) y error En los votos impugnados
	
	imprime esto :
		Presidente: 
		Votos en Blanco: 2 votos
		Traperos Unidos - Duki: 2 votos
		Rockeros para todos - Charlie: 1 voto
		Pop para divertirse - Tini: 0 votos
		Gobernador: 
		Votos en Blanco: 1 voto
		Traperos Unidos - L-Gante: 0 votos
		Rockeros para todos - El Flaco: 3 votos
		Pop para divertirse - Ale Sergi: 1 voto
		Intendente: 
		Votos en Blanco: 2 votos
		Traperos Unidos - María Becerra: 1 voto
		Rockeros para todos - El Indio: 1 voto
		Pop para divertirse - Abel: 1 voto
		Votos Impugnados: 0 
	
	deberia imprimir esto: 
		Presidente:
		Votos en Blanco: 3 votos
		Traperos Unidos - Duki: 1 voto
		Rockeros para todos - Charlie: 1 voto
		Pop para divertirse - Tini: 0 votos

		Gobernador:
		Votos en Blanco: 1 voto
		Traperos Unidos - L-Gante: 1 voto
		Rockeros para todos - El Flaco: 2 votos
		Pop para divertirse - Ale Sergi: 1 voto

		Intendente:
		Votos en Blanco: 2 votos
		Traperos Unidos - María Becerra: 1 voto
		Rockeros para todos - El Indio: 1 voto
		Pop para divertirse - Abel: 1 voto

		Votos Impugnados: 1 voto
