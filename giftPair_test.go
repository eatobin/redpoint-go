package redpoint

import (
	"reflect"
	"testing"
)

var jsonStringGP = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"
var giftPair1GP = GiftPairStruct{Givee: "GeoHar", Giver: "JohLen"}
var badJsonStringGP = "{\"givee\"\"GeoHar\",\"giver\":\"JohLen\"}"
var badJsonString2GP = "{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}"

func TestGiftPairJsonStringToGiftPair(t *testing.T) {
	t.Parallel()
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		want    GiftPairStruct
		wantErr bool
	}{
		{name: "ValidInput", args: args{jsonStringGP}, want: giftPair1GP, wantErr: false},
		{name: "InvalidInputBadJSON", args: args{badJsonStringGP}, want: GiftPairStruct{}, wantErr: true},
		{name: "InvalidInputBadFieldName", args: args{badJsonString2GP}, want: GiftPairStruct{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GiftPairJsonStringToGiftPair(tt.args.jsonString)
			if (err != nil) != tt.wantErr {
				t.Errorf("GiftPairJsonStringToGiftPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftPairJsonStringToGiftPair() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructGiftPair_GiftPairUpdateGivee(t *testing.T) {
	t.Parallel()
	type fields struct {
		Givee string
		Giver string
	}
	type args struct {
		givee string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   GiftPairStruct
	}{
		{
			name:   "ValidInput",
			fields: fields(giftPair1GP),
			args:   args{givee: "NewBee"},
			want:   GiftPairStruct{Givee: "NewBee", Giver: "JohLen"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			giftPair := GiftPairStruct{
				Givee: tt.fields.Givee,
				Giver: tt.fields.Giver,
			}
			if got := giftPair.GiftPairUpdateGivee(tt.args.givee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftPairUpdateGivee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructGiftPair_GiftPairUpdateGiver(t *testing.T) {
	t.Parallel()
	type fields struct {
		Givee string
		Giver string
	}
	type args struct {
		giver string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   GiftPairStruct
	}{
		{
			name:   "ValidInput",
			fields: fields(giftPair1GP),
			args:   args{giver: "NewBee"},
			want:   GiftPairStruct{Givee: "GeoHar", Giver: "NewBee"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			giftPair := GiftPairStruct{
				Givee: tt.fields.Givee,
				Giver: tt.fields.Giver,
			}
			if got := giftPair.GiftPairUpdateGiver(tt.args.giver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftPairUpdateGiver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructGiftPair_GiftPairString(t *testing.T) {
	t.Parallel()
	type fields struct {
		Givee string
		Giver string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "ValidInput",
			fields: fields(giftPair1GP),
			want:   "{Givee: GeoHar, Giver: JohLen}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			giftPair := GiftPairStruct{
				Givee: tt.fields.Givee,
				Giver: tt.fields.Giver,
			}
			if got := giftPair.GiftPairString(); got != tt.want {
				t.Errorf("GiftPairString() = %v, want %v", got, tt.want)
			}
		})
	}
}
