package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewRouter(motion *bool) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)

	r.HandleFunc("/watch", watchGetHandler).Methods("GET")
	r.HandleFunc("/watch", motionMiddleware(motion)).Methods("POST")
	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/watch", http.StatusFound)
}

func watchGetHandler(w http.ResponseWriter, r *http.Request) {
	/*
		w.Header().Set("Content-Type", "application/json")
		var buffer bytes.Buffer
		buffer.WriteString(`{"Response": "success", "Message": "Welcome to awesome Go service"}`)
		json.NewEncoder(w).Encode(buffer.String())
		w.WriteHeader(http.StatusOK)
		return
	*/
	mypost := Post{"15", "tit", "myBody"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mypost)
}

func motionMiddleware(motion *bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		move := r.PostForm.Get("move")
		fmt.Println(move)
		if move == "ok" {
			*motion = true
			// tempBot := bot.CreateBot(secrets.GetBotToken())
			// tempBot.MotionMessage()
		}

	}

}

/*
func watchPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	move := r.PostForm.Get("move")
	fmt.Println(move)
	if move == "ok" {
		//*motion = true
	}
}
*/
