-- Inserts estado_autor_produccion
insert into produccion_academica.estado_autor_produccion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (1, 'Autor principal', 'Autor que registra la producción', 'AutPrin', true, 1, now(), now()); 
insert into produccion_academica.estado_autor_produccion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (2, 'Autor confirmado', 'Autor que confirma que pertenece a la producción', 'AutConf', true, 2, now(), now()); 
insert into produccion_academica.estado_autor_produccion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (3, 'Autor sin confirmar', 'Autor que aún no confirma su autoria', 'AutSinConf', true, 3, now(), now()); 
insert into produccion_academica.estado_autor_produccion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (4, 'Niega participación', 'El ente niega la participación en la producción', 'AutNie', true, 4, now(), now()); 

-- inserts tipos produccion
insert into produccion_academica.tipo_produccion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (1, 'Articulo', 'Articulo', 'ART', true, 1, now(), now()); 
insert into produccion_academica.tipo_produccion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (2, 'Traducción', 'Traducción', 'TRANS', true, 2, now(), now()); 

insert into produccion_academica.subtipo_produccion (id, nombre, tipo_produccion_id, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (1, 'Articulo aplicado', 1, 'Articulo aplicado', 'ART_AP', true, 1, now(), now()); 
insert into produccion_academica.subtipo_produccion (id, nombre, tipo_produccion_id, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (2, 'Articulo de revisión', 1, 'Articulo de revisión', 'ART_REV', true, 2, now(), now()); 