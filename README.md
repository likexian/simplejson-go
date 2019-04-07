# simplejson.go

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/likexian/simplejson-go?status.svg)](https://godoc.org/github.com/likexian/simplejson-go)
[![Build Status](https://travis-ci.org/likexian/simplejson-go.svg?branch=master)](https://travis-ci.org/likexian/simplejson-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/likexian/simplejson-go)](https://goreportcard.com/report/github.com/likexian/simplejson-go)
[![Code Cover](https://codecov.io/gh/likexian/simplejson-go/graph/badge.svg)](https://codecov.io/gh/likexian/simplejson-go)

simplejson-go is a simple and powerful Go module for JSON parsing.

## Overview

This module provides the most friendly way to use JSON in golang.

## Installation

    go get -u github.com/likexian/simplejson-go

## Importing

    import (
        "github.com/likexian/simplejson-go"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/simplejson-go)

## Example

### Dump the struct data to JSON string

```go
// Define Status struct
type Status struct {
    Code    int64  `json:"code"`
    Message string `json:"message"`
}

// Init status
status := Status{1, "Success"}

// Dump status to json string
j := simplejson.New(status)
s, err := j.Dumps()
if err == nil {
    fmt.Println("Json text is:", s)
}

// OR dumps using the easy way
s, err := simplejson.Dumps(status)
if err == nil {
    fmt.Println("Json text is:", s)
}
```

### Dump the map data to JSON string

```go
// Init a map data
data := map[string]interface{}{
    "code": 1,
    "message": "success",
    "result": {
        "Name": "Li Kexian"
    }
}

// Dump to string in the easy way
s, err := simplejson.Dumps(status)
if err == nil {
    fmt.Println("Json text is:", s)
}
```

### Load the JSON string

```go
// Json strig
text := `{"Code": 1, "Message": "Success", "Result": {"Student": [{"Name": "Li Kexian"}]}}`

// Load json string
j, err := simplejson.Loads(text)
if err == nil {
    fmt.Println("Code is:", j.Get("Code").MustInt(0))
    fmt.Println("Message is:", j.Get("Message").MustString(""))
    fmt.Println("First Student name is:", j.Get("Result.Student.0.Name").MustString("-"))
}
```

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
