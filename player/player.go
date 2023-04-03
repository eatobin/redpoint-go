package player

import (
	"encoding/json"
	"errors"
	"github.com/eatobin/redpoint-go/giftHistory"
	"github.com/eatobin/redpoint-go/giftPair"
)

type Name = string

// A Struct has a PlayerName and a GiftHistory
type Struct struct {
	PlayerName  Name                `json:"playerName"`
	GiftHistory giftHistory.History `json:"giftHistory"`
}

// JsonStringToPlayer turns a JSON string into a Struct
func JsonStringToPlayer(jsonString giftPair.JsonString) (Struct, error) {
	var player Struct
	err := json.Unmarshal([]byte(jsonString), &player)
	if err != nil {
		return Struct{}, err
	}
	if player.PlayerName == "" {
		err = errors.New("missing PlayerName field value")
		return Struct{}, err
	}
	if len(player.GiftHistory) == 0 {
		err = errors.New("missing GiftHistory field value")
		return Struct{}, err
	}
	for _, v := range player.GiftHistory {
		if v.Givee == "" || v.Giver == "" {
			err = errors.New("missing a GiftHistory field value")
			return Struct{}, err
		}
	}
	return player, nil
}

//
//// UpdateGiftHistory updates a GiftHistory in a Struct
//func (player Struct) UpdateGiftHistory(gh GiftHistory) Struct {
//	player.GiftHistory = gh
//	return player
//}
//
//// String makes a Struct into a string
//func (player Struct) String() string {
//	return fmt.Sprintf("{PlayerName: %s, GiftHistory: %v}", player.PlayerName, player.GiftHistory)
//}
//

//
//// PlayerToJsonString turns a Struct into a Struct JSON string
//func (player Struct) PlayerToJsonString() (string, error) {
//	playerByte, err := json.Marshal(player)
//	return string(playerByte), err
//}
