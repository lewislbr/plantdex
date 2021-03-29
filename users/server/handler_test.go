package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"users/authenticate"
	"users/authorize"
	"users/create"
	"users/generate"
	"users/user"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	t.Run("Create should return 201 after successful request", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		csMock.On("Create", mock.AnythingOfType("user.User")).Return(nil)

		user := &user.User{Name: "test", Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(payload))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusCreated, w.Result().StatusCode)
	})

	t.Run("Create should return 400 if required data is missing", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		csMock.On("Create", mock.AnythingOfType("user.User")).Return(user.ErrMissingData)

		user := &user.User{Name: "test", Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(payload))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("Create should return 409 if a user already exists", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		csMock.On("Create", mock.AnythingOfType("user.User")).Return(user.ErrUserExists)

		user := &user.User{Name: "test", Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(payload))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusConflict, w.Result().StatusCode)
	})

	t.Run("Create should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		csMock.On("Create", mock.AnythingOfType("user.User")).Return(errors.New("error"))

		user := &user.User{Name: "test", Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(payload))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("LogIn should return 200 if the request is successful", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		nsMock.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", nil)

		user := &user.Credentials{Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(payload))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("LogIn should return 400 if required data is missing", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		nsMock.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", user.ErrMissingData)

		user := &user.Credentials{Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(payload))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("LogIn should return 404 if the user is not found", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		nsMock.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", user.ErrNotFound)

		user := &user.Credentials{Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(payload))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("LogIn should return 400 if the password is invalid", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		nsMock.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", user.ErrInvalidPassword)

		user := &user.Credentials{Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(payload))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("LogIn should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		nsMock.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", errors.New("error"))

		user := &user.Credentials{Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(payload))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Authorize should return 200 if the request is successful", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		zsMock.On("Authorize", mock.AnythingOfType("string")).Return("", nil)

		r := httptest.NewRequest(http.MethodGet, "/authorize", nil)
		w := httptest.NewRecorder()

		r.Header.Add("Authorization", "Bearer test")

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("Authorize should return 400 if the Authorization header is missing", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		zsMock.On("Authorize", mock.AnythingOfType("string")).Return("", nil)

		r := httptest.NewRequest(http.MethodGet, "/authorize", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("Authorize should return 400 if the Authorization header is empty", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		zsMock.On("Authorize", mock.AnythingOfType("string")).Return("", user.ErrMissingData)

		r := httptest.NewRequest(http.MethodGet, "/authorize", nil)
		w := httptest.NewRecorder()

		r.Header.Add("Authorization", "Bearer test")

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("Authorize should return 401 if the token is invalid", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		zsMock.On("Authorize", mock.AnythingOfType("string")).Return("", user.ErrInvalidToken)

		r := httptest.NewRequest(http.MethodGet, "/authorize", nil)
		w := httptest.NewRecorder()

		r.Header.Add("Authorization", "Bearer test")

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("Authorize should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		zsMock.On("Authorize", mock.AnythingOfType("string")).Return("", errors.New("error"))

		r := httptest.NewRequest(http.MethodGet, "/authorize", nil)
		w := httptest.NewRecorder()

		r.Header.Add("Authorization", "Bearer test")

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Refresh should return 200 if the request is successful", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		zsMock.On("Authorize", mock.AnythingOfType("string")).Return("", nil)
		gsMock.On("GenerateJWT", mock.AnythingOfType("string")).Return("", nil)

		r := httptest.NewRequest(http.MethodGet, "/refresh", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("Refresh should return 400 if the token is missing", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		zsMock.On("Authorize", mock.AnythingOfType("string")).Return("", user.ErrMissingData)
		gsMock.On("GenerateJWT", mock.AnythingOfType("string")).Return("", nil)

		r := httptest.NewRequest(http.MethodGet, "/refresh", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("Refresh should return 401 if the token is invalid", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		zsMock.On("Authorize", mock.AnythingOfType("string")).Return("", user.ErrInvalidToken)
		gsMock.On("GenerateJWT", mock.AnythingOfType("string")).Return("", nil)

		r := httptest.NewRequest(http.MethodGet, "/refresh", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("Refresh should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		csMock := &create.MockCreateService{}
		nsMock := &authenticate.MockAuthenticateService{}
		zsMock := &authorize.MockAuthorizeService{}
		gsMock := &generate.MockGenerateService{}
		handler := NewHandler(csMock, nsMock, zsMock, gsMock)
		router := setUpRouter(handler)

		zsMock.On("Authorize", mock.AnythingOfType("string")).Return("", errors.New("error"))
		gsMock.On("GenerateJWT", mock.AnythingOfType("string")).Return("", nil)

		r := httptest.NewRequest(http.MethodGet, "/refresh", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
