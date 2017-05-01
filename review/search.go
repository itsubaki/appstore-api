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

	qlimit := r.URL.Query().Get("limit")
	if qlimit == "" {
		qlimit = "50"
	}
	limit, err := strconv.Atoi(qlimit)
	if err != nil {
		log.Warningf(ctx, err.Error())
		limit = 50
	}

	query := r.URL.Query().Get("query")

	var page string
	var cached bool

	name := "Review_" + id
	key := name + "_limit_" + qlimit + "_query_" + query

	switch r.URL.Query().Get("output") {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key = key + "_json"
		page, cached = util.MemGet(ctx, key)

		if cached {
			break
		}

		list := IndexQuery(ctx, name, query, limit)
		page, err = util.ToJson(list)

	case "jsonp":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key = key + "_jsonp"
		page, cached = util.MemGet(ctx, key)

		if cached {
			break
		}

		list := IndexQuery(ctx, name, query, limit)
		page, err = util.ToJsonPretty(list)

	default:
		key = key + "_html"
		page, cached = util.MemGet(ctx, key)

		if cached {
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
