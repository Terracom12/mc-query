package fields

import (
	"testing"
)

var stringTests = []testCaseType{
	{&String{"", 0}, []byte{0x00}},
	{&String{"hello", 5}, []byte("\x05hello")},
	{&String{"abc123456", 10}, []byte("\x09abc123456")},
}

func TestStringToBytes(t *testing.T) {
	testToBytes(t, stringTests)
}

func TestStringFromBytes(t *testing.T) {
	testFromBytes(t, stringTests)
}
