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

			produccionAcademica := autor.ProduccionAcademica

			var autoresProduccion []AutorProduccionAcademica
			if _, err := o.QueryTable(new(AutorProduccionAcademica)).RelatedSel().Filter("produccion_academica",produccionAcademica.Id).All(&autoresProduccion); err != nil{
				return nil, err
			}

			var datosAdicionales []DatoAdicionalProduccionAcademica
			if _, err := o.QueryTable(new(DatoAdicionalProduccionAcademica)).RelatedSel().Filter("ProduccionAcademica__Id",produccionAcademica.Id).All(&datosAdicionales); err != nil{
				return nil, err
			}

			v = append(v,map[string]interface{}{
				"Id": produccionAcademica.Id,
				"Titulo": produccionAcademica.Titulo,
				"Resumen": produccionAcademica.Resumen,
				"Fecha": produccionAcademica.Fecha,
				"SubtipoProduccion":produccionAcademica.SubtipoProduccion,
				"Autores": &autoresProduccion,
				"DatosAdicionales": &datosAdicionales,
			})
		}

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
				return
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

// UpdateTransaccionProduccionAcademica updates ProduccionAcademica by Id and returns error if
// the record to be updated doesn't exist
func UpdateTransaccionProduccionAcademica(m *TrProduccionAcademica) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	v := ProduccionAcademica{Id: m.ProduccionAcademica.Id}
	// ascertain id exists in the database
	if errTr := o.Read(&v); errTr == nil {
		var num int64
		if num, errTr = o.Update(m.ProduccionAcademica); errTr == nil {
			fmt.Println("Number of records updated in database:", num)

			for _, v := range *m.DatosAdicionales {
					var datoAdicional DatoAdicionalProduccionAcademica
					if errTr = o.QueryTable(new(DatoAdicionalProduccionAcademica)).RelatedSel().Filter("DatoAdicionalSubtipoProduccion__Id",v.DatoAdicionalSubtipoProduccion.Id).Filter("ProduccionAcademica__Id",m.ProduccionAcademica.Id).One(&datoAdicional); err == nil{
						datoAdicional.Valor = v.Valor
						if _, errTr = o.Update(&datoAdicional,"Valor"); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}
					} else {
						err = errTr
						fmt.Println(err)
						_ = o.Rollback()
						return
					}		
			}

			_ = o.Commit()
		}	else {
			err = errTr
			fmt.Println(err)
			_ = o.Rollback()
			return
		}
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}
	return
}

// TrDeleteProduccionAcademica deletes ProduccionAcademica by Id and returns error if
// the record to be deleted doesn't exist
func TrDeleteProduccionAcademica(id int) (err error) {
	o := orm.NewOrm()
	v := ProduccionAcademica{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProduccionAcademica{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}