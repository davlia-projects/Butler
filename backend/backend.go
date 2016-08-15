package backend

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/question-service/config"
	. "github.com/question-service/models"
)

type Backend interface {
	GetRandomQuestion() (*Question, error)
}

type backend struct {
	Config config.Config
	db     *sql.DB
}

func NewBackend(conf config.Config) (Backend, error) {
	db, err := sql.Open("sqlite3", "./questions.db")
	if err != nil {
		fmt.Printf("error: could not open db connection (%+v)\n", err)
	}
	b := &backend{
		Config: conf,
		db:     db,
	}
	return b, nil
}

func (B *backend) GetRandomQuestion() (*Question, error) {
	rows, err := B.db.Query("SELECT * FROM table ORDER BY RANDOM() LIMIT 1")
	var prompt, answer, category string
	for rows.Next() {
		err = rows.Scan(&prompt, &answer, &category)
		if err != nil {
			fmt.Printf("error: could not scan from query (%+v)\n", err)
			return nil, err
		}
	}
	b := &Question{
		Prompt:   prompt,
		Answer:   answer,
		Category: category,
	}
	return b, nil
}
