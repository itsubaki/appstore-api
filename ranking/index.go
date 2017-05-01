package ranking

import (
	"strconv"
	"time"

	"github.com/itsubaki/apstlib/model"
	"golang.org/x/net/context"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
)

func IndexDrop(ctx context.Context, name string) {
	index, err := search.Open(name)
	if err != nil {
		log.Warningf(ctx, err.Error())
		return
	}

	opt := search.SearchOptions{
		IDsOnly: true,
	}

	for t := index.Search(ctx, "", &opt); ; {
		var doc model.AppDoc
		docID, err := t.Next(&doc)

		if err == search.Done {
			break
		}

		if err != nil {
			log.Errorf(ctx, err.Error())
			continue
		}

		index.Delete(ctx, docID)
	}

}

func IndexPut(ctx context.Context, name string, feed *model.AppFeed) {
	index, err := search.Open(name)
	if err != nil {
		log.Warningf(ctx, err.Error())
		return
	}

	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)

	for _, app := range feed.AppList {

		doc := model.AppDoc{
			Rank:      strconv.Itoa(app.Rank),
			ID:        app.ID,
			Name:      app.Name,
			BundleID:  app.BundleID,
			Rights:    app.Rights,
			Artist:    app.Artist,
			Timestamp: now,
		}

		id := timestamp + "_" + app.ID
		_, err = index.Put(ctx, id, &doc)
		if err != nil {
			log.Errorf(ctx, err.Error())
			return
		}
	}

}
