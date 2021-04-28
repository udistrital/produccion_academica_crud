--Actualización de datos de categorías de producción académica
update produccion_academica.categoria_produccion set nombre = 'Cambio de categoría', descripcion = 'Cambio de categoría por parte del docente', tipo_puntaje = 'Salarial' where id = 1;
update produccion_academica.categoria_produccion set nombre = 'Título de postgrado', descripcion = 'Presentación de título de postgrado obtenido', tipo_puntaje = 'Salarial' where id = 2;
update produccion_academica.categoria_produccion set nombre = 'Producción puntos salariales', descripcion = 'Producción académica para puntos salariales', tipo_puntaje = 'Salarial' where id = 3;
update produccion_academica.categoria_produccion set nombre = 'Producción puntos bonificación', descripcion = 'Producción académica para puntos de bonificación', tipo_puntaje = 'Bonificación' where id = 4;

--Actualización de los tipos de productos
update produccion_academica.tipo_produccion set nombre = 'Cambio de categoría', descripcion = 'Cambio de categoría' where codigo_abreviacion = 'CC';
update produccion_academica.tipo_produccion set nombre = 'Título de postgrado', descripcion = 'Obtención de título de postgrado' where codigo_abreviacion = 'TP';
update produccion_academica.tipo_produccion set nombre = 'Artículo', descripcion = 'Publicación de artículo' where codigo_abreviacion = 'AR';
update produccion_academica.tipo_produccion set nombre = 'Artículo corto', descripcion = 'Artículo corto' where codigo_abreviacion = 'ARTC';
update produccion_academica.tipo_produccion set nombre = 'Capítulo libro', descripcion = 'Publicación de capítulo de libro' where codigo_abreviacion = 'CLB';
update produccion_academica.tipo_produccion set nombre = 'Producción técnica' where codigo_abreviacion = 'PT';
update produccion_academica.tipo_produccion set nombre = 'Video, Cinematográfica o Fonografía', descripcion = 'Producción de videos, cinematográficas o fonografías' where codigo_abreviacion = 'PVCF';
update produccion_academica.tipo_produccion set nombre = 'Obras artísticas', descripcion = 'Realización obras artísticas' where codigo_abreviacion = 'OA';
update produccion_academica.tipo_produccion set nombre = 'Reseña Crítica', descripcion = 'Realización de reseñas críticas' where codigo_abreviacion = 'RAC';
update produccion_academica.tipo_produccion set nombre = 'Estudios postdoctorales' where codigo_abreviacion = 'PD';
update produccion_academica.tipo_produccion set nombre = 'Traducción de artículos', descripcion = 'Realización de traducción de artículos' where codigo_abreviacion = 'TA';
update produccion_academica.tipo_produccion set nombre = 'Dirección de tesis', descripcion = 'Dirección de tesis para maestría o doctorado' where codigo_abreviacion = 'DT';

