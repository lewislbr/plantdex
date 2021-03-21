package server

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	u "users/user"
)

var isDevelopment = os.Getenv("MODE") == "development"

type handler struct {
	cs u.CreateService
	ns u.AuthenticateService
	zs u.AuthorizeService
	gs u.GenerateService
}

func NewHandler(cs u.CreateService, ns u.AuthenticateService, zs u.AuthorizeService, gs u.GenerateService) *handler {
	return &handler{cs, ns, zs, gs}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var new u.User

	json.NewDecoder(r.Body).Decode(&new)

	err := h.cs.Create(new)
	if err != nil {
		switch {
		case errors.Is(err, u.ErrMissingData):
			http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, u.ErrUserExists):
			http.Error(w, u.ErrUserExists.Error(), http.StatusConflict)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) LogInUser(w http.ResponseWriter, r *http.Request) {
	var cred u.Credentials

	json.NewDecoder(r.Body).Decode(&cred)

	jwt, err := h.ns.Authenticate(cred)
	if err != nil {
		switch {
		case errors.Is(err, u.ErrMissingData):
			http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, u.ErrNotFound):
			http.Error(w, u.ErrNotFound.Error(), http.StatusNotFound)

			return
		case errors.Is(err, u.ErrInvalidPassword):
			http.Error(w, u.ErrInvalidPassword.Error(), http.StatusBadRequest)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	if isDevelopment {
		w.Header().Add("Set-Cookie", "st="+jwt+"; HttpOnly; Max-Age=604800")
		w.Header().Add("Set-Cookie", "te=true; Max-Age=604800")
	} else {
		w.Header().Add("Set-Cookie", "st="+jwt+"; Domain=plantdex.app; HttpOnly; Max-Age=604800; SameSite=Strict; Secure")
		w.Header().Add("Set-Cookie", "te=true; Domain=plantdex.app; Max-Age=604800; SameSite=Strict; Secure")
	}
}

func (h *handler) AuthorizeUser(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	jwt := strings.Split(authHeader, " ")[1]
	uid, err := h.zs.Authorize(jwt)
	if err != nil {
		switch {
		case errors.Is(err, u.ErrMissingData):
			http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, u.ErrInvalidToken):
			w.WriteHeader(http.StatusUnauthorized)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	io.WriteString(w, uid)
}

func (h *handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var jwt string

	for _, cookie := range r.Cookies() {
		if cookie.Name == "st" {
			jwt = cookie.Value
		}
	}

	uid, err := h.zs.Authorize(jwt)
	if err != nil {
		switch {
		case errors.Is(err, u.ErrMissingData):
			http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, u.ErrInvalidToken):
			w.WriteHeader(http.StatusUnauthorized)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}
	jwt, err = h.gs.GenerateJWT(uid)
	if err != nil {
		switch {
		case errors.Is(err, u.ErrMissingData):
			http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	if isDevelopment {
		w.Header().Add("Set-Cookie", "st="+jwt+"; HttpOnly; Max-Age=604800")
		w.Header().Add("Set-Cookie", "te=true; Max-Age=604800")
	} else {
		w.Header().Add("Set-Cookie", "st="+jwt+"; Domain=plantdex.app; HttpOnly; Max-Age=604800; SameSite=Strict; Secure")
		w.Header().Add("Set-Cookie", "te=true; Domain=plantdex.app; Max-Age=604800; SameSite=Strict; Secure")
	}
}
