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
    "fmt"
    "runtime"
    "encoding/json"
    "time"
    "testing"
    "github.com/bmizerany/assert"
)


type JsonResult struct {
    Result      Result      `json:"result"`
    Status      Status      `json:"status"`
}


type Result struct {
    IntList     []int64     `json:"intlist"`
    Online      bool        `json:"online"`
    Rate        float64     `json:"rate"`
}


type Status struct {
    Code        int64       `json:"code"`
    Message     string      `json:"message"`
}


var (
    json_result = JsonResult{}
    text_result = `{"result":{"intlist":[0,1,2,3,4],"online":true,"rate":0.8},"status":{"code":1,"message":"success"}}`
    text_file   = "simplejson.json"
    json_name   = "Li Kexian"
    json_link   = "https://www.likexian.com/"
)


func test_must_panic(t *testing.T, test_func func()) {
    defer func() {
        err := recover()
        if err == nil {
            _, file, line, ok := runtime.Caller(2)
            if ok {
                t.Errorf("%s: %d", file, line)
            }
        }
        assert.NotEqual(t, err, nil)
    }()

    test_func()
}


func init() {
    _ = os.Remove(text_file)

    data_result := Result{}
    data_result.IntList = []int64{0, 1, 2, 3, 4}
    data_result.Online = true
    data_result.Rate = 0.8

    data_status := Status{}
    data_status.Code = 1
    data_status.Message = "success"

    json_result.Result = data_result
    json_result.Status = data_status
}


func Test_New(t *testing.T) {
    // no init value to New
    json_data := New()
    json_text, err := json_data.Dumps()
    assert.Equal(t, err, nil)
    assert.Equal(t, json_text, `{}`)

    // pass init value to New
    json_data = New(json_result)
    json_text, err = json_data.Dumps()
    assert.Equal(t, err, nil)
    assert.Equal(t, json_text, text_result)

    // pass init map to New
    json_map := map[string]interface{}{"i": map[string]interface{}{"am": "Li Kexian", "age": 18}}
    json_data = New(json_map)
    json_text, err = json_data.Dumps()
    assert.Equal(t, err, nil)
    assert.Equal(t, json_text, `{"i":{"age":18,"am":"Li Kexian"}}`)
    name, err := json_data.Get("i").Get("am").String()
    assert.Equal(t, err, nil)
    assert.Equal(t, name, "Li Kexian")
}


func Test_Load_Dump(t *testing.T) {
    // Loads json from text
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Dumps json to text
    json_text, err := json_data.Dumps()
    assert.Equal(t, err, nil)
    assert.Equal(t, json_text, text_result)

    // Dump json to file
    b, err := json_data.Dump(text_file)
    assert.Equal(t, err, nil)
    assert.Equal(t, b, 246)

    // Load json from file
    json_data, err = Load(text_file)
    assert.Equal(t, err, nil)

    // Dumps json to text
    json_text, err = json_data.Dumps()
    assert.Equal(t, err, nil)
    assert.Equal(t, json_text, text_result)

    // Dumps json to text in pretty way
    json_text, err = json_data.PrettyDumps()
    assert.Equal(t, err, nil)
    assert.Equal(t, len(json_text), 246)
    assert.NotEqual(t, json_text, text_result)

    // Loads json from text of pretty
    json_data, err = Loads(json_text)
    assert.Equal(t, err, nil)
}


func Test_Set_Has_Get_Del(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Test key exists
    exists := json_data.Has("name")
    assert.Equal(t, exists, false)

    // Set key-value
    json_data.Set("name", json_name)
    json_data.Set("link", json_link)

    // Test dumpable
    _, err = json_data.Dump(text_file)
    assert.Equal(t, err, nil)

    // Test Set key-value
    exists = json_data.Has("name")
    assert.Equal(t, exists, true)

    // Get the Set name value
    r_name, err := json_data.Get("name").String()
    assert.Equal(t, err, nil)
    assert.Equal(t, json_name, r_name)

    // Get the Set link value
    r_link, err := json_data.Get("link").String()
    assert.Equal(t, err, nil)
    assert.Equal(t, json_link, r_link)

    // Get the not-exists key
    r_name, err = json_data.Get("not-exists").String()
    assert.NotEqual(t, err, nil)

    // Del key-value
    json_data.Del("name")
    exists = json_data.Has("name")
    assert.Equal(t, exists, false)

    // Del not-exists key
    json_data.Del("not-exists")
    exists = json_data.Has("not-exists")
    assert.Equal(t, exists, false)
}


