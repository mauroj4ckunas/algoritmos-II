import random

class Grafo:
	def __init__(self):
		self.vertices = {}

	def __str__(self):
		cadena = ""
		for v in self.vertices:
			cadena += f"{v}: ["
			for w in self.vertices[v]:
				cadena += f" {w} "
			cadena += f"]\n"
		return cadena[:len(cadena)-1]
		
	def AgregarVertice(self,vertice):
		self.vertices[vertice] = self.vertices.get(vertice,{})

	def SacarVertice(self,vertice):
		if vertice in self.vertices:
			for v in self.vertices.pop(vertice).keys():
				self.vertices[v].pop(vertice)
		else:
			raise Exception("No existe vertices")

	def SacarArista(self,desde,al):
		if desde in self.vertices and al in self.vertices:
			if al in self.vertices[desde]:
				self.vertices[desde].pop(al)
				self.vertices[al].pop(desde)
			else:
				raise Exception("No existe arista")
		else:
			raise Exception("No existe vertices")

	def AgregarArista(self,desde,al,peso = None):
		if desde in self.vertices and al in self.vertices:
			self.vertices[desde][al] = peso
			self.vertices[al][desde] = peso
		else:
			raise Exception("No existe vertices")

	def EstanUnidos(self,vertice1,vertice2):
		if vertice1 in self.vertices:
			if vertice2 in self.vertices[vertice1]:
				return (True,self.vertices[vertice1][vertice2])
		return (False,None)

	def LosVertices(self):
		return list(self.vertices.keys())

	def AdyacentesDe(self,vertice):
		if vertice in self.vertices:
			return list(self.vertices[vertice].items())
		else:
			raise Exception("No existe vertices")

	def VerticeAlAzar(self):
		return random.choice(list(self.vertices.keys()))
