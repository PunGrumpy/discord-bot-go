package bot

import (
	"fmt"
	"net/http"
)

func (bot *Bot) StartServer() {
	http.HandleFunc("/health-check", bot.healthCheckHandler)

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
