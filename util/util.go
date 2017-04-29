package apstweb

import (
	"bytes"
	"encoding/json"

	"github.com/itsubaki/apstlib/model"
)

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

func ToJson(in interface{}) (string, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ToJsonPretty(in interface{}) (string, error) {
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
