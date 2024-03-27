package wkproto

import (
	"fmt"

	"github.com/pkg/errors"
)

// RecvAckPacket 对收取包回执
type RecvAckPacket struct {
	Framer
	MessageID  int64  // 服务端的消息ID(全局唯一)
	MessageSeq uint32 // 消息序列号
}

// GetPacketType 包类型
func (s *RecvAckPacket) GetFrameType() FrameType {
	return RECVACK
}

func (s *RecvAckPacket) String() string {
	return fmt.Sprintf("Framer:%s MessageId:%d MessageSeq:%d", s.Framer.String(), s.MessageID, s.MessageSeq)
}

func decodeRecvAck(frame Frame, data []byte, version uint8) (Frame, error) {
	dec := NewDecoder(data)
	recvAckPacket := &RecvAckPacket{}
	recvAckPacket.Framer = frame.(Framer)
	var err error
	// 消息唯一ID
	if recvAckPacket.MessageID, err = dec.Int64(); err != nil {
		return nil, errors.Wrap(err, "解码MessageId失败！")
	}
	// 消息唯序列号
	if recvAckPacket.MessageSeq, err = dec.Uint32(); err != nil {
		return nil, errors.Wrap(err, "解码MessageSeq失败！")
	}
	return recvAckPacket, err
}

func encodeRecvAck(recvAckPacket *RecvAckPacket, enc *Encoder, version uint8) error {
	enc.WriteInt64(recvAckPacket.MessageID)
	enc.WriteUint32(recvAckPacket.MessageSeq)
	return nil
}

func encodeRecvAckSize(packet *RecvAckPacket, version uint8) int {
	return MessageIDByteSize + MessageSeqByteSize
}
