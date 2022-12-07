import funciones
import pila as pil
import grafo as gf

class Euler():
    
    def __init__(self, grafo: gf.Grafo):
        self.grafo = grafo
    
    def tieneCicloEuleriano(self):
        '''
            Para que un grafo tenga un ciclo euleriano debe tener estas propiedades:
            Grafo no dirigido: sus vertices deben tener grado par y ser conexo.
        '''
        return self.__tieneCicloEuleriano()

    def __tieneCicloEuleriano(self) -> bool:
        _, cant_componentes_conexas = funciones.dfs(self.grafo)
        cant_impar = self.__contarGradosImpares(funciones.grados(self.grafo))
        return cant_componentes_conexas == 1 and cant_impar == 0#Es 0 si es un ciclo Euleriano

    def __contarGradosImpares(self, gradosVertices: dict) -> int:
        impar = 0
        for v in gradosVertices:
            if gradosVertices[v] % 2 != 0:
                impar += 1
        return impar
        
    def cicloEulerianoHierholzer(self, origen):
        camino = list()
        aristasNoVisitadas = {}
        peso = 0
        pesoVisto = set()
        for v in self.grafo.verVertices():
            aristasNoVisitadas[v] = pil.Pila()
            for w in self.grafo.adyacentes(v):
                if (v, w) not in pesoVisto and (w, v) not in pesoVisto:
                    pesoVisto.add((v, w))
                    peso += int(self.grafo.peso(v, w))
                aristasNoVisitadas[v].Apilar((v, w))
        aristasVisitadas = set()
        camino = self.__algoritmoHierholzer(aristasNoVisitadas, aristasVisitadas, camino, origen)
        return camino, peso


    def __algoritmoHierholzer(self, aristasNoVisitadas: dict, aristasVisitadas: set, camino: list, vertice):
        
        camino.append(vertice)
        self.__dfsHierholzer(vertice, aristasNoVisitadas, aristasVisitadas, camino, vertice)
        # for i in range(len(camino)):
        i = 0
        while i < len(camino):
            while not aristasNoVisitadas[camino[i]].EstaVacia():
                sig = aristasNoVisitadas[camino[i]].Desapilar()
                if sig not in aristasVisitadas and (sig[1], sig[0]) not in aristasVisitadas:
                    caminoAux = []
                    aristasVisitadas.add(sig)
                    caminoAux.append(sig[0])
                    caminoAux.append(sig[1])
                    self.__dfsHierholzer(sig[1], aristasNoVisitadas, aristasVisitadas, caminoAux, sig[0])

                    a = camino[:i]
                    b = camino[i+1:]
                    camino = a + caminoAux + b
                    i = -1
            i += 1

        return camino


    def __dfsHierholzer(self, vertice, noVisitadas: dict, visitadas: set, caminoActualizado: list, inicio, seguir = True):
        while not noVisitadas[vertice].EstaVacia() and seguir:
            arista = noVisitadas[vertice].Desapilar()
            if arista not in visitadas and (arista[1], arista[0]) not in visitadas:
                visitadas.add(arista)
                if caminoActualizado[len(caminoActualizado) - 1] != arista[1]:
                    caminoActualizado.append(arista[1])
                if arista[1] == inicio:
                    return False
                seguir = self.__dfsHierholzer(arista[1], noVisitadas, visitadas, caminoActualizado, inicio)

# prueba = gf.Grafo()

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

# print(ciclo.tieneCicloEuleriano())

# a, peso = ciclo.cicloEulerianoHierholzer(5)

# print(a)

# for i in range(1, 21):
#     a, peso = ciclo.cicloEulerianoHierholzer(i)

#     print(peso)
#     print(a)

# vistos = set()
# for ar in arista:
#     if ar in vistos or (ar[1], ar[0]) in vistos:
#         print(ar)
#     else:
#         vistos.add(ar)

# for ar in arista:
#     print(ar)