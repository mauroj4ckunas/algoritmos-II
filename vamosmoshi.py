#!/usr/bin/python3
from collections import deque
import grafo as gf
import funciones as func
import euler as eu
from errores import ErrorSinRecorrido
import sys

COMANDOS = ["ir", "itinerario", "viaje", "reducir_caminos"]


def itinerarioPosible(archivo, vertices: list):
    grafo = gf.Grafo(True) #grafo dirigido
    for vertice in vertices:
        grafo.agregarVertice(vertice)

    errorLectura = False
    try:
        caminos = open(archivo)
        for linea in caminos.readlines():
            union = linea[:len(linea)-1].split(",")
            grafo.agregarArista(union[0],union[1])
    except:
        print(ErrorSinRecorrido().Error())
        errorLectura = True
    finally:
        caminos.close()

    return grafo,errorLectura



def caminosReducidos(arbol: gf.Grafo, archivo, coordenadas: dict, aristas: list):

    with open(archivo, "w") as pajek:
        pajek.writelines(f'{len(arbol.verVertices())}\n')
        for sede in arbol.verVertices():
            pajek.writelines(f'{sede},{coordenadas[sede][0]},{coordenadas[sede][1]}\n')
        pajek.writelines(f'{len(arbol.verVertices())}\n')
        for ar in aristas:
            pajek.writelines(f'{ar[0]},{ar[1]},{ar[2]}\n')

def viajeTodosLosCaminos(grafo: gf.Grafo, desde: str, nombreArchivo: str, coordenadas):
    cicloEuler = eu.Euler(grafo)
    if not cicloEuler.tieneCicloEuleriano:
        print(ErrorSinRecorrido().Error())
        return
    camino, peso, lista = cicloEuler.cicloEulerianoHierholzer(desde)
    print(f"La cantidad de arista es: {len(lista)}")
    mensajeFinal(camino, peso)
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

def mensajeFinal(lista: list, peso = None):
    mensaje = ""
    for i in range(len(lista)):
        if i == 0:
            mensaje += f"{lista[i]}"
            continue
        mensaje += f" -> {lista[i]}"
    print(mensaje)
    if peso != None:
        print(f'Tiempo total: {peso}')

def guardarCaminoMinimo(grafo: gf.Grafo, desde, hasta, nombreArchivo, coordenadas):
    dist, padres= func.dijkstra(grafo, desde)
    try:
        anterior = padres[hasta]
        camino = deque([])
        camino.appendleft(hasta)
        camino.appendleft(anterior)
        while anterior != desde:
            anterior = padres[anterior]
            camino.appendleft(anterior)
        mensajeFinal(list(camino), dist[hasta])
        crearArchivoKML(list(camino), nombreArchivo, coordenadas, desde, hasta)
    except KeyError: #Si el hasta no tiene  
        print(ErrorSinRecorrido().Error())

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

    indiceAristas = cantSedes+2

    for j in range(indiceAristas, len(listaAGrafo)):
        arista: str = listaAGrafo[j]
        aristaLista = arista.split(",")
        desde = aristaLista[0]
        hasta = aristaLista[1]
        peso = aristaLista[2]
        mundial.agregarArista(desde, hasta, peso)

    print(mundial)
    arista = func.verAristas(mundial)
    print(len(arista))


    return mundial, coordenadas

def abrirArchivo(archivo):
    try:
        qatar = open(archivo)
        listaInformacion = []
        for linea in qatar.readlines():
            listaInformacion.append(linea.replace('\n', ''))

    except:
        print("No se encontro el archivo.")

    return listaInformacion

def main():
    # archivo = sys.argv[1:]
    # if len(archivo) > 1:
    #     print("No se encontro el archivo.")
    #     return

    listaSedes = abrirArchivo("qatar.pj")
    grafoMundial, coordenadas = crearGrafoMundialista(listaSedes)

    programa = True
    while programa:
        comandoStr = input()
        comandoList = comandoStr.split(" ")
        if comandoList[0] == COMANDOS[0]:
            
            #Estas validaciones son porque existen sedes con nombres de dos palabras.
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

        elif comandoList[0] == COMANDOS[1]:
            archivo = comandoList[1]
            grafoDeRecorrido,error = itinerarioPosible(archivo, grafoMundial.verVertices())
            if error == False:
                posibleCamino, esPosible = func.bfsordenadoentrada(grafoDeRecorrido)
                if esPosible == False:
                    print(ErrorSinRecorrido().Error())
                else:
                    mensajeFinal(posibleCamino)

        elif comandoList[0] == COMANDOS[2]:
            
            if len(comandoList) == 4:
                desde = comandoList[1] + " " + comandoList[2].replace(",", "")
                archivo = comandoList[3]
            else: 
                desde = comandoList[1].replace(",", "")
                archivo = comandoList[2]

            viajeTodosLosCaminos(grafoMundial, desde, archivo, coordenadas)

        elif comandoList[0] == COMANDOS[3]:
            archivo = comandoList[1]
            arbol, peso = func.prim(grafoMundial)
            print(f'Peso total: {peso}')
            aristas = func.verAristas(arbol)
            caminosReducidos(arbol, archivo, coordenadas, aristas)

        else:
            programa = False


main()