package apstlib

import (
	"errors"
	"io/ioutil"
	"strconv"

	"golang.org/x/net/context"

	"google.golang.org/appengine/capability"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func ReviewURL(id, country string) string {
	return "https://itunes.apple.com/" + country + "/rss/customerreviews/id=" + id + "/sortBy=mostRecent/xml"
}

func RankingURL(limit int, genre, feed, country string) string {
	var slimit = strconv.Itoa(limit)
	var url = "https://itunes.apple.com/" + country + "/rss/top" + feed + "applications/limit=" + slimit
	if genre != "" {
		url = url + "/genre=" + genre
	}
	url = url + "/json"
	return url
}

func Fetch(ctx context.Context, url string) ([]byte, error) {
	if !capability.Enabled(ctx, "urlfetch", "*") {
		log.Warningf(ctx, "urlfetch is currently unavailable.")
		return nil, errors.New("urlfetch is currently unavailable.")
	}

	resp, err := urlfetch.Client(ctx).Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
