package giftHistory

import (
	"testing"
)

//var jsonStringGH = "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
var giftHistory = GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
var gotGH = addYear("NewBee", giftHistory)
var giftHistoryExtended = GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b GiftHistory) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestAddYear(t *testing.T) {
	if !equal(gotGH, giftHistoryExtended) {
		t.Fatalf("AddYear(%s, %v) == %v, want %v", "NewBee", giftHistory, gotGH, giftHistoryExtended)
	}
}
