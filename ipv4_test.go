package vo

import (
	"reflect"
	"testing"
)

func TestNewIPv4(t *testing.T) {
	type args struct {
		value string
	}

	tests := []struct {
		name    string
		args    args
		want    IPv4
		wantErr bool
	}{
		{
			name:    "Test valid IPv4",
			args:    args{value: "192.168.0.110"},
			want:    IPv4{value: "192.168.0.110"},
			wantErr: false,
		},
		{
			name:    "Test valid IPv4",
			args:    args{value: "101.101.0.101"},
			want:    IPv4{value: "101.101.0.101"},
			wantErr: false,
		},
		{
			name:    "Test invalid IPv4",
			args:    args{value: "999.999.999.999"},
			wantErr: true,
		},
		{
			name:    "Test invalid IPv4",
			args:    args{value: "ak32WKsK12SOAap1"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIPv4(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] ResultErr: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestIPv4Value(t *testing.T) {
	type fields struct {
		value string
	}

	tests := []struct {
		name   string
		fields fields
		want   IP
	}{
		{
			name:   "Return value",
			fields: fields{value: "192.168.0.110"},
			want:   IP("192.168.0.110"),
		},
		{
			name:   "Return value",
			fields: fields{value: "101.101.0.101"},
			want:   IP("101.101.0.101"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := IPv4{
				value: tt.fields.value,
			}

			if got := ip.Value(); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}