package giftPairPkg

import (
	"encoding/json"
	"errors"
	"fmt"
)

// A StructGiftPair has a Givee and a Giver
type StructGiftPair struct {
	Givee string `json:"givee"`
	Giver string `json:"giver"`
}

// JsonStringToGiftPair turns a JSON string into a GiftPairStruct
func JsonStringToGiftPair(jsonString string) (StructGiftPair, error) {
	var giftPair StructGiftPair
	err := json.Unmarshal([]byte(jsonString), &giftPair)
	if err != nil {
		return StructGiftPair{}, err
	}
	if giftPair.Givee == "" || giftPair.Giver == "" {
		err = errors.New("missing one or both field values")
		return StructGiftPair{}, err
	}
	return giftPair, nil
}

// UpdateGivee updates a Givee in a GiftPairStruct
func (giftPair StructGiftPair) UpdateGivee(givee string) StructGiftPair {
	giftPair.Givee = givee
	return giftPair
}

// UpdateGiver updates a Giver in a GiftPairStruct
func (giftPair StructGiftPair) UpdateGiver(giver string) StructGiftPair {
	giftPair.Giver = giver
	return giftPair
}

// String makes a GiftPairStruct into a string
func (giftPair StructGiftPair) String() string {
	return fmt.Sprintf("{Givee: %s, Giver: %s}", giftPair.Givee, giftPair.Giver)
}
