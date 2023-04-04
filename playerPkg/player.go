package playerPkg

import (
	"encoding/json"
	"errors"
	"github.com/eatobin/redpoint-go/giftHistoryPkg"
)

// A StructPlayer has a PlayerName and a GiftHistory
type StructPlayer struct {
	PlayerName  string                 `json:"playerName"`
	GiftHistory giftHistoryPkg.History `json:"giftHistory"`
}

// JsonStringToPlayer turns a JSON string into a StructPlayer
func JsonStringToPlayer(jsonString string) (StructPlayer, error) {
	var player StructPlayer
	err := json.Unmarshal([]byte(jsonString), &player)
	if err != nil {
		return StructPlayer{}, err
	}
	if player.PlayerName == "" {
		err = errors.New("missing PlayerName field value")
		return StructPlayer{}, err
	}
	if len(player.GiftHistory) == 0 {
		err = errors.New("missing GiftHistory field value")
		return StructPlayer{}, err
	}
	for _, v := range player.GiftHistory {
		if v.Givee == "" || v.Giver == "" {
			err = errors.New("missing a GiftHistory field value")
			return StructPlayer{}, err
		}
	}
	return player, nil
}

//// UpdateGiftHistory updates a GiftHistory in a StructPlayer
//func (player StructPlayer) UpdateGiftHistory(giftHistory giftHistory.History) StructPlayer {
//	player.GiftHistory = giftHistory
//	return player
//}

//
//// String makes a StructGiftPair into a string
//func (player StructGiftPair) String() string {
//	return fmt.Sprintf("{PlayerName: %s, GiftHistory: %v}", player.PlayerName, player.GiftHistory)
//}
//

//
//// PlayerToJsonString turns a StructGiftPair into a StructGiftPair JSON string
//func (player StructGiftPair) PlayerToJsonString() (string, error) {
//	playerByte, err := json.Marshal(player)
//	return string(playerByte), err
//}
