package fields

import (
	"errors"
	"unsafe"
)

type VarInt int32
type VarLong int64

func (v VarInt) ToBytes() ([]byte, error) {
	return toBytesIntImpl(uint64(v), uint(unsafe.Sizeof(v)))
}

func (v *VarInt) FromBytes(data []byte) (bytesRead uint, err error) {
	value, bytesRead, err := fromBytesIntImpl(data, uint(unsafe.Sizeof(v)))
	*v = VarInt(value)
	return
}

func (v VarLong) ToBytes() ([]byte, error) {
	return toBytesIntImpl(uint64(v), uint(unsafe.Sizeof(v)))
}

func (v *VarLong) FromBytes(data []byte) (bytesRead uint, err error) {
	value, bytesRead, err := fromBytesIntImpl(data, uint(unsafe.Sizeof(v)))
	*v = VarLong(value)
	return
}

func toBytesIntImpl(value uint64, numBytes uint) ([]byte, error) {
	var (
		buf []byte
	)

	value = value & ((1 << (numBytes * 8)) - 1)

	for {
		temp := byte(value & 0x7F)
		value >>= 7

		if value != 0 {
			temp |= 0x80
		}

		buf = append(buf, temp)
		if value == 0 {
			break
		}
	}

	return buf, nil
}

func fromBytesIntImpl(data []byte, maxNumBytes uint) (result uint64, bytesRead uint, err error) {
	var (
		shift uint
	)

	for _, b := range data {
		bytesRead++

		val := uint64(b & 0x7F)
		result |= val << shift

		if (b & 0x80) == 0 {
			return
		}
		shift += 7

		if shift > maxNumBytes*8 {
			err = errors.New("varInt too long")
			return
		}
	}

	err = errors.New("varInt: incomplete input")
	return
}
