//Definicion del DAO para el manejo de Planes
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
)

//Estrucutra que representa un DAO para Planes
type PlanDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de planes
func NewPlanDAO( con *ConexionBD) *PlanDAO{
	return &PlanDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *PlanDAO) RecuperaRegistros(t *entidades.Plan) []entidades.Plan{
	var vals []interface{}
	var regs []entidades.Plan
	var obj entidades.Plan
	pos:=1
	
	dao.query="SELECT id_plan, nombre, id_comercio "+
		"FROM plan "+
		"WHERE 1=1 "

	if(t!=nil && t.Comercio!=""){
		dao.query=dao.query+fmt.Sprintf(" AND id_comercio=$%d", pos)
		vals=append(vals, t.Comercio)
		pos++
	}
	dao.debug.Println("Intenta recuperar Planes");
	
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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Nombre, &obj.Comercio)
		regs=append(regs, obj)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return regs
}

func (dao *PlanDAO) RecuperaRegistroPorId(id string) entidades.Plan{
	var obj entidades.Plan

	dao.query="SELECT id_plan, nombre, id_comercio "+
		"FROM plan "+
		"WHERE id_plan=$1"
	dao.debug.Println("Intenta recuperar una Plan por Id");

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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Nombre, &obj.Comercio)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return obj
}

func (dao *PlanDAO) InsertaRegistro(t *entidades.Plan) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO plan "+
		"(id_plan, nombre, id_comercio) "+
		"VALUES($1, $2, $3)"
	dao.debug.Println("Intenta agregar un Plan");

	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//agrego registro
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	resTmp, dao.dbError=dao.dbStmt.Exec(t.Id, t.Nombre, t.Comercio)
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

func  (dao *PlanDAO) ActualizaRegistro(t *entidades.Plan) bool{
	var vals []interface{}
	var resTmp sql.Result
	var rowsAffected int64
	pos:=1

	dao.query="UPDATE plan "+
		"SET "

	if(t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf("nombre=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	if(t.Comercio!=""){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("id_comercio=$%d", pos)
		vals=append(vals, t.Comercio)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" WHERE id_plan=$%d", pos)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar una Plan con %d parametros\n", len(vals));
	
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
