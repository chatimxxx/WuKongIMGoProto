package wkproto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendackEncodeAndDecode(t *testing.T) {

	packet := &SendAckPacket{
		ClientSeq:  234,
		MessageSeq: 2,
		MessageID:  1234,
		ReasonCode: ReasonSuccess,
	}

	codec := New()
	// 编码
	packetBytes, err := codec.EncodeFrame(packet, 1)
	assert.NoError(t, err)
	// 解码
	resultPacket, _, err := codec.DecodeFrame(packetBytes, 1)
	assert.NoError(t, err)
	resultSendAckPacket, ok := resultPacket.(*SendAckPacket)
	assert.Equal(t, true, ok)

	// 比较
	assert.Equal(t, packet.ClientSeq, resultSendAckPacket.ClientSeq)
	assert.Equal(t, packet.MessageSeq, resultSendAckPacket.MessageSeq)
	assert.Equal(t, packet.MessageID, resultSendAckPacket.MessageID)
	assert.Equal(t, packet.ReasonCode, resultSendAckPacket.ReasonCode)
}
