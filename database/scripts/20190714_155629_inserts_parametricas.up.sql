-- Inserts estado_autor_produccion
insert into produccion_academica.estado_autor_produccion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion)
	values (1, 'Autor principal', 'Autor que registra la producción', 'AutPrin', true, 1, now(), now());
insert into produccion_academica.estado_autor_produccion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion)
	values (2, 'Autor confirmado', 'Autor que confirma que pertenece a la producción', 'AutConf', true, 2, now(), now());
insert into produccion_academica.estado_autor_produccion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion)
	values (3, 'Autor sin confirmar', 'Autor que aún no confirma su autoria', 'AutSinConf', true, 3, now(), now());
insert into produccion_academica.estado_autor_produccion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion)
	values (4, 'Niega participación', 'El ente niega la participación en la producción', 'AutNie', true, 4, now(), now());
