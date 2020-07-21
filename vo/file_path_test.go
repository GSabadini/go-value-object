package vo

import (
	"reflect"
	"testing"
)

func TestNewFilePath(t *testing.T) {
	type args struct {
		value string
	}

	tests := []struct {
		name    string
		args    args
		want    FilePath
		wantErr bool
	}{
		{
			name: "Test new valid file path",
			args: args{
				value: "/vo/file_path_test/",
			},
			want: FilePath{
				value: "/vo/file_path_test/",
			},
			wantErr: false,
		},
		{
			name: "Test new valid file path",
			args: args{
				value: "/tpm/file-path-test",
			},
			want: FilePath{
				value: "/tpm/file-path-test",
			},
			wantErr: false,
		},
		{
			name: "Test new valid file path",
			args: args{
				value: "invalidfilepath@.com",
			},
			wantErr: true,
		},
		{
			name: "Test new invalid file path",
			args: args{
				value: "@file-path-test@",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFilePath(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestFilePath_Equals(t *testing.T) {
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
			name: "file path value equals",
			fields: fields{
				value: "/vo/file_path_test/",
			},
			args: args{
				value: FilePath{value: "/vo/file_path_test/"},
			},
			want: true,
		},
		{
			name: "file path value equals",
			fields: fields{
				value: "/tpm/file-path-test",
			},
			args: args{
				value: FilePath{value: "/tpm/file-path-test"},
			},
			want: true,
		},
		{
			name: "file path value not equals",
			fields: fields{
				value: "/tpm/file-path-test",
			},
			args: args{
				value: FilePath{value: "/tpm/file_path_test"},
			},
			want: false,
		},
		{
			name: "file path value not equals",
			fields: fields{
				value: "/vo/file_path_test/",
			},
			args: args{
				value: FilePath{value: "/vo/file-path-test/"},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := NewFilePath(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := f.Equals(tt.args.value); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestFilePath_Value(t *testing.T) {
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
				value: "/tpm/file-path-test",
			},
			want: "/tpm/file-path-test",
		},
		{
			name: "get value",
			fields: fields{
				value: "/vo/file_path_test/",
			},
			want: "/vo/file_path_test/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := NewFilePath(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := f.Value(); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}
