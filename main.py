import grafo as gf
import recorridos as rec
import euler as eu

grafo = gf.Grafo()

grafo.agregarVertice(1)
grafo.agregarVertice(2)
grafo.agregarVertice(3)
grafo.agregarVertice(4)
grafo.agregarVertice(5)
grafo.agregarVertice(6)

grafo.agregarArista(1, 2)
grafo.agregarArista(2, 3)
grafo.agregarArista(3, 4)
grafo.agregarArista(5, 4)
grafo.agregarArista(6, 5)
grafo.agregarArista(6, 1)

grafoEuler = eu.Euler(grafo)

print(grafoEuler.cicloEuleriano(2))
print(grafoEuler.cicloEuleriano(5))