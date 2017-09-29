//Definicion del negocio para las Suscripciones
package negocio

import (
	"entidades"
	"log"
	"accesodatos"
	"os"
	"github.com/satori/go.uuid"
)

//Estrucutra que representa un BD para Suscripciones
type SuscripcionBD struct{
	debug			*log.Logger
	fatal			*log.Logger
	dao				*accesodatos.SuscripcionDAO
}
	

//Metodo que genera un BD de transacciones
func NewSuscripcionBD(con *accesodatos.ConexionBD) *SuscripcionBD{
	var obj SuscripcionBD

	obj.dao=accesodatos.NewSuscripcionDAO(con)
	obj.debug=log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	obj.fatal=log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &obj
}





// Metodos publicos ------------------------------------------------------
func (bd *SuscripcionBD) RecuperaSuscripciones(t *entidades.Suscripcion) []entidades.Suscripcion{
	return bd.dao.RecuperaRegistros(t)
}

func (bd *SuscripcionBD) RecuperaSuscripcion(id string) entidades.Suscripcion{
	return bd.dao.RecuperaRegistroPorId(id)
}

func (bd *SuscripcionBD) AgregaSuscripcion(t *entidades.Suscripcion) bool{
	//asigna datos pendientes
	t.Id=uuid.NewV4().String()

	bd.debug.Println("Agrega suscripcion con ID: "+t.Id);
	return bd.dao.InsertaRegistro(t)
}