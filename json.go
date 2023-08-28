package toolkit

import "encoding/json"

// Json2String ...
func Json2String(obj interface{}) string {
	data, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return Bytes2String(data)
}

// JsonMarshalZeroToNil ...
func JsonMarshalZeroToNil(obj interface{}) interface{} {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	str := Bytes2String(data)
	if str == "null" || str == "{}" || str == "[]" || str == "" {
		return nil
	}
	return str
}
