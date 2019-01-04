package acct

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"localhost/go-heroes/fesl-backend/backend/network"
)

type reqNuLookupUserInfo struct {
	// TXN=NuLookupUserInfo
	TXN string `fesl:"TXN"`

	// userInfo.[]=1
	// userInfo.0.userName=FirstHero
	UserInfo []userInfo `fesl:"userInfo"`
}

type ansNuLookupUserInfo struct {
	Txn      string     `fesl:"TXN"`
	UserInfo []userInfo `fesl:"userInfo"`
}

type userInfo struct {
	Namespace    string `fesl:"namespace"`
	XBoxUserID   string `fesl:"xuid,omitempty"` // int
	MasterUserID int    `fesl:"masterUserId"`
	UserID       int    `fesl:"userId"`
	UserName     string `fesl:"userName"`
}

// NuLookupUserInfo handles acct.NuLookupUserInfo command
func (acct *Account) NuLookupUserInfo(event network.EventClientCommand) {
	if event.Client.GetClientType() == clientTypeServer {
		if event.Command.Message["userInfo.0.userName"] == event.Client.PlayerData.ServerUserName {
			acct.serverNuLookupUserInfo(event)
			return
		}
	}

	acct.clientNuLookupUserInfo(event)
}

func (acct *Account) clientNuLookupUserInfo(event network.EventClientCommand) {
	heroes := []userInfo{}
	keys, _ := strconv.Atoi(event.Command.Message["userInfo.[]"])
	for i := 0; i < keys; i++ {
		heroName := event.Command.Message[fmt.Sprintf("userInfo.%d.userName", i)]

		sess := acct.DB.NewSession()

		h, err := acct.DB.FindHeroByName(sess, heroName)
		if err != nil {
			logrus.WithError(err).Warn("Cannot find Hero with name %s", heroName)
			return
		}

		// TODO: refactor it
		p, err := acct.DB.FindPlayerByID(sess, h.PlayerID)
		if err != nil {
			logrus.WithError(err).Warn("Cannot find Player %d, using Hero.ID %d", h.PlayerID, h.ID)
			return
		}

		masterHeroID := h.ID
		if p.SelectedHeroID.Valid {
			masterHeroID = int(p.SelectedHeroID.Int64)
		}

		heroes = append(heroes, userInfo{
			UserName:     h.HeroName,
			UserID:       h.ID,
			MasterUserID: masterHeroID,
			Namespace:    "MAIN",
			XBoxUserID:   "24",
		})
	}

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLookupUserInfo{Txn: acctNuLookupUserInfo, UserInfo: heroes},
	)
}

func (acct *Account) serverNuLookupUserInfo(event network.EventClientCommand) {
	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLookupUserInfo{
			Txn: acctNuLookupUserInfo,
			UserInfo: []userInfo{
				{
					Namespace:    "MAIN",
					XBoxUserID:   "24",
					MasterUserID: event.Client.PlayerData.ServerID,
					UserID:       event.Client.PlayerData.ServerID,
					UserName:     event.Client.PlayerData.ServerUserName,
				},
			},
		},
	)
}
