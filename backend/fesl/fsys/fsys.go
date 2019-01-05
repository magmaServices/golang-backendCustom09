package fsys

import (
	"gitlab.com/oiacow/fesl3/backend/network"
	"gitlab.com/oiacow/fesl3/backend/network/codec"
)

const (
	fsysGetPingSites = "GetPingSites"
	fsysHello        = "Hello"
	fsysMemCheck     = "MemCheck"
	fsysPing         = "Ping"
)

type ConnectSystem struct {
	ServerMode bool
}

func (fsys *ConnectSystem) answer(client *network.Client, pnum uint32, payload interface{}) {
	client.WriteEncode(&codec.Answer{
		Type:         codec.FeslSystem,
		PacketNumber: pnum,
		Payload:      payload,
	})
}
