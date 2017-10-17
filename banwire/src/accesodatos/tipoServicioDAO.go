//Definicion del DAO para el manejo de Tipos de servicio
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
)

//Estrucutra que representa un DAO para Tipos de Servicio
type TipoServicioDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de estado de Notificacion
func NewTipoServicioDAO( con *ConexionBD) *TipoServicioDAO{
	return &TipoServicioDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *TipoServicioDAO) RecuperaRegistros(t *entidades.TipoServicio) []entidades.TipoServicio{
	var vals []interface{}
	var regs []entidades.TipoServicio
	var obj entidades.TipoServicio
	pos:=1
	
	dao.query="SELECT T.id_tipo_servicio, T.descripcion, T.activo, T.creador, T.creacion, T.modificador, T.modificacion "+
		"FROM ctl_tipo_servicio T "+
		"WHERE 1=1 "

	if(t!=nil && t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf(" AND T.descripcion=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.debug.Println("Intenta recuperar TiposServicio");
	
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

func (dao *TipoServicioDAO) RecuperaRegistroPorId(id string) entidades.TipoServicio{
	var obj entidades.TipoServicio

	dao.query="SELECT T.id_tipo_servicio, T.descripcion, T.activo, T.creador, T.creacion, T.modificador, T.modificacion "+
		"FROM ctl_tipo_servicio T "+
		"WHERE T.id_tipo_servicio=$1"
	dao.debug.Println("Intenta recuperar un TipoServicio por Id");

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

func (dao *TipoServicioDAO) InsertaRegistro(t *entidades.TipoServicio) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO ctl_tipo_servicio "+
		"(id_tipo_servicio, descripcion, activo, creador, creacion) "+
		"VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)"
	dao.debug.Println("Intenta agregar un TipoServicio");

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

func  (dao *TipoServicioDAO) ActualizaRegistro(t *entidades.TipoServicio) bool{
	var vals []interface{}
	var resTmp sql.Result
	var rowsAffected int64
	pos:=1

	dao.query="UPDATE ctl_tipo_servicio "+
		"SET "

	if(t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf("descripcion=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" , modificador=$%d, modificacion=CURRENT_TIMESTAMP WHERE id_tipo_servicio=$%d", pos, pos+1)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar un TipoServicio con %d parametros\n", len(vals));
	
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
