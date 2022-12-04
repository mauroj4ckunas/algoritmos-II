import grafo as gf
import funciones as func
import euler as eu


COMANDOS = ["ir", "itinerario", "viaje", "reducir_caminos"]

def guardarCaminoMinimo(grafo: gf.Grafo, desde, hasta, nombreArchivo):

    dist, padres = func.dijkstra(grafo, desde)

    print(dist)
    print(padres)

    pass


def crearGrafoMundialista(listaAGrafo: list) -> gf.Grafo:

    cantSedes = int(listaAGrafo[0])
    mundial = gf.Grafo()
    coordenadas = {}
    for i in range(1, cantSedes+1):
        sede: str = listaAGrafo[i]
        sedeLista = sede.split(",")
        nombreSede: str = sedeLista[0]
        lat = sedeLista[1]
        lng = sedeLista[2]
        coordenadas[nombreSede] = [lat, lng]
        mundial.agregarVertice(nombreSede)

    indicearistas = cantSedes+2

    for j in range(indicearistas, len(listaAGrafo)):
        arista: str = listaAGrafo[j]
        aristaLista = arista.split(",")
        desde = aristaLista[0]
        hasta = aristaLista[1]
        peso = aristaLista[2]
        mundial.agregarArista(desde, hasta, peso)


    return mundial, coordenadas

def abrirArchivo():
    try:
        qatar = open("qatar.pj")
        listaInformacion = []
        for linea in qatar.readlines():
            listaInformacion.append(linea.replace('\n', ''))
    except:
        print("No se encontro el archivo.")
    finally:
        qatar.close()

    return listaInformacion

def main():

    listaSedes = abrirArchivo()
    grafoMundial, coordenadas = crearGrafoMundialista(listaSedes)

    programa = True
    while programa:
        comandoStr = input()
        comandoList = comandoStr.split(" ")
        if comandoList[0] == COMANDOS[0]:
            
            desde = comandoList[1].replace(",", "")
            hasta = comandoList[2].replace(",", "")
            archivo = comandoList[3]

            guardarCaminoMinimo(grafoMundial, desde, hasta, archivo)

            pass

        else:
            programa = False


main()





# grafo = gf.Grafo()

# grafo.agregarVertice(1)
# grafo.agregarVertice(2)
# grafo.agregarVertice(3)
# grafo.agregarVertice(4)
# grafo.agregarVertice(5)
# grafo.agregarVertice(6)

# grafo.agregarArista(1, 2)
# grafo.agregarArista(1, 3)
# grafo.agregarArista(1, 4)
# grafo.agregarArista(1, 5)
# grafo.agregarArista(2, 3)
# grafo.agregarArista(2, 4)
# grafo.agregarArista(2, 5)
# grafo.agregarArista(3, 5)
# grafo.agregarArista(3, 4)
# grafo.agregarArista(6, 4)
# grafo.agregarArista(6, 5)
# grafo.agregarArista(6, 1)


# grafoEuler = eu.Euler(grafo)

# print(grafoEuler.cicloEulerianoHierholzer(1))
# print(grafoEuler.cicloEulerianoHierholzer(2))
# print(grafoEuler.cicloEulerianoHierholzer(6))
# print(grafoEuler.cicloEulerianoHierholzer(3))