--Actualización de los subtipos de productos
update produccion_academica.subtipo_produccion set descripcion = 'Cambio de categoría de auxiliar a asistente' where codigo_abreviacion = 'CXA';
update produccion_academica.subtipo_produccion set descripcion = 'Cambio de categoría de asistente a asociado' where codigo_abreviacion = 'CAS';
update produccion_academica.subtipo_produccion set descripcion = 'Cambio de categoría de asociado a titular' where codigo_abreviacion = 'CST';
update produccion_academica.subtipo_produccion set nombre = 'Título de especialización', descripcion = 'Obtención de título de especialización' where codigo_abreviacion = 'TPE';
update produccion_academica.subtipo_produccion set nombre = 'Título de maestría', descripcion = 'Obtención de título de maestría' where codigo_abreviacion = 'TPM';
update produccion_academica.subtipo_produccion set nombre = 'Título de doctorado', descripcion = 'Obtención de título de doctorado' where codigo_abreviacion = 'TPD';
update produccion_academica.subtipo_produccion set nombre = 'Artículo completo', descripcion = 'Artículo completo en una revista' where codigo_abreviacion = 'ACO';
update produccion_academica.subtipo_produccion set nombre = 'Página editorial', descripcion = 'Producción de página editorial en revista' where codigo_abreviacion = 'PE';
update produccion_academica.subtipo_produccion set nombre = 'Artículo corto', descripcion = 'Artículo corto publicado en una revista' where codigo_abreviacion = 'ACR';
update produccion_academica.subtipo_produccion set nombre = 'Libro de investigación', descripcion = 'Producción de libro de investigación' where codigo_abreviacion = 'LI';
update produccion_academica.subtipo_produccion set nombre = 'Libro de texto', descripcion = 'Producción de libro de texto' where codigo_abreviacion = 'LT';
update produccion_academica.subtipo_produccion set nombre = 'Libro de ensayo', descripcion = 'Producción de libro de ensayo' where codigo_abreviacion = 'LE';
update produccion_academica.subtipo_produccion set nombre = 'Capítulo libro de investigación', descripcion = 'Producción de capítulo de libro de investigación' where codigo_abreviacion = 'CLI';
update produccion_academica.subtipo_produccion set nombre = 'Capítulo libro de texto', descripcion = 'Producción de capítulo de libro de texto' where codigo_abreviacion = 'CLT';
update produccion_academica.subtipo_produccion set nombre = 'Capítulo libro de ensayo', descripcion = 'Producción de capítulo de libro de ensayo' where codigo_abreviacion = 'CLE';
update produccion_academica.subtipo_produccion set nombre = 'Traducción de libro' where codigo_abreviacion = 'TL';
update produccion_academica.subtipo_produccion set nombre = 'Innovación tecnológica', descripcion = 'Producción de innovación tecnológica' where codigo_abreviacion = 'PIT';
update produccion_academica.subtipo_produccion set nombre = 'Adaptación tecnológica', descripcion = 'Producción de adaptación tecnológica' where codigo_abreviacion = 'PAT';
update produccion_academica.subtipo_produccion set nombre = 'Patente de invención', descripcion = 'Publicación de patente de invención' where codigo_abreviacion = 'PAI';
update produccion_academica.subtipo_produccion set descripcion = 'Publicación de patente de modelo de utilidad' where codigo_abreviacion = 'PAMU';
update produccion_academica.subtipo_produccion set nombre = 'Obras artísticas Internacional o Nacional', descripcion = 'Realización de obras artísticas de impacto internacional o nacional' where codigo_abreviacion = 'OAIN';
update produccion_academica.subtipo_produccion set nombre = 'Obras artísticas Regional o Local', descripcion = 'Realización de obras artísticas de impacto regional o local' where codigo_abreviacion = 'OARL';
update produccion_academica.subtipo_produccion set nombre = 'Publicación impresa de revista' where codigo_abreviacion = 'PIR';
update produccion_academica.subtipo_produccion set nombre = 'Publicación impresa de libro' where codigo_abreviacion = 'PIL';
update produccion_academica.subtipo_produccion set nombre = 'Reseña crítica', descripcion = 'Realización de reseñas críticas' where codigo_abreviacion = 'RAC';
update produccion_academica.subtipo_produccion set nombre = 'Traducción de artículos', descripcion = 'Realización de traduccion de artículos' where codigo_abreviacion = 'TA';
update produccion_academica.subtipo_produccion set nombre = 'Direccion de tesis para maestría', descripcion = 'Direccion de tesis para maestría' where codigo_abreviacion = 'DTM';
update produccion_academica.subtipo_produccion set nombre = 'Direccion de tesis para doctorado', descripcion = 'Direccion de tesis para doctorado' where codigo_abreviacion = 'DTD';
update produccion_academica.subtipo_produccion set nombre = 'Video cinematográfica o fonografía Internacional o Nacional', descripcion = 'Producción de videos, cinematográficas o fonografías internacional o nacional' where codigo_abreviacion = 'VCFIN';
update produccion_academica.subtipo_produccion set nombre = 'Video cinematográfica o fonografía Regional o Local', descripcion = 'Producción de videos, cinematográficas o fonografías Regional o local' where codigo_abreviacion = 'VCFRL';

