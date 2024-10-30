package iphhy

import (
	"encoding/json"
)

// UnmarshalText satisfies encoding.TextUnmarshaler
func (ip *IP) UnmarshalText(b []byte) error {
	x := Parse(string(b))
	ip.ip = x.ip
	ip.mask = x.mask
	return nil
}

// UnmarshalJSON is used to unmarshal json
func (ip *IP) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	x := Parse(s)
	ip.ip = x.ip
	ip.mask = x.mask
	return nil
}

// MarshalJSON used to marshal json
func (ip *IP) MarshalJSON() ([]byte, error) {
	return json.Marshal(ip.String())
}
