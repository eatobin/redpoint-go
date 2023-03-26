package giftPairPkg

import (
	"testing"
)

var giftPair1 = GiftPairStruct{Givee: "GeoHar", Giver: "JohLen"}
var giftPair2 = GiftPairStruct{Givee: "GeoHar", Giver: "JohLen"}
var giftPair3 = GiftPairStruct{Givee: "NotEven", Giver: "Close"}

var jsonString = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"
var badJsonString = "{\"givee\"\"GeoHar\",\"giver\":\"JohLen\"}"
var badJsonString2 = "{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}"

func TestGiftPairAssertEqual(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b GiftPairStruct
		want bool
	}
	testCases := []testCase{
		{a: giftPair1, b: giftPair2, want: true},
		{a: giftPair1, b: giftPair3, want: false},
	}
	for _, tc := range testCases {
		got := GiftPairAssertEqual(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("GiftPairAssertEqual(%v, %v): want %t, got %t",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestGiftPairJsonStringToGiftPair(t *testing.T) {
	t.Parallel()
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

//func TestGiftPairUpdateGivee(t *testing.T) {
//	t.Parallel()
//	got := giftPair1.GiftPairUpdateGivee("NewBee")
//	want := GiftPairStruct{Givee: "NewBee", Giver: "JohLen"}
//	if !GiftPairAssertEqual(got, want) {
//		t.Fatalf("Got: %v Want: %v", got, want)
//	}
//}
//
//func TestGiftPairUpdateGiver(t *testing.T) {
//	t.Parallel()
//	got := giftPair1.GiftPairUpdateGiver("NewBee")
//	want := GiftPairStruct{Givee: "GeoHar", Giver: "NewBee"}
//	if !GiftPairAssertEqual(got, want) {
//		t.Fatalf("Got: %v Want: %v", got, want)
//	}
//}
