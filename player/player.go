package player

import (
	"encoding/json"
	"fmt"
	"github.com/eatobin/redpoint-go/giftHistoryPkg"
	"github.com/eatobin/redpoint-go/giftPairPkg"
)

type GiftPair = giftPairPkg.GiftPair
type GiftHistory = giftHistoryPkg.GiftHistory

// A Player has a PlayerName and a GiftHistory
type Player struct {
	PlayerName  string      `json:"playerName"`
	GiftHistory GiftHistory `json:"giftHistory"`
}

// ComparePlayer compares two Players
func ComparePlayer(a, b Player) bool {
	if &a == &b {
		return true
	}
	if a.PlayerName != b.PlayerName {
		return false
	}
	if len(a.GiftHistory) != len(b.GiftHistory) {
		return false
	}
	for k, v := range a.GiftHistory {
		if b.GiftHistory[k] != v {
			return false
		}
	}
	return true
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
