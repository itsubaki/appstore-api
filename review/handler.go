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

func Handle(w http.ResponseWriter, r *http.Request) {
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

	query := ""
	content := r.URL.Query().Get("content")
	if content != "" {
		query = "Content=" + content
	}

	var page string
	var key string
	var cached bool

	switch r.URL.Query().Get("output") {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key = id + "_limit_" + qlimit + "_content_" + content + "_json"
		page, cached = util.MemGet(ctx, key)
		if cached {
			break
		}

		list := IndexQuery(ctx, id, query, limit)
		page, err = util.ToJson(list)

	case "jsonp":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		key = id + "_limit_" + qlimit + "_content_" + content + "_jsonp"
		page, cached = util.MemGet(ctx, key)
		if cached {
			break
		}

		list := IndexQuery(ctx, id, query, limit)
		page, err = util.ToJsonPretty(list)

	default:
		key = id + "_limit_" + qlimit + "_content_" + content + "_html"
		page, cached = util.MemGet(ctx, key)
		if cached {
			page = "(cache)<br>" + page
			break
		}

		list := IndexQuery(ctx, id, query, limit)
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
