package model

import "strings"

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

func (r *ReviewDoc) String() string {
	content := strings.Replace(r.Content, "\n", " ", -1)
	return "[" + r.Rating + "][" + r.Title + "] " + content + "/" + r.Author + "(" + r.Updated + ")"
}
