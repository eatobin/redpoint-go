package player

import (
	"encoding/json"
	"errors"
	"github.com/eatobin/redpoint-go/giftHistory"
	"github.com/eatobin/redpoint-go/giftPair"
)

type Name = string

// A Player has a PlayerName and a GiftHistory
type Player struct {
	PlayerName  Name                `json:"playerName"`
	GiftHistory giftHistory.History `json:"giftHistory"`
}

// JsonStringToPlayer turns a JSON string into a Player
func JsonStringToPlayer(jsonString giftPair.JsonString) (Player, error) {
	var player Player
	err := json.Unmarshal([]byte(jsonString), &player)
	if err != nil {
		return Player{}, err
	}
	if player.PlayerName == "" {
		err = errors.New("missing PlayerName field value")
		return Player{}, err
	}
	for _, v := range player.GiftHistory {
		if v.Givee == "" || v.Giver == "" {
			err = errors.New("missing a GiftHistory field value")
			return Player{}, err
		}
	}
	return player, nil
}

//
//// UpdateGiftHistory updates a GiftHistory in a Player
//func (player Player) UpdateGiftHistory(gh GiftHistory) Player {
//	player.GiftHistory = gh
//	return player
//}
//
//// String makes a Player into a string
//func (player Player) String() string {
//	return fmt.Sprintf("{PlayerName: %s, GiftHistory: %v}", player.PlayerName, player.GiftHistory)
//}
//

//
//// PlayerToJsonString turns a Player into a Player JSON string
//func (player Player) PlayerToJsonString() (string, error) {
//	playerByte, err := json.Marshal(player)
//	return string(playerByte), err
//}
