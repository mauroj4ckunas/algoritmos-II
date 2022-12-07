import cola as Tda
from grafo import Grafo
import heapq
from pila import Pila

def dfs(grafo: Grafo): #se usa
	padres = {}
	visitados = set()
	cant_comp = 0
	for v in grafo.verVertices():
		if v not in visitados:
			padres[v] = None
			cant_comp += 1
			visitados.add(v)
			_dfs(grafo,v,padres,visitados,orden)
	return padres, cant_comp
			
def _dfs(grafo: Grafo,vertice,padres,visitados):
	for adyacente in grafo.adyacentes(vertice):
		if adyacente not in visitados:
			padres[adyacente] = vertice
			visitados.add(adyacente)
			_dfs(grafo,adyacente,padres,visitados,orden)

def grados(grafo: Grafo) -> dict:
	grados = {}
	for v in grafo.verVertices():
		grados[v] = 0
	for v in grafo.verVertices():
		for w in grafo.adyacentes(v):
			grados[w] += 1
	return grados

#Orden topologico:

def bfsordenadoentrada(grafo: Grafo): # se usa
	cola = Tda.Cola()
	orden = []
	grado = grados(grafo)

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
	noHayCiclo = True
	for v in grafo.verVertices():
		if grado[v] != 0:
			noHayCiclo = False
			break

	return orden, noHayCiclo


#Camino minimo:

def reconstruirCaminoMinimo(grafo: Grafo,origen,destino):
	dist, padres = dijkstra(grafo, origen,destino)
	pila = Pila()
	actual = destino
	pila.Apilar(actual)
	while True:
		actual = padres[actual]
		pila.Apilar(actual)
		if actual == origen: 
			break
	camino = [None] * pila.Cantidad()
	i = 0
	while not pila.EstaVacia():
		camino[i] = pila.Desapilar()
		i += 1
	return camino,dist

		
def dijkstra(grafo: Grafo,origen,destino = None): #se usa
	distancia = {}
	padres = {}
	for v in grafo.verVertices():
		distancia[v] = float("inf")
	distancia[origen] = 0
	padres[origen] = None
	cola = []
	heapq.heappush(cola, (distancia[origen], origen))
	while len(cola) != 0:
		tuplaV = heapq.heappop(cola)
		v = tuplaV[1]
		if destino != None and v == destino:
			return distancia,padres
		for w in grafo.adyacentes(v):
			if distancia[v] + int(grafo.peso(v,w)) < distancia[w]:
				distancia[w] = distancia[v] + int(grafo.peso(v,w))
				padres[w] = v 
				heapq.heappush(cola, (distancia[v] + int(grafo.peso(v,w)), w))

	return distancia,padres


#arbol tendido minimo:

def verAristas(grafo: Grafo): #se usa 
	aristas = []
	visitados = set()
	for v in grafo.verVertices():
		for w in grafo.adyacentes(v):
			if (v, w) not in visitados and (w, v) not in visitados:
				visitados.add((v, w))
				aristas.append((v, w, grafo.peso(v, w)))
	return aristas
	

def prim(grafo: Grafo): # se usa
	origen = grafo.verticeAlAzar()
	visitados = set()
	visitados.add(origen)
	pesoTotal = 0
	cola = []
	for w in grafo.adyacentes(origen):
		heapq.heappush(cola, (int(grafo.peso(origen,w)), (origen,w)))

	arbol = Grafo()
	for v in grafo.verVertices():
		arbol.agregarVertice(v)

	while len(cola) != 0:
		peso, arista = heapq.heappop(cola)
		v = arista[0]
		w = arista[1]
		if w in visitados:
			continue
		arbol.agregarArista(v,w,grafo.peso(v,w))
		pesoTotal += int(grafo.peso(v,w))
		visitados.add(w)
		for x in grafo.adyacentes(w):
			if x not in visitados:
				heapq.heappush(cola, (int(grafo.peso(w,x)), (w,x)))
	
	return arbol, pesoTotal