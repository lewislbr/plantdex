package rest

import (
	"fmt"
	"net/http"
	"os"

	"users/pkg/add"
	"users/pkg/delete"
	"users/pkg/edit"
	"users/pkg/list"
	"users/pkg/storage/inmem"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

var l = list.NewService(&inmem.Storage{})
var a = add.NewService(&inmem.Storage{})
var e = edit.NewService(&inmem.Storage{})
var d = delete.NewService(&inmem.Storage{})

// Start initializes the REST API
func Start() error {
	godotenv.Load(".users-env")

	router := httprouter.New()
	port := os.Getenv("USERS_APP_PORT")
	endpointRoot := "/users"

	router.GET(endpointRoot+"/:id", listUser(l))
	router.POST(endpointRoot+"/", addUser(a))
	router.PATCH(endpointRoot+"/:id", editUser(e))
	router.PUT(endpointRoot+"/:id", editUser(e))
	router.DELETE(endpointRoot+"/:id", deleteUser(d))

	fmt.Printf("Users API ready ✅")

	err := http.ListenAndServe(":"+port, corsMiddleware(router))
	if err != nil {
		return err
	}

	return nil
}

func corsMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PATCH, PUT, DELETE",
		)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	}
}