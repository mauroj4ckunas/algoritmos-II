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
                    i = -1 #Con esto reinicio el la busqueda de adyacentes.
                    break
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