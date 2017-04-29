package model

import (
	"strconv"
	"strings"
)

type App struct {
	Rank     int
	ID       string
	Name     string
	BundleID string
	Rights   string
	Artist   string
}

func (app *App) String() string {
	rank := strconv.Itoa(app.Rank)
	id := app.ID
	appName := app.Name
	artist := app.Artist

	return rank + ": " + appName + "(" + id + ")" + " [" + artist + "]"
}

func NewApp(content interface{}, rank int) *App {
	return &App{
		Rank:     rank,
		Artist:   artist(content),
		Name:     name(content),
		BundleID: bundleID(content),
		ID:       appID(content),
		Rights:   rights(content),
	}
}

func (app *App) Contains(keyword string) bool {
	key := strings.ToLower(keyword)
	artist := strings.ToLower(app.Artist)
	name := strings.ToLower(app.Name)
	bundleID := strings.ToLower(app.BundleID)
	rights := strings.ToLower(app.Rights)

	if strings.Contains(artist, key) {
		return true
	}

	if strings.Contains(name, key) {
		return true
	}

	if strings.Contains(bundleID, key) {
		return true
	}

	if strings.Contains(rights, key) {
		return true
	}

	return false
}

func artist(content interface{}) string {
	artist := content.(map[string]interface{})["im:artist"]
	artistlabel := artist.(map[string]interface{})["label"]
	return artistlabel.(string)
}

func name(content interface{}) string {
	name := content.(map[string]interface{})["im:name"]
	namelabel := name.(map[string]interface{})["label"]
	return namelabel.(string)
}

func bundleID(content interface{}) string {
	id := content.(map[string]interface{})["id"]
	attributes := id.(map[string]interface{})["attributes"]
	bundleID := attributes.(map[string]interface{})["im:bundleId"]
	return bundleID.(string)
}

func appID(content interface{}) string {
	id := content.(map[string]interface{})["id"]
	attributes := id.(map[string]interface{})["attributes"]
	imid := attributes.(map[string]interface{})["im:id"]
	return imid.(string)
}

func rights(content interface{}) string {
	rights := content.(map[string]interface{})["rights"]
	rightslabel := rights.(map[string]interface{})["label"]
	return rightslabel.(string)
}
