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
	
/*	var ent entidades.Plan
	ent.Nombre="plan XXXX"
	ent.Comercio="bbbbbbbbbbbbbb"

	bd:=negocio.NewPlanBD(&conn)*/
	
/*	var ent entidades.Tarjeta
	ent.Digitos="0477"
	ent.Bine=30
	ent.Marca="Alguna"
	ent.Emisor="XXX"
	ent.Vigencia=time.Now()
	ent.Token="SANOASJASDO"
	ent.UltimoCobro=300.23
	ent.Creacion=time.Now()
	ent.Pais="MX"
	ent.TipoTarjeta="bbbbbbbbbbbbbb"
	ent.Cliente="cccccccccccc"
	ent.Estado="dddddddddd"
	
	bd:=negocio.NewTarjetaBD(&conn)*/
	
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


/*	res2:=bd.RecuperaPlanes(nil)
	fmt.Printf("Registros: %v\n", len(res2))
	fmt.Printf("Registros: %v\n", res2)

	res:=bd.RecuperaPlan("381a3ced-9bc8-40d6-a044-4dc518d6106b")
	fmt.Printf("Registro: %v\n", res)

	res3:=bd.AgregaPlan(&ent)
	fmt.Printf("Agrego: %v\n", res3)*/


/*	var ent entidades.TipoTarjeta
	ent.Nombre="Tipo tarjeta 1"

	bd:=negocio.NewTipoTarjetaBD(&conn)

	res2:=bd.RecuperaTiposTarjeta(nil)
	fmt.Printf("Registros: %v\n", len(res2))
	fmt.Printf("Registros: %v\n", res2)

	res:=bd.RecuperaTipoTarjeta("381a3ced-9bc8-40d6-a044-4dc518d6106b")
	fmt.Printf("Registro: %v\n", res)

	res3, er:=bd.AgregaTipoTarjeta(&ent)
	fmt.Printf("Agrego: %v\n", res3)
	if(er!=nil){
		fmt.Printf("Error: %v\n", er)
	}*/
/*	var ent entidades.Pais
	ent.Id="MX"
	ent.Nombre="México"

	bd:=negocio.NewPaisBD(&conn)

	res2:=bd.RecuperaPaises(nil)
	fmt.Printf("Registros: %v\n", len(res2))
	fmt.Printf("Registros: %v\n", res2)

	res:=bd.RecuperaPais("MX")
	fmt.Printf("Registro: %v\n", res)

	res3, er:=bd.AgregaPais(&ent)
	fmt.Printf("Agrego: %v\n", res3)
	if(er!=nil){
		fmt.Printf("Error: %v\n", er)
	}*/
	var ent entidades.EstadoTarjeta
	ent.Nombre="Activa"

	bd:=negocio.NewEstadoTarjetaBD(&conn)

	res2:=bd.RecuperaEstadosTarjeta(nil)
	fmt.Printf("Registros: %v\n", len(res2))
	fmt.Printf("Registros: %v\n", res2)

	res:=bd.RecuperaEstadoTarjeta("d55a1381-622c-45bd-924d-75396818bcbe")
	fmt.Printf("Registro: %v\n", res)

	res3, er:=bd.AgregaEstadoTarjeta(&ent)
	fmt.Printf("Agrego: %v\n", res3)
	if(er!=nil){
		fmt.Printf("Error: %v\n", er)
	}
}