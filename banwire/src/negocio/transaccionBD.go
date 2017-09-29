//Definicion del negocio para las Transacciones
package negocio

import (
	"entidades"
	"log"
	"accesodatos"
	"os"
	"time"
	"github.com/satori/go.uuid"
)

//Estrucutra que representa un BD para Transacciones
type TransaccionBD struct{
	debug			*log.Logger
	fatal			*log.Logger
	dao				*accesodatos.TransaccionDAO
}
	

//Metodo que genera un BD de transacciones
func NewTransaccionBD(con *accesodatos.ConexionBD) *TransaccionBD{
	var obj TransaccionBD

	obj.dao=accesodatos.NewTransaccionDAO(con)
	obj.debug=log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	obj.fatal=log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &obj
}





// Metodos publicos ------------------------------------------------------
func (bd *TransaccionBD) RecuperaTransacciones(t *entidades.Transaccion) []entidades.Transaccion{
	return bd.dao.RecuperaRegistros(t)
}

func (bd *TransaccionBD) RecuperaTransaccion(id string) entidades.Transaccion{
	return bd.dao.RecuperaRegistroPorId(id)
}

func (bd *TransaccionBD) AgregaTransaccion(t *entidades.Transaccion) bool{
	//asigna datos pendientes
	t.Fecha=time.Now()
	t.Id=uuid.NewV4().String()

	bd.debug.Println("Agrega transaccion con ID: "+t.Id);
	return bd.dao.InsertaRegistro(t)
}