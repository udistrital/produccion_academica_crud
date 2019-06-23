package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:AutorProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalSubtipoProduccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalSubtipoProduccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalSubtipoProduccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalSubtipoProduccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalSubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:DatoAdicionalSubtipoProduccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:EstadoAutorProduccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SoporteProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SoporteProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SoporteProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SoporteProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SoporteProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SoporteProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SoporteProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SoporteProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SoporteProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SoporteProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:SubtipoProduccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDatoAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDatoAdicionalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDatoAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDatoAdicionalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDatoAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDatoAdicionalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDatoAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDatoAdicionalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDatoAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDatoAdicionalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "GetAllByEnte",
            Router: `/:ente`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TrProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