func Test_Set_Has_Get_Del_W_Dot(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Test key exists
    exists := json_data.Has("i.am.that.who")
    assert.Equal(t, exists, false)

    // Set key-value
    json_data.Set("i.am.that.who", json_name)
    json_data.Set("name", json_name)
    json_data.Set("link", json_link)

    // Test dumpable
    _, err = json_data.Dump(text_file)
    assert.Equal(t, err, nil)

    // Test Set key-value
    exists = json_data.Has("i.am.that.who")
    assert.Equal(t, exists, true)

    // Get the Set name value
    r_name, err := json_data.Get("i.am.that.who").String()
    assert.Equal(t, err, nil)
    assert.Equal(t, json_name, r_name)

    // Get the not exists key
    r_name, err = json_data.Get("i.am.that.what").String()
    assert.NotEqual(t, err, nil)
    r_name, err = json_data.Get("i.am.this.who").String()
    assert.NotEqual(t, err, nil)

    // Get the Set name value with origin way
    r_name, err = json_data.Get("i").Get("am").Get("that").Get("who").String()
    assert.Equal(t, err, nil)
    assert.Equal(t, json_name, r_name)

    // Get the not exists key
    r_name, err = json_data.Get("i").Get("am").Get("that").Get("what").String()
    assert.NotEqual(t, err, nil)

    // Del key-value
    json_data.Del("i.am.that.who")
    exists = json_data.Has("i.am.that.who")
    assert.Equal(t, exists, false)

    // Del not-exists key
    json_data.Del("i.am.that.what")
    exists = json_data.Has("i.am.that.what")
    assert.Equal(t, exists, false)
}


func Test_Set_Has_Get_Del_W_List(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Test key exists
    exists := json_data.Has("that.is.a.list")
    assert.Equal(t, exists, false)

    // Set key-value
    json_data.Set("that.is.a.list", []interface{}{0, 1, 2, 3, 4})
    exists = json_data.Has("that.is.a.list")
    assert.Equal(t, exists, true)

    // Test N key exists
    exists = json_data.Has("that.is.a.list.3")
    assert.Equal(t, exists, true)

    // Test N key not exists
    exists = json_data.Has("that.is.a.list.666")
    assert.Equal(t, exists, false)

    // Test set dict in list
    json_data.Set("that.is.a.dict.in.list", []interface{}{map[string]interface{}{"a": 1, "b": 2, "c": 3}})
    exists = json_data.Has("that.is.a.dict.in.list")
    assert.Equal(t, exists, true)

    // Test dict in list exists
    exists = json_data.Has("that.is.a.dict.in.list.0")
    assert.Equal(t, exists, true)
    exists = json_data.Has("that.is.a.dict.in.list.0.a")
    assert.Equal(t, exists, true)
    exists = json_data.Has("that.is.a.dict.in.list.1.a")
    assert.Equal(t, exists, false)
    exists = json_data.Has("that.is.a.dict.in.list.0.z")
    assert.Equal(t, exists, false)

    // Test get dict in list
    int_data, err := json_data.Get("that.is.a.dict.in.list.0.b").Int()
    assert.Equal(t, err, nil)
    assert.Equal(t, int_data, 2)
    int_data, err = json_data.Get("that.is.a.dict.in.list.1.b").Int()
    assert.NotEqual(t, err, nil)
    int_data, err = json_data.Get("that.is.a.dict.in.list.0.z").Int()
    assert.NotEqual(t, err, nil)

    // Get the list value
    r_number, err := json_data.Get("that.is.a.list.3").Int()
    assert.Equal(t, err, nil)
    assert.Equal(t, r_number, 3)

    // Get not-exists N
    r_number, err = json_data.Get("that.is.a.list.666").Int()
    assert.NotEqual(t, err, nil)

    // Get the list value with origin way
    r_number, err = json_data.Get("that").Get("is").Get("a").Get("list").Get("4").Int()
    assert.Equal(t, err, nil)
    assert.Equal(t, r_number, 4)

    // Get the list value with GetN
    r_number, err = json_data.Get("that.is.a.list").GetN(1).Int()
    assert.Equal(t, err, nil)
    assert.Equal(t, r_number, 1)

    // Get the list value with origin way GetN
    r_number, err = json_data.Get("that").Get("is").Get("a").Get("list").GetN(2).Int()
    assert.Equal(t, err, nil)
    assert.Equal(t, r_number, 2)

    // Get not-exists N with origin way GetN
    r_number, err = json_data.Get("that").Get("is").Get("a").Get("list").GetN(666).Int()
    assert.NotEqual(t, err, nil)
}


