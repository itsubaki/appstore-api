package ranking

import (
	"fmt"
	"net/http"

	"github.com/itsubaki/appstore-api/appstoreurl"
	"github.com/itsubaki/appstore-api/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func Capture(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if len(r.Header.Get("X-Appengine-Cron")) == 0 {
		log.Warningf(ctx, "X-Appengine-Cron not found.")
		return
	}

	genre, feed, country := appstoreurl.Parse(r.URL.Query())
	url := appstoreurl.RankingURL(200, genre, feed, country)
	log.Infof(ctx, url)

	b, err := appstoreurl.Fetch(ctx, url)
	if err != nil {
		fmt.Fprint(w, err.Error()+"<br>")
		log.Warningf(ctx, err.Error())
		return
	}

	f := model.NewAppFeed(b)

	name := "Ranking_" + country + "_" + feed + "_" + genre
	Taskq(ctx, name, f)
}
