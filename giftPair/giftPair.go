package giftPair

import (
	"encoding/json"
	"errors"
	"fmt"
)

type PlayerKey = string
type Givee = PlayerKey
type Giver = PlayerKey
type JsonString = string

// A Struct has a Givee and a Giver
type Struct struct {
	Givee Givee `json:"givee"`
	Giver Giver `json:"giver"`
}

// JsonStringToGiftPair turns a JSON string into a GiftPairStruct
func JsonStringToGiftPair(jsonString JsonString) (Struct, error) {
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
func UpdateGivee(givee Givee, giftPair Struct) Struct {
	giftPair.Givee = givee
	return giftPair
}

// UpdateGiver updates a Giver in a GiftPairStruct
func UpdateGiver(giver Giver, giftPair Struct) Struct {
	giftPair.Giver = giver
	return giftPair
}

// String makes a GiftPairStruct into a string
func String(giftPair Struct) string {
	return fmt.Sprintf("{Givee: %s, Giver: %s}", giftPair.Givee, giftPair.Giver)
}
