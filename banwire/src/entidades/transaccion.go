//Definicion de entidades de trabsaccion
package entidades

import "time"


//Estructura que representa una Transaccion
type Transaccion struct{
	Id					string
	Fecha				time.Time
	Autorizacion 		string
	Digitos				string   ///validar si no es de la tarjeta
	Bin 				string	///validar si no es de la tarjeta
	Monto				float64	
	Estatus				EstadoTransaccion
	Concepto			string
	Referencia			string
	Terminal			Terminal
	Tipo 				TipoTransaccion
	Servicio 			TipoServicio
	Comercio 			Comercio
	Tarjeta 			Tarjeta
	Movimiento 			TipoMovimiento
	Iso 				TipoISO
	EstadoNotificacion 	EstadoNotificacion  //VALIDAR SI NO SE LIGA CON UNA NOTIFICACION
	NumeroTicket		string
	ComercialAsignado 	string	//que es?
}

//Estructura que representa un estado de una transaccion
type EstadoTransaccion struct{
	Id					string
	Nombre				string	
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}

//Estructura que representa un tipo de transaccion
type TipoTransaccion struct{
	Id					string
	Nombre				string	
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}

//Estructura que representa un tipo de servicio
type TipoServicio struct{
	Id					string
	Nombre				string	
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}

//Estructura que representa un tipo de movimiento
type TipoMovimiento struct{
	Id					string
	Nombre				string	
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}

//Estructura que representa un tipo de ISO
type TipoISO struct{
	Id					string
	Nombre				string	
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}


