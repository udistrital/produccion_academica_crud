-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.2-alpha1
-- PostgreSQL version: 10.0
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: core | type: DATABASE --
-- -- DROP DATABASE IF EXISTS core;
-- CREATE DATABASE core
-- 	ENCODING = 'UTF8'
-- 	LC_COLLATE = 'en_US'
-- 	LC_CTYPE = 'en_US'
-- 	TABLESPACE = pg_default
-- 	OWNER = postgres;
-- -- ddl-end --
--

-- object: produccion_academica | type: SCHEMA --
-- DROP SCHEMA IF EXISTS produccion_academica CASCADE;
CREATE SCHEMA produccion_academica;
-- ddl-end --
COMMENT ON SCHEMA produccion_academica IS 'Esquema de la base de datos de academica encargado de gestionar todo lo referente a las producciones academicas en la universidad.
';
-- ddl-end --

-- SET search_path TO produccion_academica;
-- ddl-end --

-- object: produccion_academica.subtipo_produccion | type: TABLE --
-- DROP TABLE IF EXISTS produccion_academica.subtipo_produccion CASCADE;
CREATE TABLE produccion_academica.subtipo_produccion (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	tipo_produccion_id integer,
	CONSTRAINT pk_subtipo_produccion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE produccion_academica.subtipo_produccion IS 'Entidad para almacenar los sutipos de producciones academicas.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.id IS 'Identificador de la tabla.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.nombre IS 'Nombre del subtipo de producción.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.descripcion IS 'Descripción del subtipo de producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.codigo_abreviacion IS 'Código de abreviación del subtipo de producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.activo IS 'Flag que indica si el subtipo de producción academica se encuentra activo o no.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.numero_orden IS 'Número de orden en que se deben presentar los subtipos de producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.fecha_creacion IS 'Campo para el registro de creacion de un subtipo de produccion
';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.fecha_modificacion IS 'Campo para el registro de modificación de un subtipo de producción';
-- ddl-end --

-- object: produccion_academica.tipo_produccion | type: TABLE --
-- DROP TABLE IF EXISTS produccion_academica.tipo_produccion CASCADE;
CREATE TABLE produccion_academica.tipo_produccion (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_produccion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE produccion_academica.tipo_produccion IS 'Tabla que almacena los tipos de producciones academicas';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_produccion.id IS 'Identificador de la tabla
';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_produccion.nombre IS 'Nombre del tipo de producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_produccion.descripcion IS 'Descripción del tipo de producción.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_produccion.codigo_abreviacion IS 'Código de abreviación del tipo de producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_produccion.activo IS 'Flag para indicar si el tipo de producción academica se encuentra activo o no.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_produccion.numero_orden IS 'Número de orden en el que se deben presentar los tipos de producción.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_produccion.fecha_creacion IS 'Campo para el registro de creacion de un tipo de produccion';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_produccion.fecha_modificacion IS 'Campo para el registro de modificación de un tipo de producción';
-- ddl-end --

-- object: produccion_academica.metadato_produccion_academica | type: TABLE --
-- DROP TABLE IF EXISTS produccion_academica.metadato_produccion_academica CASCADE;
CREATE TABLE produccion_academica.metadato_produccion_academica (
	id serial NOT NULL,
	valor text NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	metadato_subtipo_produccion_id integer,
	produccion_academica_id integer,
	CONSTRAINT pk_dato_adicional_produccion_academica PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE produccion_academica.metadato_produccion_academica IS 'Tabla que almacena los datos adicionales (Metadata) de una produccion academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.metadato_produccion_academica.valor IS 'Valor del dato adicinal para la producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.metadato_produccion_academica.fecha_creacion IS 'Campo para el registro de creacion de un metadato produccion academica';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.metadato_produccion_academica.fecha_modificacion IS 'Campo para el registro de modificación de un metadato de producción';
-- ddl-end --

-- object: produccion_academica.produccion_academica | type: TABLE --
-- DROP TABLE IF EXISTS produccion_academica.produccion_academica CASCADE;
CREATE TABLE produccion_academica.produccion_academica (
	id serial NOT NULL,
	titulo character varying(500) NOT NULL,
	resumen text,
	fecha date NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	subtipo_produccion_id integer,
	CONSTRAINT pk_produccion_academica PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN produccion_academica.produccion_academica.id IS 'Identificador de la tabla
';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.produccion_academica.titulo IS 'Nombre de la producción academica';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.produccion_academica.resumen IS 'Resumen o abstract de la producción academica';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.produccion_academica.fecha IS 'Fecha en la que se publica la producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.produccion_academica.fecha_creacion IS 'Campo para el registro de creacion de una produccion academica';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.produccion_academica.fecha_modificacion IS 'Campo para el registro de modificación de una producción academica
';
-- ddl-end --
COMMENT ON CONSTRAINT pk_produccion_academica ON produccion_academica.produccion_academica  IS 'llave primaria de la produccion académica ';
-- ddl-end --

-- object: produccion_academica.tipo_metadato | type: TABLE --
-- DROP TABLE IF EXISTS produccion_academica.tipo_metadato CASCADE;
CREATE TABLE produccion_academica.tipo_metadato (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	form_definition json,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_dato_adicional PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE produccion_academica.tipo_metadato IS 'Entidad parametrica para almacenar los tipos de datos adicionales que puede tener una produccion academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_metadato.id IS 'Identificador de la tabla.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_metadato.nombre IS 'Nombre del tipo de dato adicional.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_metadato.descripcion IS 'Descripción del tipo de dato adicional.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_metadato.codigo_abreviacion IS 'Código de abreviación del tipo de dato adicional.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_metadato.activo IS 'Flag que indica si el tipo de dato adicional se encuentra activo o no.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_metadato.numero_orden IS 'Número de orden en que se deben presentar los tipos de dato adicionales.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_metadato.form_definition IS 'Posibles opciones o valores que puede tomar el dato adicional y valores necesarios para el form
';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_metadato.fecha_creacion IS 'Campo para el registro de creacion de un tipo de metadato
';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.tipo_metadato.fecha_modificacion IS 'Campo para el registro de modificación de un tipo de metadato';
-- ddl-end --

-- object: produccion_academica.soporte_produccion_academica | type: TABLE --
-- DROP TABLE IF EXISTS produccion_academica.soporte_produccion_academica CASCADE;
CREATE TABLE produccion_academica.soporte_produccion_academica (
	id serial NOT NULL,
	documento_id integer NOT NULL,
	descripcion character varying(250),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	produccion_academica_id integer,
	CONSTRAINT pk_soporte_produccion_academica PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE produccion_academica.soporte_produccion_academica IS 'Tabla que almacena los documentos soporte de una producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.soporte_produccion_academica.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.soporte_produccion_academica.documento_id IS 'Documento que se referencia como soporte (Se hace referencia a a tabla documento).';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.soporte_produccion_academica.descripcion IS 'Descripción del documento soporte.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.soporte_produccion_academica.fecha_creacion IS 'Campo para el registro de creacion de un soporte de produccion academica';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.soporte_produccion_academica.fecha_modificacion IS 'Campo para el registro de modificación de un soporte de producción academica';
-- ddl-end --

-- object: produccion_academica.autor_produccion_academica | type: TABLE --
-- DROP TABLE IF EXISTS produccion_academica.autor_produccion_academica CASCADE;
CREATE TABLE produccion_academica.autor_produccion_academica (
	id serial NOT NULL,
	persona_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	produccion_academica_id integer,
	estado_autor_produccion_id integer,
	CONSTRAINT pk_autor_produccion_academica PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE produccion_academica.autor_produccion_academica IS 'Tabla que almacena los autores de una producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.autor_produccion_academica.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.autor_produccion_academica.persona_id IS 'Persona a la que se referencia como autor (Se hace referencia a a tabla persona).';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.autor_produccion_academica.fecha_creacion IS 'Campo para el registro de creacion de un autor de producion academica';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.autor_produccion_academica.fecha_modificacion IS 'Campo para el registro de modificación de un autor ';
-- ddl-end --

-- object: produccion_academica.metadato_subtipo_produccion | type: TABLE --
-- DROP TABLE IF EXISTS produccion_academica.metadato_subtipo_produccion CASCADE;
CREATE TABLE produccion_academica.metadato_subtipo_produccion (
	id serial NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	tipo_metadato_id integer,
	subtipo_produccion_id integer,
	CONSTRAINT pk_dato_adicional_subtipo_produccion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE produccion_academica.metadato_subtipo_produccion IS 'Tabla de rompimiento entre la relación muchos a muchos de los tipos de datos adicionales con los subtipos de producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.metadato_subtipo_produccion.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.metadato_subtipo_produccion.fecha_creacion IS 'Campo para el registro de creacion de un metadato subtipo ';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.metadato_subtipo_produccion.fecha_modificacion IS 'Campo para el registro de modificación de un metadato subtipo';
-- ddl-end --

-- object: produccion_academica.estado_autor_produccion | type: TABLE --
-- DROP TABLE IF EXISTS produccion_academica.estado_autor_produccion CASCADE;
CREATE TABLE produccion_academica.estado_autor_produccion (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_estado_autor_produccion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE produccion_academica.estado_autor_produccion IS 'Tabla que almacena los estados de produccion academica';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.estado_autor_produccion.id IS 'Identificador de la tabla
';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.estado_autor_produccion.nombre IS 'Nombre del estado del autor de producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.estado_autor_produccion.descripcion IS 'Descripción del estado del autor de la producción.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.estado_autor_produccion.codigo_abreviacion IS 'Código de abreviación del autor de la producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.estado_autor_produccion.activo IS 'Flag para indicar si el estado se encuentra activo o no.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.estado_autor_produccion.numero_orden IS 'Número de orden en el que se deben presentar los estados.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.estado_autor_produccion.fecha_creacion IS 'Campo para el registro de creacion de un estado de autor';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.estado_autor_produccion.fecha_modificacion IS 'Campo para el registro de modificación de un estado de auto de producción';
-- ddl-end --

-- object: fk_subtipo_produccion_tipo_produccion | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.subtipo_produccion DROP CONSTRAINT IF EXISTS fk_subtipo_produccion_tipo_produccion CASCADE;
ALTER TABLE produccion_academica.subtipo_produccion ADD CONSTRAINT fk_subtipo_produccion_tipo_produccion FOREIGN KEY (tipo_produccion_id)
REFERENCES produccion_academica.tipo_produccion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_metadato_subtipo_produccion_tipo_metadato | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.metadato_subtipo_produccion DROP CONSTRAINT IF EXISTS fk_metadato_subtipo_produccion_tipo_metadato CASCADE;
ALTER TABLE produccion_academica.metadato_subtipo_produccion ADD CONSTRAINT fk_metadato_subtipo_produccion_tipo_metadato FOREIGN KEY (tipo_metadato_id)
REFERENCES produccion_academica.tipo_metadato (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_metadato_subtipo_produccion_subtipo_produccion | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.metadato_subtipo_produccion DROP CONSTRAINT IF EXISTS fk_metadato_subtipo_produccion_subtipo_produccion CASCADE;
ALTER TABLE produccion_academica.metadato_subtipo_produccion ADD CONSTRAINT fk_metadato_subtipo_produccion_subtipo_produccion FOREIGN KEY (subtipo_produccion_id)
REFERENCES produccion_academica.subtipo_produccion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_produccion_academica_subtipo_produccion | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.produccion_academica DROP CONSTRAINT IF EXISTS fk_produccion_academica_subtipo_produccion CASCADE;
ALTER TABLE produccion_academica.produccion_academica ADD CONSTRAINT fk_produccion_academica_subtipo_produccion FOREIGN KEY (subtipo_produccion_id)
REFERENCES produccion_academica.subtipo_produccion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_metadato_produccion_academica_metadato_subtipo_pro_4208 | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.metadato_produccion_academica DROP CONSTRAINT IF EXISTS fk_metadato_produccion_academica_metadato_subtipo CASCADE;
ALTER TABLE produccion_academica.metadato_produccion_academica ADD CONSTRAINT fk_metadato_produccion_academica_metadato_subtipo FOREIGN KEY (metadato_subtipo_produccion_id)
REFERENCES produccion_academica.metadato_subtipo_produccion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_metadato_produccion_academica_produccion_academica | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.metadato_produccion_academica DROP CONSTRAINT IF EXISTS fk_metadato_produccion_academica_produccion_academica CASCADE;
ALTER TABLE produccion_academica.metadato_produccion_academica ADD CONSTRAINT fk_metadato_produccion_academica_produccion_academica FOREIGN KEY (produccion_academica_id)
REFERENCES produccion_academica.produccion_academica (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_soporte_produccion_academica_produccion_academica | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.soporte_produccion_academica DROP CONSTRAINT IF EXISTS fk_soporte_produccion_academica_produccion_academica CASCADE;
ALTER TABLE produccion_academica.soporte_produccion_academica ADD CONSTRAINT fk_soporte_produccion_academica_produccion_academica FOREIGN KEY (produccion_academica_id)
REFERENCES produccion_academica.produccion_academica (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_autor_produccion_academica_produccion_academica | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.autor_produccion_academica DROP CONSTRAINT IF EXISTS fk_autor_produccion_academica_produccion_academica CASCADE;
ALTER TABLE produccion_academica.autor_produccion_academica ADD CONSTRAINT fk_autor_produccion_academica_produccion_academica FOREIGN KEY (produccion_academica_id)
REFERENCES produccion_academica.produccion_academica (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_autor_produccion_academica_estado_autor_produccion | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.autor_produccion_academica DROP CONSTRAINT IF EXISTS fk_autor_produccion_academica_estado_autor_produccion CASCADE;
ALTER TABLE produccion_academica.autor_produccion_academica ADD CONSTRAINT fk_autor_produccion_academica_estado_autor_produccion FOREIGN KEY (estado_autor_produccion_id)
REFERENCES produccion_academica.estado_autor_produccion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- Permisos de usuario
GRANT USAGE ON SCHEMA produccion_academica TO test;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA produccion_academica TO test;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA produccion_academica TO test;
