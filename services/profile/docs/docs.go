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
        "/api/healthcheck": {
            "get": {
                "description": "Checks if the internal services used by the service are healthy",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Endpoint to check if the service is healthy",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/healthcheck.JSONHealthcheckReport"
                        }
                    }
                }
            }
        },
        "/api/v1/profiles/owners": {
            "get": {
                "description": "Get owner profile",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "owner"
                ],
                "summary": "Get owner profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Owner"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new owner profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "owner"
                ],
                "summary": "Create a new owner profile",
                "parameters": [
                    {
                        "description": "Owner Profile",
                        "name": "owner_profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Owner"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Owner"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/profiles/pets": {
            "get": {
                "description": "Get all pet profiles",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pet"
                ],
                "summary": "Get all pet profiles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Pet"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new pet profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pet"
                ],
                "summary": "Create a new pet profile",
                "parameters": [
                    {
                        "description": "Pet Profile",
                        "name": "pet_profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Pet"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Pet"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/profiles/pets/{id}": {
            "get": {
                "description": "Get pet profile by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pet"
                ],
                "summary": "Get pet profile by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pet Profile ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Pet"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "healthcheck.HealthcheckStatus": {
            "type": "string",
            "enum": [
                "healthy",
                "unhealthy",
                "degraded"
            ],
            "x-enum-varnames": [
                "HealthcheckStatusHealthy",
                "HealthcheckStatusUnhealthy",
                "HealthcheckStatusDegraded"
            ]
        },
        "healthcheck.JSONHealthcheckReport": {
            "type": "object",
            "properties": {
                "services": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/healthcheck.HealthcheckStatus"
                    }
                },
                "status": {
                    "$ref": "#/definitions/healthcheck.HealthcheckStatus"
                }
            }
        },
        "models.EGender": {
            "type": "string",
            "enum": [
                "Male",
                "Female"
            ],
            "x-enum-varnames": [
                "EGenderMale",
                "EGenderFemale"
            ]
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "request_id": {
                    "type": "string"
                }
            }
        },
        "models.Owner": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                }
            }
        },
        "models.Pet": {
            "type": "object",
            "properties": {
                "breed": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "gender": {
                    "$ref": "#/definitions/models.EGender"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "species": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
