package user

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewUser(t *testing.T) {
	type args struct {
		username       string
		email          string
		password       string
		firstName      string
		lastName       string
		profilePicture string
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				username:       "test",
				email:          "hello@example.com",
				password:       "password",
				firstName:      "John",
				lastName:       "Doe",
				profilePicture: "https://example.com/profile.jpg",
			},
			want: User{
				Username:     "test",
				EmailAddress: "hello@example.com",
				Password:     "password",
				FirstName:    "John",
				LastName:     "Doe",
				ProfilePic:   "https://example.com/profile.jpg",
			},
			wantErr: false,
		},
		{
			name: "empty username",
			args: args{
				username:       "",
				email:          "hello@example.com",
				password:       "password",
				firstName:      "John",
				lastName:       "Doe",
				profilePicture: "https://example.com/profile.jpg",
			},
			want:    User{},
			wantErr: true,
		},
		{
			name: "empty email",
			args: args{
				username:       "John",
				email:          "",
				password:       "password",
				firstName:      "John",
				lastName:       "Doe",
				profilePicture: "https://example.com/profile.jpg",
			},
			want:    User{},
			wantErr: true,
		},
		{
			name: "empty password",
			args: args{
				username:       "John",
				email:          "hello@example.com",
				password:       "",
				firstName:      "John",
				lastName:       "Doe",
				profilePicture: "https://example.com/profile.jpg",
			},
			want:    User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.username, tt.args.email, tt.args.password, tt.args.firstName, tt.args.lastName, tt.args.profilePicture)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Vérifier que l'ID n'est pas nul
				if got.ID == uuid.Nil {
					t.Errorf("NewUser() got nil ID")
				}

				// Vérifier que CreatedAt est proche du temps actuel
				if time.Since(got.CreatedAt) > time.Second {
					t.Errorf("NewUser() CreatedAt is not recent enough")
				}

				// Comparer les autres champs
				if got.Username != tt.want.Username ||
					got.EmailAddress != tt.want.EmailAddress ||
					got.Password != tt.want.Password ||
					got.FirstName != tt.want.FirstName ||
					got.LastName != tt.want.LastName ||
					got.ProfilePic != tt.want.ProfilePic {
					t.Errorf("NewUser() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
