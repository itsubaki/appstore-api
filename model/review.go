package model

import (
	"encoding/json"
	"strings"
)

type ReviewDoc struct {
	Title     string
	Content   string
	Author    string
	Rating    string
	Version   string
	Updated   string
	VoteSum   string
	VoteCount string
}

type Review struct {
	ID        string `xml:"id"`
	Title     string `xml:"title"`
	Content   string `xml:"content"`
	Author    string `xml:"author>name"`
	Rating    string `xml:"rating"`
	Version   string `xml:"version"`
	Updated   string `xml:"updated"`
	VoteSum   string `xml:"voteSum"`
	VoteCount string `xml:"voteCount"`
}

func (r *Review) Json() (string, error) {
	b, err := json.Marshal(r)
	return string(b), err
}

func (r *Review) String() string {
	content := strings.Replace(r.Content, "\n", " ", -1)
	return "[" + r.Rating + "][" + r.Title + "] " + content + "/" + r.Author + "(" + r.Updated + ")"
}

func (r *ReviewDoc) String() string {
	content := strings.Replace(r.Content, "\n", " ", -1)
	return "[" + r.Rating + "][" + r.Title + "] " + content + "/" + r.Author + "(" + r.Updated + ")"
}
