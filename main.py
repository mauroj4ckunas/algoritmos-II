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
grafo.agregarArista(1, 3)
grafo.agregarArista(1, 4)
grafo.agregarArista(1, 5)
grafo.agregarArista(2, 3)
grafo.agregarArista(2, 4)
grafo.agregarArista(2, 5)
grafo.agregarArista(3, 5)
grafo.agregarArista(3, 4)
grafo.agregarArista(6, 4)
grafo.agregarArista(6, 5)

grafoEuler = eu.Euler(grafo)

print(grafoEuler.cicloEulerianoHierholzer(2))