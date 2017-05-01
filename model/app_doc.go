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

type AppDocList []AppDoc

func (app *AppDoc) String() string {
	return app.Timestamp.String() + " [" + app.Rank + "] " + app.Name + "(" + app.ID + ")" + " [" + app.Artist + "]"
}
