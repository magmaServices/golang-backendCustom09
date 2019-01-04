package matchmaking

import (
	"localhost/go-heroes/fesl-backend/backend/network"
)

type Game struct {
	ID         int
	LobbyID    int
	GameServer *network.Client

	PlayersJoining int
	PlayersPlaying int
	PlayerSlots    int
	// VipSlots       int
}
