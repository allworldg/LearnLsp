package rpc_test

import (
	"golsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expect := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actualContent := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expect != actualContent {
		t.Fatalf("expect is %s, actual is %s", expect, actualContent)
	}
}
func TestDecode(t *testing.T) {
	_, _, error := rpc.DecodeMessage([]byte("Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"))
	if error != nil {
		t.Log(error)
		t.Fatalf("DecodeMessage is wrong")
	}
}
