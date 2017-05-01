package review

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/itsubaki/apstlib/util"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func List(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var key string
	var page string
	var cached bool
	var err error

	switch r.URL.Query().Get("output") {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key = "Review_list_json"
		if page, cached = util.MemGet(ctx, key); cached {
			break
		}

		ids := Kinds(ctx, "Review_")
		page, err = util.Json(ids)
	case "jsonp":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key = "Review_list_jsonp"
		if page, cached = util.MemGet(ctx, key); cached {
			break
		}

		ids := Kinds(ctx, "Review_")
		page, err = util.Jsonp(ids)
	default:
		key = "Review_list_html"
		if page, cached = util.MemGet(ctx, key); cached {
			page = "(cache)<br>" + page
			break
		}

		ids := Kinds(ctx, "Review_")
		for _, id := range ids {
			sid := strconv.Itoa(id)
			url := os.Getenv("external_url")
			href := "<a href=\"" + url + "/review/search?id=" + sid + "\">" + sid + "</a>"
			page = page + href + "<br>"
		}
	}

	if err != nil {
		log.Warningf(ctx, err.Error())
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, page)
	if !cached {
		util.MemPut(ctx, key, page, 10*time.Minute)
	}

}
