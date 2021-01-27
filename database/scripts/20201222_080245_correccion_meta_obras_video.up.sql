UPDATE produccion_academica.metadato_subtipo_produccion SET tipo_metadato_id=34 WHERE id=181;
UPDATE produccion_academica.metadato_subtipo_produccion SET tipo_metadato_id=34 WHERE id=196;
UPDATE produccion_academica.metadato_subtipo_produccion SET tipo_metadato_id=34 WHERE id=210;
UPDATE produccion_academica.metadato_subtipo_produccion SET tipo_metadato_id=34 WHERE id=225;

update produccion_academica.tipo_metadato set form_definition = '{"etiqueta":"input","claseGrid":"col-6","nombre":"1","label_i18n":"paginas","placeholder_i18n":"paginas","requerido":true,"tipo":"number","span":""}' where id = 1;
update produccion_academica.tipo_metadato set form_definition = '{"etiqueta":"input","claseGrid":"col-6","nombre":"21","label_i18n":"numero_autores","placeholder_i18n":"numero_autores","requerido":true,"tipo":"number","span":"El total de autores de todo el trabajo"}' where id = 21;
update produccion_academica.tipo_metadato set form_definition = '{"etiqueta":"input","claseGrid":"col-6","nombre":"24","label_i18n":"numero_capitulo","placeholder_i18n":"numero_capitulo","requerido":true,"tipo":"number","span":""}' where id = 24;
update produccion_academica.tipo_metadato set form_definition = '{"etiqueta":"input","claseGrid":"col-6","nombre":"25","label_i18n":"paginas_capitulo","placeholder_i18n":"paginas_capitulo","requerido":true,"tipo":"number","span":""}' where id = 25;
update produccion_academica.tipo_metadato set form_definition = '{"etiqueta":"input","claseGrid":"col-6","nombre":"35","label_i18n":"duracion_estudios","placeholder_i18n":"duracion_estudios","requerido":true,"tipo":"number","span":"El periodo mínimo aceptado es de 9 meses."}' where id = 35;