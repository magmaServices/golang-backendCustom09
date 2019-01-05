package gsum

import (
	"gitlab.com/oiacow/fesl3/backend/network"
	"gitlab.com/oiacow/fesl3/backend/network/codec"
)

type ansGetSessionID struct {
	Txn string `fesl:"TXN"`
	// Games  []Game  `fesl:"games"`
	// Events []Event `fesl:"events"`
}

// GetSessionID handles gsum.GetSessionID command
func (gsum *GameSummary) GetSessionID(client *network.Client, event *codec.Command) {
	gsum.answer(client, 0, ansGetSessionID{
		Txn: gsumGetSessionID,
	})
}
