package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "demo.com",
        "contact": {
            "name": "API Support",
            "url": "http://demo.com/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/protected/profile": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User By Token",
                "operationId": "GetUserByToken",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization header using the Bearer scheme",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/public/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login User",
                "operationId": "LoginUser",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "EnterDetails",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/public/signup": {
            "post": {
                "description": "Signin",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Signup User",
                "operationId": "SignupUser",
                "parameters": [
                    {
                        "description": "Signin",
                        "name": "EnterDetails",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserDetails"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
		"/pengeluaran/create": {
			"post": {
				"description": "Create Pengeluaran",
				"consumes": [
					"application/json"
				],	
				"tags": [
					"Pengeluaran"
				],
				"summary": "Create Pengeluaran",
				"operationId": "CreatePengeluaran",
				"parameters": [
					{
						"description": "Pengeluaran",
						"name": "EnterDetails",
						"in": "body",
						"required": true,
						"schema": {
							"$ref": "#/definitions/controllers.PengeluaranPayload"
						}
					}
				],
				"responses": {
					"200": {
						"description": "Success",
						"schema": {
							"type": "string"
						}
					},
					"400": {
						"description": "Error",
						"schema": {
							"type": "string"
						}
					}
				}
			}
		}
    },
    "definitions": {
        "controllers.LoginPayload": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controllers.UserDetails": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
		"controllers.PengeluaranPayload": {
			"type": "object",
			"required": [
				"jenis_pengeluaran",
				"harga",
				"tanggal",
				"keterangan"
			],
			"properties": {
				"jenis_pengeluaran": {
					"type": "string"
				},
				"harga": {
					"type": "string"
				},
				"tanggal": {
					"type": "string"
				},
				"keterangan": {
					"type": "string"
				}
			}
		}
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Swagger JWT API",
	Description:      "Create  Go REST API with JWT Authentication in Gin Framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
