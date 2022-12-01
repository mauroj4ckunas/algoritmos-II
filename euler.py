import recorridos
import cola as cl
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
        
    def cicloEuleriano(self, origen):
        if not self.__tieneCircuito():
            raise Exception("Tiene grado impar o no es conexo")

        camino = cl.Cola()
        visitados = set()

        for v in self.grafo.adyacentes(origen):
            camino.Encolar(origen)
            if v not in visitados:
                visitados.add(origen)
                camino.Encolar(v)
                visitados.add(v)
            self.__dfsEuler(v, origen, visitados, camino)
        return camino
                
    def __dfsEuler(self, vertice, origen, visitados: set, camino: cl.Cola):
        for adyacente in self.grafo.adyacentes(vertice):
            if adyacente not in visitados:
                camino.Encolar(adyacente)
                visitados.add(adyacente)
                self.__dfsEuler(adyacente,origen,visitados,camino)



