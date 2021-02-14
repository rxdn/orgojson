package orgojson

import "encoding/json"

type JsonArray struct {
	internal []interface{}
}

func BlankArray() JsonArray {
	return JsonArray{
		internal: make([]interface{}, 0),
	}
}

func ArrayFromSlice(s []interface{}) JsonArray {
	return JsonArray{s}
}

func (a *JsonArray) Clear() {
	a.internal = make([]interface{}, 0)
}

func (a *JsonArray) Length() int {
	return len(a.internal)
}

func (a *JsonArray) Marshal() ([]byte, error) {
	return json.Marshal(a.internal)
}

func (a *JsonArray) Put(value interface{}) {
	a.internal = append(a.internal, value)
}

func (a *JsonArray) Del(index int) {
	a.internal = append(a.internal[:index], a.internal[index+1:]...)
}

func (a *JsonArray) ToSlice() (dst []interface{}) {
	copy(dst, a.internal)
	return
}

func (a *JsonArray) Get(index int) (interface{}, error) {
	if index >= a.Length() {
		return nil, ErrNotFound
	}
	
	value := a.internal[index]
	return value, nil
}

func (a *JsonArray) GetBool(index int) (bool, error) {
	value, err := a.Get(index)
	if err != nil {
		return false, err
	}

	if value, ok := value.(bool); ok {
		return value, nil
	} else {
		return false, ErrWrongType
	}
}

func (a *JsonArray) GetUint8(index int) (uint8, error) {
	value, err := a.Get(index)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(uint8); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (a *JsonArray) GetUint16(index int) (uint16, error) {
	value, err := a.Get(index)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(uint16); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (a *JsonArray) GetUint32(index int) (uint32, error) {
	value, err := a.Get(index)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(uint32); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (a *JsonArray) GetUint64(index int) (uint64, error) {
	value, err := a.Get(index)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(uint64); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (a *JsonArray) GetInt8(index int) (int8, error) {
	value, err := a.Get(index)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(int8); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (a *JsonArray) GetInt16(index int) (int16, error) {
	value, err := a.Get(index)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(int16); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (a *JsonArray) GetInt32(index int) (int32, error) {
	value, err := a.Get(index)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(int32); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (a *JsonArray) GetInt64(index int) (int64, error) {
	value, err := a.Get(index)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(int64); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (a *JsonArray) GetFloat32(index int) (float32, error) {
	value, err := a.Get(index)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(float32); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (a *JsonArray) GetFloat64(index int) (float64, error) {
	value, err := a.Get(index)
	if err != nil {
		return 0, err
	}

	if value, ok := value.(float64); ok {
		return value, nil
	} else {
		return 0, ErrWrongType
	}
}

func (a *JsonArray) GetString(index int) (string, error) {
	value, err := a.Get(index)
	if err != nil {
		return "", err
	}

	if value, ok := value.(string); ok {
		return value, nil
	} else {
		return "", ErrWrongType
	}
}

func (a *JsonArray) GetJsonObject(index int) (JsonObject, error) {
	value, err := a.Get(index)
	if err != nil {
		return Blank(), err
	}

	if value, ok := value.(map[string]interface{}); ok {
		return FromMap(value), nil
	} else {
		return Blank(), ErrWrongType
	}
}

func (a *JsonArray) GetJsonArray(index int) (JsonArray, error) {
	value, err := a.Get(index)
	if err != nil {
		return BlankArray(), err
	}

	if value, ok := value.([]interface{}); ok {
		return ArrayFromSlice(value), nil
	} else {
		return BlankArray(), ErrWrongType
	}
}

func (a JsonArray) MarshalJson() ([]byte, error) {
	return json.Marshal(a.internal)
}

func (a *JsonArray) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &a.internal)
}
