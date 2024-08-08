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
                "summary": "Mensage Rabbit order/article-data",
                "parameters": [
                    {
                        "description": "Estructura general del mensage",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rabbit.ConsumeMessage"
                        }
                    },
                    {
                        "description": "Message para Type = article-data",
                        "name": "article-data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/events.ValidationEvent"
                        }
                    },
                    {
                        "description": "Message para Type = place-order",
                        "name": "place-order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/events.PlacedOrderData"
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
                            "$ref": "#/definitions/rabbit.SendValidationMessage"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/rabbit/logout": {
            "put": {
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
                            "$ref": "#/definitions/rabbit.LogoutMessage"
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
                                "$ref": "#/definitions/order_proj.Order"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrValidation"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrCustom"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrCustom"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrCustom"
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
                            "$ref": "#/definitions/order_proj.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrValidation"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrCustom"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrCustom"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrCustom"
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
                            "$ref": "#/definitions/order_proj.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrValidation"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrCustom"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrCustom"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrCustom"
                        }
                    }
                }
            }
        },
        "/v1/orders_batch/payment_defined": {
            "get": {
                "description": "Ejecuta un proceso batch que chequea ordenes en estado PAYMENT_DEFINED.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordenes"
                ],
                "summary": "Batch Payment Defined",
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
                        "description": "No Content"
                    }
                }
            }
        },
        "/v1/orders_batch/placed": {
            "get": {
                "description": "Ejecuta un proceso batch para ordenes en estado PLACED.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordenes"
                ],
                "summary": "Batch Placed",
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
                        "description": "No Content"
                    }
                }
            }
        },
        "/v1/orders_batch/validated": {
            "get": {
                "description": "Ejecuta un proceso batch para ordenes en estado VALIDATED.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordenes"
                ],
                "summary": "Batch Validated",
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
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.ErrCustom": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "errors.ErrField": {
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
        "errors.ErrValidation": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/errors.ErrField"
                    }
                }
            }
        },
        "events.PaymentEvent": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "cartId": {
                    "type": "string"
                },
                "method": {
                    "$ref": "#/definitions/events.PaymentMethod"
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
                "isValid": {
                    "type": "boolean"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "order_proj.Article": {
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
        "order_proj.Order": {
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
                        "$ref": "#/definitions/order_proj.Article"
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
                        "$ref": "#/definitions/order_proj.PaymentEvent"
                    }
                },
                "status": {
                    "$ref": "#/definitions/order_proj.OrderStatus"
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
        "order_proj.OrderStatus": {
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
        "order_proj.PaymentEvent": {
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
        "rabbit.ArticleValidationData": {
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
        "rabbit.ConsumeMessage": {
            "type": "object",
            "properties": {
                "exchange": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
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
        "rabbit.LogoutMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "rabbit.SendValidationMessage": {
            "type": "object",
            "properties": {
                "exchange": {
                    "type": "string"
                },
                "message": {
                    "$ref": "#/definitions/rabbit.ArticleValidationData"
                },
                "queue": {
                    "type": "string"
                },
                "type": {
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
