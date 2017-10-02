//Definicion del negocio para los Tipos de Tarjeta
package negocio

import (
	"entidades"
	"log"
	"accesodatos"
	"os"
	"github.com/satori/go.uuid"
	"errors"
)

//Estrucutra que representa un BD para Tipos de Tarjeta
type TipoTarjetaBD struct{
	debug			*log.Logger
	fatal			*log.Logger
	dao				*accesodatos.TipoTarjetaDAO
}
	

//Metodo que genera un BD de tipos de tarjeta
func NewTipoTarjetaBD(con *accesodatos.ConexionBD) *TipoTarjetaBD{
	var obj TipoTarjetaBD

	obj.dao=accesodatos.NewTipoTarjetaDAO(con)
	obj.debug=log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	obj.fatal=log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &obj
}





// Metodos publicos ------------------------------------------------------
func (bd *TipoTarjetaBD) RecuperaTiposTarjeta(t *entidades.TipoTarjeta) []entidades.TipoTarjeta{
	return bd.dao.RecuperaRegistros(t)
}

func (bd *TipoTarjetaBD) RecuperaTipoTarjeta(id string) entidades.TipoTarjeta{
	return bd.dao.RecuperaRegistroPorId(id)
}

func (bd *TipoTarjetaBD) AgregaTipoTarjeta(t *entidades.TipoTarjeta) (bool, error){
	//valido si no existe un registro con el mismo nombre
	var rTmp entidades.TipoTarjeta
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