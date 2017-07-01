package apstapi

import (
	"net/http"

	"github.com/itsubaki/apstapi/appinfo"
	"github.com/itsubaki/apstapi/lang"
	"github.com/itsubaki/apstapi/ranking"
	"github.com/itsubaki/apstapi/review"
)

func AppEngine() {
	http.HandleFunc("/app", appinfo.List)

	http.HandleFunc("/ranking", ranking.Latest)
	http.HandleFunc("/ranking/capture", ranking.Capture)
	http.HandleFunc("/ranking/search", ranking.Search)

	http.HandleFunc("/review", review.List)
	http.HandleFunc("/review/capture", review.Capture)
	http.HandleFunc("/review/search", review.Search)

	http.HandleFunc("/lang/annotate", lang.Annotate)
	http.HandleFunc("/lang/entities", lang.Entities)
	http.HandleFunc("/lang/sentiment", lang.Sentiment)
	http.HandleFunc("/lang/syntax", lang.Syntax)
}
