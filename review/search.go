package review

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/itsubaki/apstlib/util"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
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
	limit := util.Limit(r.URL.Query(), 50)
	name := "Review_" + id
	key := name + "_limit_" + strconv.Itoa(limit) + "_query_" + query

	var page string
	var cached bool
	var err error

	switch output {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key = key + "_json"
		if page, cached = util.MemGet(ctx, key); cached {
			break
		}

		list := IndexQuery(ctx, name, query, limit)
		page, err = util.Json(list)

	case "jsonp":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key = key + "_jsonp"
		if page, cached = util.MemGet(ctx, key); cached {
			break
		}

		list := IndexQuery(ctx, name, query, limit)
		page, err = util.Jsonp(list)

	default:
		key = key + "_html"
		if page, cached = util.MemGet(ctx, key); cached {
			page = "(cache)<br>" + page
			break
		}

		list := IndexQuery(ctx, name, query, limit)
		for _, r := range list {
			page = page + util.FontColor(r) + "<br>"
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
