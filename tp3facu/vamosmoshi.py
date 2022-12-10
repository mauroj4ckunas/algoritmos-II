#!/usr/bin/python3
from collections import deque
import grafo as gf
import funciones as func
import euler as eu
from errores import ErrorSinRecorrido
import sys

sys.setrecursionlimit(3000)
COMANDOS = ["ir", "itinerario", "viaje", "reducir_caminos"]


def itinerarioPosible(archivo, vertices: list):
    grafo = gf.Grafo(True) #grafo dirigido
    for vertice in vertices:
        grafo.agregarVertice(vertice)

    errorLectura = False
    try:
        caminos = open(archivo)
        for linea in caminos:
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
        pajek.writelines(f'{len(aristas)}\n')
        for ar in aristas:
            pajek.writelines(f'{ar[0]},{ar[1]},{ar[2]}\n')

def viajeTodosLosCaminos(grafo: gf.Grafo, desde: str, nombreArchivo: str, coordenadas):
    cicloEuler = eu.Euler(grafo)
    if not cicloEuler.tieneCicloEuleriano():
        print(ErrorSinRecorrido().Error())
        return
    camino, peso = cicloEuler.cicloEulerianoHierholzer(desde)
    mensajeFinal(camino, peso)
    crearArchivoKML(camino, nombreArchivo, coordenadas, desde)

def crearArchivoKML(listaPuntos: list, nombreArchivo: str, coordenadas: dict, desde, hasta = None):
    with open(nombreArchivo, "w", encoding="UTF-8") as nuevo:
        nuevo.writelines('<?xml version="1.0" encoding="UTF-8"?>\n')
        nuevo.writelines('<kml xmlns="http://earth.google.com/kml/2.1">\n')
        nuevo.writelines('\t<Document>\n')
        if hasta != None: nuevo.writelines(f'\t\t<name>Camino desde {desde} hacia {hasta}</name>\n')
        else: nuevo.writelines(f'\t\t<name>Viaje desde {desde}</name>\n\n')
        vistos = set()
        for sede in listaPuntos:
            if sede in vistos:
                continue
            nuevo.writelines('\t\t<Placemark>\n')
            nuevo.writelines(f'\t\t\t<name>{sede}</name>\n')
            nuevo.writelines('\t\t\t<Point>\n')
            nuevo.writelines(f'\t\t\t\t<coordinates>{coordenadas[sede][0]}, {coordenadas[sede][1]}</coordinates>\n')
            nuevo.writelines('\t\t\t</Point>\n')
            nuevo.writelines('\t\t</Placemark>\n')
            vistos.add(sede)
        nuevo.writelines('\n')
        for i in range(len(listaPuntos)-1):
            nuevo.writelines('\t\t<Placemark>\n')
            nuevo.writelines('\t\t\t<LineString>\n')
            nuevo.writelines(f'\t\t\t\t<coordinates>{coordenadas[listaPuntos[i]][0]}, {coordenadas[listaPuntos[i]][1]} {coordenadas[listaPuntos[i+1]][0]}, {coordenadas[listaPuntos[i+1]][1]}</coordinates>\n')
            nuevo.writelines('\t\t\t</LineString>\n')
            nuevo.writelines('\t\t</Placemark>\n')
        nuevo.writelines('\t</Document>\n')
        nuevo.writelines('</kml>')

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
    try:
        if grafo.pertenece(desde) and grafo.pertenece(hasta):
            camino, dist = func.reconstruirCaminoMinimo(grafo,desde,hasta)
            mensajeFinal(list(camino), dist[hasta])
            crearArchivoKML(list(camino), nombreArchivo, coordenadas, desde, hasta)
        else:
            print(ErrorSinRecorrido().Error())
    except KeyError: #Si el hasta no tiene  
        print(ErrorSinRecorrido().Error())

def reconstruirComando(comando: list):
    resultado = []
    index = 0
    palabra = ""
    while index < len(comando):
        if "," in comando[index] or index == len(comando) - 1:
            palabra += comando[index].replace(",", "")
            resultado.append(palabra)
            palabra = ""
        else:
            palabra += comando[index] + " "
        index += 1 
    return resultado

def abrirArchivo(archivo):
    try:
        qatar = open(archivo)
        mundial = gf.Grafo()
        coordenadas = {}
        cantCiudades = int(qatar.readline().replace('\n', ''))
        visitadas = 0
        while visitadas < cantCiudades:
            linea = qatar.readline().replace('\n', '')
            ciudad , lat , lng = linea.split(",")
            mundial.agregarVertice(ciudad)
            coordenadas[ciudad] = [lat,lng]
            visitadas += 1
        visitadas = 0
        cantUniones =  int(qatar.readline().replace('\n', ''))
        while visitadas < cantUniones:
            linea = qatar.readline().replace('\n', '')
            desde , hasta , tiempo = linea.split(",")
            mundial.agregarArista(desde , hasta , tiempo)
            visitadas += 1
    except:
        raise Exception("No se encontro el archivo.")
    finally:
        qatar.close()

    return mundial, coordenadas

def main():
    archivo = sys.argv[1:]
    if len(archivo) > 1:
        print("No se encontro el archivo.")
        return
    grafoMundial, coordenadas = abrirArchivo(archivo[0])

    programa = True
    while programa:
        try:
            comandoStr = input()
        except EOFError:
            return
        
        comandoList = comandoStr.split(" ")

        if comandoList[0] == COMANDOS[0]:
            escrituraUsuario = reconstruirComando(comandoList[1:])
            if len(escrituraUsuario) != 3: #3 ya que necesita el desde, hasta y el nombre del archivo
               print(ErrorSinRecorrido().Error()) 
            else:
                guardarCaminoMinimo(grafoMundial, escrituraUsuario[0], escrituraUsuario[1], escrituraUsuario[2], coordenadas)

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
            escrituraUsuario = reconstruirComando(comandoList[1:])
            if len(escrituraUsuario)!= 2:#2 ya que necesita la entrada de desde y el nombre del archivo
                print(ErrorSinRecorrido().Error())
            else:
                viajeTodosLosCaminos(grafoMundial, escrituraUsuario[0], escrituraUsuario[1], coordenadas)

        elif comandoList[0] == COMANDOS[3]:
            archivo = comandoList[1]
            arbol, peso = func.prim(grafoMundial)
            print(f'Peso total: {peso}')
            aristas = func.verAristas(arbol)
            caminosReducidos(arbol, archivo, coordenadas, aristas)

        else:
            programa = False

main()