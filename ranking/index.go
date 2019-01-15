package ranking

import (
	"context"
	"strconv"
	"time"

	"github.com/itsubaki/appstore-api/model"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
)

func IndexQuery(ctx context.Context, name, query string, limit int) model.AppDocList {
	index, err := search.Open(name)
	if err != nil {
		log.Warningf(ctx, err.Error())
		return model.AppDocList{}
	}

	opt := search.SearchOptions{
		Limit: limit,
		//	Sort: &search.SortOptions{
		//		Expressions: []search.SortExpression{
		//			{Expr: "Timestamp", Reverse: false},
		//		},
		//	},
	}

	list := model.AppDocList{}
	for t := index.Search(ctx, query, &opt); ; {
		var doc model.AppDoc
		_, err := t.Next(&doc)

		if err == search.Done {
			break
		}

		if err != nil {
			log.Errorf(ctx, err.Error())
			return list
		}

		app := model.AppDoc{
			Rank:      doc.Rank,
			ID:        doc.ID,
			Name:      doc.Name,
			BundleID:  doc.BundleID,
			Rights:    doc.Rights,
			Artist:    doc.Artist,
			Timestamp: doc.Timestamp,
		}

		list = append(list, app)
	}

	return list
}

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
