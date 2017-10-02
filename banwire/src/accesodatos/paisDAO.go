//Definicion del DAO para el manejo de Paises
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
)

//Estrucutra que representa un DAO para Paises
type PaisDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de paises
func NewPaisDAO( con *ConexionBD) *PaisDAO{
	return &PaisDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *PaisDAO) RecuperaRegistros(t *entidades.Pais) []entidades.Pais{
	var vals []interface{}
	var regs []entidades.Pais
	var obj entidades.Pais
	pos:=1
	
	dao.query="SELECT id_pais, nombre "+
		"FROM pais "+
		"WHERE 1=1 "

	if(t!=nil && t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf(" AND nombre=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.debug.Println("Intenta recuperar Paises");
	
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

func (dao *PaisDAO) RecuperaRegistroPorId(id string) entidades.Pais{
	var obj entidades.Pais

	dao.query="SELECT id_pais, nombre "+
		"FROM pais "+
		"WHERE id_pais=$1"
	dao.debug.Println("Intenta recuperar un Pais por Id");

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

func (dao *PaisDAO) InsertaRegistro(t *entidades.Pais) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO pais "+
		"(id_pais, nombre) "+
		"VALUES($1, $2)"
	dao.debug.Println("Intenta agregar un Pais");

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

func  (dao *PaisDAO) ActualizaRegistro(t *entidades.Pais) bool{
	var vals []interface{}
	var resTmp sql.Result
	var rowsAffected int64
	pos:=1

	dao.query="UPDATE pais "+
		"SET "

	if(t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf("nombre=$%d", pos)
		vals=append(vals, t.Nombre)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" WHERE id_pais=$%d", pos)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar un Pais con %d parametros\n", len(vals));
	
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
