package model

type Question struct {
	QID      string
	Category string
	Prompt   string
	Answer   string
	Tags     string
	Date     string
	Score    int
}
