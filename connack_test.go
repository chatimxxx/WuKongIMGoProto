package xoproto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnAckEncodeAndDecode(t *testing.T) {
	packet := &ConnackPacket{
		TimeDiff:      12345,
		ReasonCode:    ReasonSuccess,
		ServerKey:     "ServerKey",
		Salt:          "Salt",
		ServerVersion: 100,
	}
	packet.HasServerVersion = true
	codec := New()
	// 编码
	packetBytes, err := codec.EncodeFrame(packet, 4)
	assert.NoError(t, err)
	// 解码
	resultPacket, _, err := codec.DecodeFrame(packetBytes, 4)
	assert.NoError(t, err)
	resultConnAckPacket, ok := resultPacket.(*ConnackPacket)
	assert.Equal(t, true, ok)

	// 正确与否比较
	assert.Equal(t, packet.TimeDiff, resultConnAckPacket.TimeDiff)
	assert.Equal(t, packet.ReasonCode, resultConnAckPacket.ReasonCode)
	assert.Equal(t, packet.ServerKey, resultConnAckPacket.ServerKey)
	assert.Equal(t, packet.Salt, resultConnAckPacket.Salt)
	assert.Equal(t, packet.ServerVersion, resultConnAckPacket.ServerVersion)
}
