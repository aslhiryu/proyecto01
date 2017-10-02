//Definicion del DAO para el manejo de Tipos de Tarjeta
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
)

//Estrucutra que representa un DAO para Tipos de Tarjeta
type TipoTarjetaDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de tipos de tarjeta
func NewTipoTarjetaDAO( con *ConexionBD) *TipoTarjetaDAO{
	return &TipoTarjetaDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *TipoTarjetaDAO) RecuperaRegistros(t *entidades.TipoTarjeta) []entidades.TipoTarjeta{
	var vals []interface{}
	var regs []entidades.TipoTarjeta
	var obj entidades.TipoTarjeta
	pos:=1
	
	dao.query="SELECT id_tipo_tarjeta, nombre "+
		"FROM tipo_tarjeta "+
		"WHERE 1=1 "

	if(t!=nil && t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf(" AND nombre=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.debug.Println("Intenta recuperar TiposTarjeta");
	
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

func (dao *TipoTarjetaDAO) RecuperaRegistroPorId(id string) entidades.TipoTarjeta{
	var obj entidades.TipoTarjeta

	dao.query="SELECT id_tipo_tarjeta, nombre "+
		"FROM tipo_tarjeta "+
		"WHERE id_tipo_tarjeta=$1"
	dao.debug.Println("Intenta recuperar un TipoTarjeta por Id");

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

func (dao *TipoTarjetaDAO) InsertaRegistro(t *entidades.TipoTarjeta) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO tipo_tarjeta "+
		"(id_tipo_tarjeta, nombre) "+
		"VALUES($1, $2)"
	dao.debug.Println("Intenta agregar un TipoTarjeta");

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

func  (dao *TipoTarjetaDAO) ActualizaRegistro(t *entidades.TipoTarjeta) bool{
	var vals []interface{}
	var resTmp sql.Result
	var rowsAffected int64
	pos:=1

	dao.query="UPDATE tipo_tarjeta "+
		"SET "

	if(t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf("nombre=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" WHERE id_tipo_tarjeta=$%d", pos)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar un TipoTarjeta con %d parametros\n", len(vals));
	
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
