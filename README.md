# simplejson.go

simplejson-go is a simple Go module for JSON parser.

[![Build Status](https://secure.travis-ci.org/likexian/simplejson-go.png)](https://secure.travis-ci.org/likexian/simplejson-go)

## Overview

This module provides the most friendly way to use JSON in golang.

## Installation

    go get github.com/likexian/simplejson-go

## Importing

    import (
        "github.com/likexian/simplejson-go"
    )

## Documentation

Load string to JSON

    func Loads(text string) (result *Json, err error)

Dump JSON to string

    func Dumps(j *Json) (result string, err error)

Dump JSON to string with pretty format

    func PrettyDumps(j *Json) (result string, err error)

Load file to JSON

    func Load(file string) (result *Json, err error)

Dump JSON to file

    func Dump(file string, data *Json) (bytes int, err error)

Check key is in JSON

    func (j *Json) Exists(key string) (bool)

Get value from JSON

    func (j *Json) Get(key string) (*Json)

Set value to JSON

    func (j *Json) Set(key, value string)

Get the map value

    func (j *Json) Map() (result map[string]interface{}, err error)

Get the array value

    func (j *Json) Array() (result []interface{}, err error)

Get the bool value

    func (j *Json) Bool() (result bool, err error)

Get the string value

    func (j *Json) String() (result string, err error)

Get the int value

    func (j *Json) Int() (result int, err error)

Get the float value

    func (j *Json) Float() (result float64, err error)

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

    data_json := simplejson.Json{}
    data_json.Data = data_status
    data, err := simplejson.Dumps(&data_json)
    if err != nil {
        fmt.Println(data)
    }

Load the JSON string

    json, err := Loads(data)
    if err != nil {
        // Get the value of JSON
        fmt.Println(json.Get("Code").Int())
        fmt.Println(json.Get("Message").String())
    }

## LICENSE

Copyright 2012-2014, Kexian Li

Apache License, Version 2.0

## About

- [Kexian Li](http://github.com/likexian)
- [http://www.likexian.com/](http://www.likexian.com/)

## Thanks

- [Alex Warhawk](https://github.com/W4RH4WK)
- [John-Alan Simmons](https://github.com/jsimnz)
