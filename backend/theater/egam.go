package theater

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/oiacow/fesl3/backend/network"
	"gitlab.com/oiacow/fesl3/backend/network/codec"
)

// EGAM is sent to Game-Client
type reqEGAM struct {
	GameID int `fesl:"GID"`
	LobbyID int `fesl:"LID"`
	Port int `fesl:"PORT"`
	PlatformType int `fesl:"PTYPE"`
	RemoteIP string `fesl:"R-INT-IP"`
	RemotePort int `fesl:"R-INT-PORT"`
	AccountID int `fesl:"R-U-accid"` // TODO: Hero or PlayerID? PlayerID :(
	Category int `fesl:"R-U-category"` // TODO: What exactly it is?
	Region string `fesl:"R-U-dataCenter"`
	StatsElo int `fesl:"R-U-elo"`
	ExternalIP string `fesl:"R-U-externalIp"`
	StatsKit int `fesl:"R-U-kit"`
	StatsLevel int `fesl:"R-U-lvl"`
	StatsTeam int `fesl:"R-U-team"`
	TID int `fesl:"TID"`
}

type ansEGAM struct {
	TID     string `fesl:"TID"`
	LobbyID string `fesl:"LID"`
	GameID  int    `fesl:"GID"`
}

// a EGAM - CLIENT called when a client wants to join a gameserver
func (tm *Theater) EnterGame(event network.EventClientCommand) {
	gameID, err := event.Command.Message.IntVal("GID")
	if err != nil {
		logrus.WithError(err).Warn("Cannot parse value of GID in theater.EGAM")
		return
	}

	game, err := tm.mm.GetGame(gameID)
	if err != nil {
		logrus.
			WithError(err).
			WithField("gameID", gameID).
			Warn("Not found any server when joining game")
		return
	}

	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrEnterGame,
		Payload: ansEGAM{
			event.Command.Message["TID"],
			event.Command.Message["LID"],
			gameID,
		},
	})	

	gr := GameRequest{
		GameID:   game.ID,
	}

	tm.EnterGameRequest(&event, game.GameServer, gr)
}

type GameRequest struct {
	GameID   int
}

