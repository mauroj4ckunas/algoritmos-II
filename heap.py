from collections import deque

class Heap():

    '''Heap de minimos'''

    def __init__(self) -> None:
        self.__colaPrioridad = deque([])

    #def __str__(self):
    #    palabra = ""
    #    for i in self.__colaPrioridad:
    #        palabra += f"{str(i)} "
    #    return palabra

    def __comparador(self, elem1, elem2):
        if elem1 < elem2:
            return -1
        return 1

    def __swap(self, ind1, ind2) -> None:
        aux = self.__colaPrioridad[int(ind1)]
        self.__colaPrioridad[int(ind1)] = self.__colaPrioridad[int(ind2)]
        self.__colaPrioridad[int(ind2)] = aux

    def __upheap(self, indice) -> None:
        if indice == 0:
            return
        indicePadre = (indice - 1) / 2
        if self.__comparador(self.__colaPrioridad[int(indicePadre)], self.__colaPrioridad[int(indice)]) > 0:
            self.__swap(indicePadre, indice)
            self.__upheap(int(indicePadre))

    def __mayorEntreHijos(self, padre):
        hijoIzq = (padre * 2) + 1
        hijoDer = (padre * 2) + 2
        if hijoDer >= len(self.__colaPrioridad) or self.__comparador(self.__colaPrioridad[hijoIzq], self.__colaPrioridad[hijoDer]):
            return hijoIzq
        return hijoDer

    def __downheap(self, indicePadre):
        if indicePadre == len(self.__colaPrioridad):
            return
        mayor = self.__mayorEntreHijos(indicePadre)
        if mayor >= len(self.__colaPrioridad):
            return
        if  self.__comparador(self.__colaPrioridad[indicePadre], self.__colaPrioridad[mayor]) > 0:
            self.__swap(indicePadre, mayor)
            indicePadre = mayor
            self.__downheap(indicePadre)

    def Encolar(self, elemento) -> None:
        self.__colaPrioridad.append(elemento)
        self.__upheap(len(self.__colaPrioridad) - 1)

    def __esta_vacia(self) -> bool:
        return len(self.__colaPrioridad) == 0

    def EstaVacia(self) -> bool:
        return self.__esta_vacia()

    def Desencolar(self):
        if self.__esta_vacia():
            raise Exception("La cola esta vacia")
        self.__swap(0, len(self.__colaPrioridad) - 1)
        devolver = self.__colaPrioridad.pop()
        self.__downheap(0)
        return devolver
    
    def VerMin(self):
        if self.__esta_vacia():
            raise Exception("La cola esta vacia")
        return self.__colaPrioridad[0]