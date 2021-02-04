ALTER TABLE produccion_academica.subtipo_produccion ALTER COLUMN nombre TYPE character varying(70);

INSERT INTO produccion_academica.tipo_metadato(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, form_definition, fecha_creacion, fecha_modificacion) VALUES (86,'Tipo Produccion Audiovisual','Tipo de producción Audiovisual','TPRAV', true,86, '{"etiqueta": "select", "claseGrid":"col-6", "nombre": "86", "label_i18n": "tipo_produccion_audiovisual", "placeholder_i18n": "tipo_produccion_audiovisual", "requerido": true, "tipo": "number", "key": "Nombre", "opciones": [{"Id":"1","Nombre":"Video"},{"Id":"2","Nombre":"Cinematográfica"},{"Id":"3","Nombre":"Fonografía"}],"span":""}', now(), now());

INSERT INTO produccion_academica.metadato_subtipo_produccion(id, activo, tipo_metadato_id, subtipo_produccion_id, fecha_creacion, fecha_modificacion) VALUES (273,true,86,24, now(), now());
INSERT INTO produccion_academica.metadato_subtipo_produccion(id, activo, tipo_metadato_id, subtipo_produccion_id, fecha_creacion, fecha_modificacion) VALUES (274,true,86,25, now(), now());

update produccion_academica.subtipo_produccion set nombre = 'Video Cinematografica o Fonográfia Internacional o Nacional' where id = 24;
update produccion_academica.subtipo_produccion set nombre = 'Video Cinematografica o Fonográfia Regional o Local' where id = 25;