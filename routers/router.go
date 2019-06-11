// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/produccion_academica_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_produccion",
			beego.NSInclude(
				&controllers.TipoProduccionController{},
			),
		),

		beego.NSNamespace("/subtipo_produccion",
			beego.NSInclude(
				&controllers.SubtipoProduccionController{},
			),
		),

		beego.NSNamespace("/produccion_academica",
			beego.NSInclude(
				&controllers.ProduccionAcademicaController{},
			),
		),

		beego.NSNamespace("/dato_adicional_produccion_academica",
			beego.NSInclude(
				&controllers.DatoAdicionalProduccionAcademicaController{},
			),
		),

		beego.NSNamespace("/soporte_produccion_academica",
			beego.NSInclude(
				&controllers.SoporteProduccionAcademicaController{},
			),
		),

		beego.NSNamespace("/autor_produccion_academica",
			beego.NSInclude(
				&controllers.AutorProduccionAcademicaController{},
			),
		),

		beego.NSNamespace("/dato_adicional_subtipo_produccion",
			beego.NSInclude(
				&controllers.DatoAdicionalSubtipoProduccionController{},
			),
		),

		beego.NSNamespace("/tipo_dato_adicional",
			beego.NSInclude(
				&controllers.TipoDatoAdicionalController{},
			),
		),

		beego.NSNamespace("/estado_autor_produccion",
			beego.NSInclude(
				&controllers.EstadoAutorProduccionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
