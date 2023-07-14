package passwords

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brokentari/klefki/service/v2/api/session"
	"github.com/brokentari/klefki/service/v2/db"
	"github.com/gofrs/uuid/v5"
	log "github.com/sirupsen/logrus"
)

type PasswordHandler struct {
	DB *db.Database
	Session *session.SessionManager
}

func (p PasswordHandler) HandleAddPassword(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("userId")
	user_uuid, err := uuid.FromString(user_id)
	if err != nil {
		log.Errorln("failed to convert uuid to string: ", err.Error())
		http.Error(w, "invalid user ID", http.StatusBadRequest)
	}

	var passParams db.CreatePasswordParams

	err = json.NewDecoder(r.Body).Decode(&passParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, _ := uuid.NewV4()
	passParams.ID = u
	passParams.UserID = user_uuid

	_, err = p.DB.Queries.CreatePassword(r.Context(), passParams)
	if err != nil {
		log.Errorln("error adding password: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (p PasswordHandler) HandleGetUserPasswords(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("userId")
	user_uuid, err := uuid.FromString(user_id)
	if err != nil {
		log.Errorln("failed to convert uuid to string: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user_passwords, err := p.DB.Queries.GetUserPasswords(r.Context(), user_uuid)
	if err != nil {
		log.Errorln("failed to retrieve user passwords: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, pass := range user_passwords {
		fmt.Printf("password %d: %+v", i, pass)
	}

	user_passwords_json, err := json.Marshal(user_passwords)
	if err != nil {
		log.Error("failed to marshal the retrieved passwords")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(user_passwords) <= 0 {
		w.Write([]byte("[]"))
	} else {
		w.Write(user_passwords_json)
	}

	
}