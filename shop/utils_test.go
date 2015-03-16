package main

import (
	"runtime"
	"testing"
)

func assertEqual(t *testing.T, expect interface{}, v interface{}) {
	if v != expect {
		_, fname, lineno, ok := runtime.Caller(1)
		if !ok {
			fname, lineno = "<UNKNOWN>", -1
		}
		t.Errorf("FAIL: %s:%d\nExpected: %#v\nReceived: %#v", fname, lineno, expect, v)
	}
}
