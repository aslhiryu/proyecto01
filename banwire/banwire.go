package main

import (
	"fmt"
	"entidades"
	"negocio"
	"accesodatos"
//	"time"
)

type TransaccionLocal struct{
	*entidades.Transaccion
}

func main(){
	var conn accesodatos.ConexionBD
	conn.User="gestion"
	conn.Pass="gestion"
	conn.Database="banwire_gestion"

	//prueba de transacciones
	bd:=negocio.NewTransaccionBD(&conn)
	list, er:=bd.RecuperaTransacciones(nil)
	if(er!=nil){
		fmt.Printf("Error: %v\n", er)
	} else{
		fmt.Printf("Numero: %v\n", len(list))
		fmt.Printf("Datos: %v\n", list)
	}
}