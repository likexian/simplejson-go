/*
 * Go module for JSON parsing
 * https://www.likexian.com/
 *
 * Copyright 2012-2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package simplejson


import (
    "os"
    "io"
    "io/ioutil"
    "bytes"
    "errors"
    "encoding/json"
    "reflect"
    "strings"
    "strconv"
    "log"
)


// storing json data
type Json struct {
    Data interface{}
}


// returns package version
func Version() string {
    return "0.7.0"
}


// returns package author
func Author() string {
    return "[Li Kexian](https://www.likexian.com/)"
}


// returns package license
func License() string {
    return "Apache License, Version 2.0"
}


// returns a pointer to a new Json object
func New() (*Json) {
    return &Json {
        Data: make(map[string]interface{}),
    }
}


// loads data from a file, returns a json object
func Load(file string) (result *Json, err error) {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        return
    }

    text := string(data)
    result, err = Loads(text)

    return
}


// dumps json object to a file
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
    fd.Close()

    return
}


// unmarshal json from string, returns json object
func Loads(text string) (result *Json, err error) {
    result = new(Json)

    dec := json.NewDecoder(bytes.NewBuffer([]byte(text)))
    dec.UseNumber()
    err = dec.Decode(&result.Data)

    return
}


// marshal json object to string
func Dumps(j *Json) (result string, err error) {
    data, err := json.Marshal(&j.Data)
    if err != nil {
        return
    }

    result = string(data)

    return
}


// marshal json object to string, with identation
func PrettyDumps(j *Json) (result string, err error) {
    data, err := json.MarshalIndent(&j.Data, "", "    ")
    if err != nil {
        return
    }

    result = string(data)

    return
}


// set key-value to json object
func (j *Json) Set(key string, value interface{}) {
    result, err := j.Map()
    if err == nil {
        result[key] = value
    }
}


// check json object has key
func (j *Json) Has(key string) (bool) {
    result, err := j.Map()
    if err == nil {
        _, exists := result[key]
        return exists
    }

    return false
}


// delete key-value from json object
func (j *Json) Del(key string) {
    result, err := j.Map()
    if err == nil {
        delete(result, key)
    }
}


// returns the pointer to json object by key
//   json.Get("status").Get("code").Int()
func (j *Json) Get(key string) (*Json) {
    result, err := j.Map()
    if err == nil {
        if value, exists := result[key]; exists {
            return &Json{value}
        }
    }

    return &Json{nil}
}


// returns a pointer to the path of json object
//   json.Gets("status/code").Int()
//   json.Gets("result/intlist/3").Int()
func (j *Json)Gets(key string) (*Json) {
    result := j

    for _, v := range strings.Split(key, "/") {
        v = strings.TrimSpace(v)
        if v != "" {
            if result.Has(v) {
                result = result.Get(v)
            } else {
                i, err := strconv.Atoi(v)
                if err == nil {
                    result = result.GetIndex(i)
                }
            }
        }
    }

    return result
}


// returns a pointer to the index of json object
//   json.Get("int_list").GetIndex(1).Int()
func (j *Json)GetIndex(i int) (*Json) {
    data, err := j.Array()
    if err == nil {
        if len(data) > i {
            return &Json{data[i]}
        }
    }

    return &Json{nil}
}


// returns as map from json object
func (j *Json) Map() (result map[string]interface{}, err error) {
    result, ok := (j.Data).(map[string]interface{})
    if !ok {
        err = errors.New("assert to map failed")
    }
    return
}


// returns as array from json object
func (j *Json) Array() (result []interface{}, err error) {
    result, ok := (j.Data).([]interface{})
    if !ok {
        err = errors.New("assert to array failed")
    }
    return
}


// returns as bool from json object
func (j *Json) Bool() (result bool, err error) {
    result, ok := (j.Data).(bool)
    if !ok {
        err = errors.New("assert to bool failed")
    }
    return
}


// returns as string from json object
func (j *Json) String() (result string, err error) {
    result, ok := (j.Data).(string)
    if !ok {
        err = errors.New("assert to string failed")
    }
    return
}


// returns as string array from json object
func (j *Json) StringArray() (result []string, err error) {
    data, err := j.Array()
    if err != nil {
        return
    }

    for _, v := range data {
        if v == nil {
            result = append(result, "")
        } else {
            r, ok := v.(string)
            if !ok {
                err = errors.New("assert to []string failed")
                return
            }
            result = append(result, r)
        }
    }

    return
}


// returns as float64 from json object
func (j *Json) Float64() (result float64, err error) {
    switch j.Data.(type) {
        case json.Number:
            return j.Data.(json.Number).Float64()
        case float32, float64:
            return reflect.ValueOf(j.Data).Float(), nil
        case int, int8, int16, int32, int64:
            return float64(reflect.ValueOf(j.Data).Int()), nil
        case uint, uint8, uint16, uint32, uint64:
            return float64(reflect.ValueOf(j.Data).Uint()), nil
        default:
            return 0, errors.New("invalid value type")
    }
}


// returns as int from json object
func (j *Json) Int() (result int, err error) {
    switch j.Data.(type) {
        case json.Number:
            r, err := j.Data.(json.Number).Int64()
            return int(r), err
        case float32, float64:
            return int(reflect.ValueOf(j.Data).Float()), nil
        case int, int8, int16, int32, int64:
            return int(reflect.ValueOf(j.Data).Int()), nil
        case uint, uint8, uint16, uint32, uint64:
            return int(reflect.ValueOf(j.Data).Uint()), nil
        default:
            return 0, errors.New("invalid value type")
    }
}


// returns as int64 from json object
func (j *Json) Int64() (result int64, err error) {
    switch j.Data.(type) {
        case json.Number:
            return j.Data.(json.Number).Int64()
        case float32, float64:
            return int64(reflect.ValueOf(j.Data).Float()), nil
        case int, int8, int16, int32, int64:
            return reflect.ValueOf(j.Data).Int(), nil
        case uint, uint8, uint16, uint32, uint64:
            return int64(reflect.ValueOf(j.Data).Uint()), nil
        default:
            return 0, errors.New("invalid value type")
    }
}


// returns as uint64 from json object
func (j *Json) Uint64() (result uint64, err error) {
    switch j.Data.(type) {
        case json.Number:
            return strconv.ParseUint(j.Data.(json.Number).String(), 10, 64)
        case float32, float64:
            return uint64(reflect.ValueOf(j.Data).Float()), nil
        case int, int8, int16, int32, int64:
            return uint64(reflect.ValueOf(j.Data).Int()), nil
        case uint, uint8, uint16, uint32, uint64:
            return reflect.ValueOf(j.Data).Uint(), nil
        default:
            return 0, errors.New("invalid value type")
    }
}


// returns as bool from json object with optional default value
func (j *Json) MustBool(args ...bool) (bool) {
    var def bool

    switch len(args) {
        case 0:
        case 1:
            def = args[0]
        default:
            log.Panicf("MustBool received too many arguments %d > 1", len(args))
    }

    r, err := j.Bool()
    if err == nil {
        return r
    }

    return def
}


// returns as string from json object with optional default value
func (j *Json) MustString(args ...string) (string) {
    var def string

    switch len(args) {
        case 0:
        case 1:
            def = args[0]
        default:
            log.Panicf("MustString received too many arguments %d > 1", len(args))
    }

    r, err := j.String()
    if err == nil {
        return r
    }

    return def
}


// returns as string from json object with optional default value
func (j *Json) MustStringArray(args ...[]string) ([]string) {
    var def []string

    switch len(args) {
        case 0:
        case 1:
            def = args[0]
        default:
            log.Panicf("MustStringArray received too many arguments %d > 1", len(args))
    }

    r, err := j.StringArray()
    if err == nil {
        return r
    }

    return def
}


// returns as float64 from json object with optional default value
func (j *Json) MustFloat64(args ...float64) (float64) {
    var def float64

    switch len(args) {
        case 0:
        case 1:
            def = args[0]
        default:
            log.Panicf("MustFloat64 received too many arguments %d > 1", len(args))
    }

    r, err := j.Float64()
    if err == nil {
        return r
    }

    return def
}


// returns as int from json object with optional default value
func (j *Json) MustInt(args ...int) (int) {
    var def int

    switch len(args) {
        case 0:
        case 1:
            def = args[0]
        default:
            log.Panicf("MustInt received too many arguments %d > 1", len(args))
    }

    r, err := j.Int()
    if err == nil {
        return r
    }

    return def
}


// returns as int64 from json object with optional default value
func (j *Json) MustInt64(args ...int64) (int64) {
    var def int64

    switch len(args) {
        case 0:
        case 1:
            def = args[0]
        default:
            log.Panicf("MustInt64 received too many arguments %d > 1", len(args))
    }

    r, err := j.Int64()
    if err == nil {
        return r
    }

    return def
}


// returns as uint64 from json object with optional default value
func (j *Json) MustUint64(args ...uint64) (uint64) {
    var def uint64

    switch len(args) {
        case 0:
        case 1:
            def = args[0]
        default:
            log.Panicf("MustUint64 received too many arguments %d > 1", len(args))
    }

    r, err := j.Uint64()
    if err == nil {
        return r
    }

    return def
}
