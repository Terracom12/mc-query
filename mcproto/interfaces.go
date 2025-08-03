package mcproto

type Readable interface {
	FromBytes(data []byte) (bytesRead uint, err error)
}

type Writable interface {
	ToBytes() ([]byte, error)
}
