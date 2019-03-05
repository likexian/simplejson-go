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
	"encoding/json"
	"fmt"
	"github.com/likexian/gokit/assert"
	"os"
	"runtime"
	"testing"
	"time"
)

type JsonResult struct {
	Result Result `json:"result"`
	Status Status `json:"status"`
}

type Result struct {
	IntList []int64 `json:"intlist"`
	Online  bool    `json:"online"`
	Rate    float64 `json:"rate"`
}

type Status struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

var (
	jsonResult = JsonResult{}
	textResult = `{"result":{"intlist":[0,1,2,3,4],"online":true,"rate":0.8},"status":{"code":1,"message":"success"}}`
	textFile   = "simplejson.json"
	jsonName   = "Li Kexian"
	jsonLink   = "https://www.likexian.com/"
)

func testMustPanic(t *testing.T, testFunc func()) {
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

	testFunc()
}

func init() {
	_ = os.Remove(textFile)

	dataResult := Result{}
	dataResult.IntList = []int64{0, 1, 2, 3, 4}
	dataResult.Online = true
	dataResult.Rate = 0.8

	dataStatus := Status{}
	dataStatus.Code = 1
	dataStatus.Message = "success"

	jsonResult.Result = dataResult
	jsonResult.Status = dataStatus
}

func Test_New(t *testing.T) {
	// no init value to New
	jsonData := New()
	jsonText, err := jsonData.Dumps()
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonText, `{}`)

	// pass init value to New
	jsonData = New(jsonResult)
	jsonText, err = jsonData.Dumps()
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonText, textResult)

	// pass init map to New
	jsonMap := map[string]interface{}{"i": map[string]interface{}{"am": "Li Kexian", "age": 18}}
	jsonData = New(jsonMap)
	jsonText, err = jsonData.Dumps()
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonText, `{"i":{"age":18,"am":"Li Kexian"}}`)
	name, err := jsonData.Get("i").Get("am").String()
	assert.Equal(t, err, nil)
	assert.Equal(t, name, "Li Kexian")
}

func Test_Load_Dump(t *testing.T) {
	// Loads json from text
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Dumps json to text
	jsonText, err := jsonData.Dumps()
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonText, textResult)

	// Dump json to file
	err = jsonData.Dump(textFile)
	assert.Equal(t, err, nil)

	// Load json from file
	jsonData, err = Load(textFile)
	assert.Equal(t, err, nil)

	// Dumps json to text
	jsonText, err = jsonData.Dumps()
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonText, textResult)

	// Dumps json to text in pretty way
	jsonText, err = jsonData.PrettyDumps()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(jsonText), 246)
	assert.NotEqual(t, jsonText, textResult)

	// Loads json from text of pretty
	jsonData, err = Loads(jsonText)
	assert.Equal(t, err, nil)
}

func Test_Set_Has_Get_Del(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Test key exists
	exists := jsonData.Has("name")
	assert.Equal(t, exists, false)

	// Set key-value
	jsonData.Set("name", jsonName)
	jsonData.Set("link", jsonLink)

	// Test dumpable
	err = jsonData.Dump(textFile)
	assert.Equal(t, err, nil)

	// Test Set key-value
	exists = jsonData.Has("name")
	assert.Equal(t, exists, true)

	// Get the Set name value
	rName, err := jsonData.Get("name").String()
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonName, rName)

	// Get the Set link value
	rLink, err := jsonData.Get("link").String()
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonLink, rLink)

	// Get the not-exists key
	_, err = jsonData.Get("not-exists").String()
	assert.NotEqual(t, err, nil)

	// Del key-value
	jsonData.Del("name")
	exists = jsonData.Has("name")
	assert.Equal(t, exists, false)

	// Del not-exists key
	jsonData.Del("not-exists")
	exists = jsonData.Has("not-exists")
	assert.Equal(t, exists, false)
}

