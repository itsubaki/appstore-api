package review

import (
	"net/http"
	"strconv"
	"time"

	"github.com/itsubaki/apstapi/util"

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
		if cached, hit := util.MemGet(ctx, key); hit {
			util.Print(ctx, w, cached, nil)
			return
		}

		ids := Kinds(ctx, "Review_")
		page, err := util.Json(ids, pretty)
		util.Print(ctx, w, page, err)
		util.MemPut(ctx, key, page, 10*time.Minute)
	default:
		key := "Review_list_html"
		if cached, hit := util.MemGet(ctx, key); hit {
			util.Print(ctx, w, "(cache)<br>"+cached, nil)
			return
		}

		page := ""
		ids := Kinds(ctx, "Review_")
		for _, id := range ids {
			page = page + strconv.Itoa(id) + "<br>"
		}
		util.Print(ctx, w, page, nil)
		util.MemPut(ctx, key, page, 10*time.Minute)
	}

}