func Test_Set_Has_Get_Del_Type(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Set bool value
    json_data.Set("bool", true)
    bool_data, err := json_data.Get("bool").Bool()
    assert.Equal(t, err, nil)
    assert.Equal(t, bool_data, true)
    assert.Equal(t, "bool", fmt.Sprintf("%T", bool_data))
    assert.Equal(t, json_data.Has("bool"), true)
    json_data.Del("bool")
    assert.Equal(t, json_data.Has("bool"), false)

    // Set string value
    json_data.Set("string", "string")
    string_data, err := json_data.Get("string").String()
    assert.Equal(t, err, nil)
    assert.Equal(t, string_data, "string")
    assert.Equal(t, "string", fmt.Sprintf("%T", string_data))
    assert.Equal(t, json_data.Has("string"), true)
    json_data.Del("string")
    assert.Equal(t, json_data.Has("string"), false)

    // Set float64 value
    json_data.Set("float64", float64(999))
    float64_data, err := json_data.Get("float64").Float64()
    assert.Equal(t, err, nil)
    assert.Equal(t, float64_data, float64(999))
    assert.Equal(t, "float64", fmt.Sprintf("%T", float64_data))
    assert.Equal(t, json_data.Has("float64"), true)
    json_data.Del("float64")
    assert.Equal(t, json_data.Has("float64"), false)

    // Set int value
    json_data.Set("int", int(666))
    int_data, err := json_data.Get("int").Int()
    assert.Equal(t, err, nil)
    assert.Equal(t, int_data, int(666))
    assert.Equal(t, "int", fmt.Sprintf("%T", int_data))
    assert.Equal(t, json_data.Has("int"), true)
    json_data.Del("int")
    assert.Equal(t, json_data.Has("int"), false)

    // Set int64 value
    json_data.Set("int64", int64(666))
    int64_data, err := json_data.Get("int64").Int64()
    assert.Equal(t, err, nil)
    assert.Equal(t, int64_data, int64(666))
    assert.Equal(t, "int64", fmt.Sprintf("%T", int64_data))
    assert.Equal(t, json_data.Has("int64"), true)
    json_data.Del("int64")
    assert.Equal(t, json_data.Has("int64"), false)

    // Set uint64 value
    json_data.Set("uint64", uint64(666))
    uint64_data, err := json_data.Get("uint64").Uint64()
    assert.Equal(t, err, nil)
    assert.Equal(t, uint64_data, uint64(666))
    assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64_data))
    assert.Equal(t, json_data.Has("uint64"), true)
    json_data.Del("uint64")
    assert.Equal(t, json_data.Has("uint64"), false)

    // Set string array value
    json_data.Set("string_array", []interface{}{"a", "b", "c"})
    string_array_data, err := json_data.Get("string_array").StringArray()
    assert.Equal(t, err, nil)
    assert.Equal(t, string_array_data, []string{"a", "b", "c"})
    assert.Equal(t, "[]string", fmt.Sprintf("%T", string_array_data))
    assert.Equal(t, json_data.Has("string_array"), true)
    json_data.Del("string_array")
    assert.Equal(t, json_data.Has("string_array"), false)
}


