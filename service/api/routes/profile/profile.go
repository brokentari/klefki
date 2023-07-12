package profile

import (
	"encoding/json"
	"net/http"

	"github.com/brokentari/klefki/service/v2/api/session"
	"github.com/brokentari/klefki/service/v2/db"
	"github.com/gofrs/uuid/v5"
	log "github.com/sirupsen/logrus"
)

type ProfileHandler struct {
	DB      *db.Database
	Session *session.SessionManager
}

func (p ProfileHandler) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("userId")
	user_uuid, err := uuid.FromString(user_id)
	if err != nil {
		log.Error("failed to create uuid from string: ", err.Error())
		http.Error(w, "invalid uuid value", http.StatusBadRequest)
	}

	fetchedAuthor, err := p.DB.Queries.GetUser(r.Context(), user_uuid)
	if err != nil {
		log.Error(w, "failed to retrieve user: ", err.Error())
		http.Error(w, "unable to find user", http.StatusNotFound)
	}
	marshalledUser, err := json.Marshal(fetchedAuthor)
	if err != nil {
		log.Error("failed to marshal user: ", err.Error())
		http.Error(w, "unable to marshal user", http.StatusInternalServerError)
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(marshalledUser))
}
