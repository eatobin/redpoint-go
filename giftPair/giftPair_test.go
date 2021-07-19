package giftPair

import "testing"

var jsonStringGP = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"
var gp = GiftPair{
	Givee: "GeoHar",
	Giver: "JohLen",
}

func TestUpdateGiveeGiver(t *testing.T) {
	newGive := "NewBee"
	gpGe := GiftPair{Givee: "NewBee", Giver: "JohLen"}
	gpGr := GiftPair{Givee: "GeoHar", Giver: "NewBee"}
	gotGpEE := GiftPair.UpdateGivee(gp, newGive)
	gotGpER := GiftPair.UpdateGiver(gp, newGive)
	if gotGpEE != gpGe {
		t.Fatalf("UpdateGivee(%v, %v) == %v, want %v", gp, newGive, gotGpEE, gpGe)
	}
	if gotGpER != gpGr {
		t.Fatalf("UpdateGiver(%v, %v) == %v, want %v", gp, newGive, gotGpER, gpGr)
	}
}

func TestJsonStringToGiftPair(t *testing.T) {
	gotGP, _ := JsonStringToGiftPair(jsonStringGP)
	if gotGP != gp {
		t.Fatalf("JsonStringToGiftPair(%v) == %v, want %v", jsonStringGP, gotGP, gp)
	}
}
