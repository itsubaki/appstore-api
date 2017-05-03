package ranking

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/itsubaki/apstlib/util"

	"google.golang.org/appengine"
)

func Search(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	output := r.URL.Query().Get("output")
	pretty := r.URL.Query().Get("pretty")
	query := Query(r.URL.Query())
	limit := util.Limit(r.URL.Query(), 50)
	genre, feed, country := util.Parse(r.URL.Query())

	name := "Ranking_" + country + "_" + feed + "_" + genre
	keybase := name + "_limit_" + strconv.Itoa(limit) + "_query_" + query

	switch output {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key := keybase + "_json_pretty_" + pretty
		if cached, hit := util.MemGet(ctx, key); hit {
			util.Print(ctx, w, cached, nil)
			return
		}

		list := IndexQuery(ctx, name, query, limit)
		page, err := util.Json(list, pretty)
		util.Print(ctx, w, page, err)
		util.MemPut(ctx, key, page, 10*time.Minute)

	default:
		key := keybase + "_html"
		if cached, hit := util.MemGet(ctx, key); hit {
			util.Print(ctx, w, "(cache)<br>"+cached, nil)
			return
		}

		list := IndexQuery(ctx, name, query, limit)
		page := ""
		for _, app := range list {
			page = page + app.String() + "<br>"
		}
		util.Print(ctx, w, page, nil)
		util.MemPut(ctx, key, page, 10*time.Minute)
	}

}

func Query(values url.Values) string {

	if id := values.Get("id"); len(id) != 0 {
		return id
	}

	if q := values.Get("query"); len(q) != 0 {
		return q
	}

	return ""
}
