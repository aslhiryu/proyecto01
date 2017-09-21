//Definicion de entidades
package entidades

import "time"

//Estructura que representa una conexion a BD
type ConexionBD struct{
	User 		string
	Pass		string
	Database	string
	Server		string
}

//Estructura que representa una Transaccion
type Transaccion struct{
	Id 				string
	Autorizacion 	int32
	Monto 			float64
	Fecha 			time.Time
	Terminal 		string
	Servicio 		string
	Suscripcion 	string
}

//func NewTransaccion() *Transaccion{
//	return &Transaccion{
//		id: "",
//	}
//}