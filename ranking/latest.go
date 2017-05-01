package ranking

import (
	"fmt"
	"net/http"

	"github.com/itsubaki/apstlib/model"
	"github.com/itsubaki/apstlib/util"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func Latest(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	output := r.URL.Query().Get("output")
	query := r.URL.Query().Get("query")
	limit := util.Limit(r.URL.Query(), 20)
	genre, feed, country := util.Parse(r.URL.Query())

	url := util.RankingURL(limit, genre, feed, country)
	log.Infof(ctx, url)

	b, err := util.Fetch(ctx, url)
	if err != nil {
		fmt.Fprint(w, err.Error()+"<br>")
		log.Warningf(ctx, err.Error())
		return
	}

	f := model.NewAppFeed(b)
	list := f.Select(query)

	var content string
	switch output {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		content, err = util.Json(list)
	case "jsonp":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		content, err = util.Jsonp(list)
	default:
		for _, app := range list {
			content = content + app.String() + "<br>"
		}
	}

	if err != nil {
		log.Warningf(ctx, err.Error())
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, content)
	//IndexDrop(ctx, "Ranking_"+country+"_"+feed+"_"+genre)
}
