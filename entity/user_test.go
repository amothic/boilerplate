package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Name(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "user",
			want: "user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				name: tt.name,
			}
			assert.Equal(t, tt.want, u.Name())
		})
	}
}
