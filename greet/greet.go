// el paquete greet provee la funcionalidad

// de saludar en diferentes lenguajes

package greet

// emoji es una variable a nivel paquete, es no exportable

var emoji = "🙋"

// English retorna saludo en inglés, es exporable

func English() string {

	return "Hi" + emoji

}

// Italian retorn saludo en italiano, es exportable

func Italian() string {

	return "Ciao" + emoji

}
