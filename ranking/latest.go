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
	pretty := r.URL.Query().Get("pretty")
	limit := util.Limit(r.URL.Query(), 20)
	genre, feed, country := util.Parse(r.URL.Query())

	url := util.RankingURL(limit, genre, feed, country)
	log.Infof(ctx, url)

	b, e := util.Fetch(ctx, url)
	if e != nil {
		fmt.Fprint(w, e.Error()+"<br>")
		log.Warningf(ctx, e.Error())
		return
	}

	f := model.NewAppFeed(b)
	list := f.Select(query)

	switch output {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		page, err := util.Json(list, pretty)
		util.Print(ctx, w, page, err)
	default:
		page := ""
		for _, app := range list {
			page = page + app.String() + "<br>"
		}
		util.Print(ctx, w, page, nil)
	}

	//IndexDrop(ctx, "Ranking_"+country+"_"+feed+"_"+genre)
}
