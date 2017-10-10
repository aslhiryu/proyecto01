//Definicion de entidades del comercio
package entidades

import "time"


//Estructura que representa un comercio
type Comercio struct{
	Id					string
	RazonSocial			string	
	Rfc 				string
	Estatus				EstadoComercio
	AgenteComercial		string
	Giro				Giro
	Comercio 			string //que es esto
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}

//Estructura que representa una terminal de un comercio
type Terminal struct{
	Id					string
	Nombre				string	
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}

//Estructura que representa un estado de un comercio
type EstadoComercio struct{
	Id					string
	Nombre				string	
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}

//Estructura que representa un giro
type Giro struct{
	Id					string
	Nombre				string	
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}
