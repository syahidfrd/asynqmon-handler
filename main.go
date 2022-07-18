package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"log"
	"net/http"
	"os"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

type application struct {
	auth struct {
		username string
		password string
	}
}

func main() {
	asynqmonUser := getenv("ASYNQMON_USER", "admin")
	asynqmonPassword := getenv("ASYNQMON_PASSWORD", "admin")
	redisAddr := getenv("REDIS_ADDR", ":6379")

	app := new(application)
	app.auth.username = asynqmonUser
	app.auth.password = asynqmonPassword

	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/",
		RedisConnOpt: asynq.RedisClientOpt{Addr: redisAddr},
	})

	http.Handle(h.RootPath()+"/", app.basicAuth(h))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func (app *application) basicAuth(next *asynqmon.HTTPHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(app.auth.username))
			expectedPasswordHash := sha256.Sum256([]byte(app.auth.password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
