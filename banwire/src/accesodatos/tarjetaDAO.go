//Definicion del DAO para el manejo de Tarjetas
package accesodatos

import (
	"database/sql"
	"entidades"
	"fmt"
	_ "github.com/lib/pq"
)

//Estrucutra que representa un DAO para Tarjetas
type TarjetaDAO struct{
	*GenericDAO
}

//Metodo que genera un DAO de estado de Notificacion
func NewTarjetaDAO( con *ConexionBD) *TarjetaDAO{
	return &TarjetaDAO{NewGenericDAO(con)}
}





// Metodos publicos ------------------------------------------------------

func (dao *TarjetaDAO) RecuperaRegistros(t *entidades.Tarjeta) []entidades.Tarjeta{
	var vals []interface{}
	var regs []entidades.Tarjeta
	var obj entidades.Tarjeta
	pos:=1
	
	dao.query="SELECT T.id_tarjeta, T.codigo_autorizacion, T.tarejtahabiente, T.id_tipo_tarjeta, T.id_emisor_tarjeta, T.mail_usuario, T.digitos, T.bin, T.vigencia, T.marca, T.id_pais, T.ultimo_cargo, T.cvv, T.token, T.id_estado_tarjeta, T.creador, T.creacion, T.modificador, T.modificacion, "+
		"TT.descripcion desc_tipo_tarjeta, "
		"ET.nombre desc_emisor_tarjeta, "
		"P.nombre pais, "
		"ST.descripcion desc_estado_tarjeta "
		"LEFT JOIN ctl_tipo_tarjeta TT ON T.id_tipo_tarjeta=TT.id_tipo_tarjeta "
		"LEFT JOIN ctl_emisor_tarjeta ET ON T.id_emisor_tarjeta=ET.id_emisor_tarjeta "
		"LEFT JOIN ctl_pais P ON T.id_pais=P.id_pais "
		"LEFT JOIN ctl_estado_tarjeta ST ON T.id_estado_tarjeta=ET.id_estado_tarjeta "
		"FROM tarjeta T "+
		"WHERE 1=1 "

	if(t!=nil && t.Digitos!=""){
		dao.query=dao.query+fmt.Sprintf(" AND T.digitos=$%d", pos)
		vals=append(vals, t.Digitos)
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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.CodigoAutorizacion, &obj.Tarjetahabiente, &obj.Tipo.Id, &obj.Emisor.Id, &obj.MailUsuario, &obj.Digitos, &obj.Bin, &obj.Vencimiento, &obj.Marca, &obj.Pais.Id, &obj.UltimoCobro, &obj.Cvv, &obj.Token, &obj.Estatus.Id, &obj.Creador.Id, &obj.Creacion, &obj.Modificador.Id, &obj.Modificacion, &obj.Tipo.Nombre, &obj.Emisor.Nombre, &obj.Pais.Nombre, &obj.Estatus.Nombre)
		regs=append(regs, obj)
		dao.validaError()
	}
	dao.dbError=dao.dbResult.Err()
	dao.validaError()

	return regs
}

func (dao *TarjetaDAO) RecuperaRegistroPorId(id string) entidades.Tarjeta{
	var obj entidades.Tarjeta

	dao.query="SELECT T.id_tarjeta, T.codigo_autorizacion, T.tarejtahabiente, T.id_tipo_tarjeta, T.id_emisor_tarjeta, T.mail_usuario, T.digitos, T.bin, T.vigencia, T.marca, T.id_pais, T.ultimo_cargo, T.cvv, T.token, T.id_estado_tarjeta, T.creador, T.creacion, T.modificador, T.modificacion, "+
		"TT.descripcion desc_tipo_tarjeta, "
		"ET.nombre desc_emisor_tarjeta, "
		"P.nombre pais, "
		"ST.descripcion desc_estado_tarjeta "
		"LEFT JOIN ctl_tipo_tarjeta TT ON T.id_tipo_tarjeta=TT.id_tipo_tarjeta "
		"LEFT JOIN ctl_emisor_tarjeta ET ON T.id_emisor_tarjeta=ET.id_emisor_tarjeta "
		"LEFT JOIN ctl_pais P ON T.id_pais=P.id_pais "
		"LEFT JOIN ctl_estado_tarjeta ST ON T.id_estado_tarjeta=ET.id_estado_tarjeta "
		"FROM tarjeta T "+
		"WHERE T.id_tarjeta=$1"
	dao.debug.Println("Intenta recuperar un Tarjeta por Id");

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
		dao.dbError=dao.dbResult.Scan(&obj.Id, &obj.CodigoAutorizacion, &obj.Tarjetahabiente, &obj.Tipo.Id, &obj.Emisor.Id, &obj.MailUsuario, &obj.Digitos, &obj.Bin, &obj.Vencimiento, &obj.Marca, &obj.Pais.Id, &obj.UltimoCobro, &obj.Cvv, &obj.Token, &obj.Estatus.Id, &obj.Creador.Id, &obj.Creacion, &obj.Modificador.Id, &obj.Modificacion, &obj.Tipo.Nombre, &obj.Emisor.Nombre, &obj.Pais.Nombre, &obj.Estatus.Nombre)
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
		"(id_tarjeta, codigo_autorizacion, tarejtahabiente, id_tipo_tarjeta, id_emisor_tarjeta, mail_usuario, digitos, bin, vigencia, marca, id_pais, ultimo_cargo, cvv, token, id_estado_tarjeta, creador, creacion) "+
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, CURRENT_TIMESTAMP)"
	dao.debug.Println("Intenta agregar un Tarjeta");

	//realiza conexion
	dao.generaConexion();
	defer dao.dbConnection.Close()
	
	//agrego registro
	dao.debug.Println("Ejecuta el query: "+dao.query);
	dao.dbStmt, dao.dbError=dao.dbConnection.Prepare(dao.query)
	dao.validaError()
	resTmp, dao.dbError=dao.dbStmt.Exec(t.Id, t.CodigoAutorizacion, t.Tarjetahabiente, t.Tipo.Id, t.Emisor.Id, t.MailUsuario, t.Digitos, t.Bin, t.Vencimiento, t.Marca, t.Pais.Id, t.UltimoCobro, t.Cvv, t.Token, t.Estatus.Id, t.Creador.Id)
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

	if(t.Digitos!=""){
		dao.query=dao.query+fmt.Sprintf("digitos=$%d", pos)
		vals=append(vals, t.Digitos)
		pos++
	}
	dao.query=dao.query+fmt.Sprintf(" , modificador=$%d, modificacion=CURRENT_TIMESTAMP WHERE id_tarjeta=$%d", pos, pos+1)
	vals=append(vals, t.Id)
	dao.debug.Printf("Intenta actualizar un Tarjeta con %d parametros\n", len(vals));
	
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
