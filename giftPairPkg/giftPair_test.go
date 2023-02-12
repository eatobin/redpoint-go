package giftPairPkg

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
	gotGpEE := gpA.GiftPairUpdateGivee(newGive)
	gotGpER := gpA.GiftPairUpdateGiver(newGive)
	if gotGpEE != gpGe {
		t.Fatalf("(%v) GiftPairUpdateGivee(%v) == %v, want %v", gpA, newGive, gotGpEE, gpGe)
	}
	if gotGpER != gpGr {
		t.Fatalf("(%v) GiftPairUpdateGiver(%v) == %v, want %v", gpA, newGive, gotGpER, gpGr)
	}
}

func TestJsonStringToGiftPair(t *testing.T) {
	gotGP, _ := GiftPairJsonStringToGiftPair(jsonStringGP)
	if gotGP != gpA {
		t.Fatalf("GiftPairJsonStringToGiftPair(%v) == %v, want %v", jsonStringGP, gotGP, gpA)
	}
}
