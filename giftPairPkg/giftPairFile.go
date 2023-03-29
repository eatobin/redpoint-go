package giftPairPkg

import (
	"encoding/json"
	"errors"
	"fmt"
)

type PlayerKeyTA = string
type GiveeTA = PlayerKeyTA
type GiverTA = PlayerKeyTA
type JsonStringTA = string

// A GiftPairStruct has a Givee and a Giver
type GiftPairStruct struct {
	Givee GiveeTA `json:"givee"`
	Giver GiverTA `json:"giver"`
}

// GiftPairJsonStringToGiftPair turns a JSON string into a GiftPairStruct
func GiftPairJsonStringToGiftPair(jsonString JsonStringTA) (GiftPairStruct, error) {
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
func GiftPairUpdateGivee(givee GiveeTA, giftPair GiftPairStruct) GiftPairStruct {
	giftPair.Givee = givee
	return giftPair
}

// GiftPairUpdateGiver updates a Giver in a GiftPairStruct
func GiftPairUpdateGiver(giver GiverTA, giftPair GiftPairStruct) GiftPairStruct {
	giftPair.Giver = giver
	return giftPair
}

// GiftPairString makes a GiftPairStruct into a string
func GiftPairString(giftPair GiftPairStruct) string {
	return fmt.Sprintf("{Givee: %s, Giver: %s}", giftPair.Givee, giftPair.Giver)
}