func Test_Set_Has_Get_Del_W_Dot(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Test key exists
	exists := jsonData.Has("i.am.that.who")
	assert.Equal(t, exists, false)

	// Set key-value
	jsonData.Set("i.am.that.who", jsonName)
	jsonData.Set("name", jsonName)
	jsonData.Set("link", jsonLink)

	// Test dumpable
	err = jsonData.Dump(textFile)
	assert.Equal(t, err, nil)

	// Test Set key-value
	exists = jsonData.Has("i.am.that.who")
	assert.Equal(t, exists, true)

	// Get the Set name value
	rName, err := jsonData.Get("i.am.that.who").String()
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonName, rName)

	// Get the not exists key
	_, err = jsonData.Get("i.am.that.what").String()
	assert.NotEqual(t, err, nil)
	_, err = jsonData.Get("i.am.this.who").String()
	assert.NotEqual(t, err, nil)

	// Get the Set name value with origin way
	rName, err = jsonData.Get("i").Get("am").Get("that").Get("who").String()
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonName, rName)

	// Get the not exists key
	_, err = jsonData.Get("i").Get("am").Get("that").Get("what").String()
	assert.NotEqual(t, err, nil)

	// Del key-value
	jsonData.Del("i.am.that.who")
	exists = jsonData.Has("i.am.that.who")
	assert.Equal(t, exists, false)

	// Del not-exists key
	jsonData.Del("i.am.that.what")
	exists = jsonData.Has("i.am.that.what")
	assert.Equal(t, exists, false)
}

func Test_Set_Has_Get_Del_W_List(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Test key exists
	exists := jsonData.Has("that.is.a.list")
	assert.Equal(t, exists, false)

	// Set key-value
	jsonData.Set("that.is.a.list", []interface{}{0, 1, 2, 3, 4})
	exists = jsonData.Has("that.is.a.list")
	assert.Equal(t, exists, true)

	// Test N key exists
	exists = jsonData.Has("that.is.a.list.3")
	assert.Equal(t, exists, true)

	// Test N key not exists
	exists = jsonData.Has("that.is.a.list.666")
	assert.Equal(t, exists, false)

	// Test set dict in list
	jsonData.Set("that.is.a.dict.in.list", []interface{}{map[string]interface{}{"a": 1, "b": 2, "c": 3}})
	exists = jsonData.Has("that.is.a.dict.in.list")
	assert.Equal(t, exists, true)

	// Test dict in list exists
	exists = jsonData.Has("that.is.a.dict.in.list.0")
	assert.Equal(t, exists, true)
	exists = jsonData.Has("that.is.a.dict.in.list.0.a")
	assert.Equal(t, exists, true)
	exists = jsonData.Has("that.is.a.dict.in.list.1.a")
	assert.Equal(t, exists, false)
	exists = jsonData.Has("that.is.a.dict.in.list.0.z")
	assert.Equal(t, exists, false)

	// Test get dict in list
	intData, err := jsonData.Get("that.is.a.dict.in.list.0.b").Int()
	assert.Equal(t, err, nil)
	assert.Equal(t, intData, 2)
	_, err = jsonData.Get("that.is.a.dict.in.list.1.b").Int()
	assert.NotEqual(t, err, nil)
	_, err = jsonData.Get("that.is.a.dict.in.list.0.z").Int()
	assert.NotEqual(t, err, nil)

	// Get the list value
	rNumber, err := jsonData.Get("that.is.a.list.3").Int()
	assert.Equal(t, err, nil)
	assert.Equal(t, rNumber, 3)

	// Get not-exists N
	_, err = jsonData.Get("that.is.a.list.666").Int()
	assert.NotEqual(t, err, nil)

	// Get the list value with origin way
	rNumber, err = jsonData.Get("that").Get("is").Get("a").Get("list").Get("4").Int()
	assert.Equal(t, err, nil)
	assert.Equal(t, rNumber, 4)

	// Get the list value with GetN
	rNumber, err = jsonData.Get("that.is.a.list").GetN(1).Int()
	assert.Equal(t, err, nil)
	assert.Equal(t, rNumber, 1)

	// Get the list value with origin way GetN
	rNumber, err = jsonData.Get("that").Get("is").Get("a").Get("list").GetN(2).Int()
	assert.Equal(t, err, nil)
	assert.Equal(t, rNumber, 2)

	// Get not-exists N with origin way GetN
	_, err = jsonData.Get("that").Get("is").Get("a").Get("list").GetN(666).Int()
	assert.NotEqual(t, err, nil)
}

