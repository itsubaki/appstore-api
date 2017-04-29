package review

import (
	"net/http"

	"github.com/itsubaki/apstlib"
	"github.com/itsubaki/apstlib/model"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func capture(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if len(r.Header.Get("X-Appengine-Cron")) == 0 {
		log.Warningf(ctx, "cron job only.")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Warningf(ctx, "query[\"id\"] is empty.")
		return
	}

	country := r.URL.Query().Get("country")
	if country == "" {
		country = "jp"
	}

	url := apstlib.ReviewURL(id, country)
	log.Infof(ctx, url)

	b, err := apstlib.Fetch(ctx, url)
	if err != nil {
		log.Warningf(ctx, err.Error())
		return
	}

	f := model.NewReviewFeed(b)
	log.Infof(ctx, f.Stats())
	for _, r := range f.ReviewList {
		log.Debugf(ctx, r.String())
	}

	//	Taskq(ctx, id, f)
}
