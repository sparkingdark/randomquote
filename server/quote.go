package server

import (
	"log"
)

type Quote struct {
	Quote string `json:"quote"`
	Author string `json:"author"`
}

func(q *Quote) PostQuote() string{
	log.Println("PostQuote")
	Casewise(q, "insert")

	return "Quote Posted"
}
