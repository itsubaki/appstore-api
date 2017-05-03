package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/appengine/log"

	"github.com/itsubaki/apstlib/model"
)

func Print(ctx context.Context, w http.ResponseWriter, page string, err error) {
	if err != nil {
		log.Warningf(ctx, err.Error())
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, page)
}

func Limit(values url.Values, init int) int {
	input := values.Get("limit")
	if input == "" {
		return init
	}

	limit, err := strconv.Atoi(input)
	if err != nil {
		limit = init
	}

	if limit < 3 {
		limit = 3
	}

	return limit
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

func Json(in interface{}, pretty string) (string, error) {
	if pretty == "true" {
		return Jsonp(in)
	}

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
