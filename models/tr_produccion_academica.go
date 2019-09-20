package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type TrProduccionAcademica struct {
	ProduccionAcademica *ProduccionAcademica
	Autores             *[]AutorProduccionAcademica
	Metadatos           *[]MetadatoProduccionAcademica
}

// GetProduccionesAcademicasByPersona Transacción para consultar todas las producciones con toda la información de las mismas
func GetProduccionesAcademicasByPersona(persona int) (v []interface{}, err error) {
	o := orm.NewOrm()
	var autores []*AutorProduccionAcademica
	if _, err := o.QueryTable(new(AutorProduccionAcademica)).RelatedSel().Filter("persona_id", persona).Filter("ProduccionAcademicaId__Activo", true).All(&autores); err == nil {

		for _, autor := range autores {
			produccionAcademica := autor.ProduccionAcademicaId

			var autoresProduccion []AutorProduccionAcademica
			if _, err := o.QueryTable(new(AutorProduccionAcademica)).RelatedSel().Filter("ProduccionAcademicaId__Id", produccionAcademica.Id).All(&autoresProduccion); err != nil {
				return nil, err
			}

			var metadatos []MetadatoProduccionAcademica
			if _, err := o.QueryTable(new(MetadatoProduccionAcademica)).RelatedSel().Filter("ProduccionAcademicaId__Id", produccionAcademica.Id).All(&metadatos); err != nil {
				return nil, err
			}

			v = append(v, map[string]interface{}{
				"Id":                  produccionAcademica.Id,
				"Titulo":              produccionAcademica.Titulo,
				"Resumen":             produccionAcademica.Resumen,
				"Fecha":               produccionAcademica.Fecha,
				"SubtipoProduccionId": produccionAcademica.SubtipoProduccionId,
				"Autores":             &autoresProduccion,
				"Metadatos":           &metadatos,
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
			v.ProduccionAcademicaId.Id = int(idProduccion)
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}
		}

		for _, v := range *m.Metadatos {
			v.ProduccionAcademicaId.Id = int(idProduccion)
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
		if num, errTr = o.Update(m.ProduccionAcademica, "Titulo", "Resumen", "Fecha"); errTr == nil {
			fmt.Println("Number of records updated in database:", num)

			for _, v := range *m.Metadatos {
				fmt.Println("metadatos", m.Metadatos)
				var metadato MetadatoProduccionAcademica
				if errTr = o.QueryTable(new(MetadatoProduccionAcademica)).RelatedSel().Filter("MetadatoSubtipoProduccionId__Id", v.MetadatoSubtipoProduccionId.Id).Filter("ProduccionAcademicaId__Id", m.ProduccionAcademica.Id).One(&metadato); err == nil {

					if metadato.Valor != v.Valor {
						metadato.Valor = v.Valor
						metadato.FechaModificacion = v.FechaModificacion
					}

					if metadato.Id != 0 {
						if _, errTr = o.Update(&metadato, "Valor"); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}
					} else {
						metadato.ProduccionAcademicaId = m.ProduccionAcademica
						metadato.MetadatoSubtipoProduccionId = v.MetadatoSubtipoProduccionId
						metadato.FechaCreacion = v.FechaCreacion
						if _, errTr = o.Insert(&metadato); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}
					}

				} else {
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
		// if num, err = o.Delete(&ProduccionAcademica{Id: id}); err == nil {
		// fmt.Println("Number of records deleted in database:", num)
		if num, err = o.Update(&ProduccionAcademica{Id: id, Activo: false, FechaModificacion: time_bogota.TiempoBogotaFormato()}, "Activo", "FechaModificacion"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
