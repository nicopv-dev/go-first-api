package routes

import (
	"net/http"

	"github.com/nicopv-dev/go-first-api/common"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	common.SendResponse(w, http.StatusOK, []byte("GetAllUsers"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUser"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}