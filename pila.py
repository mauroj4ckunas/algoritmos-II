from collections import deque

class Pila:
    def __init__(self) -> None:
        self.__pila = deque([])
    
    def __str__(self):
        a = []
        if len(list(self.__pila)) != 0:
            for b in list(self.__pila):
                a.append(b)
        else:
            return "[]"
        return str(a)

    def __esta_vacia(self) -> bool:
        return len(self.__pila) == 0

    def EstaVacia(self) -> bool:
        return self.__esta_vacia()

    def VerUltimo(self):
        if self.__esta_vacia():
            raise Exception("La cola esta vacia")
        return self.__pila[len(list(self.__pila)) - 1]

    def Apilar(self, elemento):
        self.__pila.append(elemento)
    
    def Desapilar(self):
        if self.__esta_vacia():
            raise Exception("La cola esta vacia")
        return self.__pila.pop()

