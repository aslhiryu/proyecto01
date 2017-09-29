//Definicion del DAO para el manejo de Suscripciones
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
)

//Estrucutra que representa un DAO para Suscripciones
type SuscripcionDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de Suscripciones
func NewSuscripcionDAO( con *ConexionBD) *SuscripcionDAO{
	return &SuscripcionDAO{NewGenericDAO(con)}
}






// Metodos publicos ------------------------------------------------------

func (dao *SuscripcionDAO) RecuperaRegistros(t *entidades.Suscripcion) []entidades.Suscripcion{
	var vals []interface{}
	var regs []entidades.Suscripcion
	var obj entidades.Suscripcion
	pos:=1
	
	dao.query="SELECT id_suscripcion, id_plan, id_tarjeta "+
		"FROM suscripcion "+
		"WHERE 1=1 "

	if(t!=nil && t.Plan!=""){
		dao.query=dao.query+fmt.Sprintf(" AND id_plan=$%d", pos)
		vals=append(vals, t.Plan)
		pos++
	}
	if(t!=nil && t.Tarjeta!=""){
		dao.query=dao.query+fmt.Sprintf(" AND id_tarjeta=$%d", pos)
		vals=append(vals, t.Tarjeta)
		pos++
	}
	dao.debug.Println("Intenta recuperar Suscripciones");
	
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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Plan, &obj.Tarjeta)
		regs=append(regs, obj)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return regs
}

func (dao *SuscripcionDAO) RecuperaRegistroPorId(id string) entidades.Suscripcion{
	var obj entidades.Suscripcion

	dao.query="SELECT id_suscripcion, id_plan, id_tarjeta "+
		"FROM suscripcion "+
		"WHERE id_suscripcion=$1"
	dao.debug.Println("Intenta recuperar una Suscripcion por Id");

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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Plan, &obj.Tarjeta)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return obj
}

func (dao *SuscripcionDAO) InsertaRegistro(t *entidades.Suscripcion) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO suscripcion "+
		"(id_suscripcion, id_plan, id_tarjeta) "+
		"VALUES($1, $2, $3)"
	dao.debug.Println("Intenta agregar una Suscripcion");

	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//agrego registro
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	resTmp, dao.dbError=dao.dbStmt.Exec(t.Id, t.Plan, t.Tarjeta)
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

func  (dao *SuscripcionDAO) ActualizaRegistro(t *entidades.Suscripcion) bool{
	var vals []interface{}
	var resTmp sql.Result
	var rowsAffected int64
	pos:=1

	dao.query="UPDATE suscripcion "+
		"SET "

	if(t.Plan!=""){
		dao.query=dao.query+fmt.Sprintf("id_plan=$%d", pos)
		vals=append(vals, t.Plan)
		pos++
	}
	if(t.Tarjeta!=""){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("id_tarjeta=$%d", pos)
		vals=append(vals, t.Tarjeta)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" WHERE id_suscripcion=$%d", pos)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar una Suscripcion con %d parametros\n", len(vals));
	
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
