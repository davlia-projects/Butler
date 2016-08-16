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
	AddQuestion(q Question) error
}

type backend struct {
	Config config.Config
	db     *sql.DB
}

func NewBackend(conf config.Config) (Backend, error) {
	db, err := sql.Open("sqlite3", "./db/dota_trivia_bot.db")
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
	rows, err := B.db.Query("SELECT * FROM Questions ORDER BY RANDOM() LIMIT 1")
	if rows == nil {
		return nil, nil
	}
	var qid, prompt, answer, category, tags, date string
	var score int
	for rows.Next() {
		err = rows.Scan(&qid, &prompt, &answer, &category, &tags, &score, &date)
		if err != nil {
			fmt.Printf("error: could not scan from query (%+v)\n", err)
			return nil, err
		}
	}
	b := &Question{
		QID:      qid,
		Prompt:   prompt,
		Answer:   answer,
		Category: category,
		Tags:     tags,
		Date:     date,
		Score:    score,
	}
	return b, nil
}

func (B *backend) AddQuestion(q Question) error {
	cmd := "INSERT INTO Questions(prompt, answer, category, tags, score, created) values(?,?,?,?,?,CURRENT_DATE)"
	query, err := B.db.Prepare(cmd)
	if err != nil {
		fmt.Printf("error: could not prepare sqlite3 query (%+v)\n", err)
		return err
	}
	_, err = query.Exec(q.Prompt, q.Answer, q.Category, q.Tags, q.Score)
	if err != nil {
		fmt.Printf("error: could not execute sqlite3 query (%+v)\n", err)
		return err
	}
	return nil
}
