# simplejson.go

simplejson-go is a simple Go module for JSON parsing.

[![Build Status](https://secure.travis-ci.org/likexian/simplejson-go.png)](https://secure.travis-ci.org/likexian/simplejson-go)
[![GoDoc](https://godoc.org/github.com/likexian/simplejson-go?status.svg)](https://godoc.org/github.com/likexian/simplejson-go)

## Overview

This module provides the most friendly way to use JSON in golang.

## Installation

    go get github.com/likexian/simplejson-go

## Importing

    import (
        "github.com/likexian/simplejson-go"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/simplejson-go)

## Example

The struct for storing data

    type Status struct {
        Code    int64  `json:"code"`
        Message string `json:"message"`
    }

Initialize the struct and set value

    data_status := Status{}
    data_status.Code = 1
    data_status.Message = "success"

Dump the struct data to JSON string

    data_json := New(data_status)
    text, err := data_json.Dumps()
    if err == nil {
        // Get the text of JSON
        fmt.Println(text)
    }

Load the JSON string

    data_json, err := Loads(text)
    if err == nil {
        // Get the value of JSON
        fmt.Println(data_json.Get("Code").Int())
        fmt.Println(data_json.Get("Message").String())
    }

## LICENSE

Copyright 2012-2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
