package simplejson

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
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

	fd, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
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

func (j *Json) Exists(key string) bool {
	result, err := j.Map()
	if err == nil {
		_, exists := result[key]
		return exists
	}

	return false
}

func (j *Json) Set(key, value string) {
	result, err := j.Map()
	if err == nil {
		result[key] = value
	}
}

func (j *Json) Map() (result map[string]interface{}, err error) {
	result, ok := (j.data).(map[string]interface{})
	if !ok {
		err = errors.New("assert to map failed")
	}
	return
}

func (j *Json) Array() (result []interface{}, err error) {
	result, ok := (j.data).([]interface{})
	if !ok {
		err = errors.New("assert to array failed")
	}
	return
}

func (j *Json) Bool() (result bool, err error) {
	result, ok := (j.data).(bool)
	if !ok {
		err = errors.New("assert to bool failed")
	}
	return
}

func (j *Json) String() (result string, err error) {
	result, ok := (j.data).(string)
	if !ok {
		err = errors.New("assert to string failed")
	}
	return
}

func (j *Json) Int() (result int, err error) {
	f, ok := (j.data).(float64)
	result = int(f)
	if !ok {
		err = errors.New("assert to int failed")
	}
	return
}

func (j *Json) Float() (result float64, err error) {
	result, ok := (j.data).(float64)
	if !ok {
		err = errors.New("assert to float64 failed")
	}
	return
}
