package iphhy

import (
	"encoding/json"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/y0ssar1an/q"
)

func TestUnmarshal(t *testing.T) {
	const tomlString = `
[ips]
i1 = "192.168.0.10/24"
i2 = "192.168.9.0/24+4"
`

	const jsonString = `
{
	"IPs": {
	  "i1": "192.168.0.10/24",
	  "i2": "192.168.9.0/24+4"
	}
  }
`
	type TestType2 struct {
		I1 I4
		I2 I4
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
