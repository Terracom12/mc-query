package packets

import (
	"log"
	"mc-query/mcproto/fields"
	"net"
)

const handshakeID = fields.VarInt(0x0)

type HandshakeIntent fields.VarInt

const (
	IntentStatus   HandshakeIntent = 1
	IntentLogin                    = 2
	IntentTransfer                 = 3
)

const HandshakeAddrMaxLen = 255

type Handshake struct {
	protocolVersion fields.VarInt
	serverAddress   fields.String        // maxLen = 255 ; unused by vanilla
	serverPort      fields.UnsignedShort // unused by vanilla
	intent          HandshakeIntent
}

func MakeHandshake(protocolVersion int32, serverAddress string, serverPort uint16, intent HandshakeIntent) Handshake {
	return Handshake{
		fields.VarInt(protocolVersion),
		fields.String{Str: serverAddress, MaxLen: HandshakeAddrMaxLen},
		fields.UnsignedShort(serverPort),
		intent,
	}
}

func (h Handshake) Send(conn net.Conn) (err error) {
	if h.serverAddress.MaxLen > 255 {
		log.Fatalf("serverAddress MaxLen is too large! (%d)", h.serverAddress.MaxLen)
	}

	intent := fields.VarInt(h.intent)
	return sendPacket(conn, handshakeID, &h.protocolVersion, &h.serverAddress, &h.serverPort, &intent)
}