--Actualización de los metadatos
update produccion_academica.tipo_metadato set nombre = 'Páginas' where codigo_abreviacion  = 'LC';
update produccion_academica.tipo_metadato set nombre = 'Fascículo' where codigo_abreviacion  = 'FASC';
update produccion_academica.tipo_metadato set nombre = 'Número resolución convalidación', descripcion = 'Número de resolución de convalidación' where codigo_abreviacion  = 'NRCON';
update produccion_academica.tipo_metadato set nombre = 'Número autores', descripcion = 'Número de autores involucrados' where codigo_abreviacion  = 'NAUT';
update produccion_academica.tipo_metadato set nombre = 'Nombre capítulo', descripcion = 'Nombre del capítulo del libro producido' where codigo_abreviacion  = 'NOMC';
update produccion_academica.tipo_metadato set nombre = 'Número capítulo', descripcion = 'Número capítulo de libro producido' where codigo_abreviacion  = 'NUMC';
update produccion_academica.tipo_metadato set nombre = 'Páginas capítulo', descripcion = 'Páginas de capítulo del libro producido' where codigo_abreviacion  = 'PAGC';
update produccion_academica.tipo_metadato set nombre = 'Carácter producción', descripcion = 'Carácter de producción de video' where codigo_abreviacion  = 'CARP';
update produccion_academica.tipo_metadato set nombre = 'Grado participación', descripcion = 'Grado de participación de docente en obra artística o producción de video' where codigo_abreviacion  = 'GPD';
update produccion_academica.tipo_metadato set nombre = 'Teléfono productora', descripcion = 'Teléfono de productora de producción de video' where codigo_abreviacion  = 'TPRO';
update produccion_academica.tipo_metadato set nombre = 'Género musical', descripcion = 'Género musical en caso de producción de fonografía' where codigo_abreviacion  = 'GMF';
update produccion_academica.tipo_metadato set nombre = 'Número revista', descripcion = 'Número revista' where codigo_abreviacion  = 'NUMREV';
update produccion_academica.tipo_metadato set nombre = 'Tipo de difusión', descripcion = 'Tipo de difusión' where codigo_abreviacion  = 'TDSN';
update produccion_academica.tipo_metadato set nombre = 'Duración estudios', descripcion = 'Duración de estudios postdoctorales' where codigo_abreviacion  = 'DEPD';
update produccion_academica.tipo_metadato set nombre = 'Nombre estudiante', descripcion = 'Nombre de estudiante' where codigo_abreviacion  = 'NEST';
update produccion_academica.tipo_metadato set nombre = 'Número registro oficial', descripcion = 'Número de registro oficial aprobado por entidad' where codigo_abreviacion  = 'NROF';
update produccion_academica.tipo_metadato set nombre = 'Categoría revista', descripcion = 'Categoría revista de publicación' where codigo_abreviacion  = 'CRP';
update produccion_academica.tipo_metadato set nombre = 'Gran área conocimiento', descripcion = 'Gran área de conocimiento' where codigo_abreviacion  = 'GACT';
update produccion_academica.tipo_metadato set nombre = 'Categoría premio', descripcion = 'Categoría premio' where codigo_abreviacion  = 'CGP';
update produccion_academica.tipo_metadato set nombre = 'Tipo de obra audiovisual', descripcion = 'Tipo de obra de producción de video' where codigo_abreviacion  = 'TOV';
update produccion_academica.tipo_metadato set nombre = 'Tipo de obra artística', descripcion = 'Tipo de obra artística' where codigo_abreviacion  = 'TOA';
update produccion_academica.tipo_metadato set nombre = 'Categoría tesis', descripcion = 'Categoría de reconocimiento de tesis' where codigo_abreviacion  = 'CRT';
update produccion_academica.tipo_metadato set nombre = 'Área conocimiento específica', descripcion = 'Área conocimiento específica' where codigo_abreviacion  = 'ACES';
update produccion_academica.tipo_metadato set nombre = 'Fecha convalidación', descripcion = 'Fecha de convalidación de título' where codigo_abreviacion  = 'FCONV';
update produccion_academica.tipo_metadato set nombre = 'Categoría actual', descripcion = 'Categoría actual del docente' where codigo_abreviacion  = 'CAAC';
update produccion_academica.tipo_metadato set nombre = 'Fecha cambio categoría', descripcion = 'Fecha del último cambio de categoría del docente' where codigo_abreviacion  = 'FUCC';
update produccion_academica.tipo_metadato set nombre = 'Descripción obra', descripcion = 'Descripción obra' where codigo_abreviacion  = 'DOBR';
update produccion_academica.tipo_metadato set nombre = 'Certificado evaluación docente', descripcion = 'Certificación de evaluación docente del ultimo año' where codigo_abreviacion  = 'CEVD';
update produccion_academica.tipo_metadato set nombre = 'Trabajo inédito', descripcion = 'Trabajo inédito' where codigo_abreviacion  = 'TRIN';
update produccion_academica.tipo_metadato set nombre = 'Artículo', descripcion = 'Soporte de artículo' where codigo_abreviacion  = 'SART';
update produccion_academica.tipo_metadato set nombre = 'Certificado resultado investigación', descripcion = 'Certificado de resultado de investigación' where codigo_abreviacion  = 'CRIN';
update produccion_academica.tipo_metadato set nombre = 'Autorización autor original', descripcion = 'Autorización de autor original para realizar traducción' where codigo_abreviacion  = 'AUAUO';
update produccion_academica.tipo_metadato set nombre = 'Proceso selección' where codigo_abreviacion  = 'SPSP';
update produccion_academica.tipo_metadato set nombre = 'Soporte difusión', descripcion = 'Soporte de difusión' where codigo_abreviacion  = 'SDSN';
update produccion_academica.tipo_metadato set nombre = 'Soporte de evento' where codigo_abreviacion  = 'SEV';
update produccion_academica.tipo_metadato set nombre = 'Registro otorgamiento patente' where codigo_abreviacion  = 'ROP';
update produccion_academica.tipo_metadato set nombre = 'Manual técnico', descripcion = 'Manual técnico' where codigo_abreviacion  = 'MTEC';
update produccion_academica.tipo_metadato set nombre = 'Instrucciones según lenguaje', descripcion = 'Instrucciones según lenguaje' where codigo_abreviacion  = 'ISLG';
update produccion_academica.tipo_metadato set nombre = 'Certificado ponencia', descripcion = 'Certificado de presentación de ponencia' where codigo_abreviacion  = 'SPON';
update produccion_academica.tipo_metadato set nombre = 'Memorias evento o artículo ponencia', descripcion = 'Memorias del evento o artículo presentado en la ponencia' where codigo_abreviacion  = 'MEAP';
update produccion_academica.tipo_metadato set nombre = 'Publicación impresa', descripcion = 'Soporte de publicación impresa' where codigo_abreviacion  = 'SPIM';
update produccion_academica.tipo_metadato set nombre = 'Reseña crítica', descripcion = 'Soporte de reseña crítica' where codigo_abreviacion  = 'SRCR';
update produccion_academica.tipo_metadato set nombre = 'Certificado secretaría académica', descripcion = 'Certificado de secretaría académica de dirección de tesis' where codigo_abreviacion  = 'CSADT';
update produccion_academica.tipo_metadato set nombre = 'Acta grado pregrado' where codigo_abreviacion  = 'APRG';
update produccion_academica.tipo_metadato set nombre = 'Acta grado especialización' where codigo_abreviacion  = 'AGESP';
update produccion_academica.tipo_metadato set nombre = 'Acta grado maestría' where codigo_abreviacion  = 'AGMAS';
update produccion_academica.tipo_metadato set nombre = 'Acta grado doctorado' where codigo_abreviacion  = 'AGDOC';
update produccion_academica.tipo_metadato set nombre = 'Ejecutable programa', descripcion = 'Soporte de archivo ejecutable de aplicacion' where codigo_abreviacion  = 'EPRO';
update produccion_academica.tipo_metadato set nombre = 'Código fuente software', descripcion = 'Soporte de código fuente del software' where codigo_abreviacion  = 'CFS';
update produccion_academica.tipo_metadato set nombre = 'Tipo producción audiovisual' where codigo_abreviacion  = 'TPRAV';
update produccion_academica.tipo_metadato set nombre = 'Tipo universidad', descripcion = 'Tipo de universidad' where codigo_abreviacion  = 'TUNIV';
update produccion_academica.tipo_metadato set nombre = 'Título de convalidado', descripcion = 'Título con que fue convalidado' where codigo_abreviacion  = 'NTCONV';


