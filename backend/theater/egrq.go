package theater

import (
	"net"
	"strconv"

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


	//game server
	gameServer.WriteEncode(&codec.Answer{
		Type: codec.ThtrEnterGameRequest,
		Payload: ansEGRQ{
			TID:          event.Command.Message["TID"],
			Name:         gr.HeroName,
			UserID:       gr.HeroID,
			PlayerID:     gr.PlayerID,
			Ticket:       "2018751182",
			IP:           externalIP,
			Port:         strconv.Itoa(event.Client.IpAddr.(*net.TCPAddr).Port),
			IntIP:        event.Command.Message["R-INT-IP"],
			IntPort:      event.Command.Message["R-INT-PORT"],
			Ptype:        "P",
			RUser:        gr.HeroName,
			RUid:         gr.HeroID,
			RUAccid:      gr.PlayerID,
			RUElo:        gr.Stats["elo"],
			RUTeam:       gr.Stats["c_team"],
			RUKit:        gr.Stats["c_kit"],
			RULvl:        gr.Stats["level"],
			RUDataCenter: network.RegionEastCoast,
			RUExternalIP: externalIP,
			RUInternalIP: event.Command.Message["R-INT-IP"],
			RUCategory:   event.Command.Message["R-U-category"],
			RIntIP:       event.Command.Message["R-INT-IP"],
			RIntPort:     event.Command.Message["R-INT-PORT"],
			Xuid:         "24",
			RXuid:        "24",
			LobbyID:      gr.LobbyID,
			GameID:       gr.GameID,
		},
	})

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
		Port          string `fesl:"P"`
		EncryptionKey string `fesl:"EKEY"`
		// Alternatively to EKEY it is possible to use NOENCYRPTIONKEY
		NoEcryptionKey string `fesl:"NOENCYRPTIONKEY,omitempty"`
		IntIP          string `fesl:"INT-IP"`
		IntPort        string `fesl:"INT-PORT"`
		Secret         string `fesl:"SECRET,omitempty"`
		// Alternatively to SECRET it is possible to use NOSECRET
		NoSecret string `fesl:"NOSECRET,omitempty"`
		Ugid     string `fesl:"UGID,omitempty"`
		// Alternatively to UGID it is possible to use NOGUID
		NoGUID  string `fesl:"NOGUID,omitempty"`
		Huid     string `fesl:"HUID"`
		LobbyID string `fesl:"LID"`
		GameID  int    `fesl:"GID"`
	}

	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrEnterGameEntitleGame,
		Payload: ansEGEG{
			TID:           event.Command.Message["TID"],
			Platform:      "PC",
			Ticket:       "2018751182",
			PlayerID:      gr.PlayerID,
			IP:            gameServer.ServerData.Get("IP"),
			Port:          gameServer.ServerData.Get("PORT"),
			EncryptionKey: "O65zZ2D2A58mNrZw1hmuJw%3d%3d",
			IntIP:   gameServer.ServerData.Get("INT-IP"),
			IntPort: gameServer.ServerData.Get("INT-PORT"),
			Secret:  "2587913",
			Ugid:    gameServer.ServerData.Get("UGID"),
			LobbyID: gr.LobbyID,
			Huid:     "1",
			GameID:  gr.GameID,
		},
	})
}



