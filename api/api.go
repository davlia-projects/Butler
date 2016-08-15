package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/question-service/config"
	"github.com/question-service/logic"
)

type Server struct {
	*http.Server
	logic  logic.Logic
	Config config.Config
}

func NewServer(conf config.Config, logic logic.Logic) *Server {
	r := mux.NewRouter()

	s := &Server{
		Server: &http.Server{
			Handler:      r,
			Addr:         conf.Addr,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
		logic:  logic,
		Config: conf,
	}

	r.HandleFunc("/", s.health)
	r.HandleFunc("/question", s.question)
	return s
}

func (S *Server) health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func (S *Server) question(w http.ResponseWriter, r *http.Request) {
	question, err := S.logic.GetQuestion()
	if err != nil {
		fmt.Printf("error: could not get new question (%+v)\n", err)
		ServeError(w, NewErrorResponse(500, "Could not get new question"))
	}
	ServeJSON(w, *question)
}
