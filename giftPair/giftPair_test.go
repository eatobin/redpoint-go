package giftPair

import "testing"

//var jsonStringGP = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"
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

//func TestSetMaxBooks(t *testing.T) {
//	maxBooks := 1
//	gotBrMB := Borrower.String(Borrower.SetMaxBooks(badMB, maxBooks))
//	if gotBrMB != wantBr {
//		t.Fatalf("SetMaxBooks(%v, %v) == %v, want %v", badMB, maxBooks, gotBrMB, wantBr)
//	}
//}
//
//func TestBrToJsonString(t *testing.T) {
//	gotJsonString, _ := Borrower.BrToJsonString(badName)
//	if gotJsonString != jsonString {
//		t.Fatalf("BrToJsonString(%v) == %v, want %v", badName, gotJsonString, jsonString)
//	}
//}
