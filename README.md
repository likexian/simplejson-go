# simplejson.go

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/likexian/simplejson-go?status.svg)](https://godoc.org/github.com/likexian/simplejson-go)
[![Build Status](https://travis-ci.org/likexian/simplejson-go.svg?branch=master)](https://travis-ci.org/likexian/simplejson-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/likexian/simplejson-go)](https://goreportcard.com/report/github.com/likexian/simplejson-go)

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

Dump the struct data to JSON string

    type Status struct {
        Code    int64  `json:"code"`
        Message string `json:"message"`
    }

    statusData := Status{}
    statusData.Code = 1
    statusData.Message = "Success"

    jsonData := simplejson.New(statusData)
    jsonText, err := jsonData.Dumps()
    if err == nil {
        // Get the text of JSON
        fmt.Println(jsonText)
    }

Load the JSON string

    text := `{"Code": 1, "Message": "Success", "Result": {"Student": [{"Name": "Li Kexian"}]}}`
    jsonData, err := simplejson.Loads(text)
    if err == nil {
        // Get the value of JSON
        fmt.Println(jsonData.Get("Code").Int())
        fmt.Println(jsonData.Get("Message").String())
        fmt.Println(jsonData.Get("Result.Student.0.Name").MustString("-"))
    }

## LICENSE

Copyright 2012-2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
