package appinfo

import (
	"net/http"
	"strconv"
	"time"

	"github.com/itsubaki/appstore-api/appstoreurl"
	"github.com/itsubaki/appstore-api/cache"
	"github.com/itsubaki/appstore-api/format"
	"google.golang.org/appengine"
)

func List(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	output := r.URL.Query().Get("output")
	pretty := r.URL.Query().Get("pretty")
	limit := appstoreurl.Limit(r.URL.Query(), 200)

	kind := "AppInfo"
	keybase := kind + "_limit_" + strconv.Itoa(limit)

	switch output {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key := keybase + "_json_pretty_" + pretty
		if cached, hit := cache.Get(ctx, key); hit {
			format.Print(ctx, w, cached, nil)
			return
		}

		list := Select(ctx, kind, limit)
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
		list := Select(ctx, kind, limit)
		for _, info := range list {
			page = page + info.String() + "<br>"
		}
		format.Print(ctx, w, page, nil)
		cache.Put(ctx, key, page, 10*time.Minute)
	}
}
