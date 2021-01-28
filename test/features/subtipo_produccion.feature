Feature: Validate API responses
  PRODUCCION_ACADEMICA_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /subtipo_produccion
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route                 |bodyreq               |codres       |
  |GET   |/v1/subtipo_produccion|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/subtipo_produccio |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/subtipo_produccio |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/subtipo_produccio |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/subtipo_produccio |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /subtipo_produccion       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples: 
  |method|route                 |bodyreq               |codres         |bodyres                |
  |GET   |/v1/subtipo_produccion|./files/req/Vacio.json|200 OK         |./files/res7/Vok1.json |
  |POST  |/v1/subtipo_produccion|./files/req/Vacio.json|400 Bad Request|./files/res7/Ierr1.json|
 