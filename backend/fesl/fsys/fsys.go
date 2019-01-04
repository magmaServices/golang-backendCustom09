package fsys

import (
	"localhost/go-heroes/fesl-backend/backend/network"
	"localhost/go-heroes/fesl-backend/backend/network/codec"
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
