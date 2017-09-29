//Definicion del negocio para las Tarjetas
package negocio

import (
	"entidades"
	"log"
	"accesodatos"
	"os"
	"github.com/satori/go.uuid"
)

//Estrucutra que representa un BD para Tarjetas
type TarjetaBD struct{
	debug			*log.Logger
	fatal			*log.Logger
	dao				*accesodatos.TarjetaDAO
}
	

//Metodo que genera un BD de tarjetas
func NewTarjetaBD(con *accesodatos.ConexionBD) *TarjetaBD{
	var obj TarjetaBD

	obj.dao=accesodatos.NewTarjetaDAO(con)
	obj.debug=log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	obj.fatal=log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &obj
}





// Metodos publicos ------------------------------------------------------
func (bd *TarjetaBD) RecuperaTarjetas(t *entidades.Tarjeta) []entidades.Tarjeta{
	return bd.dao.RecuperaRegistros(t)
}

func (bd *TarjetaBD) RecuperaTarjeta(id string) entidades.Tarjeta{
	return bd.dao.RecuperaRegistroPorId(id)
}

func (bd *TarjetaBD) AgregaTarjeta(t *entidades.Tarjeta) bool{
	//asigna datos pendientes
	t.Id=uuid.NewV4().String()

	bd.debug.Println("Agrega tarjeta con ID: "+t.Id);
	return bd.dao.InsertaRegistro(t)
}