func Test_Get_Assert_Data(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Get data as map
    map_data, err := json_data.Get("status").Map()
    assert.Equal(t, err, nil)
    assert.Equal(t, "map[string]interface {}", fmt.Sprintf("%T", map_data))

    // Get data as array
    array_data, err := json_data.Get("result").Get("intlist").Array()
    assert.Equal(t, err, nil)
    assert.Equal(t, "[]interface {}", fmt.Sprintf("%T", array_data))
    for k, v := range array_data {
        r, _ := v.(json.Number).Int64()
        assert.Equal(t, k, int(r))
    }

    // Get data as bool
    bool_data, err := json_data.Get("result").Get("online").Bool()
    assert.Equal(t, err, nil)
    assert.Equal(t, "bool", fmt.Sprintf("%T", bool_data))
    assert.Equal(t, bool_data, true)

    // Get data as string
    string_data, err := json_data.Get("status").Get("message").String()
    assert.Equal(t, err, nil)
    assert.Equal(t, "string", fmt.Sprintf("%T", string_data))
    assert.Equal(t, string_data, "success")

    // Get data as float64
    float64_data, err := json_data.Get("result").Get("rate").Float64()
    assert.Equal(t, err, nil)
    assert.Equal(t, "float64", fmt.Sprintf("%T", float64_data))
    assert.Equal(t, float64_data, float64(0.8))

    // Get data as int
    int_data, err := json_data.Get("status").Get("code").Int()
    assert.Equal(t, err, nil)
    assert.Equal(t, "int", fmt.Sprintf("%T", int_data))
    assert.Equal(t, int_data, int(1))

    // Get data as int64
    int64_data, err := json_data.Get("status").Get("code").Int64()
    assert.Equal(t, err, nil)
    assert.Equal(t, "int64", fmt.Sprintf("%T", int64_data))
    assert.Equal(t, int64_data, int64(1))

    // Get data as uint64
    uint64_data, err := json_data.Get("status").Get("code").Uint64()
    assert.Equal(t, err, nil)
    assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64_data))
    assert.Equal(t, uint64_data, uint64(1))

    // Get data as string array
    json_data.Set("that.is.a.list", []interface{}{"a", "b", "c", "d", "e"})
    string_array_data, err := json_data.Get("that.is.a.list").StringArray()
    assert.Equal(t, err, nil)
    assert.Equal(t, "[]string", fmt.Sprintf("%T", string_array_data))
    assert.Equal(t, string_array_data, []string{"a", "b", "c", "d", "e"})
}


func Test_Get_Must_Assert_Data(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Get data as bool
    bool_data := json_data.Get("result").Get("online").MustBool()
    assert.Equal(t, "bool", fmt.Sprintf("%T", bool_data))
    assert.Equal(t, bool_data, true)

    // Get data as string
    string_data := json_data.Get("status").Get("message").MustString()
    assert.Equal(t, "string", fmt.Sprintf("%T", string_data))
    assert.Equal(t, string_data, "success")

    // Get data as float64
    float64_data := json_data.Get("result").Get("rate").MustFloat64()
    assert.Equal(t, "float64", fmt.Sprintf("%T", float64_data))
    assert.Equal(t, float64_data, float64(0.8))

    // Get data as int
    int_data := json_data.Get("status").Get("code").MustInt()
    assert.Equal(t, "int", fmt.Sprintf("%T", int_data))
    assert.Equal(t, int_data, int(1))

    // Get data as int64
    int64_data := json_data.Get("status").Get("code").MustInt64()
    assert.Equal(t, "int64", fmt.Sprintf("%T", int64_data))
    assert.Equal(t, int64_data, int64(1))

    // Get data as uint64
    uint64_data := json_data.Get("status").Get("code").MustUint64()
    assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64_data))
    assert.Equal(t, uint64_data, uint64(1))

    // Get data as string array
    json_data.Set("that.is.a.list", []interface{}{"a", "b", "c", "d", "e"})
    string_array_data := json_data.Get("that.is.a.list").MustStringArray()
    assert.Equal(t, "[]string", fmt.Sprintf("%T", string_array_data))
    assert.Equal(t, string_array_data, []string{"a", "b", "c", "d", "e"})
}


