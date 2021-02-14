package orgojson

import (
	"encoding/json"
)

type JsonObject struct {
	internal map[string]interface{}
}

func Blank() JsonObject {
	return JsonObject{
		internal: make(map[string]interface{}),
	}
}

func FromBytes(bytes []byte) (o JsonObject, err error) {
	err = json.Unmarshal(bytes, &o.internal)
	return
}

func FromMap(m map[string]interface{}) JsonObject {
	return JsonObject{m}
}

func (o *JsonObject) Clear() {
	o.internal = make(map[string]interface{})
}

func (o *JsonObject) Length() int {
	return len(o.internal)
}

func (o *JsonObject) Keys() []string {
	keys := make([]string, len(o.internal))

	i := 0
	for k := range o.internal {
		keys[i] = k
		i++
	}

	return keys
}

func (o *JsonObject) Marshal() ([]byte, error) {
	return json.Marshal(o.internal)
}

func (o *JsonObject) Put(key string, value interface{}) {
	o.internal[key] = value
}

func (o *JsonObject) Del(key string) {
	delete(o.internal, key)
}

func (o *JsonObject) Get(key string) (interface{}, error) {
	value, ok := o.internal[key]
	if !ok {
		return nil, ErrNotFound
	}

	return value, nil
}

func (o *JsonObject) GetBool(key string) (bool, error) {
	value, err := o.Get(key)
	if err != nil {
		return false, err
	}

	if value, ok := value.(bool); ok {
		return value, nil
	} else {
		return false, ErrWrongType
	}
}

func (o *JsonObject) GetUint8(key string) (uint8, error) {
	value, err := o.Get(key)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(uint8); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (o *JsonObject) GetUint16(key string) (uint16, error) {
	value, err := o.Get(key)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(uint16); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (o *JsonObject) GetUint32(key string) (uint32, error) {
	value, err := o.Get(key)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(uint32); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (o *JsonObject) GetUint64(key string) (uint64, error) {
	value, err := o.Get(key)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(uint64); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (o *JsonObject) GetInt8(key string) (int8, error) {
	value, err := o.Get(key)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(int8); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (o *JsonObject) GetInt16(key string) (int16, error) {
	value, err := o.Get(key)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(int16); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (o *JsonObject) GetInt32(key string) (int32, error) {
	value, err := o.Get(key)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(int32); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (o *JsonObject) GetInt64(key string) (int64, error) {
	value, err := o.Get(key)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(int64); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (o *JsonObject) GetFloat32(key string) (float32, error) {
	value, err := o.Get(key)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(float32); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (o *JsonObject) GetFloat64(key string) (float64, error) {
	value, err := o.Get(key)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(float64); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (o *JsonObject) GetString(key string) (string, error) {
	value, err := o.Get(key)
	if err != nil {
		return "", err
	}

	if value, ok := value.(string); ok {
		return value, nil
	} else {
		return "", ErrWrongType
	}
}

func (o *JsonObject) GetJsonObject(key string) (JsonObject, error) {
	value, err := o.Get(key)
	if err != nil {
		return Blank(), err
	}

	if value, ok := value.(map[string]interface{}); ok {
		return FromMap(value), nil
	} else {
		return Blank(), ErrWrongType
	}
}

func (o *JsonObject) GetJsonArray(key string) (JsonArray, error) {
	value, err := o.Get(key)
	if err != nil {
		return BlankArray(), err
	}

	if value, ok := value.([]interface{}); ok {
		return ArrayFromSlice(value), nil
	} else {
		return BlankArray(), ErrWrongType
	}
}

func (o JsonObject) MarshalJson() ([]byte, error) {
	return json.Marshal(o.internal)
}

func (o *JsonObject) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &o.internal)
}