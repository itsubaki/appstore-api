package util

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/itsubaki/apstlib/model"
)

func Limit(values url.Values, init int) int {
	input := values.Get("limit")
	if input == "" {
		return init
	}

	lmt, err := strconv.Atoi(input)
	if err != nil {
		lmt = init
	}

	if lmt < 3 {
		lmt = 3
	}

	return lmt
}

func Parse(values url.Values) (genre, feed, country string) {

	genre = model.Genre(values.Get("genre"))

	if country = values.Get("country"); country == "" {
		country = "jp"
	}

	if feed = values.Get("feed"); feed == "" {
		feed = "grossing"
	}

	return
}

func FontColor(r model.Review) string {
	color := "black"
	switch r.Rating {
	case "5":
		color = "green"
	case "4":
		color = "blue"
	case "2":
		color = "orange"
	case "1":
		color = "red"
	}
	return "<font color=\"" + color + "\">" + r.String() + "</font>"
}

func Json(in interface{}) (string, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func Jsonp(in interface{}) (string, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return "", err
	}

	var pretty bytes.Buffer
	err = json.Indent(&pretty, b, "", " ")
	if err != nil {
		return "", err
	}

	return string(pretty.Bytes()), nil
}
