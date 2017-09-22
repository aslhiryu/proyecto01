//Definicion del DAO para el manejo de Transacciones
package accesodatos

import (
	"database/sql"
	"entidades"
	"log"
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"time"
)

//Estrucutra que representa un DAO para Transacciones
type TransaccionDAO struct{
	conexion		*entidades.ConexionBD
	debug			*log.Logger
	fatal			*log.Logger
	query			string
	dbError			error
	dbConnection	*sql.DB
	dbStmt			*sql.Stmt
	dbResult		*sql.Rows
}

//Metodo que genera un DAO de transacciones
func NewTransaccionDAO( con *entidades.ConexionBD) *TransaccionDAO{
	var obj TransaccionDAO

	obj.conexion=con
	obj.debug=log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	obj.fatal=log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &obj
}




// Metodos privados ------------------------------------------------------

func (dao *TransaccionDAO) generaConexion(){
	dbinfo:=fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dao.conexion.User, dao.conexion.Pass, dao.conexion.Database)
	dao.dbConnection, dao.dbError=sql.Open("postgres", dbinfo)
	dao.validaError()
}

func (dao *TransaccionDAO) validaError(){
	if(dao.dbError!=nil){
		dao.fatal.Printf("Error en la conexion: %s\n", dao.dbError)
	}
}




// Metodos publicos ------------------------------------------------------

func (dao *TransaccionDAO) RecuperaRegistros(t *entidades.Transaccion) []entidades.Transaccion{
	var vals []interface{}
	var regs []entidades.Transaccion
	var obj entidades.Transaccion
	pos:=1
	
	dao.query="SELECT id_transaccion, num_autorizacion, monto, fecha, id_terminal, id_servicio, id_suscripcion "+
		"FROM transaccion "+
		"WHERE 1=1 "

	if(t!=nil && t.Autorizacion>0){
		dao.query=dao.query+fmt.Sprintf(" AND num_autorizacion=$%d", pos)
		vals=append(vals, t.Autorizacion)
		pos++
	}
	if(t!=nil && t.Terminal!=""){
		dao.query=dao.query+fmt.Sprintf(" AND id_terminal=$%d", pos)
		vals=append(vals, t.Terminal)
		pos++
	}
	if(t!=nil && t.Servicio!=""){
		dao.query=dao.query+fmt.Sprintf(" AND id_servicio=$%d", pos)
		vals=append(vals, t.Servicio)
		pos++
	}
	if(t!=nil && t.Suscripcion!=""){
		dao.query=dao.query+fmt.Sprintf(" AND id_suscripcion=$%d", pos)
		vals=append(vals, t.Suscripcion)
		pos++
	}
	dao.debug.Println("Intenta recuperar Transacciones");
	
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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Autorizacion, &obj.Monto, &obj.Fecha, &obj.Terminal, &obj.Servicio, &obj.Suscripcion)
		regs=append(regs, obj)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return regs
}

func (dao *TransaccionDAO) RecuperaRegistroPorId(id string) entidades.Transaccion{
	var obj entidades.Transaccion

	dao.query="SELECT id_transaccion, num_autorizacion, monto, fecha, id_terminal, id_servicio, id_suscripcion "+
		"FROM transaccion "+
		"WHERE id_transaccion=$1"
	dao.debug.Println("Intenta recuperar una Transaccion por Id");

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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Autorizacion, &obj.Monto, &obj.Fecha, &obj.Terminal, &obj.Servicio, &obj.Suscripcion)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return obj
}

func (dao *TransaccionDAO) InsertaRegistro(t *entidades.Transaccion) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO transaccion "+
		"(id_transaccion, num_autorizacion, monto, fecha, id_terminal, id_servicio, id_suscripcion) "+
		"VALUES($1, $2, $3, $4, $5, $6, $7)"
	dao.debug.Println("Intenta agregar una Transaccion");

	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//agrego registro
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	resTmp, dao.dbError=dao.dbStmt.Exec(t.Id, t.Autorizacion, t.Monto, t.Fecha, t.Terminal, t.Servicio, t.Suscripcion)
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

func  (dao *TransaccionDAO) ActualizaRegistro(t *entidades.Transaccion) bool{
	var vals []interface{}
	var resTmp sql.Result
	var rowsAffected int64
	pos:=1

	dao.query="UPDATE transaccion "+
		"SET "

	if(t.Autorizacion>0){
		dao.query=dao.query+fmt.Sprintf("num_autorizacion=$%d", pos)
		vals=append(vals, t.Autorizacion)
		pos++
	}
	if(t.Monto>0){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("monto=$%d", pos)
		vals=append(vals, t.Monto)
		pos++
	}
	if(!time.Time.IsZero(t.Fecha)){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("fecha=$%d", pos)
		vals=append(vals, t.Fecha)
		pos++
	}
	if(t.Terminal!=""){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("id_terminal=$%d", pos)
		vals=append(vals, t.Terminal)
		pos++
	}
	if(t.Servicio!=""){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("id_servicio=$%d", pos)
		vals=append(vals, t.Servicio)
		pos++
	}
	if(t.Suscripcion!=""){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("id_suscripcion=$%d", pos)
		vals=append(vals, t.Suscripcion)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" WHERE id_transaccion=$%d", pos)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar una Transaccion con %d parametros\n", len(vals));
	
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