func Test_Get_Must_Assert_Data_N_Default(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Get data as bool
    test_must_panic(t, func(){
        json_data.Get("not-exists").MustBool()
    })

    // Get data as string
    test_must_panic(t, func(){
        json_data.Get("not-exists").MustString()
    })

    // Get data as float64
    test_must_panic(t, func(){
        json_data.Get("not-exists").MustFloat64()
    })

    // Get data as int
    test_must_panic(t, func(){
        json_data.Get("not-exists").MustInt()
    })

    // Get data as int64
    test_must_panic(t, func(){
        json_data.Get("not-exists").MustInt64()
    })

    // Get data as uint64
    test_must_panic(t, func(){
        json_data.Get("not-exists").MustUint64()
    })

    // Get data as string array
    test_must_panic(t, func(){
        json_data.Get("not-exists").MustStringArray()
    })
}


func Test_Get_Must_Assert_Data_W_Default(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Get data as bool
    bool_data := json_data.Get("not-exists").MustBool(true)
    assert.Equal(t, "bool", fmt.Sprintf("%T", bool_data))
    assert.Equal(t, bool_data, true)

    // Get data as string
    string_data := json_data.Get("not-exists").MustString("ok")
    assert.Equal(t, "string", fmt.Sprintf("%T", string_data))
    assert.Equal(t, string_data, "ok")

    // Get data as float64
    float64_data := json_data.Get("not-exists").MustFloat64(float64(999))
    assert.Equal(t, "float64", fmt.Sprintf("%T", float64_data))
    assert.Equal(t, float64_data, float64(999))

    // Get data as int
    int_data := json_data.Get("not-exists").MustInt(int(666))
    assert.Equal(t, "int", fmt.Sprintf("%T", int_data))
    assert.Equal(t, int_data, int(666))

    // Get data as int64
    int64_data := json_data.Get("not-exists").MustInt64(int64(666))
    assert.Equal(t, "int64", fmt.Sprintf("%T", int64_data))
    assert.Equal(t, int64_data, int64(666))

    // Get data as uint64
    uint64_data := json_data.Get("not-exists").MustUint64(uint64(666))
    assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64_data))
    assert.Equal(t, uint64_data, uint64(666))

    // Get data as string array
    string_array_data := json_data.Get("not-exists").MustStringArray([]string{"i", "am", "ok"})
    assert.Equal(t, "[]string", fmt.Sprintf("%T", string_array_data))
    assert.Equal(t, string_array_data, []string{"i", "am", "ok"})
}


func Test_HTML_Escape(t *testing.T) {
    // Init json and set html
    json_data := New()
    json_data.Set("param", "a=1&b=2&c=3")
    json_data.Set("title", "<title>test escape</title>")

    // dumps not escaped html
    json_text, err := json_data.Dumps()
    assert.Equal(t, err, nil)
    assert.Equal(t, json_text, `{"param":"a=1&b=2&c=3","title":"<title>test escape</title>"}`)

    // dumps escaped html
    json_data.SetHtmlEscape(true)
    json_text, err = json_data.Dumps()
    assert.Equal(t, err, nil)
    assert.Equal(t, json_text, `{"param":"a=1\u0026b=2\u0026c=3","title":"\u003ctitle\u003etest escape\u003c/title\u003e"}`)
}


func Test_Is_Map(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Test the top level json
    is_map := json_data.IsMap()
    assert.Equal(t, is_map, true)

    // Test after get
    is_map = json_data.Get("status").IsMap()
    assert.Equal(t, is_map, true)

    // Test after twice get
    is_map = json_data.Get("result").Get("online").IsMap()
    assert.Equal(t, is_map, false)

    // Test after magic get
    is_map = json_data.Get("result.online").IsMap()
    assert.Equal(t, is_map, false)

    // Test the array
    is_map = json_data.Get("result.intlist").IsMap()
    assert.Equal(t, is_map, false)

    // Test not exists key
    is_map = json_data.Get("result.not-exists").IsMap()
    assert.Equal(t, is_map, false)
}


