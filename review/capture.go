package review

import (
	"net/http"

	"github.com/itsubaki/apstlib/model"
	"github.com/itsubaki/apstlib/util"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func Capture(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if len(r.Header.Get("X-Appengine-Cron")) == 0 {
		log.Warningf(ctx, "X-Appengine-Cron not found.")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Warningf(ctx, "query[\"id\"] is empty.")
		return
	}

	_, _, country := util.Parse(r.URL.Query())
	url := util.ReviewURL(id, country)
	log.Infof(ctx, url)

	b, err := util.Fetch(ctx, url)
	if err != nil {
		log.Warningf(ctx, err.Error())
		return
	}

	f := model.NewReviewFeed(b)
	log.Infof(ctx, f.Stats())
	for _, r := range f.ReviewList {
		log.Debugf(ctx, r.String())
	}

	name := "Review_" + id
	Taskq(ctx, name, f)
}
