//Definicion del DAO para el manejo de Giros
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
)

//Estrucutra que representa un DAO para Giros
type GiroDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de estado de Notificacion
func NewGiroDAO( con *ConexionBD) *GiroDAO{
	return &GiroDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *GiroDAO) RecuperaRegistros(t *entidades.Giro) []entidades.Giro{
	var vals []interface{}
	var regs []entidades.Giro
	var obj entidades.Giro
	pos:=1
	
	dao.query="SELECT G.id_giro, G.descripcion, G.activo, G.creador, G.creacion, G.modificador, G.modificacion "+
		"FROM ctl_giro G "+
		"WHERE 1=1 "

	if(t!=nil && t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf(" AND G.descripcion=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.debug.Println("Intenta recuperar Giros");
	
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

func (dao *GiroDAO) RecuperaRegistroPorId(id string) entidades.Giro{
	var obj entidades.Giro

	dao.query="SELECT G.id_giro, G.descripcion, G.activo, G.creador, G.creacion, G.modificador, G.modificacion "+
		"FROM ctl_giro G "+
		"WHERE G.id_giro=$1"
	dao.debug.Println("Intenta recuperar un Giro por Id");

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

func (dao *GiroDAO) InsertaRegistro(t *entidades.Giro) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO ctl_giro "+
		"(id_giro, descripcion, activo, creador, creacion) "+
		"VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)"
	dao.debug.Println("Intenta agregar un Giro");

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

func  (dao *GiroDAO) ActualizaRegistro(t *entidades.Giro) bool{
	var vals []interface{}
	var resTmp sql.Result
	var rowsAffected int64
	pos:=1

	dao.query="UPDATE ctl_giro "+
		"SET "

	if(t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf("descripcion=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" , modificador=$%d, modificacion=CURRENT_TIMESTAMP WHERE id_giro=$%d", pos, pos+1)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar un Giro con %d parametros\n", len(vals));
	
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
