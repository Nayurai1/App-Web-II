package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var _ = bufio.NewReader(os.Stdin)
var _ = errors.New
var _ = strconv.Itoa
var _ = strings.ToUpper

type cliente struct {
	id      int
	Nombre  string
	Carrera string
	saldo   float64
}

type producto struct {
	id        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type pedido struct {
	id         int
	clienteid  string
	productoid int
	cantidad   int
	total      float64
	fecha      string
}

func main() {
	var clientes []cliente
	var productos []producto
	var pedidos []pedido

	clientes = []cliente{
		{id: 1, Nombre: "Sornoza Isaac Arturo", Carrera: "TIC", saldo: 100.00},
		{id: 2, Nombre: "Sornoza Jaime Jesús", Carrera: "Psicología", saldo: 15.00},
		{id: 3, Nombre: "Sornoza Fanny April", Carrera: "Medicina", saldo: 250.00},
	}

	productos = []producto{
		{id: 1, Nombre: "Empanada", Precio: 2.00, Stock: 9, Categoria: "Desayuno"},
		{id: 2, Nombre: "Corviche", Precio: 1.00, Stock: 14, Categoria: "Desayuno"},
		{id: 3, Nombre: "Encebollado", Precio: 3.00, Stock: 10, Categoria: "Almuerzo"},
		{id: 4, Nombre: "Capuchino", Precio: 5.00, Stock: 30, Categoria: "desayuno"},
	}

	fmt.Println("=== ESTADO INICIAL ===")
	ListarClientes(clientes)
	ListarProductos(productos)

	nuevoCli := cliente{id: 4, Nombre: "Melanie Carolina", Carrera: "TI", saldo: 50.0}
	clientes = AgregarCliente(clientes, nuevoCli)

	productos = QuitarProducto(productos, 2)

	ListarClientes(clientes)
	ListarProductos(productos)

	fmt.Println("\n--- PRUEBAS DE BÚSQUEDA ---")

	idExistente := 1
	posicionReal := BuscarClientePorID(clientes, idExistente)
	fmt.Printf("Buscando ID %d: Encontrado en posicion %d\n", idExistente, posicionReal)

	idFalso := 99
	posicionFalsa := BuscarClientePorID(clientes, idFalso)
	fmt.Printf("Buscando ID %d: Resultado %d (No existe)\n", idFalso, posicionFalsa)

	_ = pedidos
}

func ListarClientes(lista []cliente) {
	if len(lista) == 0 {
		fmt.Println("No hay clientes registrados.")
		return
	}

	fmt.Println("--- LISTA DE CLIENTES ---")
	for _, c := range lista {
		fmt.Printf("ID: %d - %s (%s) - Saldo: $%.2f\n", c.id, c.Nombre, c.Carrera, c.saldo)
	}
}

func ListarProductos(lista []producto) {
	if len(lista) == 0 {
		fmt.Println("No hay productos registrados.")
		return
	}

	fmt.Println("--- LISTA DE PRODUCTOS ---")
	for _, p := range lista {
		fmt.Printf("ID: %d - %s [%s] - Precio: $%.2f - Stock: %d\n", p.id, p.Nombre, p.Categoria, p.Precio, p.Stock)
	}
}

//Checkpoint 2 — CRUD de Cliente y Producto

// CLIENTES
func BuscarClientePorID(listaDeClientes []cliente, idQueBusco int) int {
	for fila, clienteDeLaFila := range listaDeClientes {
		if clienteDeLaFila.id == idQueBusco {
			return fila
		}
	}
	return -1
}

func AgregarCliente(clientes []cliente, nuevo cliente) []cliente {
	return append(clientes, nuevo)
}

func EliminarCliente(clientes []cliente, id int) []cliente {
	idx := BuscarClientePorID(clientes, id)
	if idx == -1 {
		fmt.Println("Error: Cliente no encontrado para eliminar.")
		return clientes
	}
	return append(clientes[:idx], clientes[idx+1:]...)
}

//PRODUCTOS

func BuscarProductoPorID(lista []producto, idBuscado int) int {
	for fila, algo := range lista {
		if algo.id == idBuscado {
			return fila
		}
	}
	return -1
}

func AgregarProducto(lista []producto, nuevo producto) []producto {
	return append(lista, nuevo)
}

func QuitarProducto(lista []producto, idAEliminar int) []producto {
	fila := BuscarProductoPorID(lista, idAEliminar)

	if fila == -1 {
		fmt.Println("Error: Producto no encontrado para eliminar.")
		return lista
	}

	return append(lista[:fila], lista[fila+1:]...)
}
