Feature: Validate API responses
  PRODUCCION_ACADEMICA_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /tipo_metadato
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route            |bodyreq               |codres       |
  |GET   |/v1/tipo_metadato|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/tipo_metadat |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/tipo_metadat |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/tipo_metadat |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/tipo_metadat |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /tipo_metadato       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples:
  |method|route            |bodyreq               |codres         |bodyres                |
  |GET   |/v1/tipo_metadato|./files/req/Vacio.json|200 OK         |./files/res8/Vok1.json |
  |POST  |/v1/tipo_metadato|./files/req/Vacio.json|400 Bad Request|./files/res8/Ierr1.json|
