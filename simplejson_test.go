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
    "fmt"
    "strings"
    "encoding/json"
    "testing"
    "github.com/bmizerany/assert"
)


type JsonResult struct {
    Result Result `json:"result"`
    Status Status `json:"status"`
}


type Result struct {
    IntList     []int64     `json:"intlist"`
    StrList     []string    `json:"strlist"`
    Online      bool        `json:"online"`
    Rate        float64     `json:"rate"`
}


type Status struct {
    Code    int64  `json:"code"`
    Message string `json:"message"`
}


func TestSimplejson(t *testing.T) {
    data_result := Result{}
    data_result.IntList = []int64{0, 1, 2, 3, 4}
    data_result.StrList = []string{"0", "1", "2", "3", "4"}
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
    assert.Equal(t, data, `{"result":{"intlist":[0,1,2,3,4],"strlist":["0","1","2","3","4"],"online":true,"rate":0.8},"status":{"code":1,"message":"success"}}`)

    json_data, err := Loads(data)
    assert.NotEqual(t, nil, json_data)
    assert.Equal(t, nil, err)

    code, _ := json_data.Get("status").Get("code").Int()
    assert.Equal(t, "int", fmt.Sprintf("%T", code))
    assert.Equal(t, 1, code)

    message, _ := json_data.Get("status").Get("message").String()
    assert.Equal(t, "string", fmt.Sprintf("%T", message))
    assert.Equal(t, "success", message)

    exists := json_data.Has("result")
    assert.Equal(t, true, exists)

    deepexists := json_data.Get("result").Has("online")
    assert.Equal(t, true, deepexists)

    nexists := json_data.Has("not-exists")
    assert.Equal(t, false, nexists)

    json_data.Set("not-exists", "do-exists")
    sexists := json_data.Has("not-exists")
    assert.Equal(t, true, sexists)

    json_data.Del("not-exists")
    dexists := json_data.Has("not-exists")
    assert.Equal(t, false, dexists)

    intlist, _ := json_data.Get("result").Get("intlist").Array()
    assert.NotEqual(t, nil, intlist)
    for k, v := range intlist {
        r, _ := v.(json.Number).Int64()
        assert.Equal(t, k, int(r))
    }

    strlist, _ := json_data.Get("result").Get("strlist").StringArray()
    assert.NotEqual(t, nil, strlist)
    for k, v := range strlist {
        assert.Equal(t, v, fmt.Sprintf("%d", k))
    }

    online, _ := json_data.Get("result").Get("online").Bool()
    assert.Equal(t, true, online)

    rate, _ := json_data.Get("result").Get("rate").Float64()
    assert.Equal(t, "float64", fmt.Sprintf("%T", rate))
    assert.Equal(t, 0.80, rate)

    index_data, err := json_data.Get("result").Get("intlist").GetIndex(3).Int()
    assert.Equal(t, nil, err)
    assert.Equal(t, "int", fmt.Sprintf("%T", index_data))
    assert.Equal(t, 3, index_data)

    index_data, err = json_data.Get("result").Get("intlist").GetIndex(10).Int()
    assert.NotEqual(t, nil, err)
    assert.Equal(t, "int", fmt.Sprintf("%T", index_data))
    assert.Equal(t, 0, index_data)

    gets_data, err := json_data.Gets("status/code").Int()
    assert.Equal(t, nil, err)
    assert.Equal(t, "int", fmt.Sprintf("%T", gets_data))
    assert.Equal(t, 1, gets_data)

    gets_data, err = json_data.Gets("result/intlist/3").Int()
    assert.Equal(t, nil, err)
    assert.Equal(t, "int", fmt.Sprintf("%T", gets_data))
    assert.Equal(t, 3, gets_data)

    gets_data, err = json_data.Gets("status/not-exists").Int()
    assert.NotEqual(t, nil, err)

    result, err := Dumps(json_data)
    assert.Equal(t, nil, err)
    assert.Equal(t, true, strings.Contains(result, `"strlist":["0","1","2","3","4"]`))
    assert.Equal(t, true, strings.Contains(result, `"intlist":[0,1,2,3,4]`))

    json_data.Set("name", "Li Kexian")
    json_data.Set("link", "https://www.likexian.com/")
    name, _ := json_data.Get("name").String()
    assert.Equal(t, "Li Kexian", name)

    bytes, err := Dump("simplejson.json", json_data)
    assert.NotEqual(t, 0, bytes)
    assert.Equal(t, nil, err)

    njson, err := Load("simplejson.json")
    assert.Equal(t, json_data, njson)
    assert.Equal(t, nil, err)

    json_data.Sets("status/error/code", 666)
    gets_data, err = json_data.Gets("status/error/code").Int()
    assert.Equal(t, nil, err)
    assert.Equal(t, "int", fmt.Sprintf("%T", gets_data))
    assert.Equal(t, 666, gets_data)

    new_json := New()
    new_json.Set("new", true)

    new_exists := new_json.Has("new")
    assert.Equal(t, true, new_exists)

    new_value, _ := new_json.Get("new").Bool()
    assert.Equal(t, "bool", fmt.Sprintf("%T", new_value))
    assert.Equal(t, true, new_value)

    new_json.Set("int", 100)
    int_value, _ := new_json.Get("int").Int()
    assert.Equal(t, "int", fmt.Sprintf("%T", int_value))
    assert.Equal(t, 100, int_value)

    new_json.Set("int64", int64(100))
    int64_value, _ := new_json.Get("int64").Int64()
    assert.Equal(t, "int64", fmt.Sprintf("%T", int64_value))
    assert.Equal(t, int64(100), int64_value)

    new_json.Set("uint64", uint64(100))
    uint64_value, _ := new_json.Get("uint64").Uint64()
    assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64_value))
    assert.Equal(t, uint64(100), uint64_value)

    new_value = new_json.Get("new").MustBool()
    assert.Equal(t, true, new_value)

    new_value = new_json.Get("not-exists").MustBool()
    assert.Equal(t, false, new_value)

    new_value = new_json.Get("not-exists").MustBool(true)
    assert.Equal(t, true, new_value)

    name = json_data.Get("name").MustString()
    assert.Equal(t, "Li Kexian", name)

    name = json_data.Get("not-exists").MustString()
    assert.Equal(t, "", name)

    name = json_data.Get("not-exists").MustString("default")
    assert.Equal(t, "default", name)

    strlist = json_data.Get("result").Get("strlist").MustStringArray()
    assert.Equal(t, []string{"0","1","2","3","4"}, strlist)

    strlist = json_data.Get("result").Get("not-exists").MustStringArray()
    assert.Equal(t, 0, len(strlist))

    strlist = json_data.Get("result").Get("not-exists").MustStringArray([]string{"default"})
    assert.Equal(t, []string{"default"}, strlist)

    rate = json_data.Get("result").Get("rate").MustFloat64()
    assert.Equal(t, 0.80, rate)

    rate = json_data.Get("result").Get("not-exists").MustFloat64()
    assert.Equal(t, 0.00, rate)

    rate = json_data.Get("result").Get("not-exists").MustFloat64(0.10)
    assert.Equal(t, 0.10, rate)

    int_value = new_json.Get("int").MustInt()
    assert.Equal(t, "int", fmt.Sprintf("%T", int_value))
    assert.Equal(t, 100, int_value)

    int_value = new_json.Get("not-exists").MustInt()
    assert.Equal(t, "int", fmt.Sprintf("%T", int_value))
    assert.Equal(t, 0, int_value)

    int_value = new_json.Get("not-exists").MustInt(1)
    assert.Equal(t, "int", fmt.Sprintf("%T", int_value))
    assert.Equal(t, 1, int_value)

    int64_value = new_json.Get("int64").MustInt64()
    assert.Equal(t, "int64", fmt.Sprintf("%T", int64_value))
    assert.Equal(t, int64(100), int64_value)

    int64_value = new_json.Get("not-exists").MustInt64()
    assert.Equal(t, "int64", fmt.Sprintf("%T", int64_value))
    assert.Equal(t, int64(0), int64_value)

    int64_value = new_json.Get("not-exists").MustInt64(int64(1))
    assert.Equal(t, "int64", fmt.Sprintf("%T", int64_value))
    assert.Equal(t, int64(1), int64_value)

    uint64_value = new_json.Get("uint64").MustUint64()
    assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64_value))
    assert.Equal(t, uint64(100), uint64_value)

    uint64_value = new_json.Get("not-exists").MustUint64()
    assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64_value))
    assert.Equal(t, uint64(0), uint64_value)

    uint64_value = new_json.Get("not-exists").MustUint64(uint64(1))
    assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64_value))
    assert.Equal(t, uint64(1), uint64_value)
}
