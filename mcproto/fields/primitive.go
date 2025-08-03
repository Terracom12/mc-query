package fields

import (
	"bytes"
	"encoding/binary"
)

var ORDER = binary.BigEndian

type Boolean bool
type Byte int8
type UnsignedByte uint8
type Short int16
type UnsignedShort uint16
type Int int32
type UnsignedInt int32
type Long int64
type Float float32
type Double float64

// ********** toBytes/fromBytes implementations

func (v Boolean) ToBytes() ([]byte, error) {
	return toBytesPrimitiveImpl(v)
}
func (v *Boolean) FromBytes(data []byte) (bytesRead uint, err error) {
	return fromBytesPrimitiveImpl(data, v)
}

func (v Byte) ToBytes() ([]byte, error) {
	return toBytesPrimitiveImpl(v)
}
func (v *Byte) FromBytes(data []byte) (bytesRead uint, err error) {
	return fromBytesPrimitiveImpl(data, v)
}

func (v UnsignedByte) ToBytes() ([]byte, error) {
	return toBytesPrimitiveImpl(v)
}
func (v *UnsignedByte) FromBytes(data []byte) (bytesRead uint, err error) {
	return fromBytesPrimitiveImpl(data, v)
}

func (v Short) ToBytes() ([]byte, error) {
	return toBytesPrimitiveImpl(v)
}
func (v *Short) FromBytes(data []byte) (bytesRead uint, err error) {
	return fromBytesPrimitiveImpl(data, v)
}

func (v UnsignedShort) ToBytes() ([]byte, error) {
	return toBytesPrimitiveImpl(v)
}
func (v *UnsignedShort) FromBytes(data []byte) (bytesRead uint, err error) {
	return fromBytesPrimitiveImpl(data, v)
}

func (v Int) ToBytes() ([]byte, error) {
	return toBytesPrimitiveImpl(v)
}
func (v *Int) FromBytes(data []byte) (bytesRead uint, err error) {
	return fromBytesPrimitiveImpl(data, v)
}

func (v UnsignedInt) ToBytes() ([]byte, error) {
	return toBytesPrimitiveImpl(v)
}
func (v *UnsignedInt) FromBytes(data []byte) (bytesRead uint, err error) {
	return fromBytesPrimitiveImpl(data, v)
}

func (v Long) ToBytes() ([]byte, error) {
	return toBytesPrimitiveImpl(v)
}
func (v *Long) FromBytes(data []byte) (bytesRead uint, err error) {
	return fromBytesPrimitiveImpl(data, v)
}

func (v Float) ToBytes() ([]byte, error) {
	return toBytesPrimitiveImpl(v)
}
func (v *Float) FromBytes(data []byte) (bytesRead uint, err error) {
	return fromBytesPrimitiveImpl(data, v)
}

func (v Double) ToBytes() ([]byte, error) {
	return toBytesPrimitiveImpl(v)
}
func (v *Double) FromBytes(data []byte) (bytesRead uint, err error) {
	return fromBytesPrimitiveImpl(data, v)
}

// ********** toBytes/fromBytes helpers for all primitives

type primitiveType interface {
	Boolean | Byte | UnsignedByte | Short | UnsignedShort | Int | UnsignedInt | Long | Float | Double
}

func toBytesPrimitiveImpl[T primitiveType](value T) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, ORDER, value)
	return buf.Bytes(), err
}

func fromBytesPrimitiveImpl[T primitiveType](data []byte, out *T) (bytesRead uint, err error) {
	buf := bytes.NewReader(data)
	err = binary.Read(buf, ORDER, out)

	if err != nil {
		return
	}

	bytesRead = uint(buf.Size() - int64(buf.Len()))
	return
}
