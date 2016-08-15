package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/question-service/config"
)

type Server struct {
	*http.Server
	Config config.Config
}

func NewServer(conf config.Config) *Server {
	r := mux.NewRouter()

	r.HandleFunc("/", health)
	r.HandleFunc("/question", question)

	s := &Server{
		Server: &http.Server{
			Handler: r,
			Addr:    conf.Addr,
		},
		Config:       conf,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return s
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func question(w http.ResponseWriter, r *http.Request) {
	question, err := logic.GetNewQuestion()
	if err != nil {
		fmt.Printf("error: could not get new question (%+v)\n", err)
		ServeError(w, NewErrorResponse(500, "Could not get new question"))
	}
	ServeJSON(w, question)
}
