//para utilizar nuestro paquete 'greet' y el de terceros
//debemos crear un módulo,con el siguiente comando en terminal
//go mod init [github.com/888Leonidas888/paquetes] , lo que esta entre llaves
//la ruta de nuestro repositorio hacia nuestro paquete

package main

import (
	"fmt"

	"github.com/888Leonidas888/paquetes/greet"

	"rsc.io/quote/v3"
)

func main() {
	//Usando nuestro propio paquete
	fmt.Println("Hello")
	fmt.Println("Saludo en inglés", greet.English())
	fmt.Println("Saludo en italiano", greet.Italian())
	fmt.Println("Saludo en español", greet.Spanish())

	//Usando paquete de terceros
	//para usar el paquete de un tercero debo hacerlo con el comando en terminal
	//go get [rsc.io/quote]
	fmt.Println("saludo desde un paquete de tercero", quote.HelloV3())
	fmt.Println("Versión mayor", quote.Concurrency())
}
