package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/OleksandrZhurba-san/eve-discord-bot/config"
	"github.com/OleksandrZhurba-san/eve-discord-bot/internal/auth"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if code == "" {
		http.Error(w, "Missing code param", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received code: %s\n", code)
	fmt.Printf("Received state: %s\n", state)

	w.Write([]byte("Authorization received. You can close this tab."))
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Printf("Failed to open browser: %v\n", err)
	}
}

func main() {
	cfg := config.LoadConfig()
	fmt.Println("Using callback URL:", cfg.CallbackURL)

	state := "random-string" // TODO: replace with secure random
	scopes := "publicData"
	authUrl := auth.BuidlAuthUrl(cfg.ClientID, cfg.CallbackURL, state, scopes)
	fmt.Println("Visit URL to authorize the app...")
	openBrowser(authUrl)

	http.HandleFunc("/callback", callbackHandler)

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
