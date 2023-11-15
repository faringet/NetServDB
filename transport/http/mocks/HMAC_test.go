package mocks

import (
	_ "github.com/stretchr/testify/mock"
	"testing"
)

func TestHMAC_SignHMACSHA512(t *testing.T) {
	mock := &HMAC{}

	type fields struct {
		Mock *HMAC
	}
	type args struct {
		text string
		key  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ValidSignature",
			args: args{
				text: "test",
				key:  "797784",
			},
			want:    "5ea669064696ab961a959c5aa098cf50e4854717470cde6bdff3eaf1777dc1bb707ce048fc7c5799d108c6b53114aa01ad20c85fba04d5aefa1442812799ce2f",
			wantErr: false,
		},

		{
			name: "ValidSignature2",
			args: args{
				text: "test",
				key:  "-14",
			},
			want:    "1d5e0b0f2d4b1e730df0688f66120b61a20d0438cba6978a2e6fb9b1ae4df2ab63ed24076899fe7185e241e14c28a98bc0412eaed64818bc5dbfe5f008a4ac89",
			wantErr: false,
		},

		{
			name: "ValidSignature3",
			args: args{
				text: "test",
				key:  "",
			},
			want:    "",
			wantErr: true,
		},

		{
			name: "ValidSignature4",
			args: args{
				text: "",
				key:  "7964",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.On("SignHMACSHA512", tt.args.text, tt.args.key).Return(tt.want, nil)

			got, err := mock.SignHMACSHA512(tt.args.text, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignHMACSHA512() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignHMACSHA512() got = %v, want %v", got, tt.want)
			}
			if tt.args.key == "" {
				if err == nil {
					t.Errorf("Expected error for empty key, got nil")
					return
				}
				return
			}
			if tt.args.text == "" {
				if err == nil {
					t.Errorf("Expected error for empty text, got nil")
					return
				}
				return
			}

			mock.AssertExpectations(t)
		})
	}
}
