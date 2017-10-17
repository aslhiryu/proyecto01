//Definicion del DAO para el manejo de Transacciones
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
)

//Estrucutra que representa un DAO para Transacciones
type TransaccionDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de estado de Notificacion
func NewTransaccionDAO( con *ConexionBD) *TransaccionDAO{
	return &TransaccionDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *TransaccionDAO) RecuperaRegistros(t *entidades.Transaccion) []entidades.Transaccion{
	var vals []interface{}
	var regs []entidades.Transaccion
	var obj entidades.Transaccion
	pos:=1
	
	dao.query="SELECT T.id_transaccion, T.autorizacion, T.monto, T.id_estado_transaccion, T.concepto, T.referencia, T.id_terminal, T.id_tipo_transaccion, T.id_tipo_servicio, T.id_comercio, T.id_tarjeta, T.id_tipo_movimiento, T.id_tipo_iso, T.ticket, T.comercial_asignado, T.fecha, "+
		"E.descripcion desc_estado_transaccion, "+
		"TER.descripcion nombre_terminal, "+
		"TT.descripcion desc_tipo_transaccion, "+
		"TS.descripcion desc_tipo_servicio, "+
		"C.razon_social nombre_comercio, "+
		"TAR.digitos digitos_tarjeta, "+
		"TM.descripcion desc_tipo_movimiento, "+
		"TI.descripcion desc_tipo_iso "+
		"FROM transaccion T "+
		"LEFT JOIN ctl_estado_transaccion E ON T.id_estado_transaccion=E.id_estado_transaccion "+
		"LEFT JOIN terminal TER ON T.id_terminal=TER.id_terminal "+
		"LEFT JOIN ctl_tipo_transaccion TT ON T.id_tipo_transaccion=TT.id_tipo_transaccion "+
		"LEFT JOIN ctl_tipo_servicio TS ON T.id_tipo_servicio=TS.id_tipo_servicio "+
		"LEFT JOIN comercio C ON T.id_comercio=C.id_comercio "+
		"LEFT JOIN tarjeta TAR ON T.id_tarjeta=TAR.id_tarjeta "+
		"LEFT JOIN ctl_tipo_movimiento TM ON T.id_tipo_movimiento=TM.id_tipo_movimiento "+
		"LEFT JOIN ctl_tipo_iso TI ON T.id_tipo_iso=TI.id_tipo_iso "+
		"WHERE 1=1 "

	if(t!=nil && t.Autorizacion!=""){
		dao.query=dao.query+fmt.Sprintf(" AND T.autorizacion=$%d", pos)
		vals=append(vals, t.Autorizacion)
		pos++
	}
	dao.debug.Println("Intenta recuperar Transacciones");
	
	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//recupero los registros
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	dao.dbResult, dao.dbError=dao.dbStmt.Query(vals...)
	dao.validaError()
	for dao.dbResult.Next() {
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Autorizacion, &obj.Monto, &obj.Estatus.Id, &obj.Concepto, &obj.Referencia, &obj.Terminal.Id, &obj.Tipo.Id, &obj.Servicio.Id, &obj.Comercio.Id, &obj.Tarjeta.Id, &obj.Movimiento.Id, &obj.Iso.Id, &obj.NumeroTicket, &obj.ComercialAsignado, &obj.Fecha, &obj.Estatus.Nombre, &obj.Terminal.Nombre, &obj.Tipo.Nombre, &obj.Servicio.Nombre, &obj.Comercio.RazonSocial, &obj.Tarjeta.Digitos, &obj.Movimiento.Nombre, &obj.Iso.Nombre)
		regs=append(regs, obj)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return regs
}

