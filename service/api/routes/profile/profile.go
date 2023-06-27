package profile

import (
	"net/http"

	"github.com/brokentari/klefki/service/v2/api/session"
	"github.com/brokentari/klefki/service/v2/db"
)

type ProfileHandler struct {
	DB      *db.Database
	Session *session.SessionManager
}

func (p ProfileHandler) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	user_id := p.Session.GetString(r.Context(), "user_id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user_id))
}
