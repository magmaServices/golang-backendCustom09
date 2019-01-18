package theater

import (
	"net"
	"gitlab.com/oiacow/fesl3/backend/ranking"

	"strconv"
	"github.com/sirupsen/logrus"
	"gitlab.com/oiacow/fesl3/backend/network"
	"gitlab.com/oiacow/fesl3/backend/network/codec"
)

type reqEGRQ struct {
	reqEGAM
}

type ansEGRQ struct {
	TID      string `fesl:"TID"`
	Name     string `fesl:"NAME"`
	UserID   int    `fesl:"UID"`
	PlayerID int    `fesl:"PID"`
	Ticket   string `fesl:"TICKET"`
	IP       string `fesl:"IP"`
	Port     string `fesl:"PORT"`
	IntIP    string `fesl:"INT-IP"`
	IntPort  string `fesl:"INT-PORT"`
	// PTPE can be O or P
	Ptype        string `fesl:"PTYPE"`
	RUser        string `fesl:"R-USER"`
	RUid         int    `fesl:"R-UID"`
	RUAccid      int    `fesl:"R-U-accid"`
	RUElo        string `fesl:"R-U-elo"`
	RUTeam       string `fesl:"R-U-team"`
	Platform     string `fesl:"PL"`
	RUKit        string `fesl:"R-U-kit"`
	RULvl        string `fesl:"R-U-lvl"`
	RUDataCenter string `fesl:"R-U-dataCenter"`
	RUExternalIP string `fesl:"R-U-externalIp"`
	RUInternalIP string `fesl:"R-U-internalIp"`
	RUCategory   string `fesl:"R-U-category"`
	RIntIP       string `fesl:"R-INT-IP"`
	RIntPort     string `fesl:"R-INT-PORT"`
	Xuid         string `fesl:"XUID"`
	RXuid        string `fesl:"R-XUID"`
	LobbyID      string `fesl:"LID"`
	GameID       int    `fesl:"GID"`
}

// EnterGameRequest (EGRQ) is sent to Server to inform about the player
// who wants join server
func (tm *Theater) EnterGameRequest(event *network.EventClientCommand, gameServer *network.Client, gr GameRequest) {
	externalIP := event.Client.IpAddr.(*net.TCPAddr).IP.String()
	heroStats, err := tm.db.FindHeroStats(tm.db.NewSession(), event.Client.PlayerData.HeroID)
	if err != nil {
		logrus.
			WithError(err).
			WithField("heroID", event.Client.PlayerData.HeroID).
			Warn("Cannot fetch stats for hero when entering a game")
		return
	}	

	
	stats, err := ranking.GetStats(&heroStats, "c_kit", "c_team", "elo", "level")
	if err != nil {
		logrus.
			WithError(err).
			WithField("heroID", event.Client.PlayerData.HeroID).
			Warn("Cannot get stats for hero when entering a game")
		return
	}

	gameID, err := event.Command.Message.IntVal("GID")
	if err != nil {
		logrus.WithError(err).Warn("Cannot parse value of GID in theater.EGAM")
		return
	}

	// game, err := tm.mm.GetGame(gameID)
	// if err != nil {
	// 	logrus.
	// 		WithError(err).
	// 		WithField("gameID", gameID).
	// 		Warn("Not found any server when joining game")
	// 	return
	// }

	PlayerID := event.Client.PlayerData.PlayerID
	HeroID := event.Client.PlayerData.HeroID
	HeroName := event.Client.PlayerData.HeroName
	Stats := stats

	logrus.Println("-BEGIN EGRQ------")
	//game server
	gameServer.WriteEncode(&codec.Answer{
		Type: codec.ThtrEnterGameRequest,
		Payload: ansEGRQ{
			TID:          event.Command.Message["TID"],
			Name:         HeroName,
			UserID:       HeroID,
			PlayerID:     PlayerID,
			Ticket:       "2018751182",
			IP:           externalIP,
			Port:         strconv.Itoa(event.Client.IpAddr.(*net.TCPAddr).Port),
			IntIP:        event.Command.Message["R-INT-IP"],
			IntPort:      event.Command.Message["R-INT-PORT"],
			Ptype:        "P",
			Platform:      "PC",
			RUser:        HeroName,
			RUid:         HeroID,
			RUAccid:      PlayerID,
			RUElo:        Stats["elo"],
			RUTeam:       Stats["c_team"],
			RUKit:        Stats["c_kit"],
			RULvl:        Stats["level"],
			RUDataCenter: network.RegionEastCoast,
			RUExternalIP: externalIP,
			RUInternalIP: event.Command.Message["R-INT-IP"],
			RUCategory:   event.Command.Message["R-U-category"],
			RIntIP:       event.Command.Message["R-INT-IP"],
			RIntPort:     event.Command.Message["R-INT-PORT"],
			Xuid:         "24",
			RXuid:        "24",
			LobbyID:      "1",
			GameID:       gameID,
		},
	})
	logrus.Println("-END EGRQ------")

	//Game Client
	type reqEGEG struct {
		reqEGAM
	}

	type ansEGEG struct {
		TID           string `fesl:"TID"`
		Platform      string `fesl:"PL"`
		Ticket        string `fesl:"TICKET"`
		PlayerID      int    `fesl:"PID"`
		IP            string `fesl:"I"`
		Port          string `fesl:"P"` //in client port=p and ip= i lol
		EncryptionKey string `fesl:"EKEY"`
		IntIP          string `fesl:"INT-IP"`
		IntPort        string `fesl:"INT-PORT"`
		Secret         string `fesl:"SECRET"`
		Ugid     string `fesl:"UGID"`
		Huid     string `fesl:"HUID"`
		LobbyID string `fesl:"LID"`
		GameID  int    `fesl:"GID"`
	}

	logrus.Println("-BEGIN   E G E G ------")
	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrEnterGameEntitleGame,
		Payload: ansEGEG{
			TID:           event.Command.Message["TID"],
			Platform:      "PC",
			Ticket:       "2018751182",
			PlayerID:      PlayerID,
			IP:            gameServer.ServerData.Get("IP"),
			Port:          gameServer.ServerData.Get("PORT"),
			EncryptionKey: "TEST1234",
			IntIP:   gameServer.ServerData.Get("INT-IP"),
			IntPort: gameServer.ServerData.Get("INT-PORT"),
			Secret:  "MargeSimpson",
			Ugid:    gameServer.ServerData.Get("UGID"),
			LobbyID: "1",
			Huid:     "1",
			GameID:  gameID,
		},
	})
	logrus.Println("-END E GEG------")

}



