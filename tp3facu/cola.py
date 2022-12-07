from collections import deque

class Cola:
    def __init__(self) -> None:
        self.__cola = deque([])
    
    def __str__(self):
        a = []
        if len(list(self.__cola)) != 0:
            for b in list(self.__cola):
                a.append(b)
        else:
            return "[]"
        return str(a)

    def __esta_vacia(self) -> bool:
        return len(self.__cola) == 0

    def EstaVacia(self) -> bool:
        return self.__esta_vacia()

    def VerPrimero(self):
        if self.__esta_vacia():
            raise Exception("La cola esta vacia")
        return self.__cola[0]

    def Encolar(self, elemento):
        self.__cola.append(elemento)
    
    def Desencolar(self):
        if self.__esta_vacia():
            raise Exception("La cola esta vacia")
        return self.__cola.popleft()
