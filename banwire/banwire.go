package main

import (
	"fmt"
	"entidades"
	"accesodatos"
	"time"
	"github.com/satori/go.uuid"
)

type TransaccionLocal struct{
	*entidades.Transaccion
}

func main(){
	var tran entidades.Transaccion
	tran.Id=uuid.NewV4().String()
	tran.Autorizacion=10
	tran.Monto=25.1
	tran.Fecha=time.Now()
	tran.Terminal="aaaaaaaaaaaaaaaa"
	tran.Servicio="bbbbbbbbbbbbbb"
	tran.Suscripcion="cccccccccccc"

	var conn entidades.ConexionBD
	conn.User="gestion"
	conn.Pass="gestion"
	conn.Database="banwire_gestion"

	var dao *accesodatos.TransaccionDAO
	dao=accesodatos.NewTransaccionDAO(&conn)

	
	//fmt.Printf("Probando: %d \n", accesodatos.RecuperaNumeroRegistrosTransaccion())

//	var res bool
//	fmt.Println("Prueba Agrega")
//	res=dao.InsertaRegistro(&tran)
//	fmt.Printf("Agrego: %v\n", res)

//	tran=dao.RecuperaRegistroPorId("dasdasds")
//	fmt.Printf("Fecha del registro: %v\n", tran.Fecha)

//	var tranAct entidades.Transaccion
//	tranAct.Id="dasdasds"
//	tranAct.Fecha=time.Now()
//	res=dao.ActualizaRegistro(&tranAct)
//	fmt.Printf("Actualizo: %v\n", res)

	res2:=dao.RecuperaRegistros(nil)
	fmt.Printf("Registros: %v\n", len(res2))
	fmt.Printf("Registros: %v\n", res2)
}