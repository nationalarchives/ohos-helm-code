// Code generated by swaggo/swag. DO NOT EDIT.

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
            "name": "The National Archives"
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
        "/": {
            "get": {
                "description": "Test whether the api is running",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Test whether the API is running",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/discovery": {
            "get": {
                "description": "Requests to TNA Discovery API",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Discovery"
                ],
                "summary": "Requests to TNA Discovery API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.discoveryAllStruct"
                        }
                    }
                }
            }
        },
        "/movingImages": {
            "get": {
                "description": "Moving images queries",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MovingImages"
                ],
                "summary": "Moving images queries",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string query",
                        "name": "q",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "int page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.keywordReturnStruct"
                        }
                    },
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/sparql": {
            "post": {
                "description": "Send sparql direct to neptune",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sparql"
                ],
                "summary": "Send sparql direct to neptune",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "main.BindingsResultsTitleTopic": {
            "type": "object",
            "properties": {
                "title": {
                    "$ref": "#/definitions/main.TitleTopicStructValues"
                },
                "topics": {
                    "$ref": "#/definitions/main.TitleTopicStructValues"
                }
            }
        },
        "main.DiscoveryCodeCount": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "main.DiscoveryRecordDetails": {
            "type": "object",
            "properties": {
                "adminHistory": {
                    "type": "string"
                },
                "altName": {
                    "type": "string"
                },
                "arrangement": {
                    "type": "string"
                },
                "catalogueLevel": {
                    "type": "integer"
                },
                "closureCode": {
                    "type": "string"
                },
                "closureStatus": {
                    "type": "string"
                },
                "closureType": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "context": {
                    "type": "string"
                },
                "corpBodies": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "coveringDates": {
                    "type": "string"
                },
                "department": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "documentType": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "formerReferenceDep": {
                    "type": "string"
                },
                "formerReferencePro": {
                    "type": "string"
                },
                "heldBy": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "mapDesignation": {
                    "type": "string"
                },
                "mapScale": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "numEndDate": {
                    "type": "integer"
                },
                "numStartDate": {
                    "type": "integer"
                },
                "openingDate": {
                    "type": "string"
                },
                "physicalCondition": {
                    "type": "string"
                },
                "places": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "reference": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                },
                "source": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "taxonomies": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "urlParameters": {
                    "type": "string"
                }
            }
        },
        "main.KeywordStruct": {
            "type": "object",
            "properties": {
                "keyword": {
                    "type": "string"
                },
                "page": {
                    "type": "string"
                }
            }
        },
        "main.TitleTopicStructValues": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "main.discoveryAllStruct": {
            "type": "object",
            "properties": {
                "cagalogueLevels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DiscoveryCodeCount"
                    }
                },
                "closureStatuses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DiscoveryCodeCount"
                    }
                },
                "count": {
                    "type": "integer"
                },
                "departments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DiscoveryCodeCount"
                    }
                },
                "heldByReps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DiscoveryCodeCount"
                    }
                },
                "nextBatchMark": {
                    "type": "string"
                },
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DiscoveryRecordDetails"
                    }
                },
                "referenceFirstLetters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DiscoveryCodeCount"
                    }
                },
                "repositories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DiscoveryCodeCount"
                    }
                },
                "sources": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DiscoveryCodeCount"
                    }
                },
                "taxonomySubject": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DiscoveryCodeCount"
                    }
                },
                "timePeriods": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DiscoveryCodeCount"
                    }
                },
                "titleFirstLetters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DiscoveryCodeCount"
                    }
                }
            }
        },
        "main.keywordReturnStruct": {
            "type": "object",
            "properties": {
                "first": {
                    "type": "string"
                },
                "id": {
                    "$ref": "#/definitions/main.KeywordStruct"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.BindingsResultsTitleTopic"
                    }
                },
                "last": {
                    "type": "string"
                },
                "next": {
                    "type": "string"
                },
                "prev": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
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
	Title:            "OHOS api",
	Description:      "OHOS api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
