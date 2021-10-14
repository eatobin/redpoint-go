package player

//import "testing"
//
//var jsonStringGP = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"
//var gp = GiftPair{Givee: "GeoHar", Giver: "JohLen"}
//
//func TestUpdateGiveeGiver(t *testing.T) {
//	newGive := "NewBee"
//	gpGe := GiftPair{Givee: "NewBee", Giver: "JohLen"}
//	gpGr := GiftPair{Givee: "GeoHar", Giver: "NewBee"}
//	gotGpEE := gp.UpdateGivee(newGive)
//	gotGpER := gp.UpdateGiver(newGive)
//	if gotGpEE != gpGe {
//		t.Fatalf("(%v) UpdateGivee(%v) == %v, want %v", gp, newGive, gotGpEE, gpGe)
//	}
//	if gotGpER != gpGr {
//		t.Fatalf("(%v) UpdateGiver(%v) == %v, want %v", gp, newGive, gotGpER, gpGr)
//	}
//}
//
//func TestJsonStringToGiftPair(t *testing.T) {
//	gotGP, _ := JsonStringToGiftPair(jsonStringGP)
//	if gotGP != gp {
//		t.Fatalf("JsonStringToGiftPair(%v) == %v, want %v", jsonStringGP, gotGP, gp)
//	}
//}
//
//func TestGiftPairToJsonString(t *testing.T) {
//	gotString, _ := gp.GiftPairToJsonString()
//	if gotString != jsonStringGP {
//		t.Fatalf("(%v) JsonStringToGiftPair() == %v, want %v", gp, gotString, jsonStringGP)
//	}
//}
