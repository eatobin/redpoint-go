package giftPairPkg

import "testing"

var jsonString = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"
var giftPair1 = GiftPair{Givee: "GeoHar", Giver: "JohLen"}
var giftPair2 = GiftPair{Givee: "GeoHar", Giver: "JohLen"}
var giftPair3 = GiftPair{Givee: "NotEven", Giver: "Close"}

func TestCompareGiftPair(t *testing.T) {
	if !CompareGiftPair(giftPair1, giftPair2) {
		t.Fatalf("Got: %v Want: %v", false, true)
	}
	if CompareGiftPair(giftPair2, giftPair3) {
		t.Fatalf("Got: %v Want: %v", true, false)
	}
}

func TestGiftPairUpdateGivee(t *testing.T) {
	got := giftPair1.GiftPairUpdateGivee("NewBee")
	want := GiftPair{Givee: "NewBee", Giver: "JohLen"}
	if !CompareGiftPair(got, want) {
		t.Fatalf("Got: %v Want: %v", got, want)
	}
}

func TestGiftPairUpdateGiver(t *testing.T) {
	got := giftPair1.GiftPairUpdateGiver("NewBee")
	want := GiftPair{Givee: "GeoHar", Giver: "NewBee"}
	if !CompareGiftPair(got, want) {
		t.Fatalf("Got: %v Want: %v", got, want)
	}
}

func TestGiftPairJsonStringToGiftPair(t *testing.T) {
	got, _ := GiftPairJsonStringToGiftPair(jsonString)
	want := giftPair1
	if !CompareGiftPair(got, want) {
		t.Fatalf("Got: %v Want: %v", got, want)
	}
}
