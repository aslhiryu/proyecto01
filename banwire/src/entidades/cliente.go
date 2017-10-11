//Definicion de entidades del cliebte
package entidades

import "time"

//Estructura que representa una tarjeta bancaria
type Tarjeta struct{
	Id					string
	CodigoAutorizacion	string	//Es lo mismo que el codigo de seguridad ???
	Tarjetahabiente		string
	Tipo 				TipoTarjeta
	Emisor				Emisor
	MailUsuario			string
	Digitos				string
	Bin 				string //que es un bin ???
	Vencimiento			time.Time
	Marca				string	//que es la marca ???
	Pais				Pais
	UltimoCobro			float64
	Cvv					string
	Token				string //existe diferencia entre tarjeta y tarjeta tokenizada ???
	Estatus				EstadoTarjeta
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}

//Estructura que representa un Tipo de Tarjeta
type TipoTarjeta struct{
	Id					string
	Nombre				string
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}

//Estructura que representa un emisor de tarjetas bancarias
type Emisor struct{
	Id					string
	Nombre				string	
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}

//Estructura que representa un estado de una tarjeta bancaria
type EstadoTarjeta struct{
	Id					string
	Nombre				string
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}
