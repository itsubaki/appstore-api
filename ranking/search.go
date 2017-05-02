package ranking

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/itsubaki/apstlib/util"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func Search(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	output := r.URL.Query().Get("output")
	limit := util.Limit(r.URL.Query(), 50)
	query := Query(r.URL.Query())
	genre, feed, country := util.Parse(r.URL.Query())
	name := "Ranking_" + country + "_" + feed + "_" + genre
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
		for _, app := range list {
			page = page + app.String() + "<br>"
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

func Query(values url.Values) string {

	if id := values.Get("id"); len(id) != 0 {
		return id
	}

	if q := values.Get("query"); len(q) != 0 {
		return q
	}

	return ""
}
