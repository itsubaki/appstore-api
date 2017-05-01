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

	genre, feed, country := util.Parse(r.URL.Query())
	url := util.RankingURL(200, genre, feed, country)
	log.Infof(ctx, url)

	b, err := util.Fetch(ctx, url)
	if err != nil {
		fmt.Fprint(w, err.Error()+"<br>")
		log.Warningf(ctx, err.Error())
		return
	}

	f := model.NewAppFeed(b)

	name := "Ranking_" + country + "_" + feed + "_" + genre
	Taskq(ctx, name, f)
}
