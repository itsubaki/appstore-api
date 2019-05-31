package ranking

import (
	"fmt"
	"net/http"

	"github.com/itsubaki/appstore-api/appstoreurl"
	"github.com/itsubaki/appstore-api/format"
	"github.com/itsubaki/appstore-api/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func Latest(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	output := r.URL.Query().Get("format")
	query := r.URL.Query().Get("query")
	pretty := r.URL.Query().Get("pretty")

	limit := appstoreurl.Limit(r.URL.Query(), 200)
	genre, feed, country := appstoreurl.Parse(r.URL.Query())
	url := appstoreurl.RankingURL(limit, genre, feed, country)
	log.Infof(ctx, url)

	b, e := appstoreurl.Fetch(ctx, url)
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
		page, err := format.Json(list, pretty)
		format.Print(ctx, w, page, err)
	default:
		page := ""
		for _, app := range list {
			page = page + app.String() + "<br>"
		}
		format.Print(ctx, w, page, nil)
	}

	//IndexDrop(ctx, "Ranking_"+country+"_"+feed+"_"+genre)
}
