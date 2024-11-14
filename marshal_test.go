package iphhy

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/y0ssar1an/q"
)

func TestUnmarshal(t *testing.T) {
	const tomlString = `
[ips]
i1 = "192.168.0.10/24"
i2 = "192.168.9.0/24"
`

	const jsonString = `
{
	"IPs": {
	  "i1": "192.168.0.10/24",
	  "i2": "192.168.9.0/24",
	  "i3": "abcd::1111/128",
	  "i4": "abcd:ef01:2345:6789:abcd:ed01:2345:6789/128"
	}
  }
`
	type TestType2 struct {
		I1 IP
		I2 IP
		I3 IP
		I4 IP
	}
	type TestType1 struct {
		IPs TestType2
	}

	tp := TestType1{}

	err := json.Unmarshal([]byte(jsonString), &tp)
	if err != nil {
		t.Errorf("json decode threw error: %v", err)
	}
	q.Q(tp)

	_, err = toml.Decode(tomlString, &tp)
	if err != nil {
		t.Errorf("toml decode threw error: %v", err)
	}

	q.Q(tp)
}

func TestJSON(t *testing.T) {
	type JSONTestType struct {
		I *IP
	}

	type Test struct {
		in      *IP
		json    string
		wantErr bool
	}

	tests := []Test{
		{
			in:      Parse("0.0.0.0/0"),
			json:    `{"I":"0.0.0.0/0"}`,
			wantErr: false,
		},
	}
	for i, tt := range tests {
		// JSONTestType to json
		jtt := JSONTestType{tt.in}
		b, err := json.Marshal(jtt)
		if (err != nil) != tt.wantErr {
			t.Errorf("TestJSON#Marshal #%d wantErr = %t", i, tt.wantErr)
		}
		if bytes.Compare(b, []byte(tt.json)) != 0 {
			t.Errorf("TestJSON#Marshal #%d want = %s got = %s", i, tt.json, string(b))
		}
	}

}
