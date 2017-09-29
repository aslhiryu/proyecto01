//Definicion de entidades
package entidades

import "time"

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

//Estructura que representa una Suscripcion
type Suscripcion struct{
	Id				string
	Plan			string
	Tarjeta			string
}

//Estructura que representa un Plan de Comercio
type Plan struct{
	Id				string
	Nombre			string
	Comercio		string
}

//Estructura que representa una Tarjeta Bancaria
type Tarjeta struct{
	Id				string
	Digitos			string
	Bine			int16
	Marca			string
	Emisor			string
	Vigencia		time.Time
	Token			string
	UltimoCobro		float64
	Creacion		time.Time
	Pais			string
	TipoTarjeta		string
	Cliente			string
	Estado			string
}