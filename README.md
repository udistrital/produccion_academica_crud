# produccion_academica_crud
API para la gestión de la produccion académica

--Api de producion academica con CI--
CI deploy with lambda - S3
Drone 0.8 
produccion_academica_crud master/develop

## Requirements
Go version >= 1.8.

## Preparación:
    Para descargar el API, usar el comando:
        - go get github.com/udistrital/produccion_academica_crud

## Run

Definir los valores de las siguientes variables de entorno:

 - `PRODUCCION_ACADEMICA_CRUD__SQLCONN`: Dirección de la base de datos (formato 'postgres://user:pass@host:port/db_name?sslmode=disable&search_path=schema')
 - `API_PRODUCCION_ACADEMICA_HTTP_PORT`: Puerto en el que se expone la api.
 - `RUN_MODE`: Modo de ejecución del api.


Ejemplo: API_PRODUCCION_ACADEMICA_HTTP_PORT=8083 PRODUCCION_ACADEMICA_CRUD__SQLCONN=postgres://test:test@localhost:5432/core?sslmode=disable&search_path=produccion_academica bee run (Usar flags --downdoc=true --gendoc=true para generar documentacion)

## Dev run

Entorno de desarrollo creado con docker, para correr el proyecto modificar archivo .env con las variables de entorno, si se quiere que la base de datos se genere automaticamente dejar el esquema en blanco.

## MODELO
![produccion_academica_crud](https://user-images.githubusercontent.com/14035745/66079435-a8cdf300-e529-11e9-9f4e-7ca2b0807c06.png)
