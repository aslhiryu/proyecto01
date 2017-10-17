//Definicion del DAO para el manejo de Emisores de Tarjeta
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
)

//Estrucutra que representa un DAO para Emisores de Tarjeta
type EmisorTarjetaDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de estado de Notificacion
func NewEmisorTarjetaDAO( con *ConexionBD) *EmisorTarjetaDAO{
	return &EmisorTarjetaDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *EmisorTarjetaDAO) RecuperaRegistros(t *entidades.EmisorTarjeta) []entidades.EmisorTarjeta{
	var vals []interface{}
	var regs []entidades.EmisorTarjeta
	var obj entidades.EmisorTarjeta
	pos:=1
	
	dao.query="SELECT E.id_emisor_tarjeta, E.nombre, E.activo, E.creador, E.creacion, E.modificador, E.modificacion "+
		"FROM  ctl_emisor_tarjeta E "+
		"WHERE 1=1 "

	if(t!=nil && t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf(" AND E.nombre=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.debug.Println("Intenta recuperar EmisoresTarjeta");
	
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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Nombre, &obj.Activo, &obj.Creador.Id, &obj.Creacion, &obj.Modificador.Id, &obj.Modificacion)
		regs=append(regs, obj)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return regs
}

func (dao *EmisorTarjetaDAO) RecuperaRegistroPorId(id string) entidades.EmisorTarjeta{
	var obj entidades.EmisorTarjeta

	dao.query="SELECT E.id_emisor_tarjeta, E.nombre, E.activo, E.creador, E.creacion, E.modificador, E.modificacion "+
		"FROM  ctl_emisor_tarjeta E "+
		"WHERE E.id_emisor_tarjeta=$1"
	dao.debug.Println("Intenta recuperar un EmisorTarjeta por Id");

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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Nombre, &obj.Activo, &obj.Creador.Id, &obj.Creacion, &obj.Modificador.Id, &obj.Modificacion)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return obj
}

func (dao *EmisorTarjetaDAO) InsertaRegistro(t *entidades.EmisorTarjeta) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO  ctl_emisor_tarjeta "+
		"(id_emisor_tarjeta, nombre, activo, creador, creacion) "+
		"VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)"
	dao.debug.Println("Intenta agregar un EmisorTarjeta");

	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//agrego registro
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	resTmp, dao.dbError=dao.dbStmt.Exec(t.Id, t.Nombre, t.Activo, t.Creador.Id)
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

func  (dao *EmisorTarjetaDAO) ActualizaRegistro(t *entidades.EmisorTarjeta) bool{
	var vals []interface{}
	var resTmp sql.Result
	var rowsAffected int64
	pos:=1

	dao.query="UPDATE  ctl_emisor_tarjeta "+
		"SET "

	if(t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf("nombre=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" , modificador=$%d, modificacion=CURRENT_TIMESTAMP WHERE id_emisor_tarjeta=$%d", pos, pos+1)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar un EmisorTarjeta con %d parametros\n", len(vals));
	
	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//actualizo registro
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	resTmp, dao.dbError=dao.dbStmt.Exec(vals...)
	dao.validaError()
	rowsAffected, dao.dbError=resTmp.RowsAffected()
	dao.validaError()
	dao.debug.Printf("Actualizo %d registros\n", rowsAffected)

	if(rowsAffected>0){
		return true
	} else{
		return false
	}
}
