package iphhy

import (
	"bytes"
	"testing"
)

func TestDebug(t *testing.T) {
	SetDebug(true)
	var buf bytes.Buffer
	SetDebugOut(&buf)
	debugf("test1")
	debugln("test2")
	got := buf.String()
	want := "test1test2\n"
	if got != want {
		t.Errorf("TestDebug failed got %s want %s", got, want)
	}
	SetDebug(false)
	SetDebugOut(nil)
}
