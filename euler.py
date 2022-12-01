import recorridos

class Euler():
    
    '''
        Para que un grafo tenga un recorrido Euler debe tener estas propiedades:
        Grafo no dirigido: sus vertices deben tener grado par (2 o menos pueden tener impar) y ser conexo.
        Grafo dirigio: tiene un rastro euleriano si y solo si como máximo un vértice tiene ( grado de salida ) − ( grado de entrada ) = 1, 
            como máximo un vértice tiene (grado de entrada) − (grado de salida) = 1
    '''
    
    def __init__(self, grafo):
        self.grafo = grafo
        
    def esEuler(self):
        _, cant_componentes_conexas = recorridos.dfs(self.grafo)
        cant_impar = self.__contarGradosImpares(self.grafo)
        return cant_componentes_conexas == 1 and cant_impar > 2
    
    def __contarGradosImpares(gradosVertices: dict) -> int:
        impar = 0
        for v in gradosVertices:
            if gradosVertices[v] % 2 != 0:
                impar += 1
        return impar
