/*
 * Go module for JSON parser
 * http://www.likexian.com/
 *
 * Copyright 2012-2014, Kexian Li
 * Released under the Apache License, Version 2.0
 *
 */

package simplejson


import (
    "encoding/json"
    "errors"
    "io"
    "io/ioutil"
    "os"
)


type Json struct {
    Data interface{}
}


func Version() string {
    return "0.2.0"
}


func Author() string {
    return "[Li Kexian](http://www.likexian.com/)"
}


func License() string {
    return "Apache License, Version 2.0"
}


func Load(file string) (result *Json, err error) {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        return
    }

    text := string(data)
    result, err = Loads(text)

    return
}


func Dump(file string, data *Json) (bytes int, err error) {
    result, err := PrettyDumps(data)
    if err != nil {
        return
    }

    fd, err := os.OpenFile(file, os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0644)
    if err != nil {
        return
    }

    bytes, err = io.WriteString(fd, result)

    return
}


func Loads(text string) (result *Json, err error) {
    result = new(Json)
    err = json.Unmarshal([]byte(text), &result.Data)

    return
}


func Dumps(j *Json) (result string, err error) {
    data, err := json.Marshal(&j.Data)
    if err != nil {
        return
    }
    result = string(data)

    return
}


func PrettyDumps(j *Json) (result string, err error) {
    data, err := json.MarshalIndent(&j.Data, "", "    ")
    if err != nil {
        return
    }
    result = string(data)

    return
}


func (j *Json) Exists(key string) (bool) {
    result, err := j.Map()
    if err == nil {
        _, exists := result[key]
        return exists
    }

    return false
}


func (j *Json) Get(key string) (*Json) {
    result, err := j.Map()
    if err == nil {
        if value, exists := result[key]; exists {
            return &Json{value}
        }
    }

    return &Json{nil}
}


func (j *Json) Set(key, value string) {
    result, err := j.Map()
    if err == nil {
        result[key] = value
    }
}


func (j *Json) Map() (result map[string]interface{}, err error) {
    result, ok := (j.Data).(map[string]interface{})
    if !ok {
        err = errors.New("assert to map failed")
    }
    return
}


func (j *Json) Array() (result []interface{}, err error) {
    result, ok := (j.Data).([]interface{})
    if !ok {
        err = errors.New("assert to array failed")
    }
    return
}


func (j *Json) Bool() (result bool, err error) {
    result, ok := (j.Data).(bool)
    if !ok {
        err = errors.New("assert to bool failed")
    }
    return
}


func (j *Json) String() (result string, err error) {
    result, ok := (j.Data).(string)
    if !ok {
        err = errors.New("assert to string failed")
    }
    return
}


func (j *Json) Int() (result int, err error) {
    f, ok := (j.Data).(float64)
    result = int(f)
    if !ok {
        err = errors.New("assert to int failed")
    }
    return
}


func (j *Json) Float() (result float64, err error) {
    result, ok := (j.Data).(float64)
    if !ok {
        err = errors.New("assert to float64 failed")
    }
    return
}
