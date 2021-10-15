package giftPair

import "testing"

var jsonStringGP = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"
var gpA = GiftPair{Givee: "GeoHar", Giver: "JohLen"}
var gpB = GiftPair{Givee: "GeoHar", Giver: "JohLen"}
var gpC = GiftPair{Givee: "NotEven", Giver: "Close"}

func TestCompareGiftPair(t *testing.T) {
	if !CompareGiftPair(gpA, gpB) {
		t.Fatalf("CompareGiftPair(%v, %v) == %t, want %t", gpA, gpB, false, true)
	}
	if CompareGiftPair(gpB, gpC) {
		t.Fatalf("CompareGiftPair(%v, %v) == %t, want %t", gpB, gpC, true, false)
	}
}

func TestUpdateGiveeGiver(t *testing.T) {
	newGive := "NewBee"
	gpGe := GiftPair{Givee: "NewBee", Giver: "JohLen"}
	gpGr := GiftPair{Givee: "GeoHar", Giver: "NewBee"}
	gotGpEE := gpA.UpdateGivee(newGive)
	gotGpER := gpA.UpdateGiver(newGive)
	if gotGpEE != gpGe {
		t.Fatalf("(%v) UpdateGivee(%v) == %v, want %v", gpA, newGive, gotGpEE, gpGe)
	}
	if gotGpER != gpGr {
		t.Fatalf("(%v) UpdateGiver(%v) == %v, want %v", gpA, newGive, gotGpER, gpGr)
	}
}

func TestJsonStringToGiftPair(t *testing.T) {
	gotGP, _ := JsonStringToGiftPair(jsonStringGP)
	if gotGP != gpA {
		t.Fatalf("JsonStringToGiftPair(%v) == %v, want %v", jsonStringGP, gotGP, gpA)
	}
}

func TestGiftPairToJsonString(t *testing.T) {
	gotString, _ := gpA.GiftPairToJsonString()
	if gotString != jsonStringGP {
		t.Fatalf("(%v) JsonStringToGiftPair() == %v, want %v", gpA, gotString, jsonStringGP)
	}
}
