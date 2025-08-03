package packets

import (
	"mc-query/mcproto/fields"
	"net"
)

const statusRequestID = fields.VarInt(0x0)
const statusResponseID = fields.VarInt(0x0)

type StatusRequest struct {
}
type StatusResponse struct {
	JsonResponse fields.String
}

func (h StatusRequest) Send(conn net.Conn) error {
	return sendPacket(conn, statusRequestID)
}

func (h *StatusResponse) Receive(conn net.Conn) (readBytes uint, err error) {
	h.JsonResponse.MaxLen = 32767
	return receiveKnownPacket(conn, statusResponseID, &h.JsonResponse)
}
