package giftPairPkg

import (
	"encoding/json"
	"errors"
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

// GiftPairAssertEqual compares two GiftPairs
func GiftPairAssertEqual(a, b GiftPairStruct) bool {
	if &a == &b {
		return true
	}
	if a.Givee != b.Givee || a.Giver != b.Giver {
		return false
	}
	return true
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
	return giftPair, err
}

//// GiftPairUpdateGivee updates a Givee in a GiftPairStruct
//func (giftPair GiftPairStruct) GiftPairUpdateGivee(givee GiveeTA) GiftPairStruct {
//	giftPair.Givee = givee
//	return giftPair
//}
//
//// GiftPairUpdateGiver updates a Giver in a GiftPairStruct
//func (giftPair GiftPairStruct) GiftPairUpdateGiver(giver GiverTA) GiftPairStruct {
//	giftPair.Giver = giver
//	return giftPair
//}
//
//// String makes a GiftPairStruct into a string
//func (giftPair GiftPairStruct) String() string {
//	return fmt.Sprintf("{Givee: %s, Giver: %s}", giftPair.Givee, giftPair.Giver)
//}
