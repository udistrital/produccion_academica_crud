--DROP TABLE IF EXISTS produccion_academica.sistema_tipo_produccion CASCADE;
CREATE SEQUENCE produccion_academica.sistema_tipo_produccion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE produccion_academica.sistema_tipo_produccion (
	id integer NOT NULL DEFAULT nextval('produccion_academica.sistema_tipo_produccion_id_seq'::regclass),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	sistema_id integer,
	tipo_produccion_id integer,
	CONSTRAINT pk_sistema_tipo_produccion PRIMARY KEY (id)

); 
-- object: produccion_academica.subtipo_produccion | type: TABLE --
ALTER TABLE produccion_academica.subtipo_produccion ADD categoria_produccion_id integer;

-- ddl-end --
COMMENT ON TABLE produccion_academica.subtipo_produccion IS E'Entidad para almacenar los sutipos de producciones academicas.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.id IS E'Identificador de la tabla.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.nombre IS E'Nombre del subtipo de producción.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.descripcion IS E'Descripción del subtipo de producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.codigo_abreviacion IS E'Código de abreviación del subtipo de producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.activo IS E'Flag que indica si el subtipo de producción academica se encuentra activo o no.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.numero_orden IS E'Número de orden en que se deben presentar los subtipos de producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.fecha_creacion IS E'Campo para el registro de creacion de un subtipo de produccion';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.subtipo_produccion.fecha_modificacion IS E'Campo para el registro de modificación de un subtipo de producción';
-- ddl-end --
CREATE SEQUENCE produccion_academica.puntaje_subtipo_produccion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- object: produccion_academica.puntaje_subtipo_produccion | type: TABLE --
--DROP TABLE IF EXISTS produccion_academica.puntaje_subtipo_produccion CASCADE;
CREATE TABLE produccion_academica.puntaje_subtipo_produccion (
	id integer NOT NULL DEFAULT nextval('produccion_academica.puntaje_subtipo_produccion_id_seq'::regclass),
	nombre character varying(250) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	caracteristicas json,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	subtipo_produccion integer,
	CONSTRAINT pk_puntaje_subtipo_produccion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE produccion_academica.puntaje_subtipo_produccion IS E'Tabla que almacena los documentos soporte de una producción academica.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.puntaje_subtipo_produccion.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.puntaje_subtipo_produccion.codigo_abreviacion IS E'Descripción del documento soporte.';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.puntaje_subtipo_produccion.fecha_creacion IS E'Campo para el registro de creacion de un soporte de produccion academica';
-- ddl-end --
COMMENT ON COLUMN produccion_academica.puntaje_subtipo_produccion.fecha_modificacion IS E'Campo para el registro de modificación de un soporte de producción academica';
-- ddl-end --
-- ALTER TABLE produccion_academica.puntaje_subtipo_produccion OWNER TO postgres;
-- ddl-end --
 CREATE SEQUENCE produccion_academica.sistema_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- object: produccion_academica.sistema | type: TABLE --
--DROP TABLE IF EXISTS produccion_academica.sistema CASCADE;
CREATE TABLE produccion_academica.sistema (
	id integer NOT NULL DEFAULT nextval('produccion_academica.sistema_id_seq'::regclass),
	nombre character varying(50) NOT NULL,
	codigo_abreviacion character varying(250),
	referencia json NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_sistema_produccion_academica PRIMARY KEY (id)

);
-- ddl-end --
-- ALTER TABLE produccion_academica.sistema OWNER TO postgres;
-- ddl-end --

-- -- object: produccion_academica.sistema_produccion_academica_id_seq | type: SEQUENCE --
-- -- DROP SEQUENCE IF EXISTS produccion_academica.sistema_produccion_academica_id_seq CASCADE;
CREATE SEQUENCE produccion_academica.categoria_produccion_id_seq
	INCREMENT BY 1
	MINVALUE 1 
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- -- ddl-end --
-- -- ALTER SEQUENCE produccion_academica.sistema_produccion_academica_id_seq OWNER TO postgres;
-- -- ddl-end --
-- 
-- object: produccion_academica.categoria_produccion | type: TABLE --
DROP TABLE IF EXISTS produccion_academica.soporte_produccion_academica CASCADE;
CREATE TABLE produccion_academica.categoria_produccion (
	id integer NOT NULL  DEFAULT nextval('produccion_academica.categoria_produccion_id_seq'::regclass),
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	activo boolean NOT NULL,
	tipo_puntaje character varying(20),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_categoria_produccion PRIMARY KEY (id)

);
-- ddl-end --
-- ALTER TABLE produccion_academica.categoria_produccion OWNER TO postgres;
-- ddl-end --

-- object: fk_sistema_sistema_tipo_produccion | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.sistema_tipo_produccion DROP CONSTRAINT IF EXISTS fk_sistema_sistema_tipo_produccion CASCADE;
ALTER TABLE produccion_academica.sistema_tipo_produccion ADD CONSTRAINT fk_sistema_sistema_tipo_produccion FOREIGN KEY (sistema_id)
REFERENCES produccion_academica.sistema (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_tipo_produccion_sistema_tipo_produccion | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.sistema_tipo_produccion DROP CONSTRAINT IF EXISTS fk_tipo_produccion_sistema_tipo_produccion CASCADE;
ALTER TABLE produccion_academica.sistema_tipo_produccion ADD CONSTRAINT fk_tipo_produccion_sistema_tipo_produccion FOREIGN KEY (tipo_produccion_id)
REFERENCES produccion_academica.tipo_produccion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_subtipo_produccion_categoria_produccion | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.subtipo_produccion DROP CONSTRAINT IF EXISTS fk_subtipo_produccion_categoria_produccion CASCADE;
ALTER TABLE produccion_academica.subtipo_produccion ADD CONSTRAINT fk_subtipo_produccion_categoria_produccion FOREIGN KEY (categoria_produccion_id)
REFERENCES produccion_academica.categoria_produccion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --


-- object: fk_puntaje_subtipo_produccion_subtipo_produccion | type: CONSTRAINT --
-- ALTER TABLE produccion_academica.puntaje_subtipo_produccion DROP CONSTRAINT IF EXISTS fk_puntaje_subtipo_produccion_subtipo_produccion CASCADE;
ALTER TABLE produccion_academica.puntaje_subtipo_produccion ADD CONSTRAINT fk_puntaje_subtipo_produccion_subtipo_produccion FOREIGN KEY (subtipo_produccion)
REFERENCES produccion_academica.subtipo_produccion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --