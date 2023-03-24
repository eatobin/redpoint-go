package giftPairPkg

import (
	"encoding/json"
	"fmt"
)

type PlayerKeyTA = string
type GiveeTA = PlayerKeyTA
type GiverTA = PlayerKeyTA
type JsonStringTA = string

// A GiftPair has a Givee and a Giver
type GiftPair struct {
	Givee GiveeTA `json:"givee"`
	Giver GiverTA `json:"giver"`
}

// GiftPairAssertEqual compares two GiftPairs
func GiftPairAssertEqual(a, b GiftPair) bool {
	if &a == &b {
		return true
	}
	if a.Givee != b.Givee || a.Giver != b.Giver {
		return false
	}
	return true
}

// GiftPairJsonStringToGiftPair turns a JSON string into a GiftPair
func GiftPairJsonStringToGiftPair(jsonString JsonStringTA) (GiftPair, error) {
	var giftPair GiftPair
	err := json.Unmarshal([]byte(jsonString), &giftPair)
	return giftPair, err
}

// GiftPairUpdateGivee updates a Givee in a GiftPair
func (giftPair GiftPair) GiftPairUpdateGivee(givee GiveeTA) GiftPair {
	giftPair.Givee = givee
	return giftPair
}

// GiftPairUpdateGiver updates a Giver in a GiftPair
func (giftPair GiftPair) GiftPairUpdateGiver(giver GiverTA) GiftPair {
	giftPair.Giver = giver
	return giftPair
}

// String makes a GiftPair into a string
func (giftPair GiftPair) String() string {
	return fmt.Sprintf("{Givee: %s, Giver: %s}", giftPair.Givee, giftPair.Giver)
}
