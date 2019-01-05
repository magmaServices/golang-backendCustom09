package pnow

import (
	"gitlab.com/oiacow/fesl3/backend/matchmaking"
	"gitlab.com/oiacow/fesl3/backend/network"
	"gitlab.com/oiacow/fesl3/backend/network/codec"
)

const (
	pnowStart  = "Start"
	pnowStatus = "Status"
)

// PlayNow probably stands for PlayNow
type PlayNow struct {
	MM *matchmaking.Pool
}

func (pnow *PlayNow) answer(client *network.Client, pnum uint32, payload interface{}) {
	client.WriteEncode(&codec.Answer{
		Type:         codec.FeslPlayNow,
		PacketNumber: pnum,
		Payload:      payload,
	})
}
