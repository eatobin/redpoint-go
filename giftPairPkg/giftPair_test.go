package giftPairPkg

import "testing"

var giftPair1 = GiftPair{Givee: "GeoHar", Giver: "JohLen"}
var giftPair2 = GiftPair{Givee: "GeoHar", Giver: "JohLen"}
var giftPair3 = GiftPair{Givee: "NotEven", Giver: "Close"}
var jsonString = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"

func TestGiftPairAssertEqualTrue(t *testing.T) {
	t.Parallel()
	got := GiftPairAssertEqual(giftPair1, giftPair2)
	if got != true {
		t.Fatalf("Got: %v Want: %v", got, true)
	}
}

func TestGiftPairAssertEqualFalse(t *testing.T) {
	t.Parallel()
	got := GiftPairAssertEqual(giftPair1, giftPair3)
	if got != false {
		t.Fatalf("Got: %v Want: %v", got, false)
	}
}

func TestGiftPairJsonStringToGiftPair(t *testing.T) {
	t.Parallel()
	got, _ := GiftPairJsonStringToGiftPair(jsonString)
	want := giftPair1
	if !GiftPairAssertEqual(got, want) {
		t.Fatalf("Got: %v Want: %v", got, want)
	}
}

func TestGiftPairUpdateGivee(t *testing.T) {
	t.Parallel()
	got := giftPair1.GiftPairUpdateGivee("NewBee")
	want := GiftPair{Givee: "NewBee", Giver: "JohLen"}
	if !GiftPairAssertEqual(got, want) {
		t.Fatalf("Got: %v Want: %v", got, want)
	}
}

func TestGiftPairUpdateGiver(t *testing.T) {
	t.Parallel()
	got := giftPair1.GiftPairUpdateGiver("NewBee")
	want := GiftPair{Givee: "GeoHar", Giver: "NewBee"}
	if !GiftPairAssertEqual(got, want) {
		t.Fatalf("Got: %v Want: %v", got, want)
	}
}
