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
		
	def agregarVertice(self,vertice):
		self.vertices[vertice] = self.vertices.get(vertice,{})

	def sacarVertice(self,vertice):
		if vertice in self.vertices:
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

	def sacarArista(self,desde,al):
		if desde in self.vertices and al in self.vertices:
			if al in self.vertices[desde]:
				self.vertices[desde].pop(al)
				if self.direccion == None:
					self.vertices[al].pop(desde)
			else:
				raise Exception("La arista no existe")
		else:
			raise Exception("No existe vertices")

	def agregarArista(self,desde,al,peso = None):
		if desde in self.vertices and al in self.vertices:
			self.vertices[desde][al] = peso
			if self.direccion == None:
				self.vertices[al][desde] = peso
		else:
			raise Exception("No existe vertices")

	def pertenece(self,vertice):
		return vertice in self.vertices

	def peso(self,vertice1,vertice2):
		if vertice1 in self.vertices and vertice2 in self.vertices: return self.vertices[vertice1][vertice2]
		else: raise Exception("No existe vertices")

	def estanUnidos(self,vertice1,vertice2):
		if vertice1 in self.vertices:
			if vertice2 in self.vertices[vertice1]:
				return True
			return False
		raise Exception("No existe vertices")


	def verVertices(self):
		return list(self.vertices.keys())

	def adyacentes(self,vertice):
		if vertice in self.vertices:
			if self.direccion == None:
				ady = []
				for a in list(self.vertices[vertice].items()):
					ady.append(a[0])
			else:
				ady = list(self.vertices[vertice].items())
			return ady
		else:
			raise Exception("No existe vertices")

	def verticeAlAzar(self):
		return random.choice(list(self.vertices.keys()))


	