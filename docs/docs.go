// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Nestor Marsollier",
            "email": "nmarsollier@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/rabbit/article-data": {
            "get": {
                "description": "Cuando se consume place-order se genera la orden y se inicia el proceso.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Mensage Rabbit order/article-data",
                "parameters": [
                    {
                        "description": "Message para Type = place-order",
                        "name": "place-order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/consume.consumePlaceDataMessage"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/rabbit/cart/article-data": {
            "put": {
                "description": "Antes de iniciar las operaciones se validan los artículos contra el catalogo.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Emite Validar Artículos a Cart cart/article-data",
                "parameters": [
                    {
                        "description": "Mensage de validacion",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/emit.SendValidationMessage"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/rabbit/logout": {
            "get": {
                "description": "Escucha de mensajes logout desde auth.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Mensage Rabbit",
                "parameters": [
                    {
                        "description": "Estructura general del mensage",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/consume.logoutMessage"
                        }
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "SendOrderPlaced envía un broadcast a rabbit con logout. Esto no es Rest es RabbitMQ.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Mensage Rabbit",
                "parameters": [
                    {
                        "description": "Order Placed Event",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/emit.message"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/orders": {
            "get": {
                "description": "Busca todas las ordenes del usuario logueado.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordenes"
                ],
                "summary": "Ordenes de Usuario",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ordenes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/rest.OrderListData"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/orders/:orderId": {
            "get": {
                "description": "Busca una order del usuario logueado, por su id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordenes"
                ],
                "summary": "Buscar Orden",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de orden",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ordenes",
                        "schema": {
                            "$ref": "#/definitions/order_projection.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/orders/:orderId/payment": {
            "post": {
                "description": "Agrega un Pago",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordenes"
                ],
                "summary": "Agrega un Pago",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de orden",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Informacion del pago",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/events.PaymentEvent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ordenes",
                        "schema": {
                            "$ref": "#/definitions/order_projection.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/orders/:orderId/update": {
            "get": {
                "description": "Actualiza las proyecciones en caso que hayamos roto algo.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordenes"
                ],
                "summary": "Actualiza la proyeccion",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID de orden",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "consume.consumeArticleDataMessage": {
            "type": "object",
            "properties": {
                "exchange": {
                    "type": "string"
                },
                "message": {
                    "$ref": "#/definitions/events.ValidationEvent"
                },
                "queue": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "consume.consumePlaceDataMessage": {
            "type": "object",
            "properties": {
                "exchange": {
                    "type": "string"
                },
                "message": {
                    "$ref": "#/definitions/events.PlacedOrderData"
                },
                "queue": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "example": "place-order"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "consume.logoutMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg"
                },
                "type": {
                    "type": "string",
                    "example": "logout"
                }
            }
        },
        "emit.ArticleValidationData": {
            "type": "object",
            "properties": {
                "articleId": {
                    "type": "string"
                },
                "referenceId": {
                    "type": "string"
                }
            }
        },
        "emit.SendValidationMessage": {
            "type": "object",
            "properties": {
                "exchange": {
                    "type": "string"
                },
                "message": {
                    "$ref": "#/definitions/emit.ArticleValidationData"
                },
                "queue": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "emit.articlePlacedData": {
            "type": "object",
            "properties": {
                "articleId": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "emit.message": {
            "type": "object",
            "properties": {
                "exchange": {
                    "type": "string"
                },
                "message": {
                    "$ref": "#/definitions/emit.orderPlacedData"
                },
                "queue": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "emit.orderPlacedData": {
            "type": "object",
            "properties": {
                "articles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/emit.articlePlacedData"
                    }
                },
                "cartId": {
                    "type": "string"
                },
                "orderId": {
                    "type": "string"
                }
            }
        },
        "engine.ErrorData": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "errs.ValidationErr": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/errs.errField"
                    }
                }
            }
        },
        "errs.errField": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "events.PaymentEvent": {
            "type": "object",
            "required": [
                "amount",
                "method",
                "orderId"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "method": {
                    "$ref": "#/definitions/events.PaymentMethod"
                },
                "orderId": {
                    "type": "string"
                }
            }
        },
        "events.PaymentMethod": {
            "type": "string",
            "enum": [
                "CASH",
                "CREDIT",
                "DEBIT"
            ],
            "x-enum-varnames": [
                "Cash",
                "Credit",
                "Debit"
            ]
        },
        "events.PlacePrderArticleData": {
            "type": "object",
            "required": [
                "id",
                "quantity"
            ],
            "properties": {
                "id": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "quantity": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "events.PlacedOrderData": {
            "type": "object",
            "required": [
                "articles",
                "cartId",
                "userId"
            ],
            "properties": {
                "articles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/events.PlacePrderArticleData"
                    }
                },
                "cartId": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "userId": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                }
            }
        },
        "events.ValidationEvent": {
            "type": "object",
            "properties": {
                "articleId": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "referenceId": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "order_projection.Article": {
            "type": "object",
            "required": [
                "articleId",
                "quantity"
            ],
            "properties": {
                "articleId": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "isValid": {
                    "type": "boolean"
                },
                "isValidated": {
                    "type": "boolean"
                },
                "quantity": {
                    "type": "integer",
                    "minimum": 1
                },
                "unitaryPrice": {
                    "type": "number"
                }
            }
        },
        "order_projection.Order": {
            "type": "object",
            "required": [
                "cartId",
                "orderId",
                "status",
                "userId"
            ],
            "properties": {
                "articles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order_projection.Article"
                    }
                },
                "cartId": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "orderId": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "payments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order_projection.PaymentEvent"
                    }
                },
                "status": {
                    "$ref": "#/definitions/order_projection.OrderStatus"
                },
                "updated": {
                    "type": "string"
                },
                "userId": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                }
            }
        },
        "order_projection.OrderStatus": {
            "type": "string",
            "enum": [
                "placed",
                "invalid",
                "validated",
                "payment_defined"
            ],
            "x-enum-varnames": [
                "Placed",
                "Invalid",
                "Validated",
                "Payment_Defined"
            ]
        },
        "order_projection.PaymentEvent": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "method": {
                    "$ref": "#/definitions/events.PaymentMethod"
                }
            }
        },
        "rest.OrderListData": {
            "type": "object",
            "properties": {
                "articles": {
                    "type": "integer"
                },
                "cartId": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/order_projection.OrderStatus"
                },
                "totalPayment": {
                    "type": "number"
                },
                "totalPrice": {
                    "type": "number"
                },
                "updated": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3004",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "OrdersGo",
	Description:      "Microservicio de Ordenes.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
