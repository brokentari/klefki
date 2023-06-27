package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"net/mail"

	"github.com/gofrs/uuid/v5"
	log "github.com/sirupsen/logrus"

	"github.com/brokentari/klefki/service/v2/api/session"
	"github.com/brokentari/klefki/service/v2/db"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB      *db.Database
	Session *session.SessionManager
}

type authComponents struct {
	password string
	salt     string
}

func (a AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if !validEmail(email) {
		http.Error(w, "Invalid email.", http.StatusBadRequest)
		return
	}

	password := r.FormValue("password")

	userToAuthenticate, err := a.DB.Queries.GetUserViaEmail(r.Context(), email)
	if err != nil {
		http.Error(w, "The email that was entered does not exist.", http.StatusUnauthorized)
		return
	}

	err = authenticateUser(userToAuthenticate, password)
	if err != nil {
		http.Error(w, "The email or password entered was incorrect.", http.StatusUnauthorized)
		return
	} else {
		token, err := generateAuthToken(16)
		if err != nil {
			http.Error(w, "Something wrong happened. Oops.", http.StatusInternalServerError)
		}
		w.Header().Set("Authorization", token)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	}

	w.WriteHeader(http.StatusAccepted)
}

func generateAuthToken(length int) (string, error) {
	token := make([]byte, length)
	_, err := rand.Read(token)
	if err != nil {
		log.Error("unable to read token variable: ", err.Error())
		return "", err
	}

	return base64.URLEncoding.EncodeToString(token)[:length], nil
}

func (a AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if !validEmail(email) {
		http.Error(w, "Invalid email.", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	password := r.FormValue("password")

	hashedAuth, err := hashWithSalt(password)
	if err != nil {
		http.Error(w, "Unable to register new user. Try again later.", http.StatusInternalServerError)
		return
	}

	u, _ := uuid.NewV4()

	newUser, err := a.DB.Queries.CreateUser(r.Context(), db.CreateUserParams{
		ID:       u,
		Name:     name,
		Email:    email,
		Password: hashedAuth.password,
		Salt:     hashedAuth.salt,
	})

	if err != nil {
		log.Errorln("error creating new user: ", err.Error())
		http.Error(w, "Unable to register new user. Try again later.", http.StatusInternalServerError)
		return
	}

	a.Session.Put(r.Context(), "user_id", newUser.ID.String())

	w.WriteHeader(http.StatusCreated)

}

func hashWithSalt(plaintext string) (authComponents, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Errorln("unable to read the salt value: ", err.Error())
		return authComponents{"", ""}, err
	}

	passwordWithSalt := append([]byte(plaintext), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)
	if err != nil {
		log.Errorln("unable to hash password: ", err.Error())
		return authComponents{"", ""}, err
	}

	return authComponents{string(hashedPassword), hex.EncodeToString(salt)}, nil
}

func authenticateUser(user db.User, passwordToCheck string) error {

	storedPassword := []byte(user.Password)
	storedSalt, err := hex.DecodeString(user.Salt)
	if err != nil {
		log.Errorln("failed to decode salt from user info")
		return err
	}

	inputPasswordWithSalt := append([]byte(passwordToCheck), storedSalt...)
	err = bcrypt.CompareHashAndPassword(storedPassword, inputPasswordWithSalt)
	if err != nil {
		return err
	}

	return nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
