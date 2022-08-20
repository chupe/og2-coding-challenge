// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://chupe.ba/terms/",
        "contact": {
            "name": "Adnan",
            "email": "chupe@chupe.ba"
        },
        "license": {
            "name": "GPLv3",
            "url": "https://www.gnu.org/licenses/gpl-3.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/User": {
            "get": {
                "tags": [
                    "User"
                ],
                "summary": "Return all Users from the DB",
                "operationId": "get-Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/UserResponse"
                                }
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "Send User URL to create a new shortened User",
                "operationId": "create-User",
                "parameters": [
                    {
                        "description": "JSON with a 'url' field that contains full URL",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateUserBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/User/{id}": {
            "get": {
                "tags": [
                    "User"
                ],
                "summary": "Return User by ID from the DB",
                "operationId": "get-User",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "User"
                ],
                "summary": "Delete User object from the DB",
                "operationId": "delete-User",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "tags": [
                    "Health"
                ],
                "summary": "Check the status of the service",
                "operationId": "check-health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{code}": {
            "get": {
                "tags": [
                    "Code"
                ],
                "summary": "Return User by ID from the DB",
                "operationId": "get-code",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User Code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateUserBody": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "description": "Full url",
                    "type": "string",
                    "example": "http://chupe.ba"
                }
            }
        },
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "Actual error thrown",
                    "type": "string",
                    "example": "No Users found in the DB"
                },
                "message": {
                    "description": "User friendly message",
                    "type": "string",
                    "example": "Review input"
                },
                "status": {
                    "description": "Http status",
                    "type": "integer",
                    "example": 404
                }
            }
        },
        "UserResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Short alphanumeric 6 letter code that is used for redirection",
                    "type": "string",
                    "example": "a1b2c3"
                },
                "created": {
                    "description": "Date the User was stored",
                    "type": "string",
                    "format": "date-time",
                    "example": "2021-05-25T00:00:00.0Z"
                },
                "hitCount": {
                    "description": "Number of times the redirection took place",
                    "type": "integer",
                    "example": 42
                },
                "id": {
                    "description": "ObjectID represented as a string",
                    "type": "string",
                    "example": "62fbfaa5f79e97a5501979f3"
                },
                "url": {
                    "description": "Full URL",
                    "type": "string",
                    "example": "http://chupe.ba"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "User Shortener API Demo",
	Description:      "This is a practice project for Go",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
