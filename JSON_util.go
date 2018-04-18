package describe

import "encoding/json"

// JSONBytes get json bytes from interface{}, ignore marshal error as empty []byte
func JSONBytes(v interface{}) []byte {
	b, err := json.Marshal(v)
	if IsErr(err) {
		return []byte{}
	}
	return b
}

// JSONString get json string from interface{}, ignore marshal error as empty string
func JSONString(v interface{}) string {
	return string(JSONBytes(v))
}
