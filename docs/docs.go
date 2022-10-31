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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/search": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Пошук по параметрам",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "number",
                        "description": "Шукати по  transactionid",
                        "name": "transactionid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Шукати по  terminalid шукати за декількома одночасно можна через кому наприклад 3507,3508,3509....",
                        "name": "terminalid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Шукати по  status accepted/declined",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Шукати по  payment type cash/card ",
                        "name": "paymenttype",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Шукати по  date post  рік-місяць-день з,по. Наприклад: 2022-08-18,2022-09-28",
                        "name": "datepost",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Шукати по  Payment narrative",
                        "name": "paymentnarrative",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/searchcsv": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Пошук по параметрам з відповіддю у форматі CSV",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "number",
                        "description": "Шукати по  transactionid",
                        "name": "transactionid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Шукати по  terminalid шукати за декількома одночасно можна через кому наприклад 3507,3508,3509....",
                        "name": "terminalid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Шукати по  status accepted/declined",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Шукати по  payment type cash/card ",
                        "name": "paymenttype",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Шукати по  date post  рік-місяць-день з,по. Наприклад: 2022-08-18,2022-09-28",
                        "name": "datepost",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Шукати по  Payment narrative",
                        "name": "paymentnarrative",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/uploadfile": {
            "put": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "uploadfile"
                ],
                "summary": "Пошук по параметрам з відповіддю у форматі CSV",
                "operationId": "file.upload",
                "parameters": [
                    {
                        "type": "file",
                        "description": "тест файл",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
