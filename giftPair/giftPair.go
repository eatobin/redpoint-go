package giftPair

import (
	"encoding/json"
	"fmt"
)

type Givee = string
type Giver = string

// A GiftPair has a Givee and a Giver
type GiftPair struct {
	Givee Givee `json:"givee"`
	Giver Giver `json:"giver"`
}

// CompareGiftPair compares two GiftPairs
func CompareGiftPair(a, b GiftPair) bool {
	if &a == &b {
		return true
	}
	if a.Givee != b.Givee || a.Giver != b.Giver {
		return false
	}
	return true
}

// UpdateGivee updates a Givee in a GiftPair
func (gp GiftPair) UpdateGivee(givee string) GiftPair {
	gp.Givee = givee
	return gp
}

// UpdateGiver updates a Giver in a GiftPair
func (gp GiftPair) UpdateGiver(giver string) GiftPair {
	gp.Giver = giver
	return gp
}

// String makes a GiftPair into a string
func (gp GiftPair) String() string {
	return fmt.Sprintf("{Givee: %s, Giver: %s}", gp.Givee, gp.Giver)
}

// JsonStringToGiftPair turns a GiftPair JSON string into a GiftPair
func JsonStringToGiftPair(gpString string) (GiftPair, error) {
	var giftPair GiftPair
	err := json.Unmarshal([]byte(gpString), &giftPair)
	return giftPair, err
}

// GiftPairToJsonString turns a GiftPair into a GiftPair JSON string
func (gp GiftPair) GiftPairToJsonString() (string, error) {
	gpByte, err := json.Marshal(gp)
	return string(gpByte), err
}
