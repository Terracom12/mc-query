package fields

import (
	"bytes"
	"reflect"
	"testing"
)

type testCaseType struct {
	field  Field
	binary []byte
}

func testToBytes(t *testing.T, tests []testCaseType) {
	for _, tt := range tests {
		b, err := tt.field.ToBytes()

		var typeName = reflect.TypeOf(tt.field).String()

		if err != nil {
			t.Errorf("%s.toBytes(%v) returned error: %v", typeName, tt.field, err)
			continue
		}
		if !bytes.Equal(b, tt.binary) {
			t.Errorf("%s.toBytes(%v) = %v; want %v", typeName, tt.field, b, tt.binary)
		}
	}
}

func testFromBytes(t *testing.T, tests []testCaseType) {
	for _, tt := range tests {
		res := tt.field
		_, err := res.FromBytes(tt.binary)

		var typeName = reflect.TypeOf(tt.field).String()

		if err != nil {
			t.Errorf("%s.fromBytes(%v) returned error: %v", typeName, tt.binary, err)
			continue
		}
		if res != tt.field {
			t.Errorf("%s.fromBytes(%v) = %v; want %v", typeName, tt.binary, res, tt.field)
		}
	}
}
