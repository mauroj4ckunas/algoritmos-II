import funciones
import pila as pil
import grafo as gf

class Euler():
    
    def __init__(self, grafo: gf.Grafo):
        self.grafo = grafo
    
    def tieneCaminoEuleriano(self):
        '''
            Para que un grafo tenga un ciclo euleriano debe tener estas propiedades:
            Grafo no dirigido: sus vertices deben tener grado par y ser conexo.
        '''
        return self.__tieneCaminoEuleriano()

    def __tieneCaminoEuleriano(self) -> bool:
        _, cant_componentes_conexas = funciones.dfs(self.grafo)
        cant_impar = self.__contarGradosImpares(funciones.gradosNoDirigido(self.grafo))
        return cant_componentes_conexas == 1 and cant_impar == 0#Es 0 si es un ciclo Euleriano

    def __contarGradosImpares(self, gradosVertices: dict) -> int:
        impar = 0
        for v in gradosVertices:
            if gradosVertices[v] % 2 != 0:
                impar += 1
        return impar
        
    def cicloEulerianoHierholzer(self, origen):
        if not self.__tieneCaminoEuleriano():
            return []
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
        for i in camino:
            while not aristasNoVisitadas[i].EstaVacia():
                sig = aristasNoVisitadas[i].Desapilar()
                if sig not in aristasVisitadas and (sig[1], sig[0]) not in aristasVisitadas:
                    caminoAux = []
                    caminoAux.append(sig[0])
                    self.__dfsHierholzer(sig[0], aristasNoVisitadas, aristasVisitadas, caminoAux, sig[0])

                    if i == camino[len(camino) - 1]:
                        camino = camino[:len(camino) - 2] + caminoAux
                    else:
                        for j in range(len(camino) - 2, -1, -1):
                            if camino[j] == i:
                                a = camino[:j]
                                b = camino[j+1:]
                                camino = a + caminoAux + b
                                break

        return camino


    def __dfsHierholzer(self, vertice, noVisitadas: dict, visitadas: set, caminoActualizado: list, inicio, seguir = True):
        while not noVisitadas[vertice].EstaVacia() and seguir:
            arista = noVisitadas[vertice].Desapilar()
            if arista not in visitadas and (arista[1], arista[0]) not in visitadas:
                visitadas.add(arista)
                if len(caminoActualizado) != 0:
                    if caminoActualizado[len(caminoActualizado) - 1] != arista[1]:
                        caminoActualizado.append(arista[1])
                if arista[1] == inicio:
                    return False
                seguir = self.__dfsHierholzer(arista[1], noVisitadas, visitadas, caminoActualizado, inicio)

