package giftPair

import (
	"encoding/json"
	"errors"
	"fmt"
)

type PlayerKeyTA = string
type GiveeTA = PlayerKeyTA
type GiverTA = PlayerKeyTA
type JsonStringTA = string

// A Struct has a Givee and a Giver
type Struct struct {
	Givee GiveeTA `json:"givee"`
	Giver GiverTA `json:"giver"`
}

// JsonStringToGiftPair turns a JSON string into a GiftPairStruct
func JsonStringToGiftPair(jsonString JsonStringTA) (Struct, error) {
	var giftPair Struct
	err := json.Unmarshal([]byte(jsonString), &giftPair)
	if err != nil {
		return Struct{}, err
	}
	if giftPair.Givee == "" || giftPair.Giver == "" {
		err = errors.New("missing one or both field values")
		return Struct{}, err
	}
	return giftPair, nil
}

// UpdateGivee updates a Givee in a GiftPairStruct
func UpdateGivee(givee GiveeTA, giftPair Struct) Struct {
	giftPair.Givee = givee
	return giftPair
}

// UpdateGiver updates a Giver in a GiftPairStruct
func UpdateGiver(giver GiverTA, giftPair Struct) Struct {
	giftPair.Giver = giver
	return giftPair
}

// String makes a GiftPairStruct into a string
func String(giftPair Struct) string {
	return fmt.Sprintf("{Givee: %s, Giver: %s}", giftPair.Givee, giftPair.Giver)
}
