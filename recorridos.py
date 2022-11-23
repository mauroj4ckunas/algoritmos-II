import cola as Tda

def bfs_generico(grafo):
	padres = {}
	visitados = set()
	orden = {}
	for v in grafo.LosVertices():
		if v not in visitados:
			padres[v] = None
			orden[v] = 0
			visitados.add(v)
			bfs(grafo,v,padres,visitados,orden)

def bfs(grafo,origen,padres,visitados,orden):
	cola = Tda.Cola()
	cola.Encolar(origen)
	while not cola.EstaVacia() :
		vertice = cola.Desencolar()
		for adyacente in grafo.AdyacentesDe(vertice):
			if adyacente not in visitados:
				padres[adyacente] = vertice
				orden[adyacente] = orden[vertice] + 1
				visitados.add(adyacente)
				cola.Encolar(adyacente)

def dfs(grafo):
	padres = {}
	visitados = set()
	orden = {}
	for v in grafo.LosVertices():
		if v not in visitados:
			padres[v] = None
			orden[v] = 0
			visitados.add(v)
			_dfs(grafo,v,padres,visitados,orden)

def _dfs(grafo,vertice,padres,visitados,orden):
	for adyacentes in grafo.AdyacentesDe(vertice):
		if adyacentes not in visitados:
			padres[adyacente] = vertice
			orden[adyacente] = orden[vertice] + 1
			visitados.add(adyacente)
			_dfs(grafo,adyacentes,padres,visitados,orden)
