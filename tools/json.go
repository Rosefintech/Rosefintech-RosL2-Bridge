/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2020/11/16 16:57
*/

package tools

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var (
	ErrNil    = errors.New("error key not exist")
	ErrFormat = errors.New("error formart")
)

type JsonObject struct {
	data map[string]interface{}
}

// NewJsonObject
func NewJsonObject(bytes []byte) (*JsonObject, error) {
	jsonObject := new(JsonObject)
	err := json.Unmarshal(bytes, &jsonObject.data)
	if err != nil {
		return nil, err
	}
	return jsonObject, nil
}

func getArrMapInterface(m interface{}) (map[string]interface{}, bool) {
	temp, ok := m.([]interface{})
	if !ok {
		return nil, false
	}
	if len(temp) == 0 {
		return nil, false
	}
	result, ok := temp[0].(map[string]interface{})
	if !ok {
		return nil, false
	}
	return result, ok
}

func (o *JsonObject) getValue(key string) (interface{}, error) {
	m := o.data
	keys := strings.Split(key, ".")

	var (
		value interface{}
		ok    = false
	)

	for i := 0; i < len(keys)-1; i++ {
		value, ok = m[keys[i]]
		if !ok {
			return nil, ErrNil
		}
		m, ok = value.(map[string]interface{})
		if !ok {
			m, ok = getArrMapInterface(value)
			if !ok {
				return nil, ErrFormat
			}
		}
	}

	value, ok = m[keys[len(keys)-1]]
	if !ok {
		return nil, ErrNil
	}
	return value, nil
}

// GetBool
func (o *JsonObject) GetBool(key string) (bool, error) {
	value, err := o.getValue(key)
	if err != nil {
		return false, err
	}

	result, ok := value.(bool)
	if !ok {
		return false, ErrFormat
	}
	return result, nil
}

// GetFloat64
func (o *JsonObject) GetFloat64(key string) (float64, error) {
	value, err := o.getValue(key)
	if err != nil {
		return 0, err
	}

	result, ok := value.(float64)
	if !ok {
		return 0, ErrFormat
	}
	return result, nil
}

func (o *JsonObject) GetTime(key string) (time.Time, error) {
	value, err := o.getValue(key)
	if err != nil {
		return time.Time{}, err
	}

	result, ok := value.(time.Time)
	if !ok {
		return time.Time{}, errors.New("there's no such type time")
	}
	return result, nil
}

func (o *JsonObject) GetInt64(key string) (int64, error) {
	value, err := o.getValue(key)
	if err != nil {
		return 0, err
	}
	result, ok := value.(int64)
	if !ok {
		return 0, ErrFormat
	}
	return result, nil
}

// GetString
func (o *JsonObject) GetString(key string) (string, error) {
	value, err := o.getValue(key)
	if err != nil {
		return "", err
	}

	result, ok := value.(string)
	if !ok {
		return "", ErrFormat
	}
	return result, nil
}

