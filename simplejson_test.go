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
    "fmt"
    "testing"
    "github.com/bmizerany/assert"
)


type JsonResult struct {
    Result Result `json:"result"`
    Status Status `json:"status"`
}


type Result struct {
    List    []int64 `json:"list"`
    Online  bool    `json:"online"`
    Rate    float64 `json:"rate"`
}


type Status struct {
    Code    int64  `json:"code"`
    Message string `json:"message"`
}


func TestSimplejson(t *testing.T) {
    data_result := Result{}
    data_result.List = []int64{0, 1, 2, 3, 4}
    data_result.Online = true
    data_result.Rate = 0.8

    data_status := Status{}
    data_status.Code = 1
    data_status.Message = "success"

    data_json_result := JsonResult{}
    data_json_result.Result = data_result
    data_json_result.Status = data_status

    data_json := Json{}
    data_json.Data = data_json_result

    data, err := Dumps(&data_json)
    assert.Equal(t, nil, err)
    assert.Equal(t, data, `{"result":{"list":[0,1,2,3,4],"online":true,"rate":0.8},"status":{"code":1,"message":"success"}}`)

    json, err := Loads(data)
    assert.NotEqual(t, nil, json)
    assert.Equal(t, nil, err)

    code, _ := json.Get("status").Get("code").Int()
    assert.Equal(t, 1, code)

    message, _ := json.Get("status").Get("message").String()
    assert.Equal(t, "success", message)

    exists := json.Has("result")
    assert.Equal(t, true, exists)

    deepexists := json.Get("result").Has("online")
    assert.Equal(t, true, deepexists)

    nexists := json.Has("not-exists")
    assert.Equal(t, false, nexists)

    json.Set("not-exists", "do-exists")
    sexists := json.Has("not-exists")
    assert.Equal(t, true, sexists)

    json.Del("not-exists")
    dexists := json.Has("not-exists")
    assert.Equal(t, false, dexists)

    list, _ := json.Get("result").Get("list").Array()
    assert.NotEqual(t, nil, list)
    for k, v := range list {
        assert.Equal(t, k, int(v.(float64)))
    }

    online, _ := json.Get("result").Get("online").Bool()
    assert.Equal(t, true, online)

    rate, _ := json.Get("result").Get("rate").Float64()
    assert.Equal(t, 0.80, rate)

    result, err := Dumps(json)
    assert.Equal(t, nil, err)
    assert.Equal(t, data, result)

    json.Set("name", "Li Kexian")
    json.Set("link", "http://www.likexian.com/")
    name, _ := json.Get("name").String()
    assert.Equal(t, "Li Kexian", name)

    bytes, err := Dump("simplejson.json", json)
    assert.NotEqual(t, 0, bytes)
    assert.Equal(t, nil, err)

    result, bytes2, err := DumpPretty("simplejson.json", json)
    assert.NotEqual(t, 0,  bytes2)
    assert.Equal(t, nil, err)
    fmt.Print(result)

    njson, err := Load("simplejson.json")
    assert.Equal(t, json, njson)
    assert.Equal(t, nil, err)

    new_json := New()
    new_json.Set("new", true)

    new_exists := new_json.Has("new")
    assert.Equal(t, true, new_exists)

    new_value, _ := new_json.Get("new").Bool()
    assert.Equal(t, true, new_value)

    new_json.Set("int", 100)
    int_value, _ := new_json.Get("int").Int()
    assert.Equal(t, "int", fmt.Sprintf("%T", int_value))

    uint_value, _ := new_json.Get("int").Uint()
    assert.Equal(t, "uint", fmt.Sprintf("%T", uint_value))

    int64_value, _ := new_json.Get("int").Int64()
    assert.Equal(t, "int64", fmt.Sprintf("%T", int64_value))

    uint64_value, _ := new_json.Get("int").Uint64()
    assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64_value))
}
