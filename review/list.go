package review

import (
	"net/http"
	"strconv"
	"time"

	"github.com/itsubaki/appstore-api/cache"
	"github.com/itsubaki/appstore-api/format"
	"google.golang.org/appengine"
)

func List(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	output := r.URL.Query().Get("output")
	pretty := r.URL.Query().Get("pretty")

	switch output {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key := "Review_list_json_pretty_" + pretty
		if cached, hit := cache.Get(ctx, key); hit {
			format.Print(ctx, w, cached, nil)
			return
		}

		ids := Kinds(ctx, "Review_")
		page, err := format.Json(ids, pretty)
		format.Print(ctx, w, page, err)
		cache.Put(ctx, key, page, 10*time.Minute)
	default:
		key := "Review_list_html"
		if cached, hit := cache.Get(ctx, key); hit {
			format.Print(ctx, w, "(cache)<br>"+cached, nil)
			return
		}

		page := ""
		ids := Kinds(ctx, "Review_")
		for _, id := range ids {
			page = page + strconv.Itoa(id) + "<br>"
		}
		format.Print(ctx, w, page, nil)
		cache.Put(ctx, key, page, 10*time.Minute)
	}

}
