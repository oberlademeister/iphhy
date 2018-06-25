package iphhy

import (
	"encoding/json"
)

// UnmarshalText satisfies encoding.TextUnmarshaler
func (i4 *I4) UnmarshalText(b []byte) error {
	err := i4.FromString(string(b))
	return err
}

// UnmarshalJSON satisfies json.UnmarshalJSON
func (i4 *I4) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	err = i4.FromString(s)
	return err
}
