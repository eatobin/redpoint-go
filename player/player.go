package player

import (
	"encoding/json"
	"fmt"
	"github.com/eatobin/redpointGo/giftHistory"
	"github.com/eatobin/redpointGo/giftPair"
)

type GiftPair = giftPair.GiftPair
type GiftHistory = giftHistory.GiftHistory

// A Player has a PlayerName and a GiftHistory
type Player struct {
	PlayerName  string      `json:"playerName"`
	GiftHistory GiftHistory `json:"giftHistory"`
}

// UpdateGiftHistory updates a GiftHistory in a Player
func (player Player) UpdateGiftHistory(gh GiftHistory) Player {
	player.GiftHistory = gh
	return player
}

// String makes a Player into a string
func (player Player) String() string {
	return fmt.Sprintf("{PlayerName: %s, GiftHistory: %v}", player.PlayerName, player.GiftHistory)
}

// JsonStringToPlayer turns a Player JSON string into a Player
func JsonStringToPlayer(playerString string) (Player, error) {
	var player Player
	err := json.Unmarshal([]byte(playerString), &player)
	return player, err
}

// PlayerToJsonString turns a Player into a Player JSON string
func (player Player) PlayerToJsonString() (string, error) {
	playerByte, err := json.Marshal(player)
	return string(playerByte), err
}
