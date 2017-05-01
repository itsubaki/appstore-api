package ranking

import (
	"fmt"
	"net/http"

	"github.com/itsubaki/apstlib/model"
	"github.com/itsubaki/apstlib/util"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func Capture(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	country := r.URL.Query().Get("country")
	if country == "" {
		country = "jp"
	}

	feed := r.URL.Query().Get("feed")
	if feed == "" {
		feed = "grossing"
	}

	genre := model.Genre(r.URL.Query().Get("genre"))

	url := util.RankingURL(200, genre, feed, country)
	log.Infof(ctx, url)

	b, err := util.Fetch(ctx, url)
	if err != nil {
		fmt.Fprint(w, err.Error()+"<br>")
		log.Warningf(ctx, err.Error())
		return
	}

	name := "Ranking_" + country + "_" + feed + "_" + genre
	f := model.NewAppFeed(b)

	Taskq(ctx, name, f)
}