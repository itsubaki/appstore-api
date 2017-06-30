package model

import "strconv"

type AppInfo struct {
	ID   string
	Name string
}

type AppInfoList []AppInfo

func (app *AppInfo) String() string {
	return app.Name + "(" + app.ID + ")"
}

func (s AppInfoList) Len() int {
	return len(s)
}

func (s AppInfoList) Less(i, j int) bool {
	left, _ := strconv.Atoi(s[i].ID)
	right, _ := strconv.Atoi(s[j].ID)
	return left < right
}

func (s AppInfoList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