func Test_Set_Has_Get_Del_Type(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Set bool value
	jsonData.Set("bool", true)
	boolData, err := jsonData.Get("bool").Bool()
	assert.Equal(t, err, nil)
	assert.Equal(t, boolData, true)
	assert.Equal(t, "bool", fmt.Sprintf("%T", boolData))
	assert.Equal(t, jsonData.Has("bool"), true)
	jsonData.Del("bool")
	assert.Equal(t, jsonData.Has("bool"), false)

	// Set string value
	jsonData.Set("string", "string")
	stringData, err := jsonData.Get("string").String()
	assert.Equal(t, err, nil)
	assert.Equal(t, stringData, "string")
	assert.Equal(t, "string", fmt.Sprintf("%T", stringData))
	assert.Equal(t, jsonData.Has("string"), true)
	jsonData.Del("string")
	assert.Equal(t, jsonData.Has("string"), false)

	// Set float64 value
	jsonData.Set("float64", float64(999))
	float64Data, err := jsonData.Get("float64").Float64()
	assert.Equal(t, err, nil)
	assert.Equal(t, float64Data, float64(999))
	assert.Equal(t, "float64", fmt.Sprintf("%T", float64Data))
	assert.Equal(t, jsonData.Has("float64"), true)
	jsonData.Del("float64")
	assert.Equal(t, jsonData.Has("float64"), false)

	// Set int value
	jsonData.Set("int", int(666))
	intData, err := jsonData.Get("int").Int()
	assert.Equal(t, err, nil)
	assert.Equal(t, intData, int(666))
	assert.Equal(t, "int", fmt.Sprintf("%T", intData))
	assert.Equal(t, jsonData.Has("int"), true)
	jsonData.Del("int")
	assert.Equal(t, jsonData.Has("int"), false)

	// Set int64 value
	jsonData.Set("int64", int64(666))
	int64Data, err := jsonData.Get("int64").Int64()
	assert.Equal(t, err, nil)
	assert.Equal(t, int64Data, int64(666))
	assert.Equal(t, "int64", fmt.Sprintf("%T", int64Data))
	assert.Equal(t, jsonData.Has("int64"), true)
	jsonData.Del("int64")
	assert.Equal(t, jsonData.Has("int64"), false)

	// Set uint64 value
	jsonData.Set("uint64", uint64(666))
	uint64Data, err := jsonData.Get("uint64").Uint64()
	assert.Equal(t, err, nil)
	assert.Equal(t, uint64Data, uint64(666))
	assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64Data))
	assert.Equal(t, jsonData.Has("uint64"), true)
	jsonData.Del("uint64")
	assert.Equal(t, jsonData.Has("uint64"), false)

	// Set string array value
	jsonData.Set("string_array", []interface{}{"a", "b", "c"})
	stringArrayData, err := jsonData.Get("string_array").StringArray()
	assert.Equal(t, err, nil)
	assert.Equal(t, stringArrayData, []string{"a", "b", "c"})
	assert.Equal(t, "[]string", fmt.Sprintf("%T", stringArrayData))
	assert.Equal(t, jsonData.Has("string_array"), true)
	jsonData.Del("string_array")
	assert.Equal(t, jsonData.Has("string_array"), false)
}

