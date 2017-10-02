//Definicion del negocio para los Estados de Tarjeta
package negocio

import (
	"entidades"
	"log"
	"accesodatos"
	"os"
	"github.com/satori/go.uuid"
	"errors"
)

//Estrucutra que representa un BD para Estados de Tarjeta
type EstadoTarjetaBD struct{
	debug			*log.Logger
	fatal			*log.Logger
	dao				*accesodatos.EstadoTarjetaDAO
}
	

//Metodo que genera un BD de estados de tarjeta
func NewEstadoTarjetaBD(con *accesodatos.ConexionBD) *EstadoTarjetaBD{
	var obj EstadoTarjetaBD

	obj.dao=accesodatos.NewEstadoTarjetaDAO(con)
	obj.debug=log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	obj.fatal=log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &obj
}





// Metodos publicos ------------------------------------------------------
func (bd *EstadoTarjetaBD) RecuperaEstadosTarjeta(t *entidades.EstadoTarjeta) []entidades.EstadoTarjeta{
	return bd.dao.RecuperaRegistros(t)
}

func (bd *EstadoTarjetaBD) RecuperaEstadoTarjeta(id string) entidades.EstadoTarjeta{
	return bd.dao.RecuperaRegistroPorId(id)
}

func (bd *EstadoTarjetaBD) AgregaEstadoTarjeta(t *entidades.EstadoTarjeta) (bool, error){
	//valido si no existe un registro con el mismo nombre
	var rTmp entidades.EstadoTarjeta
	rTmp.Nombre=t.Nombre

	bd.debug.Println("Intento agregar un tipo de tarjeta");
	lTmp:=bd.dao.RecuperaRegistros(&rTmp)
	if( lTmp!=nil && len(lTmp)>0 ){
		return false, errors.New("Ya existe un tipo con ese nombre")
	}

	//asigna datos pendientes
	t.Id=uuid.NewV4().String()

	bd.debug.Println("Agrega tipo de tarjeta con ID: "+t.Id);
	return bd.dao.InsertaRegistro(t), nil
}