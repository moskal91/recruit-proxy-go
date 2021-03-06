// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Recruit Proxy Server",
    "version": "1.0.0"
  },
  "paths": {
    "/clients.json": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Get all clients",
        "operationId": "getClients",
        "responses": {
          "200": {
            "description": "Client list JSON data",
            "schema": {
              "$ref": "#/definitions/Clients"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/invoices.json": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Get invoice by specified client ID",
        "operationId": "getInvoices",
        "parameters": [
          {
            "description": "Client ID",
            "name": "client_id",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "client_id"
              ],
              "properties": {
                "client_id": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Invoices JSON data",
            "schema": {
              "$ref": "#/definitions/Invoices"
            }
          },
          "400": {
            "description": "Invalid data",
            "schema": {
              "description": "Error message body",
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "Clients": {
      "type": "object",
      "properties": {
        "clients": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Invoices": {
      "type": "object",
      "properties": {
        "invoices": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "customer": {
                "type": "string"
              },
              "services": {
                "type": "string"
              },
              "total": {
                "type": "string"
              }
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Recruit Proxy Server",
    "version": "1.0.0"
  },
  "paths": {
    "/clients.json": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Get all clients",
        "operationId": "getClients",
        "responses": {
          "200": {
            "description": "Client list JSON data",
            "schema": {
              "$ref": "#/definitions/Clients"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/invoices.json": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Get invoice by specified client ID",
        "operationId": "getInvoices",
        "parameters": [
          {
            "description": "Client ID",
            "name": "client_id",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/getInvoicesParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Invoices JSON data",
            "schema": {
              "$ref": "#/definitions/Invoices"
            }
          },
          "400": {
            "description": "Invalid data",
            "schema": {
              "description": "Error message body",
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "Clients": {
      "type": "object",
      "properties": {
        "clients": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Invoices": {
      "type": "object",
      "properties": {
        "invoices": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/invoicesInvoicesItems"
          }
        }
      }
    },
    "getInvoicesParamsBody": {
      "type": "object",
      "required": [
        "client_id"
      ],
      "properties": {
        "client_id": {
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "invoicesInvoicesItems": {
      "type": "object",
      "properties": {
        "customer": {
          "type": "string"
        },
        "services": {
          "type": "string"
        },
        "total": {
          "type": "string"
        }
      },
      "x-go-gen-location": "models"
    }
  }
}`))
}
