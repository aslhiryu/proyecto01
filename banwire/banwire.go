package main

import (
	"fmt"
	"entidades"
	"negocio"
)

type TransaccionLocal struct{
	*entidades.Transaccion
}

func main(){
	var conn entidades.ConexionBD
	conn.User="gestion"
	conn.Pass="gestion"
	conn.Database="banwire_gestion"

/*	var ent entidades.Transaccion
	ent.Autorizacion=10
	ent.Monto=25.1
	ent.Terminal="aaaaaaaaaaaaaaaa"
	ent.Servicio="bbbbbbbbbbbbbb"
	ent.Suscripcion="cccccccccccc"

	var bd *negocio.TransaccionBD
	bd=negocio.NewTransaccionBD(&conn)*/

/*	var ent entidades.Suscripcion
	ent.Plan="aaaaaaaaaaaaaaaa"
	ent.Tarjeta="bbbbbbbbbbbbbb"

	bd:=negocio.NewSuscripcionBD(&conn)*/
	
	var ent entidades.Plan
	ent.Nombre="plan XXXX"
	ent.Comercio="bbbbbbbbbbbbbb"

	bd:=negocio.NewPlanBD(&conn)
	
	//fmt.Printf("Probando: %d \n", accesodatos.RecuperaNumeroRegistrosTransaccion())

//	var res bool
//	fmt.Println("Prueba Agrega")
//	res=dao.InsertaRegistro(&ent)
//	fmt.Printf("Agrego: %v\n", res)

//	ent=dao.RecuperaRegistroPorId("dasdasds")
//	fmt.Printf("Fecha del registro: %v\n", ent.Fecha)

//	var entAct entidades.Transaccion
//	entAct.Id="dasdasds"
//	entAct.Fecha=time.Now()
//	res=dao.ActualizaRegistro(&entAct)
//	fmt.Printf("Actualizo: %v\n", res)

/*	res2:=bd.RecuperaTransacciones(nil)
	fmt.Printf("Registros: %v\n", len(res2))
	//fmt.Printf("Registros: %v\n", res2)

	res:=bd.RecuperaTransaccion("dasdasds")
	fmt.Printf("Registro: %v\n", res)

	res3:=bd.AgregaTransaccion(&ent)
	fmt.Printf("Agrego: %v\n", res3)*/


/*	res2:=bd.RecuperaSuscripciones(nil)
	fmt.Printf("Registros: %v\n", len(res2))
	fmt.Printf("Registros: %v\n", res2)

	res:=bd.RecuperaSuscripcion("24664c4b-e7d4-4f8f-9745-d54b283a3ce1")
	fmt.Printf("Registro: %v\n", res)

	res3:=bd.AgregaSuscripcion(&ent)
	fmt.Printf("Agrego: %v\n", res3)*/


	res2:=bd.RecuperaPlanes(nil)
	fmt.Printf("Registros: %v\n", len(res2))
	fmt.Printf("Registros: %v\n", res2)

	res:=bd.RecuperaPlan("381a3ced-9bc8-40d6-a044-4dc518d6106b")
	fmt.Printf("Registro: %v\n", res)

	res3:=bd.AgregaPlan(&ent)
	fmt.Printf("Agrego: %v\n", res3)
}