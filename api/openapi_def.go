package api

var OpenAPIDefinition = `
{
  "openapi": "3.0.2",
  "info": {
    "title": "WHIDS API documentation",
    "version": "1.0"
  },
  "servers": [
    {
      "url": "https://localhost:1520"
    }
  ],
  "paths": {
    "/endpoints": {
      "get": {
        "tags": [
          "Endpoint Management"
        ],
        "summary": "Get endpoints",
        "parameters": [
          {
            "name": "showkey",
            "in": "query",
            "description": "Show or not key",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "name": "group",
            "in": "query",
            "description": "Filter by group",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "status",
            "in": "query",
            "description": "Filter by status",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "criticality",
            "in": "query",
            "description": "Filter by criticality",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "criticality": 0,
                      "group": "",
                      "hostname": "OpenHappy",
                      "ip": "127.0.0.1",
                      "key": "arfYYaebYFBDfUlBhQXvQOaxY2bOSLpoHeHJWqsyxlaYygs1fLdSEc3WKp3VyN36",
                      "last-connection": "2022-01-06T13:44:02.293480175Z",
                      "last-detection": "2022-01-06T14:44:01.229835936+01:00",
                      "score": 0,
                      "status": "",
                      "uuid": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d"
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "put": {
        "tags": [
          "Endpoint Management"
        ],
        "summary": "Create a new endpoint",
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "criticality": 0,
                    "group": "",
                    "hostname": "",
                    "ip": "",
                    "key": "png9DJ0MN2phhj2yR6GivcQLG28sQlDO1dLv557EX7LNQilflOW8ewRSCjbWaFzF",
                    "last-connection": "0001-01-01T00:00:00Z",
                    "last-detection": "0001-01-01T00:00:00Z",
                    "score": 0,
                    "status": "",
                    "uuid": "f7a49e1a-8d73-1c60-e442-561d2e115fca"
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/endpoints/artifacts": {
      "get": {
        "tags": [
          "Artifact Search and Retrieval"
        ],
        "summary": "Artifacts on all endpoints",
        "parameters": [
          {
            "name": "since",
            "in": "query",
            "description": "Retrieve artifacts received since date (RFC3339)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string",
              "format": "date"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d": [
                      {
                        "base-url": "/endpoints/5a92baeb-9384-47d3-92b4-a0db6f9b8c6d/artifacts/5a92baeb-9384-47d3-92b4-a0db6f9b8c6d/3d8441643c204ba9b9dcb5c414b25a3129f66f6c/",
                        "creation": "2022-01-06T13:44:06.841824842Z",
                        "event-hash": "3d8441643c204ba9b9dcb5c414b25a3129f66f6c",
                        "files": [
                          {
                            "name": "bar.txt",
                            "size": 4,
                            "timestamp": "2022-01-06T13:44:06.851824872Z"
                          },
                          {
                            "name": "foo.txt",
                            "size": 4,
                            "timestamp": "2022-01-06T13:44:06.841824842Z"
                          }
                        ],
                        "modification": "2022-01-06T13:44:06.851824872Z",
                        "process-guid": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d"
                      }
                    ]
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/endpoints/reports": {
      "get": {
        "tags": [
          "Detection Reports"
        ],
        "summary": "Get all detection reports",
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d": {
                      "alert-count": 50,
                      "alert-criticality-metric": 0,
                      "avg-alert-criticality": 0,
                      "avg-signature-criticality": 0,
                      "bounded-score": 0,
                      "count-by-signature": {
                        "DefenderConfigChanged": 5,
                        "NewAutorun": 21,
                        "SuspiciousService": 5,
                        "UnknownServices": 9,
                        "UntrustedDriverLoaded": 10
                      },
                      "count-uniq-signatures": 5,
                      "identifier": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d",
                      "median-time": "2022-01-06T14:44:03.495817905+01:00",
                      "score": 0,
                      "signature-count": 50,
                      "signature-criticality-metric": 0,
                      "signature-diversity": 100,
                      "signatures": [
                        "NewAutorun",
                        "DefenderConfigChanged",
                        "UnknownServices",
                        "SuspiciousService",
                        "UntrustedDriverLoaded"
                      ],
                      "start-time": "2022-01-06T14:44:03.493541699+01:00",
                      "std-dev-alert-criticality": 0,
                      "std-dev-signature-criticality": -92233720368547760,
                      "stop-time": "2022-01-06T14:44:03.498094111+01:00",
                      "sum-alert-criticality": 0,
                      "sum-rule-criticality": 0,
                      "tactics": null,
                      "techniques": null
                    }
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/endpoints/{uuid}": {
      "get": {
        "tags": [
          "Endpoint Management"
        ],
        "summary": "Get information about a single endpoint",
        "parameters": [
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "criticality": 0,
                    "group": "",
                    "hostname": "OpenHappy",
                    "ip": "127.0.0.1",
                    "last-connection": "2022-01-06T13:44:02.293480175Z",
                    "last-detection": "2022-01-06T14:44:01.229835936+01:00",
                    "score": 0,
                    "status": "",
                    "uuid": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d"
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "Endpoint Management"
        ],
        "summary": "Modify an existing endpoint",
        "parameters": [
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "showkey",
            "in": "query",
            "description": "Show endpoint key in response",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "name": "newkey",
            "in": "query",
            "description": "Generate a new key for endpoint",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "boolean"
            }
          }
        ],
        "requestBody": {
          "description": "Fields to modify. NB: Not all the fields can be modified",
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "Item": {
                    "type": "object"
                  },
                  "command": {
                    "type": "object",
                    "properties": {
                      "args": {
                        "type": "array",
                        "items": {
                          "type": "string"
                        }
                      },
                      "background": {
                        "type": "boolean"
                      },
                      "completed": {
                        "type": "boolean"
                      },
                      "drop": {
                        "type": "array",
                        "items": {
                          "type": "object",
                          "properties": {
                            "data": {
                              "type": "string",
                              "format": "binary"
                            },
                            "error": {
                              "type": "string"
                            },
                            "name": {
                              "type": "string"
                            },
                            "uuid": {
                              "type": "string"
                            }
                          }
                        }
                      },
                      "error": {
                        "type": "string"
                      },
                      "expect-json": {
                        "type": "boolean"
                      },
                      "fetch": {
                        "type": "object",
                        "properties": {
                          "key(string)": {
                            "type": "object",
                            "properties": {
                              "data": {
                                "type": "string",
                                "format": "binary"
                              },
                              "error": {
                                "type": "string"
                              },
                              "name": {
                                "type": "string"
                              },
                              "uuid": {
                                "type": "string"
                              }
                            }
                          }
                        }
                      },
                      "json": {
                        "type": "object"
                      },
                      "name": {
                        "type": "string"
                      },
                      "sent": {
                        "type": "boolean"
                      },
                      "sent-time": {
                        "type": "string",
                        "format": "date"
                      },
                      "stderr": {
                        "type": "string",
                        "format": "binary"
                      },
                      "stdout": {
                        "type": "string",
                        "format": "binary"
                      },
                      "timeout": {
                        "type": "object"
                      },
                      "uuid": {
                        "type": "string"
                      }
                    }
                  },
                  "criticality": {
                    "type": "integer",
                    "format": "int64"
                  },
                  "group": {
                    "type": "string"
                  },
                  "hostname": {
                    "type": "string"
                  },
                  "ip": {
                    "type": "string"
                  },
                  "key": {
                    "type": "string"
                  },
                  "last-connection": {
                    "type": "string",
                    "format": "date"
                  },
                  "last-detection": {
                    "type": "string",
                    "format": "date"
                  },
                  "score": {
                    "type": "number",
                    "format": "double"
                  },
                  "status": {
                    "type": "string"
                  },
                  "uuid": {
                    "type": "string"
                  }
                }
              },
              "example": {
                "uuid": "",
                "hostname": "",
                "ip": "",
                "group": "New Group",
                "criticality": 0,
                "key": "New Key",
                "score": 0,
                "status": "New Status",
                "last-detection": "0001-01-01T00:00:00Z",
                "last-connection": "0001-01-01T00:00:00Z"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "criticality": 0,
                    "group": "New Group",
                    "hostname": "OpenHappy",
                    "ip": "127.0.0.1",
                    "last-connection": "2022-01-06T13:44:02.293480175Z",
                    "last-detection": "2022-01-06T14:44:01.229835936+01:00",
                    "score": 0,
                    "status": "New Status",
                    "uuid": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d"
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "delete": {
        "tags": [
          "Endpoint Management"
        ],
        "summary": "Delete an existing endpoint",
        "parameters": [
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "criticality": 0,
                    "group": "New Group",
                    "hostname": "OpenHappy",
                    "ip": "127.0.0.1",
                    "last-connection": "2022-01-06T13:44:02.293480175Z",
                    "last-detection": "2022-01-06T14:44:01.229835936+01:00",
                    "score": 0,
                    "status": "New Status",
                    "uuid": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d"
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/endpoints/{uuid}/artifacts": {
      "get": {
        "tags": [
          "Artifact Search and Retrieval"
        ],
        "summary": "Artifacts for a single endpoint",
        "parameters": [
          {
            "name": "since",
            "in": "query",
            "description": "Retrieve artifacts received since date (RFC3339)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string",
              "format": "date"
            }
          },
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "base-url": "/endpoints/5a92baeb-9384-47d3-92b4-a0db6f9b8c6d/artifacts/5a92baeb-9384-47d3-92b4-a0db6f9b8c6d/3d8441643c204ba9b9dcb5c414b25a3129f66f6c/",
                      "creation": "2022-01-06T13:44:06.841824842Z",
                      "event-hash": "3d8441643c204ba9b9dcb5c414b25a3129f66f6c",
                      "files": [
                        {
                          "name": "bar.txt",
                          "size": 4,
                          "timestamp": "2022-01-06T13:44:06.851824872Z"
                        },
                        {
                          "name": "foo.txt",
                          "size": 4,
                          "timestamp": "2022-01-06T13:44:06.841824842Z"
                        }
                      ],
                      "modification": "2022-01-06T13:44:06.851824872Z",
                      "process-guid": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d"
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/endpoints/{uuid}/artifacts/{pguid}/{ehash}/{filename}": {
      "get": {
        "tags": [
          "Artifact Search and Retrieval"
        ],
        "summary": "Retrieve the content of an artifact",
        "parameters": [
          {
            "name": "raw",
            "in": "query",
            "description": "Retrieve raw file content",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "name": "gunzip",
            "in": "query",
            "description": "Serve gunziped file content",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "pguid",
            "in": "path",
            "description": "pguid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "ehash",
            "in": "path",
            "description": "ehash path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "filename",
            "in": "path",
            "description": "filename path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": "QmxhaA==",
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/endpoints/{uuid}/command": {
      "get": {
        "tags": [
          "Endpoint Execution"
        ],
        "summary": "Get the result of a command executed on endpoint",
        "parameters": [
          {
            "name": "wait",
            "in": "query",
            "description": "Wait command to end before responding, making the call blocking",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "args": [
                      "printf",
                      "Hello World"
                    ],
                    "background": false,
                    "completed": true,
                    "drop": [],
                    "error": "",
                    "expect-json": false,
                    "fetch": {},
                    "json": null,
                    "name": "/usr/bin/printf",
                    "sent": true,
                    "sent-time": "2022-01-06T14:44:03.406035777+01:00",
                    "stderr": null,
                    "stdout": "SGVsbG8gV29ybGQ=",
                    "timeout": 0,
                    "uuid": "cc6e3d18-ee3a-49c6-49ae-53133b211775"
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "Endpoint Execution"
        ],
        "summary": "Send a command to be executed by the endpoint",
        "parameters": [
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "requestBody": {
          "description": "Command to be executed. One can also specify files \n\t\t\t\tto drop from the manager to the endpoint prior to command execution \n\t\t\t\tand files to fetch after execution. A timeout for the can also \n\t\t\t\tbe specified, if zero there will be no timeout.",
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "command-line": {
                    "type": "string"
                  },
                  "drop-files": {
                    "type": "array",
                    "items": {
                      "type": "string"
                    }
                  },
                  "fetch-files": {
                    "type": "array",
                    "items": {
                      "type": "string"
                    }
                  },
                  "timeout": {
                    "type": "object"
                  }
                }
              },
              "example": {
                "command-line": "printf \"Hello World\"",
                "fetch-files": null,
                "drop-files": null,
                "timeout": 0
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "command": {
                      "args": [
                        "Hello World"
                      ],
                      "background": false,
                      "completed": false,
                      "drop": [],
                      "error": "",
                      "expect-json": false,
                      "fetch": {},
                      "json": null,
                      "name": "printf",
                      "sent": false,
                      "sent-time": "0001-01-01T00:00:00Z",
                      "stderr": null,
                      "stdout": null,
                      "timeout": 0,
                      "uuid": "cc6e3d18-ee3a-49c6-49ae-53133b211775"
                    },
                    "criticality": 0,
                    "group": "",
                    "hostname": "OpenHappy",
                    "ip": "127.0.0.1",
                    "key": "Dww8vYgBztjICjj5cBWP5DTJmh0OsINnEFe65g6lYw853L52NVPndE5T1cLaKgwK",
                    "last-connection": "2022-01-06T13:44:03.395443028Z",
                    "last-detection": "2022-01-06T14:44:02.322396378+01:00",
                    "score": 0,
                    "status": "",
                    "uuid": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d"
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/endpoints/{uuid}/command/{field}": {
      "get": {
        "tags": [
          "Endpoint Execution"
        ],
        "summary": "Retrieve only a field of the command structure",
        "parameters": [
          {
            "name": "wait",
            "in": "query",
            "description": "Wait command to end before responding, making the call blocking",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "field",
            "in": "path",
            "description": "Field of the Command structure to return",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": "SGVsbG8gV29ybGQ=",
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/endpoints/{uuid}/detections": {
      "get": {
        "tags": [
          "Endpoint Log Retrieval"
        ],
        "summary": "Retrieve detections logs",
        "parameters": [
          {
            "name": "since",
            "in": "query",
            "description": "Retrieve logs since date (RFC3339)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string",
              "format": "date"
            }
          },
          {
            "name": "until",
            "in": "query",
            "description": "Retrieve logs until date (RFC3339)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string",
              "format": "date"
            }
          },
          {
            "name": "last",
            "in": "query",
            "description": "Return last logs from duration (ex: '1d' for last day)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "pivot",
            "in": "query",
            "description": "Timestamp to pivot around (RFC3339)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string",
              "format": "date"
            }
          },
          {
            "name": "delta",
            "in": "query",
            "description": "Delta duration used to pivot (ex: '5m' to get logs 5min around pivot) ",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "limit",
            "in": "query",
            "description": "Maximum number of reports to return",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          },
          {
            "name": "skip",
            "in": "query",
            "description": "Skip number of events",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          },
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "Event": {
                        "Detection": {
                          "Actions": [],
                          "Criticality": 8,
                          "Signature": [
                            "NewAutorun"
                          ]
                        },
                        "EdrData": {
                          "Endpoint": {
                            "Group": "",
                            "Hostname": "OpenHappy",
                            "IP": "127.0.0.1",
                            "UUID": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d"
                          },
                          "Event": {
                            "Detection": true,
                            "Hash": "bb3d916eabc676d02cb64d7154e1539801500d73",
                            "ReceiptTime": "2022-01-06T13:44:01.124856033Z"
                          }
                        },
                        "EventData": {
                          "CommandLine": "\"C:\\ProgramData\\Microsoft\\Windows Defender\\Platform\\4.18.2106.6-0\\MsMpEng.exe\"",
                          "CurrentDirectory": "C:\\Windows\\system32\\",
                          "Details": "Both",
                          "EventType": "SetValue",
                          "Image": "C:\\ProgramData\\Microsoft\\Windows Defender\\Platform\\4.18.2106.6-0\\MsMpEng.exe",
                          "ImageHashes": "SHA1=FBF03B5D6DC1A7EDAB0BA8D4DD27291C739E5813,MD5=B1C15F9DB942B373B2FC468B7048E63F,SHA256=1DC05B6DD6281840CEB822604B0E403E499180D636D02EC08AD77B4EB56F1B9C,IMPHASH=8AA2B8727E6858A3557A4C09970B9A5D",
                          "ImageSignature": "?",
                          "ImageSignatureStatus": "?",
                          "ImageSigned": "false",
                          "IntegrityLevel": "System",
                          "ProcessGuid": "{515cd0d1-7669-6123-4e00-000000007300}",
                          "ProcessId": "3276",
                          "ProcessThreatScore": "16",
                          "RuleName": "-",
                          "Services": "WinDefend",
                          "TargetObject": "HKCR\\CLSID\\{2DCD7FDB-8809-48E4-8E4F-3157C57CF987}\\InprocServer32\\ThreadingModel",
                          "User": "NT AUTHORITY\\SYSTEM",
                          "UtcTime": "2021-08-23 10:20:25.878"
                        },
                        "System": {
                          "Channel": "Microsoft-Windows-Sysmon/Operational",
                          "Computer": "DESKTOP-LJRVE06",
                          "EventID": 13,
                          "Execution": {
                            "ProcessID": 3220,
                            "ThreadID": 3848
                          },
                          "Keywords": {
                            "Name": "",
                            "Value": 9223372036854776000
                          },
                          "Level": {
                            "Name": "Information",
                            "Value": 4
                          },
                          "Opcode": {
                            "Name": "Info",
                            "Value": 0
                          },
                          "Provider": {
                            "Guid": "{5770385F-C22A-43E0-BF4C-06F5698FFBD9}",
                            "Name": "Microsoft-Windows-Sysmon"
                          },
                          "Task": {
                            "Name": "",
                            "Value": 0
                          },
                          "TimeCreated": {
                            "SystemTime": "2022-01-06T14:44:00.098502889+01:00"
                          }
                        }
                      }
                    },
                    {
                      "Event": {
                        "Detection": {
                          "Actions": [],
                          "Criticality": 8,
                          "Signature": [
                            "NewAutorun"
                          ]
                        },
                        "EdrData": {
                          "Endpoint": {
                            "Group": "",
                            "Hostname": "OpenHappy",
                            "IP": "127.0.0.1",
                            "UUID": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d"
                          },
                          "Event": {
                            "Detection": true,
                            "Hash": "bb3d916eabc676d02cb64d7154e1539801500d73",
                            "ReceiptTime": "2022-01-06T13:44:01.124509569Z"
                          }
                        },
                        "EventData": {
                          "CommandLine": "\"C:\\ProgramData\\Microsoft\\Windows Defender\\Platform\\4.18.2106.6-0\\MsMpEng.exe\"",
                          "CurrentDirectory": "C:\\Windows\\system32\\",
                          "Details": "Both",
                          "EventType": "SetValue",
                          "Image": "C:\\ProgramData\\Microsoft\\Windows Defender\\Platform\\4.18.2106.6-0\\MsMpEng.exe",
                          "ImageHashes": "SHA1=FBF03B5D6DC1A7EDAB0BA8D4DD27291C739E5813,MD5=B1C15F9DB942B373B2FC468B7048E63F,SHA256=1DC05B6DD6281840CEB822604B0E403E499180D636D02EC08AD77B4EB56F1B9C,IMPHASH=8AA2B8727E6858A3557A4C09970B9A5D",
                          "ImageSignature": "?",
                          "ImageSignatureStatus": "?",
                          "ImageSigned": "false",
                          "IntegrityLevel": "System",
                          "ProcessGuid": "{515cd0d1-7669-6123-4e00-000000007300}",
                          "ProcessId": "3276",
                          "ProcessThreatScore": "16",
                          "RuleName": "-",
                          "Services": "WinDefend",
                          "TargetObject": "HKCR\\CLSID\\{2DCD7FDB-8809-48E4-8E4F-3157C57CF987}\\InprocServer32\\ThreadingModel",
                          "User": "NT AUTHORITY\\SYSTEM",
                          "UtcTime": "2021-08-23 10:20:25.878"
                        },
                        "System": {
                          "Channel": "Microsoft-Windows-Sysmon/Operational",
                          "Computer": "DESKTOP-LJRVE06",
                          "EventID": 13,
                          "Execution": {
                            "ProcessID": 3220,
                            "ThreadID": 3848
                          },
                          "Keywords": {
                            "Name": "",
                            "Value": 9223372036854776000
                          },
                          "Level": {
                            "Name": "Information",
                            "Value": 4
                          },
                          "Opcode": {
                            "Name": "Info",
                            "Value": 0
                          },
                          "Provider": {
                            "Guid": "{5770385F-C22A-43E0-BF4C-06F5698FFBD9}",
                            "Name": "Microsoft-Windows-Sysmon"
                          },
                          "Task": {
                            "Name": "",
                            "Value": 0
                          },
                          "TimeCreated": {
                            "SystemTime": "2022-01-06T14:44:00.098502889+01:00"
                          }
                        }
                      }
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/endpoints/{uuid}/logs": {
      "get": {
        "tags": [
          "Endpoint Log Retrieval"
        ],
        "summary": "Retrieve any kind of logs (detections + filtered)",
        "parameters": [
          {
            "name": "since",
            "in": "query",
            "description": "Retrieve logs since date (RFC3339)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string",
              "format": "date"
            }
          },
          {
            "name": "until",
            "in": "query",
            "description": "Retrieve logs until date (RFC3339)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string",
              "format": "date"
            }
          },
          {
            "name": "last",
            "in": "query",
            "description": "Return last logs from duration (ex: '1d' for last day)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "pivot",
            "in": "query",
            "description": "Timestamp to pivot around (RFC3339)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string",
              "format": "date"
            }
          },
          {
            "name": "delta",
            "in": "query",
            "description": "Delta duration used to pivot (ex: '5m' to get logs 5min around pivot) ",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "limit",
            "in": "query",
            "description": "Maximum number of reports to return",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          },
          {
            "name": "skip",
            "in": "query",
            "description": "Skip number of events",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          },
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "Event": {
                        "EdrData": {
                          "Endpoint": {
                            "Group": "",
                            "Hostname": "OpenHappy",
                            "IP": "127.0.0.1",
                            "UUID": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d"
                          },
                          "Event": {
                            "Detection": false,
                            "Hash": "05f430438c37dd8d6f70d74ab4c826235216046a",
                            "ReceiptTime": "2022-01-06T13:44:01.120537178Z"
                          }
                        },
                        "EventData": {
                          "CommandLine": "C:\\Windows\\system32\\svchost.exe -k appmodel -p -s StateRepository",
                          "CurrentDirectory": "C:\\Windows\\system32\\",
                          "Details": "Binary Data",
                          "EventType": "SetValue",
                          "Image": "C:\\Windows\\system32\\svchost.exe",
                          "ImageHashes": "SHA1=75C5A97F521F760E32A4A9639A653EED862E9C61,MD5=9520A99E77D6196D0D09833146424113,SHA256=DD191A5B23DF92E12A8852291F9FB5ED594B76A28A5A464418442584AFD1E048,IMPHASH=247B9220E5D9B720A82B2C8B5069AD69",
                          "ImageSignature": "?",
                          "ImageSignatureStatus": "?",
                          "ImageSigned": "false",
                          "IntegrityLevel": "System",
                          "ProcessGuid": "{515cd0d1-7668-6123-3c00-000000007300}",
                          "ProcessId": "2556",
                          "ProcessThreatScore": "0",
                          "RuleName": "-",
                          "Services": "StateRepository",
                          "TargetObject": "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\AppModel\\StateRepository\\Cache\\Package\\Data\\1eb\\_IndexKeys",
                          "User": "NT AUTHORITY\\SYSTEM",
                          "UtcTime": "2021-08-23 10:20:29.923"
                        },
                        "System": {
                          "Channel": "Microsoft-Windows-Sysmon/Operational",
                          "Computer": "DESKTOP-LJRVE06",
                          "EventID": 13,
                          "Execution": {
                            "ProcessID": 3220,
                            "ThreadID": 3848
                          },
                          "Keywords": {
                            "Name": "",
                            "Value": 9223372036854776000
                          },
                          "Level": {
                            "Name": "Information",
                            "Value": 4
                          },
                          "Opcode": {
                            "Name": "Info",
                            "Value": 0
                          },
                          "Provider": {
                            "Guid": "{5770385F-C22A-43E0-BF4C-06F5698FFBD9}",
                            "Name": "Microsoft-Windows-Sysmon"
                          },
                          "Task": {
                            "Name": "",
                            "Value": 0
                          },
                          "TimeCreated": {
                            "SystemTime": "2022-01-06T14:44:00.097624306+01:00"
                          }
                        }
                      }
                    },
                    {
                      "Event": {
                        "EdrData": {
                          "Endpoint": {
                            "Group": "",
                            "Hostname": "OpenHappy",
                            "IP": "127.0.0.1",
                            "UUID": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d"
                          },
                          "Event": {
                            "Detection": false,
                            "Hash": "1fe4b98ef33f53f23e694e69eb5720f8270fbe9a",
                            "ReceiptTime": "2022-01-06T13:44:01.120890623Z"
                          }
                        },
                        "EventData": {
                          "CommandLine": "C:\\Windows\\system32\\svchost.exe -k appmodel -p -s StateRepository",
                          "CurrentDirectory": "C:\\Windows\\system32\\",
                          "Details": "DWORD (0x0000001a)",
                          "EventType": "SetValue",
                          "Image": "C:\\Windows\\system32\\svchost.exe",
                          "ImageHashes": "SHA1=75C5A97F521F760E32A4A9639A653EED862E9C61,MD5=9520A99E77D6196D0D09833146424113,SHA256=DD191A5B23DF92E12A8852291F9FB5ED594B76A28A5A464418442584AFD1E048,IMPHASH=247B9220E5D9B720A82B2C8B5069AD69",
                          "ImageSignature": "?",
                          "ImageSignatureStatus": "?",
                          "ImageSigned": "false",
                          "IntegrityLevel": "System",
                          "ProcessGuid": "{515cd0d1-7668-6123-3c00-000000007300}",
                          "ProcessId": "2556",
                          "ProcessThreatScore": "0",
                          "RuleName": "-",
                          "Services": "StateRepository",
                          "TargetObject": "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\AppModel\\StateRepository\\Cache\\ApplicationExtension\\Data\\7e\\Application",
                          "User": "NT AUTHORITY\\SYSTEM",
                          "UtcTime": "2021-08-23 10:20:30.629"
                        },
                        "System": {
                          "Channel": "Microsoft-Windows-Sysmon/Operational",
                          "Computer": "DESKTOP-LJRVE06",
                          "EventID": 13,
                          "Execution": {
                            "ProcessID": 3220,
                            "ThreadID": 3848
                          },
                          "Keywords": {
                            "Name": "",
                            "Value": 9223372036854776000
                          },
                          "Level": {
                            "Name": "Information",
                            "Value": 4
                          },
                          "Opcode": {
                            "Name": "Info",
                            "Value": 0
                          },
                          "Provider": {
                            "Guid": "{5770385F-C22A-43E0-BF4C-06F5698FFBD9}",
                            "Name": "Microsoft-Windows-Sysmon"
                          },
                          "Task": {
                            "Name": "",
                            "Value": 0
                          },
                          "TimeCreated": {
                            "SystemTime": "2022-01-06T14:44:00.097624726+01:00"
                          }
                        }
                      }
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/endpoints/{uuid}/report": {
      "get": {
        "tags": [
          "Detection Reports"
        ],
        "summary": "Retrieve report for a single endpoint",
        "parameters": [
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "alert-count": 50,
                    "alert-criticality-metric": 0,
                    "avg-alert-criticality": 0,
                    "avg-signature-criticality": 0,
                    "bounded-score": 0,
                    "count-by-signature": {
                      "DefenderConfigChanged": 5,
                      "NewAutorun": 21,
                      "SuspiciousService": 5,
                      "UnknownServices": 9,
                      "UntrustedDriverLoaded": 10
                    },
                    "count-uniq-signatures": 5,
                    "identifier": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d",
                    "median-time": "2022-01-06T14:44:03.495817905+01:00",
                    "score": 0,
                    "signature-count": 50,
                    "signature-criticality-metric": 0,
                    "signature-diversity": 100,
                    "signatures": [
                      "UntrustedDriverLoaded",
                      "NewAutorun",
                      "DefenderConfigChanged",
                      "UnknownServices",
                      "SuspiciousService"
                    ],
                    "start-time": "2022-01-06T14:44:03.493541699+01:00",
                    "std-dev-alert-criticality": 0,
                    "std-dev-signature-criticality": -92233720368547760,
                    "stop-time": "2022-01-06T14:44:03.498094111+01:00",
                    "sum-alert-criticality": 0,
                    "sum-rule-criticality": 0,
                    "tactics": null,
                    "techniques": null
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "delete": {
        "tags": [
          "Detection Reports"
        ],
        "summary": "Delete and archive a report for a single endpoint",
        "parameters": [
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "alert-count": 50,
                    "alert-criticality-metric": 0,
                    "avg-alert-criticality": 0,
                    "avg-signature-criticality": 0,
                    "bounded-score": 0,
                    "count-by-signature": {
                      "DefenderConfigChanged": 5,
                      "NewAutorun": 21,
                      "SuspiciousService": 5,
                      "UnknownServices": 9,
                      "UntrustedDriverLoaded": 10
                    },
                    "count-uniq-signatures": 5,
                    "identifier": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d",
                    "median-time": "2022-01-06T14:44:03.495817905+01:00",
                    "score": 0,
                    "signature-count": 50,
                    "signature-criticality-metric": 0,
                    "signature-diversity": 100,
                    "signatures": [
                      "UntrustedDriverLoaded",
                      "NewAutorun",
                      "DefenderConfigChanged",
                      "UnknownServices",
                      "SuspiciousService"
                    ],
                    "start-time": "2022-01-06T14:44:03.493541699+01:00",
                    "std-dev-alert-criticality": 0,
                    "std-dev-signature-criticality": -92233720368547760,
                    "stop-time": "2022-01-06T14:44:03.498094111+01:00",
                    "sum-alert-criticality": 0,
                    "sum-rule-criticality": 0,
                    "tactics": null,
                    "techniques": null
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/endpoints/{uuid}/report/archive": {
      "get": {
        "tags": [
          "Detection Reports"
        ],
        "summary": "Get archived reports",
        "parameters": [
          {
            "name": "since",
            "in": "query",
            "description": "Retrieve report since date (RFC3339)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string",
              "format": "date"
            }
          },
          {
            "name": "until",
            "in": "query",
            "description": "Retrieve report until date (RFC3339)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string",
              "format": "date"
            }
          },
          {
            "name": "last",
            "in": "query",
            "description": "Return last reports from duration (ex: '1d' for last day)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "limit",
            "in": "query",
            "description": "Maximum number of reports to return",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          },
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "alert-count": 50,
                      "alert-criticality-metric": 0,
                      "archived-time": "2022-01-06T14:44:04.593119393+01:00",
                      "avg-alert-criticality": 0,
                      "avg-signature-criticality": 0,
                      "bounded-score": 0,
                      "count-by-signature": {
                        "DefenderConfigChanged": 5,
                        "NewAutorun": 21,
                        "SuspiciousService": 5,
                        "UnknownServices": 9,
                        "UntrustedDriverLoaded": 10
                      },
                      "count-uniq-signatures": 5,
                      "identifier": "5a92baeb-9384-47d3-92b4-a0db6f9b8c6d",
                      "median-time": "2022-01-06T14:44:03.495817905+01:00",
                      "score": 0,
                      "signature-count": 50,
                      "signature-criticality-metric": 0,
                      "signature-diversity": 100,
                      "signatures": [
                        "UntrustedDriverLoaded",
                        "NewAutorun",
                        "DefenderConfigChanged",
                        "UnknownServices",
                        "SuspiciousService"
                      ],
                      "start-time": "2022-01-06T14:44:03.493541699+01:00",
                      "std-dev-alert-criticality": 0,
                      "std-dev-signature-criticality": -92233720368547760,
                      "stop-time": "2022-01-06T14:44:03.498094111+01:00",
                      "sum-alert-criticality": 0,
                      "sum-rule-criticality": 0,
                      "tactics": null,
                      "techniques": null
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/iocs": {
      "get": {
        "tags": [
          "IoC Management (control IoCs pushed on Endpoints)"
        ],
        "summary": "Query IoCs loaded on manager and currently pushed to endpoints.\n\t\t\t\tQuery parameters can be used to restrict the search. Search criteria are\n\t\t\t\tORed together.",
        "parameters": [
          {
            "name": "source",
            "in": "query",
            "description": "Filter by source",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "key",
            "in": "query",
            "description": "Filter by key",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "value",
            "in": "query",
            "description": "Filter by value",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "type",
            "in": "query",
            "description": "Filter by type",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "guuid": "918cf5ba-636c-a565-7eb7-4a4028b62625",
                      "source": "XyzTIProvider",
                      "type": "domain",
                      "value": "some.random.domain"
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "IoC Management (control IoCs pushed on Endpoints)"
        ],
        "summary": "Add IoCs to be pushed on endpoints for detection",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "array",
                "items": {
                  "type": "object",
                  "properties": {
                    "Item": {
                      "type": "object"
                    },
                    "guuid": {
                      "type": "string"
                    },
                    "source": {
                      "type": "string"
                    },
                    "type": {
                      "type": "string"
                    },
                    "value": {
                      "type": "string"
                    }
                  }
                }
              },
              "example": [
                {
                  "source": "XyzTIProvider",
                  "guuid": "918cf5ba-636c-a565-7eb7-4a4028b62625",
                  "value": "some.random.domain",
                  "type": "domain"
                }
              ]
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "guuid": "918cf5ba-636c-a565-7eb7-4a4028b62625",
                      "source": "XyzTIProvider",
                      "type": "domain",
                      "value": "some.random.domain"
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "delete": {
        "tags": [
          "IoC Management (control IoCs pushed on Endpoints)"
        ],
        "summary": "Delete IoCs from manager, modulo a synchronization delay, endpoints should \n\t\t\tstop using those for detection. Query parameters can be used to select IoCs to delete.\n\t\t\tDeletion criteria are ANDed together.",
        "parameters": [
          {
            "name": "source",
            "in": "query",
            "description": "Filter by source",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "key",
            "in": "query",
            "description": "Filter by key",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "value",
            "in": "query",
            "description": "Filter by value",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "type",
            "in": "query",
            "description": "Filter by type",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "guuid": "918cf5ba-636c-a565-7eb7-4a4028b62625",
                      "source": "XyzTIProvider",
                      "type": "domain",
                      "value": "some.random.domain"
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/rules": {
      "get": {
        "tags": [
          "Rules Management"
        ],
        "summary": "Get rules loaded on endpoints",
        "parameters": [
          {
            "name": "name",
            "in": "query",
            "description": "Regex matching the names of the rules to retrieve",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "filters",
            "in": "query",
            "description": "Show only filters (rules used to filter-in logs)",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "boolean"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "Actions": [
                        "memdump",
                        "kill"
                      ],
                      "Condition": "$foo or $bar",
                      "Matches": [
                        "$foo: Image ~= 'C:\\\\Malware.exe'",
                        "$bar: TargetFilename ~= 'C:\\\\config.txt'"
                      ],
                      "Meta": {
                        "Computers": null,
                        "Criticality": 10,
                        "Disable": false,
                        "Events": {
                          "Microsoft-Windows-Sysmon/Operational": [
                            11,
                            23,
                            26
                          ]
                        },
                        "Filter": false,
                        "Schema": "2.0.0"
                      },
                      "Name": "TestRule",
                      "Tags": null
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "Rules Management"
        ],
        "summary": "Add or modify a rule",
        "parameters": [
          {
            "name": "update",
            "in": "query",
            "description": "Update rule if already existing",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "boolean"
            }
          }
        ],
        "requestBody": {
          "description": "Rule to add to the manager",
          "content": {
            "application/json": {
              "schema": {
                "type": "array",
                "items": {
                  "type": "object",
                  "properties": {
                    "Actions": {
                      "type": "array",
                      "items": {
                        "type": "string"
                      }
                    },
                    "Condition": {
                      "type": "string"
                    },
                    "Matches": {
                      "type": "array",
                      "items": {
                        "type": "string"
                      }
                    },
                    "Meta": {
                      "type": "object",
                      "properties": {
                        "ATTACK": {
                          "type": "array",
                          "items": {
                            "type": "object",
                            "properties": {
                              "": {
                                "type": "string"
                              },
                              "ID": {
                                "type": "string"
                              },
                              "Reference": {
                                "type": "string"
                              },
                              "Tactic": {
                                "type": "string"
                              }
                            }
                          }
                        },
                        "Computers": {
                          "type": "array",
                          "items": {
                            "type": "string"
                          }
                        },
                        "Criticality": {
                          "type": "integer",
                          "format": "int64"
                        },
                        "Disable": {
                          "type": "boolean"
                        },
                        "Events": {
                          "type": "object",
                          "properties": {
                            "key(string)": {
                              "type": "array",
                              "items": {
                                "type": "integer",
                                "format": "int64"
                              }
                            }
                          }
                        },
                        "Filter": {
                          "type": "boolean"
                        },
                        "Schema": {
                          "type": "object",
                          "properties": {
                            "Major": {
                              "type": "integer",
                              "format": "int64"
                            },
                            "Minor": {
                              "type": "integer",
                              "format": "int64"
                            },
                            "Patch": {
                              "type": "integer",
                              "format": "int64"
                            }
                          }
                        }
                      }
                    },
                    "Name": {
                      "type": "string"
                    },
                    "Tags": {
                      "type": "array",
                      "items": {
                        "type": "string"
                      }
                    }
                  }
                }
              },
              "example": [
                {
                  "Name": "TestRule",
                  "Tags": null,
                  "Meta": {
                    "Events": {
                      "Microsoft-Windows-Sysmon/Operational": [
                        11,
                        23,
                        26
                      ]
                    },
                    "Computers": null,
                    "Criticality": 10,
                    "Disable": false,
                    "Filter": false,
                    "Schema": "2.0.0"
                  },
                  "Matches": [
                    "$foo: Image ~= 'C:\\\\Malware.exe'",
                    "$bar: TargetFilename ~= 'C:\\\\config.txt'"
                  ],
                  "Condition": "$foo or $bar",
                  "Actions": [
                    "memdump",
                    "kill"
                  ]
                }
              ]
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "Actions": [
                        "memdump",
                        "kill"
                      ],
                      "Condition": "$foo or $bar",
                      "Matches": [
                        "$foo: Image ~= 'C:\\\\Malware.exe'",
                        "$bar: TargetFilename ~= 'C:\\\\config.txt'"
                      ],
                      "Meta": {
                        "Computers": null,
                        "Criticality": 10,
                        "Disable": false,
                        "Events": {
                          "Microsoft-Windows-Sysmon/Operational": [
                            11,
                            23,
                            26
                          ]
                        },
                        "Filter": false,
                        "Schema": "2.0.0"
                      },
                      "Name": "TestRule",
                      "Tags": null
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "delete": {
        "tags": [
          "Rules Management"
        ],
        "summary": "Delete rules from manager",
        "parameters": [
          {
            "name": "name",
            "in": "query",
            "description": "Name of the rule to delete. To avoid mistakes, this\n\t\t\t\tparameter cannot be a regex.",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "Actions": [
                        "memdump",
                        "kill"
                      ],
                      "Condition": "$foo or $bar",
                      "Matches": [
                        "$foo: Image ~= 'C:\\\\Malware.exe'",
                        "$bar: TargetFilename ~= 'C:\\\\config.txt'"
                      ],
                      "Meta": {
                        "Computers": null,
                        "Criticality": 10,
                        "Disable": false,
                        "Events": {
                          "Microsoft-Windows-Sysmon/Operational": [
                            11,
                            23,
                            26
                          ]
                        },
                        "Filter": false,
                        "Schema": "2.0.0"
                      },
                      "Name": "TestRule",
                      "Tags": null
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/stats": {
      "get": {
        "tags": [
          "Statistics about the manager"
        ],
        "summary": "Get statistics",
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "endpoint-count": 1,
                    "rule-count": 0
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "tags": [
          "Admin API User's Management"
        ],
        "summary": "List all users",
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": [
                    {
                      "description": "",
                      "group": "",
                      "identifier": "test",
                      "key": "plM6aQxHK74KL4MGlHL8OmYgRHiofhuKdid0PRYX7fTi6OrfQfmVUOOjhlnYBPCX",
                      "uuid": ""
                    }
                  ],
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "put": {
        "tags": [
          "Admin API User's Management"
        ],
        "summary": "Create a new user with identifier",
        "parameters": [
          {
            "name": "identifier",
            "in": "query",
            "description": "identifier query parameter",
            "required": true,
            "allowEmptyValue": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "description": "",
                    "group": "",
                    "identifier": "TestAdminUser",
                    "key": "iQTqHs5REEexuqJMhbeFRPbl3NOp9dRYeoF4vn3JyzkiKsM7Y4XCdkxET16HcMqr",
                    "uuid": "76603bcd-a3e4-8279-ec90-459c2d6785ce"
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "Admin API User's Management"
        ],
        "summary": "Create a new user from POST data",
        "requestBody": {
          "description": "Data to create the user with. Fields uuid and key if empty will be generated.",
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "Item": {
                    "type": "object"
                  },
                  "description": {
                    "type": "string"
                  },
                  "group": {
                    "type": "string"
                  },
                  "identifier": {
                    "type": "string"
                  },
                  "key": {
                    "type": "string"
                  },
                  "uuid": {
                    "type": "string"
                  }
                }
              },
              "example": {
                "uuid": "191033db-5727-d1c6-284e-802f291e61d1",
                "identifier": "SecondTestAdmin",
                "key": "ChangeMe",
                "group": "CSIRT",
                "description": "Second admin user"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "description": "Second admin user",
                    "group": "CSIRT",
                    "identifier": "SecondTestAdmin",
                    "key": "ChangeMe",
                    "uuid": "191033db-5727-d1c6-284e-802f291e61d1"
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    },
    "/users/{uuid}": {
      "post": {
        "tags": [
          "Admin API User's Management"
        ],
        "summary": "Modify existing admin API user",
        "parameters": [
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "newkey",
            "in": "query",
            "description": "Generate a new random key for user",
            "required": false,
            "allowEmptyValue": true,
            "schema": {
              "type": "boolean"
            }
          }
        ],
        "requestBody": {
          "description": "Data to update user with",
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "Item": {
                    "type": "object"
                  },
                  "description": {
                    "type": "string"
                  },
                  "group": {
                    "type": "string"
                  },
                  "identifier": {
                    "type": "string"
                  },
                  "key": {
                    "type": "string"
                  },
                  "uuid": {
                    "type": "string"
                  }
                }
              },
              "example": {
                "uuid": "",
                "identifier": "",
                "key": "NewWeakKey",
                "group": "SOC",
                "description": "Second admin user changed"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "description": "Second admin user changed",
                    "group": "SOC",
                    "identifier": "SecondTestAdmin",
                    "key": "NewWeakKey",
                    "uuid": "191033db-5727-d1c6-284e-802f291e61d1"
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      },
      "delete": {
        "tags": [
          "Admin API User's Management"
        ],
        "summary": "Delete an existing admin API user",
        "parameters": [
          {
            "name": "uuid",
            "in": "path",
            "description": "uuid path parameter",
            "required": true,
            "allowEmptyValue": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "HTTP 200 response",
            "content": {
              "application/json": {
                "example": {
                  "data": {
                    "description": "Second admin user changed",
                    "group": "SOC",
                    "identifier": "SecondTestAdmin",
                    "key": "NewWeakKey",
                    "uuid": "191033db-5727-d1c6-284e-802f291e61d1"
                  },
                  "error": "",
                  "message": "OK"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "securitySchemes": {
      "ApiKeyAuth": {
        "type": "apiKey",
        "name": "X-Api-Key",
        "in": "header"
      }
    }
  },
  "security": [
    {
      "ApiKeyAuth": []
    }
  ]
}
`