func Test_Get_Assert_Data(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Get data as map
	mapData, err := jsonData.Get("status").Map()
	assert.Equal(t, err, nil)
	assert.Equal(t, "map[string]interface {}", fmt.Sprintf("%T", mapData))

	// Get data as array
	ArrayData, err := jsonData.Get("result").Get("intlist").Array()
	assert.Equal(t, err, nil)
	assert.Equal(t, "[]interface {}", fmt.Sprintf("%T", ArrayData))
	for k, v := range ArrayData {
		r, _ := v.(json.Number).Int64()
		assert.Equal(t, k, int(r))
	}

	// Get data as bool
	boolData, err := jsonData.Get("result").Get("online").Bool()
	assert.Equal(t, err, nil)
	assert.Equal(t, "bool", fmt.Sprintf("%T", boolData))
	assert.Equal(t, boolData, true)

	// Get data as string
	stringData, err := jsonData.Get("status").Get("message").String()
	assert.Equal(t, err, nil)
	assert.Equal(t, "string", fmt.Sprintf("%T", stringData))
	assert.Equal(t, stringData, "success")

	// Get data as float64
	float64Data, err := jsonData.Get("result").Get("rate").Float64()
	assert.Equal(t, err, nil)
	assert.Equal(t, "float64", fmt.Sprintf("%T", float64Data))
	assert.Equal(t, float64Data, float64(0.8))

	// Get data as int
	intData, err := jsonData.Get("status").Get("code").Int()
	assert.Equal(t, err, nil)
	assert.Equal(t, "int", fmt.Sprintf("%T", intData))
	assert.Equal(t, intData, int(1))

	// Get data as int64
	int64Data, err := jsonData.Get("status").Get("code").Int64()
	assert.Equal(t, err, nil)
	assert.Equal(t, "int64", fmt.Sprintf("%T", int64Data))
	assert.Equal(t, int64Data, int64(1))

	// Get data as uint64
	uint64Data, err := jsonData.Get("status").Get("code").Uint64()
	assert.Equal(t, err, nil)
	assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64Data))
	assert.Equal(t, uint64Data, uint64(1))

	// Get data as string array
	jsonData.Set("that.is.a.list", []interface{}{"a", "b", "c", "d", "e"})
	stringArrayData, err := jsonData.Get("that.is.a.list").StringArray()
	assert.Equal(t, err, nil)
	assert.Equal(t, "[]string", fmt.Sprintf("%T", stringArrayData))
	assert.Equal(t, stringArrayData, []string{"a", "b", "c", "d", "e"})
}

func Test_Get_Must_Assert_Data(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Get data as bool
	boolData := jsonData.Get("result").Get("online").MustBool()
	assert.Equal(t, "bool", fmt.Sprintf("%T", boolData))
	assert.Equal(t, boolData, true)

	// Get data as string
	stringData := jsonData.Get("status").Get("message").MustString()
	assert.Equal(t, "string", fmt.Sprintf("%T", stringData))
	assert.Equal(t, stringData, "success")

	// Get data as float64
	float64Data := jsonData.Get("result").Get("rate").MustFloat64()
	assert.Equal(t, "float64", fmt.Sprintf("%T", float64Data))
	assert.Equal(t, float64Data, float64(0.8))

	// Get data as int
	intData := jsonData.Get("status").Get("code").MustInt()
	assert.Equal(t, "int", fmt.Sprintf("%T", intData))
	assert.Equal(t, intData, int(1))

	// Get data as int64
	int64Data := jsonData.Get("status").Get("code").MustInt64()
	assert.Equal(t, "int64", fmt.Sprintf("%T", int64Data))
	assert.Equal(t, int64Data, int64(1))

	// Get data as uint64
	uint64Data := jsonData.Get("status").Get("code").MustUint64()
	assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64Data))
	assert.Equal(t, uint64Data, uint64(1))

	// Get data as string array
	jsonData.Set("that.is.a.list", []interface{}{"a", "b", "c", "d", "e"})
	stringArrayData := jsonData.Get("that.is.a.list").MustStringArray()
	assert.Equal(t, "[]string", fmt.Sprintf("%T", stringArrayData))
	assert.Equal(t, stringArrayData, []string{"a", "b", "c", "d", "e"})
}

func Test_Get_Must_Assert_Data_N_Default(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Get data as bool
	testMustPanic(t, func() {
		jsonData.Get("not-exists").MustBool()
	})

	// Get data as string
	testMustPanic(t, func() {
		jsonData.Get("not-exists").MustString()
	})

	// Get data as float64
	testMustPanic(t, func() {
		jsonData.Get("not-exists").MustFloat64()
	})

	// Get data as int
	testMustPanic(t, func() {
		jsonData.Get("not-exists").MustInt()
	})

	// Get data as int64
	testMustPanic(t, func() {
		jsonData.Get("not-exists").MustInt64()
	})

	// Get data as uint64
	testMustPanic(t, func() {
		jsonData.Get("not-exists").MustUint64()
	})

	// Get data as string array
	testMustPanic(t, func() {
		jsonData.Get("not-exists").MustStringArray()
	})
}

