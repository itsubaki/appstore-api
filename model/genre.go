package model

var m = map[string]string{
	"business":         "6000",
	"weather":          "6001",
	"utilities":        "6002",
	"travel":           "6003",
	"sports":           "6004",
	"socialnetworking": "6005",
	"reference":        "6006",
	"productivity":     "6007",
	"photo_video":      "6008",
	"news":             "6009",
	"navigation":       "6010",
	"music":            "6011",
	"lifestyle":        "6012",
	"health_fitness":   "6013",
	"games":            "6014",
	"finance":          "6015",
	"entertainment":    "6016",
	"education":        "6017",
	"books":            "6018",
	"medical":          "6020",
	"newsstand":        "6021",
	"catalogs":         "6022",
	"food_drink":       "6023",
}

func Genre(name string) string {
	v, ok := m[name]
	if ok {
		return v
	}
	return ""
}
