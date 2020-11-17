package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoSubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoSubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoSubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoSubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MetadatoSubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoMetadatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoMetadatoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoMetadatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoMetadatoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoMetadatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoMetadatoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoMetadatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoMetadatoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoMetadatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoMetadatoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "GetAllByPersona",
			Router:           `/:persona`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaTipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaTipoProduccionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaTipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaTipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaTipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaTipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaTipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaTipoProduccionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaTipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaTipoProduccionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SistemaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:PuntajeSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:PuntajeSubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:PuntajeSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:PuntajeSubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:PuntajeSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:PuntajeSubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:PuntajeSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:PuntajeSubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:PuntajeSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:PuntajeSubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:CategoriaProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:CategoriaProduccionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:CategoriaProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:CategoriaProduccionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:CategoriaProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:CategoriaProduccionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:CategoriaProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:CategoriaProduccionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:CategoriaProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:CategoriaProduccionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}
