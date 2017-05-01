package model

import "time"

type AppDoc struct {
	Rank      string
	ID        string
	Name      string
	BundleID  string
	Rights    string
	Artist    string
	Timestamp time.Time
}
