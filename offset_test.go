package iph

import "testing"

func TestBaseLast(t *testing.T) {
	tests := []struct {
		input         string
		wantedHostInt uint32
		wantedBase    uint32
		wantedLast    uint32
	}{
		{"192.168.21.18/24", 3232240914, 3232240896, 3232241151},
	}

	for i, tt := range tests {
		aip, _ := NewI4(tt.input)
		gotHostInt := aip.ip
		gotBase := IPToInt(aip.BaseIP())
		gotLast := IPToInt(aip.LastIP())
		switch {
		case gotHostInt != tt.wantedHostInt:
			t.Errorf("TestBaseLast#%d: tt: %v IpToInt failed: %d", i, tt, gotHostInt)
		case gotBase != tt.wantedBase:
			t.Errorf("TestBaseLast#%d: tt: %v BaseIP failed: %d", i, tt, gotBase)
		case gotLast != tt.wantedLast:
			t.Errorf("TestBaseLast#%d: tt: %v LastIP failed: %d", i, tt, gotLast)
		}
	}
}

func TestOffset(t *testing.T) {
	type args struct {
		AIP    string
		Offset int
	}
	tests := []struct {
		args           args
		wanted         string
		wantederrisnil bool
	}{
		{args{"192.168.0.1/24", 0}, "192.168.0.0/24", true},
	}

	for i, tt := range tests {
		aip := MustNewI4(tt.args.AIP)
		resultAip, err := aip.Offset(tt.args.Offset)
		errisnil := err == nil
		if errisnil != tt.wantederrisnil {
			t.Errorf("TestOffset#%d: wrong error conditon args:%v ", i, tt.args)
			continue
		}
		if errisnil && resultAip.String() != tt.wanted {
			t.Errorf("TestOffset#%d: wrong result args:%v got: %s wanted:%s ", i, tt.args, resultAip.String(), tt.wanted)
			continue
		}
	}
}
