package giftPairPkg

import (
	"encoding/json"
	"fmt"
)

type PlayerKeyTA = string
type GiveeTA = PlayerKeyTA
type GiverTA = PlayerKeyTA

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

// GiftPairUpdateGivee updates a GiveeTA in a GiftPair
func (giftPair GiftPair) GiftPairUpdateGivee(givee GiveeTA) GiftPair {
	giftPair.Givee = givee
	return giftPair
}

// GiftPairUpdateGiver updates a GiverTA in a GiftPair
func (giftPair GiftPair) GiftPairUpdateGiver(giver GiverTA) GiftPair {
	giftPair.Giver = giver
	return giftPair
}

// String makes a GiftPair into a string
func (giftPair GiftPair) String() string {
	return fmt.Sprintf("{Givee: %s, Giver: %s}", giftPair.Givee, giftPair.Giver)
}

// GiftPairJsonStringToGiftPair turns a GiftPair JSON string into a GiftPair
func GiftPairJsonStringToGiftPair(gpString string) (GiftPair, error) {
	var giftPair GiftPair
	err := json.Unmarshal([]byte(gpString), &giftPair)
	return giftPair, err
}
