package appinfo

import (
	"net/http"
	"strconv"
	"time"

	"github.com/itsubaki/apstlib/util"
	"google.golang.org/appengine"
)

func List(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	output := r.URL.Query().Get("output")
	pretty := r.URL.Query().Get("pretty")
	limit := util.Limit(r.URL.Query(), 200)

	kind := "AppInfo"
	keybase := kind + "_limit_" + strconv.Itoa(limit)

	switch output {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key := keybase + "_json_pretty_" + pretty
		if cached, hit := util.MemGet(ctx, key); hit {
			util.Print(ctx, w, cached, nil)
			return
		}

		list := Select(ctx, kind, limit)
		page, err := util.Json(list, pretty)
		util.Print(ctx, w, page, err)
		util.MemPut(ctx, key, page, 10*time.Minute)
	default:
		key := keybase + "_html"
		if cached, hit := util.MemGet(ctx, key); hit {
			util.Print(ctx, w, "(cache)<br>"+cached, nil)
			return
		}

		page := ""
		list := Select(ctx, kind, limit)
		for _, info := range list {
			page = page + info.String() + "<br>"
		}
		util.Print(ctx, w, page, nil)
		util.MemPut(ctx, key, page, 10*time.Minute)
	}
}
