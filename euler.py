import recorridos
import cola as col
import pila as pil
import grafo as gf

class Euler():
    
    def __init__(self, grafo: gf.Grafo):
        self.grafo = grafo
    
    def esEuler(self):
        '''
            Para que un grafo tenga un ciclo euleriano debe tener estas propiedades:
            Grafo no dirigido: sus vertices deben tener grado par y ser conexo.
        '''
        return self.__tieneCircuito()

    def __tieneCircuito(self) -> bool:
        _, cant_componentes_conexas = recorridos.dfs(self.grafo)
        cant_impar = self.__contarGradosImpares(recorridos.gradosNoDirigido(self.grafo))
        return cant_componentes_conexas == 1 and cant_impar == 0 #Es 0 si es un ciclo Euleriano

    def __contarGradosImpares(self, gradosVertices: dict) -> int:
        impar = 0
        for v in gradosVertices:
            if gradosVertices[v] % 2 != 0:
                impar += 1
        return impar
        
    def cicloEulerianoHierholzer(self, origen):
        if not self.__tieneCircuito():
            raise Exception("Tiene grado impar o no es conexo")

        aristasNoVisitadas = pil.Pila()
        aristasVisitadas = set()
        camino = pil.Pila()
        for v in self.grafo.adyacentes(origen):
            aristasNoVisitadas.Apilar((origen, v))
        

        self.__algoritmoHierholzer(aristasNoVisitadas, aristasVisitadas, camino, origen)

        return camino


    def __algoritmoHierholzer(self, aristasNoVisitadas: pil.Pila, aristasVisitadas: set, camino: pil.Pila, origen):
        
        while not aristasNoVisitadas.EstaVacia():
            arista = aristasNoVisitadas.Desapilar()
            aristasVisitadas.add(arista)
            camino.Apilar(origen)
            if arista[0] == origen:
                visitado = arista[1]
            else:
                visitado = arista[0]
            camino.Apilar(visitado)
            self.__dfsHierholzer(visitado, origen, aristasVisitadas, camino)


    def __dfsHierholzer(self, inicio, origen, visitadas: set, camino: pil.Pila):
        for adyacente in self.grafo.adyacentes(inicio):
            if ((adyacente, inicio) not in visitadas) and ((inicio, adyacente) not in visitadas):
                print((inicio, adyacente))
                visitadas.add((inicio, adyacente))
                if adyacente == origen:
                    if camino.VerUltimo() != origen:
                        camino.Apilar(adyacente)
                    break
                camino.Apilar(adyacente)
                self.__dfsHierholzer(adyacente, origen, visitadas, camino)   



