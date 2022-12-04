import grafo as gf
import recorridos as rec
import euler as eu


COMANDOS = ["ir", "itinerario", "viaje", "reducir_caminos"]

# class Sede():
#     def __init__(self, nombre, lat, lng) -> None:
#         self.nombre = nombre
#         self.lat = lat
#         self.lng = lng


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

    cantAristas = cantSedes+2

    for j in range(cantAristas, len(listaAGrafo)):
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

    print(grafoMundial)
    print(coordenadas)

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