update produccion_academica.tipo_metadato
set form_definition = '{"etiqueta": "select", "claseGrid":"col-6", "nombre": "39", "label_i18n": "gran_area_conocimiento", "placeholder_i18n": "gran_area_conocimiento", "requerido": true, "tipo": "number", "key": "Nombre", "opciones": [{"Id":"1","Nombre":"Ciencias agrícolas"},{"Id":"2","Nombre":"Ciencias sociales"},{"Id":"3","Nombre":"Humanidades"},{"Id":"4","Nombre":"Ciencias naturales"},{"Id":"5","Nombre":"Ingeniería y tecnología"},{"Id":"6","Nombre":"Ciencias médicas y de la salud"}],"span":""}'
where codigo_abreviacion = 'GACT';

update produccion_academica.tipo_metadato
set form_definition = '{"etiqueta": "select", "claseGrid":"col-6", "nombre": "42", "label_i18n": "tipo_obra_artistica", "placeholder_i18n": "tipo_obra_artistica", "requerido": true, "tipo": "number", "key": "Nombre", "opciones": [{"Id":"1","Nombre":"Creación original"},{"Id":"2","Nombre":"Creación complementaria o de apoyo"},{"Id":"3","Nombre":"Interpretación de obra"}],"span":""}'
where codigo_abreviacion = 'TOA';

--Solicitado por Mélida
update produccion_academica.tipo_metadato
set form_definition = '{"etiqueta":"file","claseGrid":"col-lg-6 col-md-6 col-sm-6 col-xs-6","clase":"form-control","nombre":"70","label_i18n":"certificado_ponencia","placeholder_i18n":"certificado_ponencia","requerido":true,"tipo":"pdf","tipoDocumento":15,"formatos":"pdf","url":"","tamanoMaximo":2,"span":"El archivo debe especificar si es ponente, expositor o conferencista"}'
where codigo_abreviacion = 'SPON';