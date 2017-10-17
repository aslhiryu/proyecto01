//Definicion del DAO para el manejo de Paises
package accesodatos

import (
//	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
)

//Estrucutra que representa un DAO para Paises
type PaisDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de pais
func NewPaisDAO( con *ConexionBD) *PaisDAO{
	return &PaisDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *PaisDAO) RecuperaRegistros(t *entidades.Pais) []entidades.Pais{
	var vals []interface{}
	var regs []entidades.Pais
	var obj entidades.Pais
	pos:=1
	
	dao.query="SELECT P.id_pais, P.nombre "+
		"FROM ctl_pais P "+
		"WHERE 1=1 "

	if(t!=nil && t.Nombre!=""){
		dao.query=dao.query+fmt.Sprintf(" AND P.nombre=$%d", pos)
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

	dao.query="SELECT SELECT P.id_pais, P.nombre "+
		"FROM ctl_pais P "+
		"WHERE P.id_pais=$1"
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
