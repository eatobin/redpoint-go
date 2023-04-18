package redpoint

import (
	"encoding/json"
	"errors"
	"fmt"
)

// A GiftPairStruct has a Givee and a Giver
type GiftPairStruct struct {
	Givee string `json:"givee"`
	Giver string `json:"giver"`
}

// GiftPairJsonStringToGiftPair turns a JSON string into a GiftPairStruct
func GiftPairJsonStringToGiftPair(jsonString string) (GiftPairStruct, error) {
	var giftPair GiftPairStruct
	err := json.Unmarshal([]byte(jsonString), &giftPair)
	if err != nil {
		return GiftPairStruct{}, err
	}
	if giftPair.Givee == "" || giftPair.Giver == "" {
		err = errors.New("missing one or both field values")
		return GiftPairStruct{}, err
	}
	return giftPair, nil
}

// GiftPairUpdateGivee updates a Givee in a GiftPairStruct
func (giftPair GiftPairStruct) GiftPairUpdateGivee(givee string) GiftPairStruct {
	giftPair.Givee = givee
	return giftPair
}

// GiftPairUpdateGiver updates a Giver in a GiftPairStruct
func (giftPair GiftPairStruct) GiftPairUpdateGiver(giver string) GiftPairStruct {
	giftPair.Giver = giver
	return giftPair
}

// GiftPairString makes a GiftPairStruct into a string
func (giftPair GiftPairStruct) GiftPairString() string {
	return fmt.Sprintf("{Givee: %s, Giver: %s}", giftPair.Givee, giftPair.Giver)
}
