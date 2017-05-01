package review

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"google.golang.org/appengine"
)

func List(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ids := Kinds(ctx, "Review_")

	for _, id := range ids {
		sid := strconv.Itoa(id)
		url := os.Getenv("external_url")
		href := "<a href=\"" + url + "/review/search?id=" + sid + "\">" + sid + "</a><br>"
		fmt.Fprintf(w, href)
	}
}
