// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://fastfood.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.fastfood.io/support",
            "email": "support@fastfood.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/clientes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Get all clients",
                "operationId": "get-all-clients",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Optional Filter by CPF",
                        "name": "cpf",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cliente.Cliente"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "New client",
                "operationId": "create-client",
                "parameters": [
                    {
                        "description": "Client data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/cliente.Cliente"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cliente.Cliente"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/clientes/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Get a client by ID",
                "operationId": "get-client-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cliente.Cliente"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Update a client",
                "operationId": "update-client",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Client data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/cliente.Cliente"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cliente.Cliente"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Delete a client by ID",
                "operationId": "delete-client-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "cliente.Cliente": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Orders API",
	Description:      "Here you will find everything you need to have the best possible integration with our APIs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
