package main

import (
	"fmt"
)

func main() {
	lista := ListaE{}

	for {

		var elg string
		fmt.Println("Elegir con un numero:")
		fmt.Println("--------------------------")
		fmt.Println("1. insertar")
		fmt.Println("2. buscar")
		fmt.Println("3. eliminar")
		fmt.Println("4. Mostrar")
		fmt.Println("5. salir")
		fmt.Println("--------------------------")
		fmt.Scan(&elg)
		if elg == "1" {
			var dato Estudiante
			fmt.Println("Ingrese carnet")
			fmt.Scan(&dato.carnet)
			fmt.Println("Ingrese nombre")
			fmt.Scan(&dato.nombre)
			lista.insertar(dato)
		} else if elg == "2" {
			var carnet string
			fmt.Println("Ingrese carnet que desea buscar")
			fmt.Scan(&carnet)
			lista.buscar(carnet)
		} else if elg == "3" {
			var carnet string
			fmt.Println("Ingrese carnet que desea eliminar")
			fmt.Scan(&carnet)
			lista.eliminar(carnet)
		} else if elg == "4" {
			lista.mostrar()
		} else if elg == "5" {
			fmt.Println("no salio")
			break
		}

	}

}

type Estudiante struct {
	carnet string
	nombre string
}

type Nodo struct {
	dato      Estudiante
	siguiente *Nodo
}

type ListaE struct {
	inicio *Nodo
	fin    *Nodo
}

func (lista *ListaE) insertar(dato Estudiante) {
	aux := Nodo{
		dato:      dato,
		siguiente: nil,
	}
	if lista.inicio == nil {
		lista.inicio = &aux
		lista.fin = lista.inicio
	} else {
		lista.fin.siguiente = &aux
		lista.fin = lista.fin.siguiente
	}
}
func (l *ListaE) mostrar() {
	var auxiliar *Nodo
	auxiliar = l.inicio
	fmt.Println("Carnet--------Nombre")
	fmt.Println("_____________________")

	for auxiliar != nil {
		fmt.Println(auxiliar.dato.carnet + "-------" + auxiliar.dato.nombre)
		auxiliar = auxiliar.siguiente

	}
}
func (l *ListaE) buscar(dato string) *Nodo {
	var auxiliar *Nodo
	auxiliar = l.inicio
	fmt.Println("Carnet--------Nombre")
	fmt.Println("_____________________")
	for auxiliar != nil {
		if auxiliar.dato.carnet == dato {
			fmt.Println(auxiliar.dato.carnet + "-------" + auxiliar.dato.nombre)
			return auxiliar
		}
		auxiliar = auxiliar.siguiente

	}
	var null *Nodo
	return null
}
func (l *ListaE) eliminar(dato string) {
	var auxiliar1, actual, siguiente *Nodo
	auxiliar1 = l.buscar(dato)
	actual = l.inicio

	for {
		if auxiliar1 == l.inicio && l.inicio.siguiente == nil {
			l.inicio = nil
			break
		} else if auxiliar1 == l.inicio {
			auxiliar1 = l.inicio.siguiente
			l.inicio = nil
			l.inicio = auxiliar1
			break
		} else if auxiliar1 == actual.siguiente {
			siguiente = auxiliar1.siguiente
			l.fin = actual
			l.fin.siguiente = siguiente
			fmt.Println("simon")
			break
		}

		if actual == l.fin {
			break
		}
		actual = actual.siguiente

	}
}
