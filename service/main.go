package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/brokentari/klefki/service/v2/api"
	"github.com/brokentari/klefki/service/v2/api/session"
	"github.com/brokentari/klefki/service/v2/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Errorln("unable to read environment variables")
	}

	db := db.GetDB()
	sess := session.GetSession()
	r := api.InitRouter(http.Dir("../build"), db, sess)

	listenAddr := fmt.Sprintf(":%s", os.Getenv("KLEFKI_PORT"))
	log.Infoln("server listening on port " + listenAddr)
	if err := http.ListenAndServe(listenAddr, sess.LoadAndSave(r)); err != nil {
		log.Errorln("failed to initialize api server")
	}

	http.ListenAndServe(":3000", r)
}

