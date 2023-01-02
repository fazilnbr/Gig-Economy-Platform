// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
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
        "/admin/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Login for admin",
                "operationId": "admin login authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Name : ",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password : ",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/domain.UserResponse"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/account/verifyJWT": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Varify JWT of users",
                "operationId": "Varify JWT authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email : ",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Login for users",
                "operationId": "login authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Name : ",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password : ",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/domain.UserResponse"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/send/verification": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Send OTP varification mail to users",
                "operationId": "SendVerificationMail authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email : ",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/domain.UserResponse"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "SignUp for users",
                "operationId": "SignUp authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Name : ",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password : ",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/domain.Login"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/verify/account": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Varify OTP of users",
                "operationId": "Varify OTP authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email : ",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OTP : ",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/woker/signup": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "SignUp for Workers",
                "operationId": "Worker SignUp authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Name : ",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password : ",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/domain.Login"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/worker/account/verifyJWT": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Varify JWT of users",
                "operationId": "Varify worker JWT authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email : ",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/worker/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Login for worker",
                "operationId": "worker login authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Name : ",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password : ",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/domain.UserResponse"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/worker/send/verification": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Send OTP varification mail to worker",
                "operationId": "Worker SendVerificationMail authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email : ",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/domain.UserResponse"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/worker/verify/account": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Varify OTP of users",
                "operationId": "Varify worker OTP authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email : ",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OTP : ",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "string"
                                        },
                                        "Errors": {
                                            "type": "string"
                                        },
                                        "Message": {
                                            "type": "string"
                                        },
                                        "Status": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Login": {
            "type": "object",
            "properties": {
                "id_login": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "usertype": {
                    "type": "string"
                },
                "verification": {
                    "type": "string"
                }
            }
        },
        "domain.UserResponse": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go + Gin Workey API",
	Description:      "This is a sample server Job Portal server. You can visit the GitHub repository at https://github.com/fazil/Job_Portal_Project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
