package fields

import "mc-query/mcproto"

type Field interface {
	mcproto.Readable
	mcproto.Writable
}
