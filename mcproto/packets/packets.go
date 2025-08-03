package packets

import (
	"errors"
	"mc-query/mcproto/fields"
	"net"
)

func sendPacket(conn net.Conn, packetID fields.VarInt, field ...fields.Field) error {
	data, err := serializePacket(packetID, field...)

	if err != nil {
		return err
	}

	numWritten, err := conn.Write(data)

	if err != nil {
		return err
	}
	if numWritten != len(data) {
		return errors.New("Partial packet written!")
	}

	return nil
}

func receiveKnownPacket(conn net.Conn, packetID fields.VarInt, field ...fields.Field) (readBytes uint, err error) {
	data := make([]byte, 32767+100) // Max json data size for status packet + some padding

	numBytesRecv, err := conn.Read(data)

	if err != nil {
		return
	}

	readBytes, err = deserializeKnownPacket(data, packetID, field...)

	if err != nil {
		return
	}
	if numBytesRecv != int(readBytes) {
		err = errors.New("Packet data not fully processed")
	}

	return
}

func serializePacket(packetID fields.VarInt, field ...fields.Field) (data []byte, err error) {
	var (
		buf       []byte
		beforeLen []byte
	)

	buf, err = packetID.ToBytes()
	if err != nil {
		return
	}

	beforeLen = append(beforeLen, buf...)

	for _, f := range field {
		buf, err = f.ToBytes()
		if err != nil {
			return
		}

		beforeLen = append(beforeLen, buf...)
	}

	length := len(beforeLen)

	buf, err = fields.VarInt(length).ToBytes()
	if err != nil {
		return
	}

	data = append(data, buf...)
	data = append(data, beforeLen...)

	return
}

func deserializeKnownPacket(data []byte, packetID fields.VarInt, field ...fields.Field) (bytesRead uint, err error) {
	var (
		bytesReadTemp uint
	)

	var packetLen fields.VarInt
	var readPacketID fields.VarInt

	bytesReadTemp, err = packetLen.FromBytes(data)
	if err != nil {
		return
	}
	bytesRead += bytesReadTemp

	data = data[bytesReadTemp:]

	bytesReadTemp, err = readPacketID.FromBytes(data)
	if err != nil {
		return
	}
	if readPacketID != packetID {
		err = errors.New("Packet IDs do not match")
		return
	}
	bytesRead += bytesReadTemp
	data = data[bytesReadTemp:]

	for _, f := range field {
		bytesReadTemp, err = f.FromBytes(data)
		if err != nil {
			return
		}

		bytesRead += bytesReadTemp
		data = data[bytesReadTemp:]
	}

	return
}
