package giftPair

import (
	"encoding/json"
	"fmt"
)

// A GiftPair has a Givee and a Giver
type GiftPair struct {
	Givee string `json:"givee"`
	Giver string `json:"giver"`
}

// JsonStringToGiftPair turns a GiftPair JSON string into a GiftPair
func JsonStringToGiftPair(gpString string) (GiftPair, error) {
	var giftPair GiftPair
	err := json.Unmarshal([]byte(gpString), &giftPair)
	return giftPair, err
}

// SetGivee updates a Givee in a GiftPair
func (gp GiftPair) SetGivee(givee string) GiftPair {
	gp.Givee = givee
	return gp
}

// SetGiver updates a Giver in a GiftPair
func (gp GiftPair) SetGiver(giver string) GiftPair {
	gp.Giver = giver
	return gp
}

// String makes a GiftPair into a string
func (gp GiftPair) String() string {
	return fmt.Sprintf("[Givee: %s, Giver: %s]", gp.Givee, gp.Giver)
}

// GiftPairToJsonString turns a GiftPair into a GiftPair JSON string
func (gp GiftPair) GiftPairToJsonString() (string, error) {
	gpByte, err := json.Marshal(gp)
	return string(gpByte), err
}