func Test_Is_Array(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Test the top level json
    is_map := json_data.IsArray()
    assert.Equal(t, is_map, false)

    // Test after get
    is_map = json_data.Get("status").IsArray()
    assert.Equal(t, is_map, false)

    // Test after twice get
    is_map = json_data.Get("result").Get("intlist").IsArray()
    assert.Equal(t, is_map, true)

    // Test the array
    is_map = json_data.Get("result.intlist").IsArray()
    assert.Equal(t, is_map, true)

    // Test the array element
    is_map = json_data.Get("result.intlist.0").IsArray()
    assert.Equal(t, is_map, false)

    // Test not exists key
    is_map = json_data.Get("result.not-exists").IsArray()
    assert.Equal(t, is_map, false)
}


func Test_Len(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Get len of top level json
    n := json_data.Len()
    assert.Equal(t, n, 2)

    // Get len of map
    n = json_data.Get("status").Len()
    assert.Equal(t, n, 2)

    // Get len of not exists map
    n = json_data.Get("status.not-exists").Len()
    assert.Equal(t, n, -1)

    // Get len of int
    n = json_data.Get("status.code").Len()
    assert.Equal(t, n, -1)

    // Get len of string
    n = json_data.Get("status.message").Len()
    assert.Equal(t, n, 7)

    // Get len of not exists string
    n = json_data.Get("status.message.not-exists").Len()
    assert.Equal(t, n, -1)

    // Get len of array
    n = json_data.Get("result.intlist").Len()
    assert.Equal(t, n, 5)

    // Get len of not exists array
    n = json_data.Get("result.intlist.not-exists").Len()
    assert.Equal(t, n, -1)
}


func Test_Time(t *testing.T) {
    // Loads json for Set
    json_data, err := Loads(text_result)
    assert.Equal(t, err, nil)

    // Time for comparing
    test_time, err := time.Parse(time.RFC3339, "2019-01-31T12:11:10+08:00")
    assert.Equal(t, err, nil)
    assert.NotEqual(t, test_time.Unix(), int64(0))

    // Test get rfc3339 time
    json_data.Set("time", "2019-01-31T12:11:10+08:00")
    time_data, err := json_data.Get("time").Time()
    assert.Equal(t, err, nil)
    assert.Equal(t, time_data, test_time)

    // Test get format time
    json_data.Set("time", "2019-01-31 12:11:10")
    time_data, err = json_data.Get("time").Time("2006-01-02 15:04:05")
    assert.Equal(t, err, nil)
    assert.Equal(t, time_data, test_time)

    // Test get rfc3339 time with not exists key
    time_data, err = json_data.Get("not-exists").Time()
    assert.NotEqual(t, err, nil)
    assert.Equal(t, time_data.Unix(), int64(-62135596800))

    // Test get time from int
    json_data.Set("time", int(1548907870))
    time_data, err = json_data.Get("time").Time()
    assert.Equal(t, err, nil)
    assert.Equal(t, time_data, test_time)

    // Test get time from int64
    json_data.Set("time", int64(1548907870))
    time_data, err = json_data.Get("time").Time()
    assert.Equal(t, err, nil)
    assert.Equal(t, time_data, test_time)

    // Test get time from uint64
    json_data.Set("time", uint64(1548907870))
    time_data, err = json_data.Get("time").Time()
    assert.Equal(t, err, nil)
    assert.Equal(t, time_data, test_time)

    // Test get time from float64
    json_data.Set("time", float64(1548907870))
    time_data, err = json_data.Get("time").Time()
    assert.Equal(t, err, nil)
    assert.Equal(t, time_data, test_time)

    // Date for comparing
    test_time, err = time.ParseInLocation("2006-01-02", "2019-01-31", time.Local)
    assert.Equal(t, err, nil)
    assert.NotEqual(t, test_time.Unix(), int64(0))

    // Test get format date
    json_data.Set("time", "2019-01-31")
    time_data, err = json_data.Get("time").Time("2006-01-02")
    assert.Equal(t, err, nil)
    assert.Equal(t, time_data, test_time)
}
