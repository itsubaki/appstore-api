package format

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/itsubaki/appstore-api/model"

	"google.golang.org/appengine/log"
)

func Print(ctx context.Context, w http.ResponseWriter, page string, err error) {
	if err != nil {
		log.Warningf(ctx, err.Error())
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, page)
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
