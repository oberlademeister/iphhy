package iphhy

import (
	"encoding/json"
)

// UnmarshalText satisfies encoding.TextUnmarshaler
func (i4 *I4) UnmarshalText(b []byte) error {
	err := i4.FromString(string(b))
	return err
}

// UnmarshalJSON is used to unmarshal json
func (i4 *I4) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i4.FromString(s)
	return nil
}

// MarshalJSON used to marshal json
func (i4 I4) MarshalJSON() ([]byte, error) {
	return json.Marshal(i4.CIDRString())
}
