// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users/login/": {
            "post": {
                "description": "Autentica a un usuario y devuelve un token de acceso y un token de refresco",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Iniciar sesión",
                "parameters": [
                    {
                        "description": "Credenciales del usuario",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Respuesta exitosa con tokens y detalles del usuario",
                        "schema": {
                            "$ref": "#/definitions/swagger.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Error: Datos inválidos",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Error: Usuario o contraseña incorrectos",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/users/refresh": {
            "post": {
                "description": "Refresca un token de acceso utilizando un token de refresco válido",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Refrescar el token de acceso",
                "parameters": [
                    {
                        "description": "Token de refresco",
                        "name": "refreshToken",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.RefreshTokenInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Respuesta exitosa con un nuevo token de acceso",
                        "schema": {
                            "$ref": "#/definitions/swagger.accessTokenResponse"
                        }
                    },
                    "400": {
                        "description": "Error: Datos inválidos",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Error: Token inválido o expirado",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/users/resources/{roleid}": {
            "get": {
                "description": "Retrieve resources associated with a specific role ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resources"
                ],
                "summary": "Get resources by role ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "roleid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved resources",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ResourceResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid role ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error fetching resources",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ResourceResponse": {
            "type": "object",
            "properties": {
                "actions": {
                    "type": "object",
                    "additionalProperties": true
                },
                "resourceName": {
                    "type": "string"
                }
            }
        },
        "swagger.LoginResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "id_company": {
                    "type": "integer"
                },
                "id_customer": {
                    "type": "integer"
                },
                "id_username": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "permission": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/swagger.PermissionResponse"
                    }
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "swagger.PermissionResponse": {
            "type": "object",
            "properties": {
                "delete": {
                    "type": "boolean"
                },
                "fk_module": {
                    "type": "integer"
                },
                "fk_role": {
                    "type": "integer"
                },
                "fk_username": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "read": {
                    "type": "boolean"
                },
                "update": {
                    "type": "boolean"
                },
                "write": {
                    "type": "boolean"
                }
            }
        },
        "swagger.RefreshTokenInput": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "swagger.accessTokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string",
                    "example": "tu_nuevo_token_de_acceso_ejemplo"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:60030",
	BasePath:         "/Vehicle",
	Schemes:          []string{},
	Title:            "Mi API",
	Description:      "Esta es mi API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}