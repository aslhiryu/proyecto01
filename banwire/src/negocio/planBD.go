//Definicion del negocio para las Planes
package negocio

import (
	"entidades"
	"log"
	"accesodatos"
	"os"
	"github.com/satori/go.uuid"
)

//Estrucutra que representa un BD para Planes
type PlanBD struct{
	debug			*log.Logger
	fatal			*log.Logger
	dao				*accesodatos.PlanDAO
}
	

//Metodo que genera un BD de transacciones
func NewPlanBD(con *entidades.ConexionBD) *PlanBD{
	var obj PlanBD

	obj.dao=accesodatos.NewPlanDAO(con)
	obj.debug=log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	obj.fatal=log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &obj
}





// Metodos publicos ------------------------------------------------------
func (bd *PlanBD) RecuperaPlanes(t *entidades.Plan) []entidades.Plan{
	return bd.dao.RecuperaRegistros(t)
}

func (bd *PlanBD) RecuperaPlan(id string) entidades.Plan{
	return bd.dao.RecuperaRegistroPorId(id)
}

func (bd *PlanBD) AgregaPlan(t *entidades.Plan) bool{
	//asigna datos pendientes
	t.Id=uuid.NewV4().String()

	bd.debug.Println("Agrega plan con ID: "+t.Id);
	return bd.dao.InsertaRegistro(t)
}