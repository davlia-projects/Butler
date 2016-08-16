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
	r.HandleFunc("/question", s.randomQuestion)

	qr := r.PathPrefix("/question").Subrouter()

	qr.HandleFunc("/random", s.randomQuestion)
	qr.HandleFunc("/add", s.addQuestion)
	qr.HandleFunc("/delete", s.deleteQuestion)
	return s
}

func (S *Server) health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func (S *Server) randomQuestion(w http.ResponseWriter, r *http.Request) {
	question, err := S.logic.GetQuestion()
	if err != nil || question == nil {
		fmt.Printf("error: could not get new question (%+v)\n", err)
		ServeError(w, NewErrorResponse(http.StatusInternalServerError, "Could not get new question"))
		return
	}
	ServeJSON(w, *question)
}

func (S *Server) addQuestion(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	prompt := params.Get("prompt")
	answer := params.Get("answer")
	category := params.Get("category")
	tags := params.Get("tags")
	if prompt == "" || answer == "" {
		ServeError(w, NewErrorResponse(http.StatusBadRequest, "bad input format"))
		return
	}
	err := S.logic.AddQuestion(prompt, answer, category, tags)
	if err != nil {
		ServeError(w, NewErrorResponse(http.StatusInternalServerError, "could not add question"))
		return
	}
	ServeJSON(w, "OK")
}

func (S *Server) deleteQuestion(w http.ResponseWriter, r *http.Request) {

}
