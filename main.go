package main

import "fmt"

func crearUsuario() {
	fmt.Println("Usuario creado exitosamente!")
}

func listarUsuarios() {
	fmt.Println("Listado de usuarios")
}
func actualizarUsuarios() {
	fmt.Println("A usuarios")
}

func eliminarUsuario() {
	fmt.Println("Usuario creado exitosamente!")
}

func main() {

	var option string

	fmt.Println("A) Crear")
	fmt.Println("B) Listar")
	fmt.Println("C) Actualizar")
	fmt.Println("D) Eliminar")

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
