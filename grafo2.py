import random

class Grafo:
	def __init__(self,direccion = None):
		self.vertices = {}
		self.direccion = direccion #si la direccion es None significa que es no dirigido

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
		if self.ExisteVertice(vertice):
			if self.direccion == None:
				for v in self.vertices.pop(vertice).keys():
					self.vertices[v].pop(vertice)
			else:
				self.vertices.pop(vertice)
				for v in self.vertices:
					if vertice in self.vertices[v]:
						self.vertices[v].pop(vertice)
		else:
			raise Exception("No existe vertices")

	def SacarArista(self,desde,al):
		if self.ExisteVertice(desde) and self.ExisteVertice(al):
			if al in self.vertices[desde]:
				self.vertices[desde].pop(al)
				if self.direccion == None:
					self.vertices[al].pop(desde)
			else:
				raise Exception("La arista no existe")
		else:
			raise Exception("No existe vertices")

	def AgregarArista(self,desde,al,peso = None):
		if self.ExisteVertice(desde) and self.ExisteVertice(al):
			self.vertices[desde][al] = peso
			if self.direccion == None:
				self.vertices[al][desde] = peso
		else:
			raise Exception("No existe vertices")

	def ExisteVertice(self,vertice):
		return vertice in self.vertices

	def EstanUnidos(self,vertice1,vertice2):
		if self.ExisteVertice(vertice1):
			if vertice2 in self.vertices[vertice1]:
				return (True,self.vertices[vertice1][vertice2])
			return (False,None)
		raise Exception("No existe vertices")


	def LosVertices(self):
		return list(self.vertices.keys())

	def AdyacentesDe(self,vertice):
		if self.ExisteVertice(vertice):
			return list(self.vertices[vertice].items())
		else:
			raise Exception("No existe vertices")

	def VerticeAlAzar(self):
		return random.choice(list(self.vertices.keys()))