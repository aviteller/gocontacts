package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/avi/go-contacts/models"
	u "github.com/avi/go-contacts/utils"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	contact.UserId = user
	res := contact.Create()
	u.Respond(w, res)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "there was an error in your request"))
		return
	}

	data := models.GetContacts(uint(id))
	res := u.Message(true, "success")
	res["data"] = data
	u.Respond(w, res)
}
