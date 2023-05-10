package bot

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 5) // allows 1 request per second with a burst of 5

func rateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next(w, r)
	}
}

func (bot *Bot) StartServer() {
	http.HandleFunc("/health-check", rateLimit(bot.healthCheckHandler))

	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func (bot *Bot) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if bot.Status == "up" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "up"}`))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"status": "down"}`))
	}
}