func (dao *TransaccionDAO) RecuperaRegistroPorId(id string) entidades.Transaccion{
	var obj entidades.Transaccion

	dao.query="SELECT T.id_transaccion, T.autorizacion, T.monto, T.id_estado_transaccion, T.concepto, T.referencia, T.id_terminal, T.id_tipo_transaccion, T.id_tipo_servicio, T.id_comercio, T.id_tarjeta, T.id_tipo_movimiento, T.id_tipo_iso, T.ticket, T.comercial_asignado, T.fecha, "+
		"E.descripcion desc_estado_transaccion, "+
		"TER.descripcion nombre_terminal, "+
		"TT.descripcion desc_tipo_transaccion, "+
		"TS.descripcion desc_tipo_servicio, "+
		"C.razon_social nombre_comercio, "+
		"TAR.digitos digitos_tarjeta, "+
		"TM.descripcion desc_tipo_movimiento, "+
		"TI.descripcion desc_tipo_iso "+
		"FROM transaccion T "+
		"LEFT JOIN ctl_estado_transaccion E ON T.id_estado_transaccion=E.id_estado_transaccion "+
		"LEFT JOIN terminal TER ON T.id_terminal=TER.id_terminal "+
		"LEFT JOIN ctl_tipo_transaccion TT ON T.id_tipo_transaccion=TT.id_tipo_transaccion "+
		"LEFT JOIN ctl_tipo_servicio TS ON T.id_tipo_servicio=TS.id_tipo_servicio "+
		"LEFT JOIN comercio C ON T.id_comercio=C.id_comercio "+
		"LEFT JOIN tarjeta TAR ON T.id_tarjeta=TAR.id_tarjeta "+
		"LEFT JOIN ctl_tipo_movimiento TM ON T.id_tipo_movimiento=TM.id_tipo_movimiento "+
		"LEFT JOIN ctl_tipo_iso TI ON T.id_tipo_iso=TI.id_tipo_iso "+
		"WHERE T.id_transaccion=$1"
	dao.debug.Println("Intenta recuperar un Transaccion por Id");

	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//recupero los datos del registro
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	dao.dbResult, dao.dbError=dao.dbStmt.Query(id)
	dao.validaError()

	if dao.dbResult.Next() {
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Autorizacion, &obj.Monto, &obj.Estatus.Id, &obj.Concepto, &obj.Referencia, &obj.Terminal.Id, &obj.Tipo.Id, &obj.Servicio.Id, &obj.Comercio.Id, &obj.Tarjeta.Id, &obj.Movimiento.Id, &obj.Iso.Id, &obj.NumeroTicket, &obj.ComercialAsignado, &obj.Fecha, &obj.Estatus.Nombre, &obj.Terminal.Nombre, &obj.Tipo.Nombre, &obj.Servicio.Nombre, &obj.Comercio.RazonSocial, &obj.Tarjeta.Digitos, &obj.Movimiento.Nombre, &obj.Iso.Nombre)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return obj
}

func (dao *TransaccionDAO) InsertaRegistro(t *entidades.Transaccion) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO transaccion "+
		"(id_transaccion, autorizacion, monto, id_estado_transaccion, concepto, referencia, id_terminal, id_tipo_transaccion, id_tipo_servicio, id_comercio, id_tarjeta, id_tipo_movimiento, id_tipo_iso, ticket, comercial_asignado, fecha) "+
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, CURRENT_TIMESTAMP)"
	dao.debug.Println("Intenta agregar un Transaccion");

	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//agrego registro
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	resTmp, dao.dbError=dao.dbStmt.Exec(t.Id, t.Autorizacion, t.Monto, t.Estatus.Id, t.Concepto, t.Referencia, t.Terminal.Id, t.Tipo.Id, t.Servicio.Id, t.Comercio.Id, t.Tarjeta.Id, t.Movimiento.Id, t.Iso.Id, t.NumeroTicket, t.ComercialAsignado, t.Fecha)
	dao.validaError()
	rowsAffected, dao.dbError=resTmp.RowsAffected()
	dao.validaError()
	dao.debug.Printf("Agrego %d registros\n", rowsAffected)

	if(rowsAffected>0){
		return true
	} else{
		return false
	}
}

