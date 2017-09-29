//Definicion del DAO base
package accesodatos

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
)

//Estructura que representa una conexion a BD
type ConexionBD struct{
	User 		string
	Pass		string
	Database	string
	Server		string
}


type BaseDAO interface{
	queryParaSelectId() string
}

//Estrucutra que representa un DAO
type GenericDAO struct{
	conexion		*ConexionBD
	debug			*log.Logger
	fatal			*log.Logger
	query			string
	dbError			error
	dbConnection	*sql.DB
	dbStmt			*sql.Stmt
	dbResult		*sql.Rows
	BaseDAO
}	

func NewGenericDAO(con *ConexionBD) *GenericDAO{
	var obj GenericDAO
	
	obj.conexion=con
	obj.debug=log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	obj.fatal=log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &obj
}

// Metodos privados ------------------------------------------------------

func (dao *GenericDAO) generaConexion(){
	dbinfo:=fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dao.conexion.User, dao.conexion.Pass, dao.conexion.Database)
	dao.dbConnection, dao.dbError=sql.Open("postgres", dbinfo)
	dao.validaError()
}

func (dao *GenericDAO) validaError(){
	if(dao.dbError!=nil){
		dao.fatal.Printf("Error en la conexion: %s\n", dao.dbError)
	}
}



