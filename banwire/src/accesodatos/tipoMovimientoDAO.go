//Definicion del DAO para el manejo de Tipos de Movimiento
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
)

//Estrucutra que representa un DAO para Tipos de Movimiento
type TipoMovimientoDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de estado de Notificacion
func NewTipoMovimientoDAO( con *ConexionBD) *TipoMovimientoDAO{
	return &TipoMovimientoDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *TipoMovimientoDAO) RecuperaRegistros(t *entidades.TipoMovimiento) []entidades.TipoMovimiento{
	var vals []interface{}
	var regs []entidades.TipoMovimiento
	var obj entidades.TipoMovimiento
	pos:=1
	
	dao.query="SELECT T.id_tipo_movimiento, T.descripcion, T.activo, T.creador, T.creacion, T.modificador, T.modificacion "+
		"FROM ctl_tipo_movimiento T "+
		"WHERE 1=1 "

	if(t!=nil && t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf(" AND T.descripcion=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.debug.Println("Intenta recuperar TiposMovimiento");
	
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

func (dao *TipoMovimientoDAO) RecuperaRegistroPorId(id string) entidades.TipoMovimiento{
	var obj entidades.TipoMovimiento

	dao.query="SELECT T.id_tipo_movimiento, T.descripcion, T.activo, T.creador, T.creacion, T.modificador, T.modificacion "+
		"FROM ctl_tipo_movimiento T "+
		"WHERE T.id_tipo_movimiento=$1"
	dao.debug.Println("Intenta recuperar un TipoMovimiento por Id");

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

func (dao *TipoMovimientoDAO) InsertaRegistro(t *entidades.TipoMovimiento) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO ctl_tipo_movimiento "+
		"(id_tipo_movimiento, descripcion, activo, creador, creacion) "+
		"VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)"
	dao.debug.Println("Intenta agregar un TipoMovimiento");

	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//agrego registro
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	resTmp, dao.dbError=dao.dbStmt.Exec(t.Id, t.Nombre, t.Activo, T.Creador.Id)
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

func  (dao *TipoMovimientoDAO) ActualizaRegistro(t *entidades.TipoMovimiento) bool{
	var vals []interface{}
	var resTmp sql.Result
	var rowsAffected int64
	pos:=1

	dao.query="UPDATE ctl_tipo_movimiento "+
		"SET "

	if(t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf("descripcion=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" , modificador=$%d, modificacion=CURRENT_TIMESTAMP WHERE id_tipo_movimiento=$%d", pos, pos+1)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar un TipoMovimiento con %d parametros\n", len(vals));
	
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