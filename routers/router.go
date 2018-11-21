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

		beego.NSNamespace("/medio_divulgacion",
			beego.NSInclude(
				&controllers.MedioDivulgacionController{},
			),
		),

		beego.NSNamespace("/medio_publicacion",
			beego.NSInclude(
				&controllers.MedioPublicacionController{},
			),
		),

		beego.NSNamespace("/produccion_artes_arqu_diseno",
			beego.NSInclude(
				&controllers.ProduccionArtesArquDisenoController{},
			),
		),

		beego.NSNamespace("/tipo_publicacion_libro",
			beego.NSInclude(
				&controllers.TipoPublicacionLibroController{},
			),
		),

		beego.NSNamespace("/articulo",
			beego.NSInclude(
				&controllers.ArticuloController{},
			),
		),

		beego.NSNamespace("/otra_publicacion",
			beego.NSInclude(
				&controllers.OtraPublicacionController{},
			),
		),

		beego.NSNamespace("/otro_documento",
			beego.NSInclude(
				&controllers.OtroDocumentoController{},
			),
		),

		beego.NSNamespace("/tipo_disciplina",
			beego.NSInclude(
				&controllers.TipoDisciplinaController{},
			),
		),

		beego.NSNamespace("/traduccion",
			beego.NSInclude(
				&controllers.TraduccionController{},
			),
		),

		beego.NSNamespace("/libro",
			beego.NSInclude(
				&controllers.LibroController{},
			),
		),

		beego.NSNamespace("/produccion_tecnica",
			beego.NSInclude(
				&controllers.ProduccionTecnicaController{},
			),
		),

		beego.NSNamespace("/tipo_articulo",
			beego.NSInclude(
				&controllers.TipoArticuloController{},
			),
		),

		beego.NSNamespace("/tipo_otra_publicacion",
			beego.NSInclude(
				&controllers.TipoOtraPublicacionController{},
			),
		),

		beego.NSNamespace("/tipo_produccion_tecnica",
			beego.NSInclude(
				&controllers.TipoProduccionTecnicaController{},
			),
		),

		beego.NSNamespace("/tipo_traduccion",
			beego.NSInclude(
				&controllers.TipoTraduccionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
