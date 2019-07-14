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

		beego.NSNamespace("/metadato_produccion_academica",
			beego.NSInclude(
				&controllers.MetadatoProduccionAcademicaController{},
			),
		),

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

		beego.NSNamespace("/metadato_subtipo_produccion",
			beego.NSInclude(
				&controllers.MetadatoSubtipoProduccionController{},
			),
		),

		beego.NSNamespace("/tipo_metadato",
			beego.NSInclude(
				&controllers.TipoMetadatoController{},
			),
		),

		beego.NSNamespace("/produccion_academica",
			beego.NSInclude(
				&controllers.ProduccionAcademicaController{},
			),
		),

		beego.NSNamespace("/estado_autor_produccion",
			beego.NSInclude(
				&controllers.EstadoAutorProduccionController{},
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
	)
	beego.AddNamespace(ns)
}
