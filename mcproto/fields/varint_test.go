package fields

import (
	"testing"
)

// Examples from: https://minecraft.wiki/w/Java_Edition_protocol/Packets#VarInt_and_VarLong
var intTests = []testCaseType{
	{ptr(VarInt(0)), []byte{0x00}},
	{ptr(VarInt(1)), []byte{0x01}},
	{ptr(VarInt(2)), []byte{0x02}},
	{ptr(VarInt(127)), []byte{0x7f}},
	{ptr(VarInt(128)), []byte{0x80, 0x01}},
	{ptr(VarInt(255)), []byte{0xff, 0x01}},
	{ptr(VarInt(25565)), []byte{0xdd, 0xc7, 0x01}},
	{ptr(VarInt(2097151)), []byte{0xff, 0xff, 0x7f}},
	{ptr(VarInt(2147483647)), []byte{0xff, 0xff, 0xff, 0xff, 0x07}},
	{ptr(VarInt(-1)), []byte{0xff, 0xff, 0xff, 0xff, 0x0f}},
	{ptr(VarInt(-2147483648)), []byte{0x80, 0x80, 0x80, 0x80, 0x08}},
}

var longTests = []testCaseType{
	{ptr(VarLong(0)), []byte{0x00}},
	{ptr(VarLong(1)), []byte{0x01}},
	{ptr(VarLong(2)), []byte{0x02}},
	{ptr(VarLong(127)), []byte{0x7f}},
	{ptr(VarLong(128)), []byte{0x80, 0x01}},
	{ptr(VarLong(255)), []byte{0xff, 0x01}},
	{ptr(VarLong(2147483647)), []byte{0xff, 0xff, 0xff, 0xff, 0x07}},
	{ptr(VarLong(9223372036854775807)), []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}},
	{ptr(VarLong(-1)), []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01}},
	{ptr(VarLong(-2147483648)), []byte{0x80, 0x80, 0x80, 0x80, 0xf8, 0xff, 0xff, 0xff, 0xff, 0x01}},
	{ptr(VarLong(-9223372036854775808)), []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}},
}

func ptr[T any](v T) *T {
	return &v
}

func TestIntToBytes(t *testing.T) {
	testToBytes(t, intTests)
}

func TestIntFromBytes(t *testing.T) {
	testFromBytes(t, intTests)
}

func TestLongToBytes(t *testing.T) {
	testToBytes(t, longTests)
}

func TestLongFromBytes(t *testing.T) {
	testFromBytes(t, longTests)
}
