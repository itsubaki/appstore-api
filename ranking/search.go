package ranking

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/itsubaki/appstore-api/appstoreurl"
	"github.com/itsubaki/appstore-api/cache"
	"github.com/itsubaki/appstore-api/format"
	"google.golang.org/appengine"
)

func Search(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	output := r.URL.Query().Get("format")
	pretty := r.URL.Query().Get("pretty")
	query := Query(r.URL.Query())
	limit := appstoreurl.Limit(r.URL.Query(), 50)
	genre, feed, country := appstoreurl.Parse(r.URL.Query())

	name := "Ranking_" + country + "_" + feed + "_" + genre
	keybase := name + "_limit_" + strconv.Itoa(limit) + "_query_" + query

	switch output {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key := keybase + "_json_pretty_" + pretty
		if cached, hit := cache.Get(ctx, key); hit {
			format.Print(ctx, w, cached, nil)
			return
		}

		list := IndexQuery(ctx, name, query, limit)
		page, err := format.Json(list, pretty)
		format.Print(ctx, w, page, err)
		cache.Put(ctx, key, page, 10*time.Minute)

	default:
		key := keybase + "_html"
		if cached, hit := cache.Get(ctx, key); hit {
			format.Print(ctx, w, "(cache)<br>"+cached, nil)
			return
		}

		list := IndexQuery(ctx, name, query, limit)
		page := ""
		for _, app := range list {
			page = page + app.String() + "<br>"
		}
		format.Print(ctx, w, page, nil)
		cache.Put(ctx, key, page, 10*time.Minute)
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
