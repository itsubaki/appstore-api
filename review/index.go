package review

import (
	"github.com/itsubaki/apstlib/model"
	"golang.org/x/net/context"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
)

func IndexQuery(ctx context.Context, id, query string, limit int) model.ReviewList {
	index, err := search.Open("Review_" + id)
	if err != nil {
		log.Warningf(ctx, err.Error())
		return model.ReviewList{}
	}

	opt := search.SearchOptions{
		Limit: limit,
		Sort: &search.SortOptions{
			Expressions: []search.SortExpression{
				{Expr: "Updated", Reverse: false},
			},
		},
	}

	list := model.ReviewList{}
	for t := index.Search(ctx, query, &opt); ; {
		var doc model.ReviewDoc
		docID, err := t.Next(&doc)

		if err == search.Done {
			break
		}

		if err != nil {
			log.Errorf(ctx, err.Error())
			return list
		}

		r := model.Review{
			ID:        docID,
			Title:     doc.Title,
			Content:   doc.Content,
			Author:    doc.Author,
			Rating:    doc.Rating,
			Version:   doc.Version,
			Updated:   doc.Updated,
			VoteSum:   doc.VoteSum,
			VoteCount: doc.VoteCount,
		}
		list = append(list, r)
	}

	return list
}

func IndexPut(ctx context.Context, id string, feed *model.ReviewFeed) {
	index, err := search.Open("Review_" + id)
	if err != nil {
		log.Warningf(ctx, err.Error())
		return
	}

	for _, r := range feed.ReviewList {
		doc := model.ReviewDoc{
			Title:     r.Title,
			Content:   r.Content,
			Author:    r.Author,
			Rating:    r.Rating,
			Version:   r.Version,
			Updated:   r.Updated,
			VoteSum:   r.VoteSum,
			VoteCount: r.VoteCount,
		}

		_, err = index.Put(ctx, r.ID, &doc)
		if err != nil {
			log.Errorf(ctx, err.Error())
			return
		}
	}

}
