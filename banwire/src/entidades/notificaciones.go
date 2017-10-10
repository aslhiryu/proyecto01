//Definicion de entidades de notificacion
package entidades

import "time"


//Estructura que representa un estado de una notificacion
type EstadoNotificacion struct{
	Id					string
	Nombre				string	
	Activo				bool
	Creador				Usuario
	Creacion			time.Time
	Modificador			Usuario
	Modificacion		time.Time
}