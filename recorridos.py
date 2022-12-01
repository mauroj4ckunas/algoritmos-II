import cola as Tda
from heap import Heap
from grafo import Grafo

def bfs_generico(grafo: Grafo):
	padres = {}
	visitados = set()
	orden = {}
	for v in grafo.verVertices():
		if v not in visitados:
			padres[v] = None
			orden[v] = 0
			visitados.add(v)
			bfs(grafo,v,padres,visitados,orden)

def bfs(grafo: Grafo,origen,padres,visitados,orden):
	cola = Tda.Cola()
	cola.Encolar(origen)
	while not cola.EstaVacia() :
		vertice = cola.Desencolar()
		for adyacente in grafo.adyacentes(vertice):
			if adyacente not in visitados:
				padres[adyacente] = vertice
				orden[adyacente] = orden[vertice] + 1
				visitados.add(adyacente)
				cola.Encolar(adyacente)

def dfs(grafo: Grafo):
	padres = {}
	visitados = set()
	orden = {}
	cant_comp = 0
	for v in grafo.verVertices():
		if v not in visitados:
			padres[v] = None
			cant_comp += 1
			orden[v] = 0
			visitados.add(v)
			_dfs(grafo,v,padres,visitados,orden)
	return padres, cant_comp
			
def _dfs(grafo: Grafo,vertice,padres,visitados: set,orden):
	for adyacente in grafo.adyacentes(vertice):
		if adyacente not in visitados:
			padres[adyacente] = vertice
			orden[adyacente] = orden[vertice] + 1
			visitados.add(adyacente)
			_dfs(grafo,adyacente,padres,visitados,orden)


def gradosNoDirigido(grafo: Grafo) -> dict:
	cola = Tda.Cola()
	origen = grafo.verticeAlAzar()
	visitados = set()
	grados = {}
	cola.Encolar(origen)
	while not cola.EstaVacia():
		v = cola.Desencolar()
		grados[v] = len(grafo.adyacentes(v))
		for adyacente in grafo.adyacentes(v):
			if adyacente not in visitados:
				visitados.add(adyacente)
				cola.Encolar(adyacente)
	return grados			
			
			
			
#Orden topologico:

def bfsordenadoentrada(grafo: Grafo):
	grado = {}
	cola = Tda.Cola()
	orden = []
	for v in grafo.verVertices():
		grado[v] = 0

	for v in grafo.verVertices():
		for w in grado.adyacentes(v):
			grado[w] += 1

	for v in grafo.verVertices():
		if grado[v]==0:
			cola.Encolar(v)

	while not cola.EstaVacia():
		v = cola.Desencolar()
		orden.append(v)
		for w in grafo.adyacentes(v):
			grado[w] -= 1
			if grado[w]==0:
				cola.Encolar(w)

	return orden


"""def dfsorden(grafo):
	pila = Pila()
	visitados = set()

	for v in grafo.vertices():
		if v not in visitados:
			visitados.add(v)
			_dfsorden(grafo,visitados,pila,v)

	result = []
	while not pila.EstaVacia():
		result.append(pila.Desapilar())

	return result


def _dfsorden(grafo,visitados,pila,vertice):
	for w in grafo.adyacentes(vertice):
		if w not in visitados:
			visitados.add(w)
			_dfsorden(grafo,visitados,pila,w)
	pila.Apilar(vertice)"""


#Camino minimo:
		
def dijkstra(grafo: Grafo,origen):
	distancia = {}
	padres = {}
	for v in grafo.verVertices():
		distancia[v] = float("inf")
	distancia[origen] = 0
	padres[origen] = None
	cola = Heap()
	cola.Encolar((origen,distancia[origen]))
	while not cola.EstaVacia():
		v , _ = cola.Desencolar()
		for w in grafo.adyacentes(v):
			if distancia[v] + grafo.peso(v,w) < distancia[w]:
				distancia[w] = distancia[v] + grafo.peso(v,w)
				padres[w] = v 
				cola.Encolar(w,distancia[v] + grafo.peso(v,w))

	return distancia,padres

def belmanford(grafo: Grafo,origen):
	distancia = {}
	padres = {}
	for v in grafo.verVertices():
		distancia[v] = float("inf")
	distancia[origen] = 0
	padres[origen] = None
	aristas = []
	visitado = set()
	for v in grafo.verVertices():
		for w in grafo.adyacentes(v):
			if (v,w) not in visitado:
				visitado.add((v,w))
				aristas.append((v,w,grafo.peso(v,w)))

	for i in range(len(grafo.verVertices())):
		cambio = False
		for origen,destino,peso in aristas:
			if distancia[origen]+peso < distancia[destino]:
				cambio = True
				padres[destino] = origen
				distancia[destino] = distancia[origen] + peso
		if not cambio:
			break

	for v,w,peso in aristas:
		if distancia[v] + peso < distancia[w]:
			raise Exception("Hay un ciclo")

	return padres,distancia


#arbol tendido minimo:

def prim(grafo: Grafo,origen):
	visitados = set()
	visitados.add(origen)
	cola = Heap()
	for w in grafo.adyacentes(origen):
		cola.Encolar((origen,w,grafo.peso(origen,w)))

	arbol = Grafo()
	for v in grafo.verVertices():
		arbol.AgregarVertice(v)

	while not cola.EstaVacia():
		v,w,peso = cola.Desencolar()
		if w not in visitados:
			arbol.agregararista(v,w,grafo.peso(v,w))
			visitados.add(w)
			for x in grafo.adyacentes(w):
				if x not in visitados:
					cola.Encolar((w,x,grafo.peso(w,x)))
	return arbol

class UnionFind:
	
	def __init__(self,vertices):
		self.groups = {}
		for x in vertices:
			self.groups[x] = x

	def union(self,vertice1,vertice2):
		new_group = self.find(vertice1)
		other = self.find(vertice2)
		self.groups[other] = new_group

	def find(self,vertice):
		if self.groups[vertice] == vertice:
			return vertice

		real_groups = self.find(self.groups[vertice])
		self.groups[vertice] = real_groups
		return real_groups



def kruskal(grafo: Grafo):
	conjunto = UnionFind(grafo.verVertices())
	aristas = []
	visitados = set()
	for v in grafo.verVertices():
		if v not in visitados: visitados.add(v)
		for w in grafo.adyacentes(v):
			if w not in visitados:
				aristas.append((v,w,grafo.peso(v,w)))

	arbol = Grafo()
	for v in grafo.verVertices():
		arbol.AgregarVertice(v)

	aristas = sorted(aristas)

	for a in aristas:
		v,w,peso = a
		if conjunto.find(v) != conjunto.find(w):
			arbol.agregararista(v,w,peso)
			conjunto.union(v,w)

	return arbol
