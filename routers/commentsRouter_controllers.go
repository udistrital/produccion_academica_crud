package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ArticuloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ArticuloController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ArticuloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ArticuloController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ArticuloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ArticuloController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ArticuloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ArticuloController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ArticuloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ArticuloController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:LibroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:LibroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:LibroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:LibroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:LibroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:LibroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:LibroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:LibroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:LibroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:LibroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioDivulgacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioDivulgacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioDivulgacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioDivulgacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioDivulgacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioDivulgacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioDivulgacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioDivulgacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioDivulgacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioDivulgacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioPublicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioPublicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioPublicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioPublicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:MedioPublicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtraPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtraPublicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtraPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtraPublicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtraPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtraPublicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtraPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtraPublicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtraPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtraPublicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtroDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtroDocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtroDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtroDocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtroDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtroDocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtroDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtroDocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtroDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:OtroDocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionArtesArquDisenoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionArtesArquDisenoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionArtesArquDisenoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionArtesArquDisenoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionArtesArquDisenoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionArtesArquDisenoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionArtesArquDisenoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionArtesArquDisenoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionArtesArquDisenoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionArtesArquDisenoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionTecnicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionTecnicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionTecnicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionTecnicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:ProduccionTecnicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoArticuloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoArticuloController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoArticuloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoArticuloController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoArticuloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoArticuloController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoArticuloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoArticuloController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoArticuloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoArticuloController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDisciplinaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDisciplinaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDisciplinaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDisciplinaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDisciplinaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDisciplinaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDisciplinaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDisciplinaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDisciplinaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoDisciplinaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoOtraPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoOtraPublicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoOtraPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoOtraPublicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoOtraPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoOtraPublicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoOtraPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoOtraPublicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoOtraPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoOtraPublicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionTecnicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionTecnicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionTecnicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionTecnicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoProduccionTecnicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoPublicacionLibroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoPublicacionLibroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoPublicacionLibroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoPublicacionLibroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoPublicacionLibroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoPublicacionLibroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoPublicacionLibroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoPublicacionLibroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoPublicacionLibroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoPublicacionLibroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoTraduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoTraduccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoTraduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoTraduccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoTraduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoTraduccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoTraduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoTraduccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoTraduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TipoTraduccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TraduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TraduccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TraduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TraduccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TraduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TraduccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TraduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TraduccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TraduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/produccion_academica_crud/controllers:TraduccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
