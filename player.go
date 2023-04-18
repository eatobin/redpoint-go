package redpoint

import (
	"encoding/json"
	"errors"
	"fmt"
)

// A PlayerStruct has a PlayerName and a GiftHistory
type PlayerStruct struct {
	PlayerName  string  `json:"playerName"`
	GiftHistory History `json:"giftHistory"`
}

// PlayerJsonStringToPlayer turns a JSON string into a PlayerStruct
func PlayerJsonStringToPlayer(jsonString string) (PlayerStruct, error) {
	var player PlayerStruct
	err := json.Unmarshal([]byte(jsonString), &player)
	if err != nil {
		return PlayerStruct{}, err
	}
	if player.PlayerName == "" {
		err = errors.New("missing PlayerName field value")
		return PlayerStruct{}, err
	}
	if len(player.GiftHistory) == 0 {
		err = errors.New("missing GiftHistory field value")
		return PlayerStruct{}, err
	}
	for _, v := range player.GiftHistory {
		if v.Givee == "" || v.Giver == "" {
			err = errors.New("missing a GiftHistory field value")
			return PlayerStruct{}, err
		}
	}
	return player, nil
}

// PlayerUpdateGiftHistory updates a GiftHistory in a PlayerStruct
func (player PlayerStruct) PlayerUpdateGiftHistory(giftHistory History) PlayerStruct {
	player.GiftHistory = giftHistory
	return player
}

// PlayerString makes a StructGiftPair into a string
func (player PlayerStruct) PlayerString() string {
	return fmt.Sprintf("{PlayerName: %s, GiftHistory: %v}", player.PlayerName, player.GiftHistory)
}
