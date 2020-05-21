## A Simple Reference Data API written in GO
REST API to serve reference data (words, tags, names, cities etc etc).

``go get github.com/saiprasadkrishnamurthy/reference-data-api
``

## Swagger UI
``
go get -u github.com/swaggo/swag/cmd/swag
``

### 
``
../../../../bin/swag init --output swaggerui
``

The above must be run everytime you want to regenerate the swagger documentation.

SwaggerUI:
http://localhost:8082/swaggerui