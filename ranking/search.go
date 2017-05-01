package ranking

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/itsubaki/apstlib/model"
	"github.com/itsubaki/apstlib/util"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func Search(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	qlimit := r.URL.Query().Get("limit")
	if qlimit == "" {
		qlimit = "50"
	}
	limit, err := strconv.Atoi(qlimit)
	if err != nil {
		log.Warningf(ctx, err.Error())
		limit = 50
	}

	country := r.URL.Query().Get("country")
	if country == "" {
		country = "jp"
	}

	feed := r.URL.Query().Get("feed")
	if feed == "" {
		feed = "grossing"
	}

	genre := model.Genre(r.URL.Query().Get("genre"))
	query := r.URL.Query().Get("query")

	name := "Ranking_" + country + "_" + feed + "_" + genre
	key := name + "_limit_" + qlimit + "_query_" + query

	var page string
	var cached bool

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
