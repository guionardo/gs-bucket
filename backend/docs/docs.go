// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Guionardo Furlan",
            "url": "https://github.com/guionardo/gs-bucket",
            "email": "guionardo@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "List all users allowed to publish",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key (master key)",
                        "name": "api-key",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.AuthResponse"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    }
                }
            }
        },
        "/auth/{user}": {
            "post": {
                "description": "Post a file to a pad, accessible for anyone",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Create a key for a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "api-key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.AuthResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Delete all keys of user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "api-key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    }
                }
            }
        },
        "/pads": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pads"
                ],
                "summary": "List pads",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.File"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Post a file to a pad, accessible for anyone",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pads"
                ],
                "summary": "Create a pad",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "api-key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "File name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Slug or easy name (if not informed, will be used a hash value)",
                        "name": "slug",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Time to live",
                        "name": "ttl",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "If informed, the file will be deleted after first download. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as ",
                        "name": "delete-after-read",
                        "in": "query"
                    },
                    {
                        "description": "Content",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.File"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    },
                    "507": {
                        "description": "Insufficient Storage",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    }
                }
            }
        },
        "/pads/{code}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pads"
                ],
                "summary": "Download a pad",
                "parameters": [
                    {
                        "type": "string",
                        "description": "File code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pads"
                ],
                "summary": "Delete a pad",
                "parameters": [
                    {
                        "type": "string",
                        "description": "File code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.AuthResponse": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "domain.File": {
            "type": "object",
            "properties": {
                "content_type": {
                    "type": "string"
                },
                "creation_date": {
                    "type": "string"
                },
                "delete_after_read": {
                    "type": "boolean"
                },
                "last_seen": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                },
                "seen_count": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "slug": {
                    "type": "string"
                },
                "valid_until": {
                    "type": "string"
                }
            }
        },
        "server.ErrResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "application-specific error code",
                    "type": "integer"
                },
                "error": {
                    "description": "application-level error message, for debugging",
                    "type": "string"
                },
                "status": {
                    "description": "user-level status message",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.5",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "GS-Bucket API",
	Description:      "This application will run a HTTP server to store files",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
