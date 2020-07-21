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
			name:    "Test new valid IPv4",
			args:    args{value: "192.168.0.110"},
			want:    IPv4{value: "192.168.0.110"},
			wantErr: false,
		},
		{
			name:    "Test new valid IPv4",
			args:    args{value: "101.101.0.101"},
			want:    IPv4{value: "101.101.0.101"},
			wantErr: false,
		},
		{
			name:    "Test new invalid IPv4",
			args:    args{value: "999.999.999.999"},
			wantErr: true,
		},
		{
			name:    "Test new invalid IPv4",
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

func TestIPv4_Equals(t *testing.T) {
	type fields struct {
		value string
	}

	type args struct {
		value Value
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "ipv4 value equals",
			fields: fields{
				value: "192.168.0.110",
			},
			args: args{
				value: IPv4{"192.168.0.110"},
			},
			want: true,
		},
		{
			name: "ipv4 value equals",
			fields: fields{
				value: "192.168.0.0",
			},
			args: args{
				value: IPv4{"192.168.0.0"},
			},
			want: true,
		},
		{
			name: "ipv4 value not equals",
			fields: fields{
				value: "192.168.1.1",
			},
			args: args{
				value: IPv4{"192.168.0.0"},
			},
			want: false,
		},
		{
			name: "ipv4 value not equals",
			fields: fields{
				value: "172.168.1.1",
			},
			args: args{
				value: IPv4{"172.168.1.0"},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip, err := NewIPv4(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := ip.Equals(tt.args.value); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestIPv4_Value(t *testing.T) {
	type fields struct {
		value string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get value",
			fields: fields{
				value: "172.168.1.1",
			},
			want: "172.168.1.1",
		},
		{
			name: "get value",
			fields: fields{
				value: "101.101.0.101",
			},
			want: "101.101.0.101",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip, err := NewIPv4(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := ip.Value(); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}
