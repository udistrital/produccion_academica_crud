# produccion_academica_crud
API para la gestión de la produccion académica, publicaciones y afines

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `produccion_academica_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/planesticud/produccion_academica_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `API_PRODUCCION_ACADEMICA_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `PRODUCCION_ACADEMICA_CRUD__PGUSER`: Usuario de la base de datos
 - `PRODUCCION_ACADEMICA_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `PRODUCCION_ACADEMICA_CRUD__PGURLS`: Host de conexión
 - `PRODUCCION_ACADEMICA_CRUD__PGDB`: Nombre de la base de datos
 - `PRODUCCION_ACADEMICA_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
API_PRODUCCION_ACADEMICA_HTTP_PORT=9012 PRODUCCION_ACADEMICA_CRUD__PGUSER=user PRODUCCION_ACADEMICA_CRUD__PGPASS=password PRODUCCION_ACADEMICA_CRUD__PGURLS=localhost PRODUCCION_ACADEMICA_CRUD__PGDB=bd PRODUCCION_ACADEMICA_CRUD__SCHEMA=schema_new bee run

## MODELO
![image](https://github.com/planesticud/produccion_academica_crud/blob/develop/modelo_produccion_academica_crud.png).
