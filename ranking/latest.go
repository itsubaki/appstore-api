package ranking

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/itsubaki/apstlib/model"
	"github.com/itsubaki/apstlib/util"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func Latest(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	qlimit := r.URL.Query().Get("limit")
	if qlimit == "" {
		qlimit = "20"
	}
	limit, err := strconv.Atoi(qlimit)
	if err != nil {
		log.Warningf(ctx, err.Error())
		limit = 20
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

	url := util.RankingURL(limit, genre, feed, country)
	log.Infof(ctx, url)

	b, err := util.Fetch(ctx, url)
	if err != nil {
		fmt.Fprint(w, err.Error()+"<br>")
		log.Warningf(ctx, err.Error())
		return
	}

	query := r.URL.Query().Get("query")
	f := model.NewAppFeed(b)
	list := f.Select(query)

	var content string
	switch r.URL.Query().Get("output") {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		content, err = util.Json(list)
	case "jsonp":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		content, err = util.Jsonp(list)
	default:
		for _, app := range list {
			content = content + app.String() + "<br>"
		}
	}

	if err != nil {
		log.Warningf(ctx, err.Error())
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, content)
	//IndexDrop(ctx, "Ranking_"+country+"_"+feed+"_"+genre)
}
