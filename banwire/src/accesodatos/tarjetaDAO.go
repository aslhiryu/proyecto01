//Definicion del DAO para el manejo de Tarjetas
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

//Estrucutra que representa un DAO para Tarjetas
type TarjetaDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de tarjetas
func NewTarjetaDAO( con *ConexionBD) *TarjetaDAO{
	return &TarjetaDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *TarjetaDAO) RecuperaRegistros(t *entidades.Tarjeta) []entidades.Tarjeta{
	var vals []interface{}
	var regs []entidades.Tarjeta
	var obj entidades.Tarjeta
	pos:=1
	
	dao.query="SELECT id_tarjeta, digitos, bine, marca, emisor, vigencia, token, ultimo_cobro, creacion, pais, tipo_tarjeta, cliente, estado "+
		"FROM tarjeta "+
		"WHERE 1=1 "

	if(t!=nil && t.TipoTarjeta!=""){
		dao.query=dao.query+fmt.Sprintf(" AND id_tipo_tarjeta=$%d", pos)
		vals=append(vals, t.TipoTarjeta)
		pos++
	}
	if(t!=nil && t.Cliente!=""){
		dao.query=dao.query+fmt.Sprintf(" AND cliente=$%d", pos)
		vals=append(vals, t.Cliente)
		pos++
	}
	dao.debug.Println("Intenta recuperar Tarjetas");
	
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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Digitos, &obj.Bine, &obj.Marca, &obj.Emisor, &obj.Vigencia, &obj.Token, &obj.UltimoCobro, &obj.Creacion, &obj.Pais, &obj.TipoTarjeta, &obj.Cliente, &obj.Estado)
		regs=append(regs, obj)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return regs
}

func (dao *TarjetaDAO) RecuperaRegistroPorId(id string) entidades.Tarjeta{
	var obj entidades.Tarjeta

	dao.query="SELECT id_tarjeta, digitos, bine, marca, emisor, vigencia, token, ultimo_cobro, creacion, pais, tipo_tarjeta, cliente, estado "+
		"FROM tarjeta "+
		"WHERE id_tarjeta=$1"
	dao.debug.Println("Intenta recuperar una Tarjeta por Id");

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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.Digitos, &obj.Bine, &obj.Marca, &obj.Emisor, &obj.Vigencia, &obj.Token, &obj.UltimoCobro, &obj.Creacion, &obj.Pais, &obj.TipoTarjeta, &obj.Cliente, &obj.Estado)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return obj
}

func (dao *TarjetaDAO) InsertaRegistro(t *entidades.Tarjeta) bool{
	var resTmp sql.Result
	var rowsAffected int64

	dao.query="INSERT INTO tarjeta "+
		"(id_tarjeta, digitos, bine, marca, emisor, vigencia, token, ultimo_cobro, creacion, pais, tipo_tarjeta, cliente, estado) "+
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
	dao.debug.Println("Intenta agregar un Tarjeta");

	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//agrego registro
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	resTmp, dao.dbError=dao.dbStmt.Exec(t.Id, t.Digitos, t.Bine, t.Marca, t.Emisor, t.Vigencia, t.Token, t.UltimoCobro, t.Creacion, t.Pais, t.TipoTarjeta, t.Cliente, t.Estado)
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

func  (dao *TarjetaDAO) ActualizaRegistro(t *entidades.Tarjeta) bool{
	var vals []interface{}
	var resTmp sql.Result
	var rowsAffected int64
	pos:=1

	dao.query="UPDATE tarjeta "+
		"SET "

	if(t.Bine>0){
		dao.query=dao.query+fmt.Sprintf("bine=$%d", pos)
		vals=append(vals, t.Bine)
		pos++
	}
	if(t.Marca!=""){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("marca=$%d", pos)
		vals=append(vals, t.Marca)
		pos++
	}
	if(t.Emisor!=""){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("emisor=$%d", pos)
		vals=append(vals, t.Emisor)
		pos++
	}
	if(!time.Time.IsZero(t.Vigencia)){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("vigencia=$%d", pos)
		vals=append(vals, t.Vigencia)
		pos++
	}
	if(t.Token!=""){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("token=$%d", pos)
		vals=append(vals, t.Token)
		pos++
	}
	if(t.UltimoCobro>0){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("ultimo_cobro=$%d", pos)
		vals=append(vals, t.UltimoCobro)
		pos++
	}
	if(t.TipoTarjeta!=""){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("tipo_tarjeta=$%d", pos)
		vals=append(vals, t.TipoTarjeta)
		pos++
	}
	if(t.Cliente!=""){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("cliente=$%d", pos)
		vals=append(vals, t.Cliente)
		pos++
	}
	if(t.Estado!=""){
		if(pos>1){
			dao.query=dao.query+", "
		}
		dao.query=dao.query+fmt.Sprintf("estado=$%d", pos)
		vals=append(vals, t.Estado)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" WHERE id_tarjeta=$%d", pos)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar una Tarjeta con %d parametros\n", len(vals));
	
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
