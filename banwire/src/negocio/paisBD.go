//Definicion del negocio para los Paises
package negocio

import (
	"entidades"
	"log"
	"accesodatos"
	"os"
	"errors"
)

//Estrucutra que representa un BD para Paises
type PaisBD struct{
	debug			*log.Logger
	fatal			*log.Logger
	dao				*accesodatos.PaisDAO
}
	

//Metodo que genera un BD de transacciones
func NewPaisBD(con *accesodatos.ConexionBD) *PaisBD{
	var obj PaisBD

	obj.dao=accesodatos.NewPaisDAO(con)
	obj.debug=log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	obj.fatal=log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &obj
}





// Metodos publicos ------------------------------------------------------
func (bd *PaisBD) RecuperaPaises(t *entidades.Pais) []entidades.Pais{
	return bd.dao.RecuperaRegistros(t)
}

func (bd *PaisBD) RecuperaPais(id string) entidades.Pais{
	return bd.dao.RecuperaRegistroPorId(id)
}

func (bd *PaisBD) AgregaPais(t *entidades.Pais) (bool, error){
	//valido si no existe un registro con el mismo nombre
	var rTmp entidades.Pais
	rTmp.Nombre=t.Nombre

	bd.debug.Println("Intento agregar un pais");
	lTmp:=bd.dao.RecuperaRegistros(&rTmp)
	if( lTmp!=nil && len(lTmp)>0 ){
		return false, errors.New("Ya existe un pais con ese nombre")
	}

	bd.debug.Println("Agrega pais con ID: "+t.Id);
	return bd.dao.InsertaRegistro(t), nil
}