package giftPairPkg

import (
	"testing"
)

func TestGiftPairJsonStringToGiftPair(t *testing.T) {
	t.Parallel()
	var jsonString = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"
	var giftPair1 = GiftPairStruct{Givee: "GeoHar", Giver: "JohLen"}
	got, err := GiftPairJsonStringToGiftPair(jsonString)
	want := giftPair1
	if err != nil {
		t.Fatalf("want no error for valid input, got: %v", err)
	}
	if want != got {
		t.Errorf("GiftPairJsonStringToGiftPair(%s): want %v, got %v",
			giftPair1, want, got)
	}
}

func TestGiftPairJsonStringToGiftPairInvalid(t *testing.T) {
	t.Parallel()
	var badJsonString = "{\"givee\"\"GeoHar\",\"giver\":\"JohLen\"}"
	var badJsonString2 = "{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}"
	type testCase struct {
		a    JsonStringTA
		want GiftPairStruct
	}
	testCases := []testCase{
		{a: badJsonString},
		{a: badJsonString2},
	}
	for _, tc := range testCases {
		_, err := GiftPairJsonStringToGiftPair(tc.a)
		if err == nil {
			t.Error("want error for invalid input, got nil")
		}
	}
}

func TestGiftPairUpdateGivee(t *testing.T) {
	t.Parallel()
	var giftPair1 = GiftPairStruct{Givee: "GeoHar", Giver: "JohLen"}
	got := GiftPairUpdateGivee("NewBee", giftPair1)
	want := GiftPairStruct{Givee: "NewBee", Giver: "JohLen"}
	if want != got {
		t.Errorf("GiftPairUpdateGivee(%s, %v): want %v, got %v",
			"NewBee", giftPair1, want, got)
	}
}

func TestGiftPairUpdateGiver(t *testing.T) {
	t.Parallel()
	var giftPair1 = GiftPairStruct{Givee: "GeoHar", Giver: "JohLen"}
	got := GiftPairUpdateGiver("NewBee", giftPair1)
	want := GiftPairStruct{Givee: "GeoHar", Giver: "NewBee"}
	if want != got {
		t.Errorf("GiftPairUpdateGiver(%s, %v): want %v, got %v",
			"NewBee", giftPair1, want, got)
	}
}

func TestGiftPairString(t *testing.T) {
	t.Parallel()
	var giftPair1 = GiftPairStruct{Givee: "GeoHar", Giver: "JohLen"}
	got := GiftPairString(giftPair1)
	want := "{Givee: GeoHar, Giver: JohLen}"
	if want != got {
		t.Errorf("GiftPairString(%v): want %s, got %s",
			giftPair1, want, got)
	}
}
