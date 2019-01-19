package review

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/itsubaki/appstore-api/appstoreurl"
	"github.com/itsubaki/appstore-api/cache"
	"github.com/itsubaki/appstore-api/format"
	"google.golang.org/appengine"
)

func Search(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Fprintf(w, "query[id] is empty.")
		return
	}

	output := r.URL.Query().Get("output")
	query := r.URL.Query().Get("query")
	pretty := r.URL.Query().Get("pretty")
	limit := appstoreurl.Limit(r.URL.Query(), 50)

	name := "Review_" + id
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

		page := ""
		list := IndexQuery(ctx, name, query, limit)
		for _, r := range list {
			page = page + format.FontColor(r) + "<br>"
		}
		format.Print(ctx, w, page, nil)
		cache.Put(ctx, key, page, 10*time.Minute)
	}
}
