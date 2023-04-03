package giftPair

import (
	"reflect"
	"testing"
)

var jsonString = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"
var giftPair1 = Struct{Givee: "GeoHar", Giver: "JohLen"}
var badJsonString = "{\"givee\"\"GeoHar\",\"giver\":\"JohLen\"}"
var badJsonString2 = "{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}"

//func TestGiftPairJsonStringToGiftPair(t *testing.T) {
//	t.Parallel()
//	got, err := JsonStringToGiftPair(jsonString)
//	want := giftPair1
//	if err != nil {
//		t.Fatalf("want no error for valid input, got: %v", err)
//	}
//	if want != got {
//		t.Errorf("JsonStringToGiftPair(%s): want %v, got %v",
//			giftPair1, want, got)
//	}
//}
//
//func TestGiftPairJsonStringToGiftPairInvalid(t *testing.T) {
//	t.Parallel()
//	type testCase struct {
//		a    JsonStringTA
//		want Struct
//	}
//	testCases := []testCase{
//		{a: badJsonString},
//		{a: badJsonString2},
//	}
//	for _, tc := range testCases {
//		_, err := JsonStringToGiftPair(tc.a)
//		if err == nil {
//			t.Error("want error for invalid input, got nil")
//		}
//	}
//}
//
//func TestGiftPairUpdateGivee(t *testing.T) {
//	t.Parallel()
//	got := UpdateGivee("NewBee", giftPair1)
//	want := Struct{Givee: "NewBee", Giver: "JohLen"}
//	if want != got {
//		t.Errorf("UpdateGivee(%s, %v): want %v, got %v",
//			"NewBee", giftPair1, want, got)
//	}
//}
//
//func TestGiftPairUpdateGiver(t *testing.T) {
//	t.Parallel()
//	got := UpdateGiver("NewBee", giftPair1)
//	want := Struct{Givee: "GeoHar", Giver: "NewBee"}
//	if want != got {
//		t.Errorf("UpdateGiver(%s, %v): want %v, got %v",
//			"NewBee", giftPair1, want, got)
//	}
//}
//
//func TestGiftPairString(t *testing.T) {
//	t.Parallel()
//	got := String(giftPair1)
//	want := "{Givee: GeoHar, Giver: JohLen}"
//	if want != got {
//		t.Errorf("String(%v): want %s, got %s",
//			giftPair1, want, got)
//	}
//}

func TestJsonStringToGiftPair(t *testing.T) {
	type args struct {
		jsonString JsonStringTA
	}
	tests := []struct {
		name    string
		args    args
		want    Struct
		wantErr bool
	}{
		{name: "ValidInput", args: args{jsonString}, want: giftPair1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonStringToGiftPair(tt.args.jsonString)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonStringToGiftPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonStringToGiftPair() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		giftPair Struct
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.giftPair); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateGivee(t *testing.T) {
	type args struct {
		givee    GiveeTA
		giftPair Struct
	}
	tests := []struct {
		name string
		args args
		want Struct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateGivee(tt.args.givee, tt.args.giftPair); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateGivee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateGiver(t *testing.T) {
	type args struct {
		giver    GiverTA
		giftPair Struct
	}
	tests := []struct {
		name string
		args args
		want Struct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateGiver(tt.args.giver, tt.args.giftPair); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateGiver() = %v, want %v", got, tt.want)
			}
		})
	}
}
