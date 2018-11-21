package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Traduccion struct {
	Id               int               `orm:"column(id);pk"`
	TipoTraduccion   *TipoTraduccion   `orm:"column(tipo_traduccion);rel(fk)"`
	Titulo           string               `orm:"column(titulo)"`
	NombreOriginal   string               `orm:"column(nombre_original)"`
	Autor            string               `orm:"column(autor)"`
	IdiomaOriginal   string               `orm:"column(idioma_original)"`
	IdiomaTraducido  string               `orm:"column(idioma_traducido)"`
	Ano              int               `orm:"column(ano)"`
	Mes              int               `orm:"column(mes)"`
	MedioDivulgacion *MedioDivulgacion `orm:"column(medio_divulgacion);rel(fk)"`
	Edicion          int               `orm:"column(edicion)"`
	Serie            int               `orm:"column(serie)"`
}

func (t *Traduccion) TableName() string {
	return "traduccion"
}

func init() {
	orm.RegisterModel(new(Traduccion))
}

// AddTraduccion insert a new Traduccion into database and returns
// last inserted Id on success.
func AddTraduccion(m *Traduccion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTraduccionById retrieves Traduccion by Id. Returns error if
// Id doesn't exist
func GetTraduccionById(id int) (v *Traduccion, err error) {
	o := orm.NewOrm()
	v = &Traduccion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTraduccion retrieves all Traduccion matches certain condition. Returns empty list if
// no records exist
func GetAllTraduccion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Traduccion)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Traduccion
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTraduccion updates Traduccion by Id and returns error if
// the record to be updated doesn't exist
func UpdateTraduccionById(m *Traduccion) (err error) {
	o := orm.NewOrm()
	v := Traduccion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTraduccion deletes Traduccion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTraduccion(id int) (err error) {
	o := orm.NewOrm()
	v := Traduccion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Traduccion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
