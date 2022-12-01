import cola as Tda
from heap import Heap
from grafo import Grafo

def bfs_generico(grafo):
	padres = {}
	visitados = set()
	orden = {}
	for v in grafo.vertices():
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
		for adyacente in grafo.adyacentes(vertice):
			if adyacente not in visitados:
				padres[adyacente] = vertice
				orden[adyacente] = orden[vertice] + 1
				visitados.add(adyacente)
				cola.Encolar(adyacente)

def dfs(grafo):
	padres = {}
	visitados = set()
	orden = {}
	cant_comp = 0
	for v in grafo.vertices():
		if v not in visitados:
			padres[v] = None
			cant_comp += 1
			orden[v] = 0
			visitados.add(v)
			_dfs(grafo,v,padres,visitados,orden)
	return padres, cant_comp
			
def _dfs(grafo,vertice,padres,visitados,orden):
	for adyacente in grafo.adyacentes(vertice):
		if adyacente not in visitados:
			padres[adyacente] = vertice
			orden[adyacente] = orden[vertice] + 1
			visitados.add(adyacente)
			_dfs(grafo,adyacente,padres,visitados,orden)


def gradosNoDirigido(grafo: Grafo, dirigido: bool) -> dict:
	cola = Tda.Cola()
	origen = grafo.verticeAlAzar()
	visitados = set()
	grados = {}
	cola.Encolar(origen)
	while not cola.EstaVacia():
		v = cola.Desencolar()
		grados[v] = grafo.adyacentes(v)
		for adyacente in grafo.adyacentes(v):
			if adyacente not in visitados:
				visitados.add(adyacente)
				cola.Encolar(adyacente)
	return grados			
			
			
			
#Orden topologico:

def bfsordenadoentrada(grafo):
	grado = {}
	cola = Tda.Cola()
	orden = []
	for v in grafo.vertices():
		grado[v] = 0

	for v in grado.vertices():
		for w in grado.adyacentes(v):
			grado[w] += 1

	for v in grafo.vertices():
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
		
def dijkstra(grafo,origen):
	distancia = {}
	padres = {}
	for v in grafo.vertices():
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
				padre[w] = v 
				cola.Encolar(w,distancia[v] + grafo.peso(v,w))

	return distancia,padres

def belmanford(grafo,origen):
	distancia = {}
	padres = {}
	for v in grafo.vertices():
		distancia[v] = float("inf")
	distancia[origen] = 0
	padres[origen] = None
	aristas = []
	visitado = set()
	for v in grafo.vertices():
		for w in adyacentes():
			if (v,w) not in visitado:
				visitado.add((v,w))
				aristas.append((v,w,grafo.peso(v,w)))

	for i in range(len(grafo.vertices())):
		cambio = False
		for origen,destino,peso in aristas:
			if distancia[origen]+peso < distancia[destino]:
				cambio = True
				padre[destino] = origen
				distancia[destino] = distancia[origen] + peso
		if not cambio:
			break

	for v,w,peso in aristas:
		if distancia[v] + peso < distancia[w]:
			raise Exception("Hay un ciclo")

	return padre,distancia


#arbol tendido minimo:

def prim(grafo,origen):
	visitados = set()
	visitados.add(origen)
	cola = Heap()
	for w in grafo.adyacentes(origen):
		cola.Encolar((origen,w,grafo.peso(origen,w)))

	arbol = Grafo()
	for v in grafo.vertices():
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



def kruskal(grafo):
	conjunto = UnionFind(grafo.vertices())
	aristas = []
	visitados = set()
	for v in grafo.vertices():
		if v not in visitados: visitados.add(v)
		for w in grafo.adyacentes(v)
			if w not in visitados:
				aristas.append((v,w,grafo.peso(v,w)))

	arbol = Grafo()
	for v in grafo.vertices():
		arbol.AgregarVertice(v)

	aristas = sorted(aristas)

	for a in aristas:
		v,w,peso = a
		if conjuntos.find(v) != conjuntos.find(w):
			arbol.agregararista(v,w,peso)
			conjuntos.union(v,w)

	return arbol
