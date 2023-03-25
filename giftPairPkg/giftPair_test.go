package giftPairPkg

import (
	"testing"
)

var giftPair1 = GiftPair{Givee: "GeoHar", Giver: "JohLen"}
var giftPair2 = GiftPair{Givee: "GeoHar", Giver: "JohLen"}
var giftPair3 = GiftPair{Givee: "NotEven", Giver: "Close"}

//var jsonString = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"

//var badJsonString = "{\"givee\"\"GeoHar\",\"giver\":\"JohLen\"}"

//var badJsonString2 = "{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}"

func TestGiftPairAssertEqual(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b GiftPair
		want bool
	}
	testCases := []testCase{
		{a: giftPair1, b: giftPair2, want: true},
		{a: giftPair1, b: giftPair3, want: false},
	}
	for _, tc := range testCases {
		got := GiftPairAssertEqual(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("GiftPairAssertEqual(%v, %v): want %v, got %v",
				tc.a, tc.b, tc.want, got)
		}
	}
}

//func TestGiftPairAssertEqualTrue(t *testing.T) {
//	t.Parallel()
//	got := GiftPairAssertEqual(giftPair1, giftPair2)
//	if got != true {
//		t.Fatalf("Got: %v Want: %v", got, true)
//	}
//}
//
//func TestGiftPairAssertEqualFalse(t *testing.T) {
//	t.Parallel()
//	got := GiftPairAssertEqual(giftPair1, giftPair3)
//	if got != false {
//		t.Fatalf("Got: %v Want: %v", got, false)
//	}
//}
//
//func TestGiftPairJsonStringToGiftPair(t *testing.T) {
//	t.Parallel()
//	got, _ := GiftPairJsonStringToGiftPair(jsonString)
//	want := giftPair1
//	if !GiftPairAssertEqual(got, want) {
//		t.Fatalf("Got: %v Want: %v", got, want)
//	}
//}

//func TestGiftPairJsonStringToGiftPairFromBADJSON(t *testing.T) {
//	t.Parallel()
//	got, err := GiftPairJsonStringToGiftPair(jsonString)
//	want := giftPair1
//	if err != nil {
//		t.Fatalf("want no error for valid input, got: %v", err)
//	}
//	if want != got {
//		t.Errorf("%v, %v", got, want)
//	}
//}

//func Test_jsonStringToBorrowersPass(t *testing.T) {
//	js := jsonStringBorrowers
//	wantBrs := brs1
//	wantError := error(nil)
//
//	got, err := JsonStringToBorrowers(js)
//	if !reflect.DeepEqual(got, wantBrs) || err != wantError {
//		t.Errorf("JsonStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
//			js, got, err, wantBrs, wantError)
//	}
//}

//func Test_jsonStringToBorrowersFail(t *testing.T) {
//	cases := []struct {
//		js        string
//		wantBrs   Borrowers
//		wantError error
//	}{
//		{jsonStringBorrowersBadParse, ZeroBorrowers, errors.New("invalid character '\"' after object key")},
//		{jsonStringBorrowersBadNameField, ZeroBorrowers, errors.New("missing Borrower field value - borrowers list is empty")},
//		{jsonStringBorrowersBadMaxBooksField, ZeroBorrowers, errors.New("missing Borrower field value - borrowers list is empty")},
//	}
//	for _, c := range cases {
//		got, err := JsonStringToBorrowers(c.js)
//		if err != nil {
//			if err.Error() != c.wantError.Error() {
//				t.Errorf("JsonStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
//					c.js, got, err, c.wantBrs, c.wantError)
//			}
//		}
//	}
//}

//func TestGiftPairUpdateGivee(t *testing.T) {
//	t.Parallel()
//	got := giftPair1.GiftPairUpdateGivee("NewBee")
//	want := GiftPair{Givee: "NewBee", Giver: "JohLen"}
//	if !GiftPairAssertEqual(got, want) {
//		t.Fatalf("Got: %v Want: %v", got, want)
//	}
//}
//
//func TestGiftPairUpdateGiver(t *testing.T) {
//	t.Parallel()
//	got := giftPair1.GiftPairUpdateGiver("NewBee")
//	want := GiftPair{Givee: "GeoHar", Giver: "NewBee"}
//	if !GiftPairAssertEqual(got, want) {
//		t.Fatalf("Got: %v Want: %v", got, want)
//	}
//}
