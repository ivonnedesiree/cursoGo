package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var id int
var users map[int]User

func crearUsuario() {
	fmt.Print("Ingresa un valor para username: ")
	username := readLine()
	fmt.Print("Ingresa un valor para email: ")
	email := readLine()
	fmt.Print("Ingresa un valor para edad: ")
	age, err := strconv.Atoi(readLine())
	if err != nil {
		panic("No es posible convertir a un tipo entero")
	}

	id++
	user := User{id, username, email, age}
	users[id] = user
	fmt.Println(user)
	fmt.Println("Usuario creado exitosamente!")
}

func listarUsuarios() {
	fmt.Println("Listado de usuarios")
}
func actualizarUsuarios() {
	fmt.Println("Usuario actualizado exitosamente!")
}

func eliminarUsuario() {
	fmt.Println("Usuario creado exitosamente!")
}

func readLine() string {
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor!")
	} else {
		return strings.TrimSpace(option) //eliminamos los espacios si no, no funciona
		//option = strings.TrimSuffix(option, "\n") //eliminamos el salto de linea
	}

}

func main() {
	var option string
	users = make(map[int]User)
	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("A) Crear")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")

		fmt.Print("Ingresa una opcion válida: ")
		option = readLine()
		if option == "quit" || option == "q" {
			break
		}

		//fmt.Println(option)

		switch option {
		case "a", "crear":
			crearUsuario()
		case "b", "listar":
			listarUsuarios()
		case "c", "actualizar":
			actualizarUsuarios()
		case "d", "eliminar":
			eliminarUsuario()
		default:
			fmt.Println("Opcion no valida")
		}

	}

}
