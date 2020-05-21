## A Sample REST API with Go
Demonstrates a decent MVC pattern, configuration management, messaging using NATS io, database connectivity.

``go get github.com/saiprasadkrishnamurthy/web-api
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