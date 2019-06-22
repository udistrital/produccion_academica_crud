package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrProduccionAcademica struct {
	ProduccionAcademica            *ProduccionAcademica
	Autores  *[]AutorProduccionAcademica
	DatosAdicionales        *[]DatoAdicionalProduccionAcademica
}

// GetProduccionesAcademicas Transacción para consultar todas las producciones con toda la información de las mismas
func GetProduccionesAcademicasByEnte(ente int) (v []interface{}, err error) {
	o := orm.NewOrm()
	var autores []*AutorProduccionAcademica
	if _, err := o.QueryTable(new(AutorProduccionAcademica)).RelatedSel().Filter("ente",ente).All(&autores); err == nil{
		for _, autor := range autores {
			v = append(v,TrProduccionAcademica{
				ProduccionAcademica: autor.ProduccionAcademica,
				Autores: nil,
				DatosAdicionales: nil,
			})
		}
		fmt.Println(v)
		return v, nil
	}
	return nil, err
}

// AddTransaccionProduccionAcademica Transacción para registrar toda la información de una producción
func AddTransaccionProduccionAcademica(m *TrProduccionAcademica) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	
	if idProduccion, errTr := o.Insert(m.ProduccionAcademica); errTr == nil {
		fmt.Println(idProduccion)

		for _, v := range *m.Autores {
			v.ProduccionAcademica.Id = int(idProduccion)
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
			}
		}

		for _, v := range *m.DatosAdicionales {
			v.ProduccionAcademica.Id = int(idProduccion)
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}
		}

		_ = o.Commit()
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}

	return
}