import funciones
from pila import Pila
from grafo import Grafo
import random

def _contarGradosImpares(gradosVertices: dict) -> int:
        impar = 0
        for v in gradosVertices:
            if gradosVertices[v] % 2 != 0:
                impar += 1
        return impar

def dfsHierholzer(grafo:Grafo,visitados,recorrido,vertice,origen):
    for v in grafo.adyacentes(vertice):
        if (vertice,v) not in visitados:
            visitados.add((vertice,v))
            visitados.add((v,vertice))
            recorrido.Apilar(v)
            if v == origen:
                return
            else:
                return dfsHierholzer(grafo,visitados,recorrido,v,origen)
        

class Euler():
    
    def __init__(self, grafo: Grafo):
        self.grafo = grafo
        _, cant_componentes_conexas = funciones.dfs(self.grafo)
        self.cant_impar = _contarGradosImpares(funciones.grados(self.grafo))
        self.tieneCiclo = (cant_componentes_conexas == 1 and self.cant_impar == 0)  # or self.cant_impar == 2
    
    def tieneCicloEuleriano(self):
        '''
            Para que un grafo tenga un ciclo euleriano debe tener estas propiedades:
            Grafo no dirigido: sus vertices deben tener grado par y ser conexo.
        '''
        return self.tieneCiclo

    def cicloEulerianoHierholzer(self, origen):
        pilaRecorrido = Pila()
        resultado = []
        aristasVisitados = set()


        pilaRecorrido.Apilar(origen)
        primerCiclo = random.choice(self.grafo.adyacentes(origen))
        aristasVisitados.add((origen,primerCiclo))
        aristasVisitados.add((primerCiclo,origen))
        pilaRecorrido.Apilar(primerCiclo)

        dfsHierholzer(self.grafo,aristasVisitados,pilaRecorrido,primerCiclo,origen)
        
        while not pilaRecorrido.EstaVacia():
            v = pilaRecorrido.VerUltimo()
            for w in self.grafo.adyacentes(v):
                if (v,w) not in aristasVisitados:
                    aristasVisitados.add((v,w))
                    aristasVisitados.add((w,v))
                    pilaRecorrido.Apilar(w)
                    dfsHierholzer(self.grafo,aristasVisitados,pilaRecorrido,w,v)
            resultado.append(pilaRecorrido.Desapilar())

        return resultado

# prueba = Grafo()

# prueba.agregarVertice(1)
# prueba.agregarVertice(2)
# prueba.agregarVertice(3)
# prueba.agregarVertice(4)
# prueba.agregarVertice(5)
# prueba.agregarVertice(6)
# prueba.agregarVertice(7)
# prueba.agregarVertice(8)
# prueba.agregarVertice(9)
# prueba.agregarVertice(10)
# prueba.agregarVertice(11)
# prueba.agregarVertice(12)
# prueba.agregarVertice(13)
# prueba.agregarVertice(14)
# prueba.agregarVertice(15)
# prueba.agregarVertice(16)
# prueba.agregarVertice(17)
# prueba.agregarVertice(18)
# prueba.agregarVertice(19)
# prueba.agregarVertice(20)

# prueba.agregarArista(1, 2, 3)
# prueba.agregarArista(1, 4, 4)
# prueba.agregarArista(2, 3, 3)
# prueba.agregarArista(3, 4, 5)
# prueba.agregarArista(3, 6, 2)
# prueba.agregarArista(4, 6, 3)
# prueba.agregarArista(4, 5, 5)
# prueba.agregarArista(3, 5, 10)
# prueba.agregarArista(1, 2, 10)
# prueba.agregarArista(1, 3, 10)
# prueba.agregarArista(1, 4, 10)
# prueba.agregarArista(1, 5, 10)
# prueba.agregarArista(15, 2, 10)
# prueba.agregarArista(15, 6, 10)
# prueba.agregarArista(15, 7, 10)
# prueba.agregarArista(15, 8, 10)
# prueba.agregarArista(16, 8, 10)
# prueba.agregarArista(16, 9, 10)
# prueba.agregarArista(16, 10, 10)
# prueba.agregarArista(16, 11, 10)
# prueba.agregarArista(14, 11, 10)
# prueba.agregarArista(14, 12, 10)
# prueba.agregarArista(14, 13, 10)
# prueba.agregarArista(14, 5, 10)
# prueba.agregarArista(2, 3, 10)
# prueba.agregarArista(2, 6, 10)
# prueba.agregarArista(3, 17, 10)
# prueba.agregarArista(3, 4, 10)
# prueba.agregarArista(4, 18, 10)
# prueba.agregarArista(4, 5, 10)
# prueba.agregarArista(5, 13, 10)
# prueba.agregarArista(6, 17, 10)
# prueba.agregarArista(6, 7, 10)
# prueba.agregarArista(7, 19, 10)
# prueba.agregarArista(7, 8, 10)
# prueba.agregarArista(8, 9, 10)
# prueba.agregarArista(17, 19, 10)
# prueba.agregarArista(17, 18, 10)
# prueba.agregarArista(18, 20, 10)
# prueba.agregarArista(18, 13, 10)
# prueba.agregarArista(19, 20, 10)
# prueba.agregarArista(19, 9, 10)
# prueba.agregarArista(18, 20, 10)
# prueba.agregarArista(13, 18, 10)
# prueba.agregarArista(20, 12, 10)
# prueba.agregarArista(20, 10, 10)
# prueba.agregarArista(10, 11, 10)
# prueba.agregarArista(12, 11, 10)
# prueba.agregarArista(12, 13, 10)
# prueba.agregarArista(9, 10, 10)

# ciclo = Euler(prueba)

# print(ciclo.cicloEulerianoHierholzer(4))