package giftPairPkg

import (
	"reflect"
	"testing"
)

var jsonString = "{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}"
var giftPair1 = StructGiftPair{Givee: "GeoHar", Giver: "JohLen"}
var badJsonString = "{\"givee\"\"GeoHar\",\"giver\":\"JohLen\"}"
var badJsonString2 = "{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}"

func TestJsonStringToGiftPair(t *testing.T) {
	t.Parallel()
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		want    StructGiftPair
		wantErr bool
	}{
		{name: "ValidInput", args: args{jsonString}, want: giftPair1, wantErr: false},
		{name: "InvalidInputBadJSON", args: args{badJsonString}, want: StructGiftPair{}, wantErr: true},
		{name: "InvalidInputBadFieldName", args: args{badJsonString2}, want: StructGiftPair{}, wantErr: true},
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

func TestStructGiftPair_UpdateGivee(t *testing.T) {
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
		want   StructGiftPair
	}{
		{
			name:   "ValidInput",
			fields: fields(giftPair1),
			args:   args{givee: "NewBee"},
			want:   StructGiftPair{Givee: "NewBee", Giver: "JohLen"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			giftPair := StructGiftPair{
				Givee: tt.fields.Givee,
				Giver: tt.fields.Giver,
			}
			if got := giftPair.UpdateGivee(tt.args.givee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateGivee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructGiftPair_UpdateGiver(t *testing.T) {
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
		want   StructGiftPair
	}{
		{
			name:   "ValidInput",
			fields: fields(giftPair1),
			args:   args{giver: "NewBee"},
			want:   StructGiftPair{Givee: "GeoHar", Giver: "NewBee"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			giftPair := StructGiftPair{
				Givee: tt.fields.Givee,
				Giver: tt.fields.Giver,
			}
			if got := giftPair.UpdateGiver(tt.args.giver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateGiver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructGiftPair_String(t *testing.T) {
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
			fields: fields(giftPair1),
			want:   "{Givee: GeoHar, Giver: JohLen}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			giftPair := StructGiftPair{
				Givee: tt.fields.Givee,
				Giver: tt.fields.Giver,
			}
			if got := giftPair.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
