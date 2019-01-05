package gsum

import (
	"gitlab.com/oiacow/fesl3/backend/network"
	"gitlab.com/oiacow/fesl3/backend/network/codec"
)

const (
	gsumGetSessionID = "GetSessionId"
)

// GameSummary probably stands for Game Summary
type GameSummary struct {
	//
}

func (gsum *GameSummary) answer(client *network.Client, pnum uint32, payload interface{}) {
	client.WriteEncode(&codec.Answer{
		Type:         codec.FeslGameSummary,
		PacketNumber: pnum,
		Payload:      payload,
	})
}
