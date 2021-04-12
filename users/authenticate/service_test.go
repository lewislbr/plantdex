package authenticate

import (
	"errors"
	"testing"

	"users/generate"
	"users/user"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticate(t *testing.T) {
	t.Run("should error when there are missing fields", func(t *testing.T) {
		t.Parallel()

		f := &MockFinder{}

		f.On("FindOne", mock.AnythingOfType("string")).Return(user.User{}, nil)

		gs := generate.NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V")
		ns := NewService(gs, f)
		creds := user.Credentials{
			Email: "test@test.com",
		}
		token, err := ns.Authenticate(creds)

		require.Empty(t, token)
		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should error when user does not exist", func(t *testing.T) {
		t.Parallel()

		f := &MockFinder{}

		f.On("FindOne", mock.AnythingOfType("string")).Return(user.User{}, errors.New("user not found"))

		gs := generate.NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V")
		ns := NewService(gs, f)
		creds := user.Credentials{
			Email:    "test@test.com",
			Password: "123",
		}
		token, err := ns.Authenticate(creds)

		require.Empty(t, token)
		require.EqualError(t, err, user.ErrNotFound.Error())
	})

	t.Run("should error when password is incorrect", func(t *testing.T) {
		t.Parallel()

		f := &MockFinder{}
		password := "123"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

		f.On("FindOne", mock.AnythingOfType("string")).Return(
			user.User{
				ID:    "1",
				Name:  "test",
				Email: "test@test.com",
				Hash:  string(hash),
			},
			nil,
		)

		gs := generate.NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V")
		ns := NewService(gs, f)
		creds := user.Credentials{
			Email:    "test@test.com",
			Password: "321",
		}
		token, err := ns.Authenticate(creds)

		require.Empty(t, token)
		require.EqualError(t, err, user.ErrInvalidPassword.Error())
	})

	t.Run("should return a token on correct authentication", func(t *testing.T) {
		t.Parallel()

		f := &MockFinder{}
		password := "123"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

		f.On("FindOne", mock.AnythingOfType("string")).Return(
			user.User{
				ID:    "123",
				Name:  "test",
				Email: "test@test.com",
				Hash:  string(hash),
			},
			nil,
		)

		gs := generate.NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V")
		ns := NewService(gs, f)
		creds := user.Credentials{
			Email:    "test@test.com",
			Password: password,
		}
		token, err := ns.Authenticate(creds)

		require.NotEmpty(t, token)
		require.NoError(t, err)
	})
}