func Test_Get_Must_Assert_Data_W_Default(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Get data as bool
	boolData := jsonData.Get("not-exists").MustBool(true)
	assert.Equal(t, "bool", fmt.Sprintf("%T", boolData))
	assert.Equal(t, boolData, true)

	// Get data as string
	stringData := jsonData.Get("not-exists").MustString("ok")
	assert.Equal(t, "string", fmt.Sprintf("%T", stringData))
	assert.Equal(t, stringData, "ok")

	// Get data as float64
	float64Data := jsonData.Get("not-exists").MustFloat64(float64(999))
	assert.Equal(t, "float64", fmt.Sprintf("%T", float64Data))
	assert.Equal(t, float64Data, float64(999))

	// Get data as int
	intData := jsonData.Get("not-exists").MustInt(int(666))
	assert.Equal(t, "int", fmt.Sprintf("%T", intData))
	assert.Equal(t, intData, int(666))

	// Get data as int64
	int64Data := jsonData.Get("not-exists").MustInt64(int64(666))
	assert.Equal(t, "int64", fmt.Sprintf("%T", int64Data))
	assert.Equal(t, int64Data, int64(666))

	// Get data as uint64
	uint64Data := jsonData.Get("not-exists").MustUint64(uint64(666))
	assert.Equal(t, "uint64", fmt.Sprintf("%T", uint64Data))
	assert.Equal(t, uint64Data, uint64(666))

	// Get data as string array
	stringArrayData := jsonData.Get("not-exists").MustStringArray([]string{"i", "am", "ok"})
	assert.Equal(t, "[]string", fmt.Sprintf("%T", stringArrayData))
	assert.Equal(t, stringArrayData, []string{"i", "am", "ok"})
}

func Test_HTML_Escape(t *testing.T) {
	// Init json and set html
	jsonData := New()
	jsonData.Set("param", "a=1&b=2&c=3")
	jsonData.Set("title", "<title>test escape</title>")

	// dumps not escaped html
	jsonText, err := jsonData.Dumps()
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonText, `{"param":"a=1&b=2&c=3","title":"<title>test escape</title>"}`)

	// dumps escaped html
	jsonData.SetHtmlEscape(true)
	jsonText, err = jsonData.Dumps()
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonText, `{"param":"a=1\u0026b=2\u0026c=3","title":"\u003ctitle\u003etest escape\u003c/title\u003e"}`)
}

func Test_isMap(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Test the top level json
	isMap := jsonData.IsMap()
	assert.Equal(t, isMap, true)

	// Test after get
	isMap = jsonData.Get("status").IsMap()
	assert.Equal(t, isMap, true)

	// Test after twice get
	isMap = jsonData.Get("result").Get("online").IsMap()
	assert.Equal(t, isMap, false)

	// Test after magic get
	isMap = jsonData.Get("result.online").IsMap()
	assert.Equal(t, isMap, false)

	// Test the array
	isMap = jsonData.Get("result.intlist").IsMap()
	assert.Equal(t, isMap, false)

	// Test not exists key
	isMap = jsonData.Get("result.not-exists").IsMap()
	assert.Equal(t, isMap, false)
}

func Test_Is_Array(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Test the top level json
	isMap := jsonData.IsArray()
	assert.Equal(t, isMap, false)

	// Test after get
	isMap = jsonData.Get("status").IsArray()
	assert.Equal(t, isMap, false)

	// Test after twice get
	isMap = jsonData.Get("result").Get("intlist").IsArray()
	assert.Equal(t, isMap, true)

	// Test the array
	isMap = jsonData.Get("result.intlist").IsArray()
	assert.Equal(t, isMap, true)

	// Test the array element
	isMap = jsonData.Get("result.intlist.0").IsArray()
	assert.Equal(t, isMap, false)

	// Test not exists key
	isMap = jsonData.Get("result.not-exists").IsArray()
	assert.Equal(t, isMap, false)
}

func Test_Len(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Get len of top level json
	n := jsonData.Len()
	assert.Equal(t, n, 2)

	// Get len of map
	n = jsonData.Get("status").Len()
	assert.Equal(t, n, 2)

	// Get len of not exists map
	n = jsonData.Get("status.not-exists").Len()
	assert.Equal(t, n, -1)

	// Get len of int
	n = jsonData.Get("status.code").Len()
	assert.Equal(t, n, -1)

	// Get len of string
	n = jsonData.Get("status.message").Len()
	assert.Equal(t, n, 7)

	// Get len of not exists string
	n = jsonData.Get("status.message.not-exists").Len()
	assert.Equal(t, n, -1)

	// Get len of array
	n = jsonData.Get("result.intlist").Len()
	assert.Equal(t, n, 5)

	// Get len of not exists array
	n = jsonData.Get("result.intlist.not-exists").Len()
	assert.Equal(t, n, -1)
}

