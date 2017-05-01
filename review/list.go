package review

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/itsubaki/apst/util"

	"google.golang.org/appengine"
)

func List(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ids := Kinds(ctx, "Review_")

	var page string

	switch r.URL.Query().Get("output") {
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		page = util.ToJson(ids)
	case "jsonp":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		page = util.ToJsonPretty(ids)
	default:
		for _, id := range ids {
			sid := strconv.Itoa(id)
			url := os.Getenv("external_url")
			href := "<a href=\"" + url + "/review/search?id=" + sid + "\">" + sid + "</a>"
			page = page + href + "<br>"
		}
	}
	fmt.Fprintf(w, page)

}
