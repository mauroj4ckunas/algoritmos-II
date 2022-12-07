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
        pajek.writelines(f'{len(arbol.verVertices())}\n')
        for ar in aristas:
            pajek.writelines(f'{ar[0]},{ar[1]},{ar[2]}\n')

    pass

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
    try:
        camino, dist = func.reconstruirCaminoMinimo(grafo,desde,hasta)
        mensajeFinal(list(camino), dist[hasta])
        crearArchivoKML(list(camino), nombreArchivo, coordenadas, desde, hasta)
    except KeyError: #Si el hasta no tiene  
        print(ErrorSinRecorrido().Error())

def reconstruirComando(grafoMundial:gf.Grafo,comando: list):
    index = 0
    primerElemento = comando[index].replace(",", "")
    while not grafoMundial.pertenece(primerElemento):
        index += 1
        primerElemento += " " + comando[index].replace(",", "")
    index += 1
    segundoElemento = comando[index].replace(",", "")
    if len(comando) - 1 == index:
        terceroYultimo = ""
    else:
        while not grafoMundial.pertenece(segundoElemento):
           index += 1
           segundoElemento += " " + comando[index].replace(",", "")
        index += 1
        terceroYultimo = comando[index] 
    
    return primerElemento,segundoElemento,terceroYultimo

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
            desde , hasta , archivo = reconstruirComando(grafoMundial,comandoList[1:])
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
            desde,archivo,_ = reconstruirComando(grafoMundial,comandoList[1:])
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