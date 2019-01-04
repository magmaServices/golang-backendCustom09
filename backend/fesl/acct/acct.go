package acct

import (
	"localhost/go-heroes/fesl-backend/backend/network"
	"localhost/go-heroes/fesl-backend/backend/network/codec"
	"localhost/go-heroes/fesl-backend/backend/storage/database"
)

const (
	acctGetTelemetryToken = "GetTelemetryToken"
	acctNuGetAccount      = "NuGetAccount"
	acctNuGetPersonas     = "NuGetPersonas"
	acctNuLogin           = "NuLogin"
	acctNuLoginPersona    = "NuLoginPersona"
	acctNuLookupUserInfo  = "NuLookupUserInfo"
)

const (
	clientTypeServer = "server"
)

// Account probably stands for "Account"
type Account struct {
	DB database.Adapter
}

func (acct *Account) answer(client *network.Client, pnum uint32, payload interface{}) {
	client.WriteEncode(&codec.Answer{
		Type:         codec.FeslAccount,
		PacketNumber: pnum,
		Payload:      payload,
	})
}
