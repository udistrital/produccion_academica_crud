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