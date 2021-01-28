Feature: Validate API responses
  PRODUCCION_ACADEMICA_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /metadato_subtipo_produccion
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route                          |bodyreq               |codres       |
  |GET   |/v1/metadato_subtipo_produccion|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/metadato_subtipo_produccio |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/metadato_subtipo_produccio |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/metadato_subtipo_produccio |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/metadato_subtipo_produccio |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /metadato_subtipo_produccion       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples: 
  |method|route                          |bodyreq               |codres         |bodyres                |
  |GET   |/v1/metadato_subtipo_produccion|./files/req/Vacio.json|200 OK         |./files/res4/Vok1.json |
  |POST  |/v1/metadato_subtipo_produccion|./files/req/Vacio.json|400 Bad Request|./files/res4/Ierr1.json|
 