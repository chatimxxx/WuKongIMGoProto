package wkproto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecvAckEncodeAndDecode(t *testing.T) {

	packet := &RecvAckPacket{
		Framer: Framer{
			RedDot: true,
		},
		MessageID:  1234,
		MessageSeq: 2334,
	}

	codec := New()
	// 编码
	packetBytes, err := codec.EncodeFrame(packet, 1)
	assert.NoError(t, err)

	// 解码
	resultPacket, _, err := codec.DecodeFrame(packetBytes, 1)
	assert.NoError(t, err)
	resultRecvAckPacket, ok := resultPacket.(*RecvAckPacket)
	assert.Equal(t, true, ok)

	// 比较
	assert.Equal(t, packet.MessageID, resultRecvAckPacket.MessageID)
	assert.Equal(t, packet.MessageSeq, resultRecvAckPacket.MessageSeq)
	assert.Equal(t, packet.RedDot, resultRecvAckPacket.RedDot)
}
