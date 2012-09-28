package simplejson

import (
    "os"
    "errors"
    "io"
    "io/ioutil"
    "encoding/json"
)

type Json struct {
    data interface{}
}

func Version() string {
    return "0.1"
}

func Author() string {
    return "[Li Kexian](http://www.zhetenga.com/)"
}

func License() string {
    return "Apache"
}

func Load(file string) (*Json, error) {
    result, err := ioutil.ReadFile(file)
    if err != nil {
        return nil, err
    }
    
    return Loads(string(result))
}

func Dump(file string, data *Json) int {
    result, err := Dumps(data)
    if err != nil {
        return 0
    }
    
    fd, err := os.OpenFile(file, os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        return 0
    }
    
    bytes, err := io.WriteString(fd, result)
    if err != nil {
        return 0
    }
    
    return bytes
}

func Loads(data string) (*Json, error) {
    result := new(Json)
    err := json.Unmarshal([]byte(data), &result.data)
    if err != nil {
        return nil, err
    }
    
    return result, nil
}

func Dumps(data *Json) (string, error) {
    result, err := json.Marshal(&data.data)
    if err != nil {
        return "", err
    }
    
    return string(result), nil
}

func (j *Json) Get(key string) *Json {
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

func (j *Json) Map() (map[string]interface{}, error) {
    if result, ok := (j.data).(map[string]interface{}); ok {
        return result, nil
    }
    
    return nil, errors.New("assert to map failed")
}

func (j *Json) Array() ([]interface{}, error) {
    if result, ok := (j.data).([]interface{}); ok {
        return result, nil
    }
    
    return nil, errors.New("assert to array failed")
}

func (j *Json) Bool() (bool, error) {
    if result, ok := (j.data).(bool); ok {
        return result, nil
    }
    
    return false, errors.New("assert to bool failed")
}

func (j *Json) String() (string, error) {
    if result, ok := (j.data).(string); ok {
        return result, nil
    }
    
    return "", errors.New("assert to string failed")
}

func (j *Json) Int() (int, error) {
    if result, ok := (j.data).(float64); ok {
        return int(result), nil
    }
    
    return -1, errors.New("assert to float64 failed")
}

func (j *Json) Float() (float64, error) {
    if result, ok := (j.data).(float64); ok {
        return result, nil
    }
    
    return -1, errors.New("assert to float64 failed")
}

