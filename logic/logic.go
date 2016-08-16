package logic

import (
	"github.com/question-service/backend"
	"github.com/question-service/config"
	. "github.com/question-service/models"
)

type Logic interface {
	GetQuestion() (*Question, error)
	AddQuestion(prompt, answer, category, tags string) error
}
type logic struct {
	backend backend.Backend
	Config  config.Config
}

func NewLogic(conf config.Config, backend backend.Backend) (Logic, error) {
	l := &logic{
		backend: backend,
		Config:  conf,
	}
	return l, nil
}

func (L *logic) GetQuestion() (*Question, error) {
	question, err := L.backend.GetRandomQuestion()
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (L *logic) AddQuestion(prompt, answer, category, tags string) error {
	q := Question{
		Prompt:   prompt,
		Answer:   answer,
		Category: category,
		Tags:     tags,
	}
	err := L.backend.AddQuestion(q)
	return err
}
