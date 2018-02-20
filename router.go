package teledisq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func SetupRouter() {
	http.HandleFunc("/", healthHandler)
	http.HandleFunc(fmt.Sprintf("/hook/test/"), hookHandler)
	http.HandleFunc(fmt.Sprintf("/telegram/%s/", os.Getenv("TELEGRAM_WEBHOOK")), telegramHandler)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "All your base are belong to us!")
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func telegramHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	u := Update{}
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		log.Errorf(ctx, "Telegram update decoding error: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HandleTelegramUpdate(ctx, u)
	w.WriteHeader(http.StatusOK)
}
