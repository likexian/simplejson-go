package simplejson

import (
    "testing"
    "github.com/bmizerany/assert"
)

func TestSimplejson(t *testing.T) {
    data := `{"result":{"list":[0,1,2,3,4],"online":true,"rate":0.8},"status":{"code":1,"message":"success"}}`
    json, err := Loads(data)
    assert.NotEqual(t, nil, json)
    assert.Equal(t, nil, err)
    
    code, _ := json.Get("status").Get("code").Int()
    assert.Equal(t, 1, code)
    message, _ := json.Get("status").Get("message").String()
    assert.Equal(t, "success", message)
    
    online, _ := json.Get("result").Get("online").Bool()
    assert.Equal(t, true, online)
    rate, _ := json.Get("result").Get("rate").Float()
    assert.Equal(t, 0.80, rate)
    
    list, _ := json.Get("result").Get("list").Array()
    assert.NotEqual(t, nil, list)
    for k, v := range list {
        assert.Equal(t, k, int(v.(float64)))
    }
    
    result, err := Dumps(json)
    assert.Equal(t, nil, err)
    assert.Equal(t, data, result)
    
    json.Set("name", "Li Kexian")
    name, _ := json.Get("name").String()
    assert.Equal(t, "Li Kexian", name)
    
    bytes := Dump("simplejson.json", json)
    assert.NotEqual(t, 0, bytes)
    
    njson, err := Load("simplejson.json")
    assert.Equal(t, json, njson)
    assert.Equal(t, nil, err)
}

