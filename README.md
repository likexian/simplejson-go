# simplejson.go

simplejson-go is a simple and powerful Go module for JSON parsing.

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

    data_json := simplejson.New(data_status)
    text, err := data_json.Dumps()
    if err == nil {
        // Get the text of JSON
        fmt.Println(text)
    }

Load the JSON string

    text := `{"Code": 1, "Message": "Ok", "Result": {"Student": [{"Name": "Li Kexian"}]}}`
    data_json, err := simplejson.Loads(text)
    if err == nil {
        // Get the int value of JSON
        int_code, err := data_json.Get("Code").Int()
        fmt.Println(int_code, err)
        // Get the string value of JSON
        str_message, err := data_json.Get("Message").String()
        fmt.Println(str_message, err)
        // Magic get the string value of JSON with default
        fmt.Println(data_json.Get("Result.Student.0.Name").MustString("-"))
    }

## LICENSE

Copyright 2012-2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
