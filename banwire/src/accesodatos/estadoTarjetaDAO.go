//Definicion del DAO para el manejo de Estados de Tarjeta
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
)

//Estrucutra que representa un DAO para Estados de Tarjeta
type EstadoTarjetaDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de estado de tarjeta
func NewEstadoTarjetaDAO( con *ConexionBD) *EstadoTarjetaDAO{
	return &EstadoTarjetaDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *EstadoTarjetaDAO) RecuperaRegistros(t *entidades.EstadoTarjeta) []entidades.EstadoTarjeta{
	var vals []interface{}
	var regs []entidades.EstadoTarjeta
	var obj entidades.EstadoTarjeta
	pos:=1
	
	dao.query="SELECT id_estado_tarjeta, nombre "+
		"FROM estado_tarjeta "+
		"WHERE 1=1 "

	if(t!=nil && t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf(" AND nombre=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.debug.Println("Intenta recuperar EstadosTarjeta");
	
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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Nombre)
		regs=append(regs, obj)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return regs
}

func (dao *EstadoTarjetaDAO) RecuperaRegistroPorId(id string) entidades.EstadoTarjeta{
	var obj entidades.EstadoTarjeta

	dao.query="SELECT id_estado_tarjeta, nombre "+
		"FROM estado_tarjeta "+
		"WHERE id_estado_tarjeta=$1"
	dao.debug.Println("Intenta recuperar un EstadoTarjeta por Id");

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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Nombre)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return obj
}

func (dao *EstadoTarjetaDAO) InsertaRegistro(t *entidades.EstadoTarjeta) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO estado_tarjeta "+
		"(id_estado_tarjeta, nombre) "+
		"VALUES($1, $2)"
	dao.debug.Println("Intenta agregar un EstadoTarjeta");

	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//agrego registro
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	resTmp, dao.dbError=dao.dbStmt.Exec(t.Id, t.Nombre)
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

func  (dao *EstadoTarjetaDAO) ActualizaRegistro(t *entidades.EstadoTarjeta) bool{
	var vals []interface{}
	var resTmp sql.Result
	var rowsAffected int64
	pos:=1

	dao.query="UPDATE estado_tarjeta "+
		"SET "

	if(t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf("nombre=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" WHERE id_estado_tarjeta=$%d", pos)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar un EstadoTarjeta con %d parametros\n", len(vals));
	
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