func Test_Time_Assert_Data(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Time for comparing
	testTime, err := time.Parse(time.RFC3339, "2019-01-31T12:11:10+08:00")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, testTime.Unix(), int64(0))

	// Test get rfc3339 time
	jsonData.Set("time", "2019-01-31T12:11:10+08:00")
	timeData, err := jsonData.Get("time").Time()
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)

	// Test get format time
	jsonData.Set("time", "2019-01-31 12:11:10")
	timeData, err = jsonData.Get("time").Time("2006-01-02 15:04:05")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, timeData.Unix(), int64(0))

	// Test get rfc3339 time with not exists key
	timeData, err = jsonData.Get("not-exists").Time()
	assert.NotEqual(t, err, nil)
	assert.Equal(t, timeData.Unix(), int64(-62135596800))

	// Test get time from int
	testTime = time.Unix(1548907870, 0)
	jsonData.Set("time", int(1548907870))
	timeData, err = jsonData.Get("time").Time()
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)

	// Test get time from int64
	jsonData.Set("time", int64(1548907870))
	timeData, err = jsonData.Get("time").Time()
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)

	// Test get time from uint64
	jsonData.Set("time", uint64(1548907870))
	timeData, err = jsonData.Get("time").Time()
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)

	// Test get time from float64
	jsonData.Set("time", float64(1548907870))
	timeData, err = jsonData.Get("time").Time()
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)

	// Date for comparing
	testTime, err = time.ParseInLocation("2006-01-02", "2019-01-31", time.Local)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, testTime.Unix(), int64(0))

	// Test get format date
	jsonData.Set("time", "2019-01-31")
	timeData, err = jsonData.Get("time").Time("2006-01-02")
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)
}

func Test_Time_Must_Assert_Data(t *testing.T) {
	// Loads json for Set
	jsonData, err := Loads(textResult)
	assert.Equal(t, err, nil)

	// Time for comparing
	testTime, err := time.Parse(time.RFC3339, "2019-01-31T12:11:10+08:00")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, testTime.Unix(), int64(0))

	// Test get rfc3339 time
	jsonData.Set("time", "2019-01-31T12:11:10+08:00")
	timeData := jsonData.Get("time").MustTime()
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)

	// Test get format time
	jsonData.Set("time", "2019-01-31 12:11:10")
	timeData = jsonData.Get("time").MustTime("2006-01-02 15:04:05")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, timeData.Unix(), int64(0))

	// No format, no default
	jsonData.Set("time", "i-am-not-the-time")
	testMustPanic(t, func() {
		jsonData.Get("time").MustTime()
	})

	// Has format, no default
	testMustPanic(t, func() {
		jsonData.Get("time").MustTime("2006-01-02 15:04:05")
	})

	// No format, has default
	testTime = time.Unix(1548907870, 0)
	timeData = jsonData.Get("time").MustTime(time.Unix(1548907870, 0))
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)

	// Has format, has default
	timeData = jsonData.Get("time").MustTime("2006-01-02 15:04:05", time.Unix(1548907870, 0))
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)

	// No format, no default, key not exists
	testMustPanic(t, func() {
		jsonData.Get("not-exists").MustTime()
	})

	// No format, has default, key not exists
	timeData = jsonData.Get("not-exists").MustTime(time.Unix(1548907870, 0))
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)

	// Test must get time from int
	jsonData.Set("time", int(1548907870))
	timeData = jsonData.Get("time").MustTime()
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)

	// Get from int, No default
	jsonData.Set("time", true)
	testMustPanic(t, func() {
		jsonData.Get("time").MustTime()
	})

	// Get from int, Has default
	timeData = jsonData.Get("time").MustTime(time.Unix(1548907870, 0))
	assert.Equal(t, err, nil)
	assert.Equal(t, timeData, testTime)
}
