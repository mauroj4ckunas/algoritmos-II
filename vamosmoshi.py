from collections import deque
import grafo as gf
import funciones as func
import euler as eu

COMANDOS = ["ir", "itinerario", "viaje", "reducir_caminos"]




def viajeTodosLosCaminos(grafo: gf.Grafo, desde: str, nombreArchivo: str, coordenadas):
    cicloEuler = eu.Euler(grafo)

    camino, peso = cicloEuler.cicloEulerianoHierholzer(desde)

    mensaje = ""
    for i in range(len(camino)):
        if i == 0:
            mensaje += camino[i]
            continue
        mensaje += " -> " + camino[i]

    print(mensaje)
    print(f'Tiempo total: {peso}')
    crearArchivoKML(camino, nombreArchivo, coordenadas, desde)

def crearArchivoKML(listaPuntos: list, nombreArchivo: str, coordenadas: dict, desde, hasta = None):
    with open(nombreArchivo, "w", encoding="UTF-8") as nuevo:
        nuevo.writelines('<?xml version="1.0" encoding="UTF-8"?>\n')
        nuevo.writelines('<kml xmlns="http://earth.google.com/kml/2.1">\n')
        nuevo.writelines('\t<Document>\n')
        if hasta != None: nuevo.writelines(f'\t\t<name>Camino desde {desde} hacia {hasta}</name>\n')
        else: nuevo.writelines(f'\t\t<name>Viaje desde {desde}</name>\n\n')
        for sede in listaPuntos:
            nuevo.writelines('\t\t<Placemark>\n')
            nuevo.writelines(f'\t\t\t<name>{sede}</name>\n')
            nuevo.writelines('\t\t\t<Point>\n')
            nuevo.writelines(f'\t\t\t\t<coordinates>{coordenadas[sede][0]}, {coordenadas[sede][1]}</coordinates>\n')
            nuevo.writelines('\t\t\t</Point>\n')
            nuevo.writelines('\t\t</Placemark>\n')
        nuevo.writelines('\n')
        for i in range(len(listaPuntos)):
            if i+1 == len(listaPuntos):
                break
            nuevo.writelines('\t\t<Placemark>\n')
            nuevo.writelines('\t\t\t<LineString>\n')
            nuevo.writelines(f'\t\t\t\t<coordinates>{coordenadas[listaPuntos[i]][0]}, {coordenadas[listaPuntos[i]][1]} {coordenadas[listaPuntos[i+1]][0]}, {coordenadas[listaPuntos[i+1]][1]}</coordinates>\n')
            nuevo.writelines('\t\t\t</LineString>\n')
            nuevo.writelines('\t\t</Placemark>\n')

def guardarCaminoMinimo(grafo: gf.Grafo, desde, hasta, nombreArchivo, coordenadas):

    dist, padres= func.dijkstra(grafo, desde)

    anterior = padres[hasta]
    camino = deque([])
    camino.appendleft(hasta)
    camino.appendleft(anterior)
    while anterior != desde:
        anterior = padres[anterior]
        camino.appendleft(anterior)

    mensaje = ""
    for i in range(len(list(camino))):
        if i == 0:
            mensaje += list(camino)[i]
            continue
        mensaje += " -> " + list(camino)[i]

    print(mensaje)
    print(f'Tiempo total: {dist[hasta]}')
    crearArchivoKML(list(camino), nombreArchivo, coordenadas, desde, hasta)

def crearGrafoMundialista(listaAGrafo: list):

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
            
            if len(comandoList) > 4:
                if len(comandoList) == 6:
                    desde = comandoList[1] + " " + comandoList[2].replace(",", "")
                    hasta = comandoList[3] + " " + comandoList[4].replace(",", "")
                    archivo = comandoList[5]

                elif "," in comandoList[2]:
                    desde = comandoList[1] + " " + comandoList[2].replace(",", "")
                    hasta = comandoList[3].replace(",", "")
                    archivo = comandoList[4]

                elif "," in comandoList[3]:
                    desde = comandoList[1].replace(",", "")
                    hasta = comandoList[2] + " " + comandoList[3].replace(",", "")
                    archivo = comandoList[4]
            elif len(comandoList) <= 4: 
                desde = comandoList[1].replace(",", "")
                hasta = comandoList[2].replace(",", "")
                archivo = comandoList[3]

            guardarCaminoMinimo(grafoMundial, desde, hasta, archivo, coordenadas)

        if comandoList[0] == COMANDOS[2]:
            
            if len(comandoList) == 4:
                desde = comandoList[1] + " " + comandoList[2].replace(",", "")
                archivo = comandoList[3]
            else: 
                desde = comandoList[1].replace(",", "")
                archivo = comandoList[2]

            viajeTodosLosCaminos(grafoMundial, desde, archivo, coordenadas)

        if comandoList[0] == COMANDOS[3]:
            archivo = comandoList[1]
            grafo, peso = func.prim(grafoMundial)

            print(peso)

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


