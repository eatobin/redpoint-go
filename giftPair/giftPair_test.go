package giftPair

import "testing"

var badGivee, _ = JsonStringToGiftPair("{\"givee\":\"GeoHaX\",\"giver\":\"JohLen\"}")
var badGiver, _ = JsonStringToGiftPair("{\"givee\":\"GeoHar\",\"giver\":\"JohLeX\"}")
var wantGP = "[Givee: GeoHar, Giver: JohLen]"
var jsonString = "{\"givee\":\"GeoHaX\",\"giver\":\"JohLen\"}"

func TestUpdateGiveeGiver(t *testing.T) {
	goodGivee := "GeoHar"
	goodGiver := "JohLen"
	gotGpEE := badGivee.UpdateGivee(goodGivee).String()
	gotGpER := badGiver.UpdateGiver(goodGiver).String()
	if gotGpEE != wantGP {
		t.Fatalf("(%v) UpdateGivee(%v) == %v, want %v", badGivee, goodGivee, gotGpEE, wantGP)
	}
	if gotGpER != wantGP {
		t.Fatalf("(%v) UpdateGiver(%v) == %v, want %v", badGiver, goodGiver, gotGpER, wantGP)
	}
}

func TestGiftPairToJsonString(t *testing.T) {
	gotJsonString, _ := badGivee.GiftPairToJsonString()
	if gotJsonString != jsonString {
		t.Fatalf("(%v) JsonStringToGiftPair() == %v, want %v", badGivee, gotJsonString, jsonString)
	}
}
