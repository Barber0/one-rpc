package one

import (
	"one/protocol/res/requestf"
	"time"
)

type Message struct {
	Rsp		*requestf.RspPacket
	start	time.Time
}

func (m *Message) Start() {
	m.start = time.Now()
}

func (m *Message) Finish() time.Duration {
	return time.Now().Sub(m.start)
}