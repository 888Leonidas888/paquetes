package greet

//para crear paquetes este debe llamarse igual que la carpeta que lo contiene
//ejemplo 'greet'

//para declarar una función,variable, constante o estructura como privada o pública
//quiere decir que puede ser llamada desde otro paquete o no
//dependera de la 'de la primera letra' ejm:
//'nombre' como es minúscula es privada solo puede ser llamada en este paquete
//'Nombre' esta forma de declaración indica que se puede llamar desde fuera
var emoji = "🙋‍♀️"

func English() string {
	return "Hi" + emoji
}

func Italian() string {
	return "Ciao" + emoji
}
