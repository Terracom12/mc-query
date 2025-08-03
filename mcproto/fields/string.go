package fields

import (
	"errors"
)

type String struct {
	Str    string
	MaxLen uint
}

func (s String) ToBytes() ([]byte, error) {
	if len(s.Str) > int(s.MaxLen) {
		return nil, errors.New("String length is > maxLen")
	}

	var (
		length = VarInt(len(s.Str))
		buf    []byte
	)

	lengthBytes, err := length.ToBytes()
	if err != nil {
		return nil, err
	}

	buf = append(buf, lengthBytes...)
	buf = append(buf, s.Str...)

	return buf, nil
}

func (s *String) FromBytes(data []byte) (bytesRead uint, err error) {
	length, intBytesRead, err := fromBytesIntImpl(data, 4)

	if err != nil {
		return
	}

	if length > uint64(s.MaxLen) {
		err = errors.New("String length is > maxLen")
		return
	}

	stringBytes := data[intBytesRead : intBytesRead+uint(length)]

	(*s).Str = string(stringBytes)

	bytesRead = intBytesRead + uint(length)
	return
}
