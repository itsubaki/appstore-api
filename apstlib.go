package apstlib

import (
	"net/http"

	"github.com/itsubaki/apstlib/ranking"
	"github.com/itsubaki/apstlib/review"
)

func AppEngine() {
	http.HandleFunc("/ranking", ranking.Latest)
	http.HandleFunc("/ranking/capture", ranking.Capture)
	http.HandleFunc("/ranking/search", ranking.Search)

	http.HandleFunc("/review", review.List)
	http.HandleFunc("/review/capture", review.Capture)
	http.HandleFunc("/review/search", review.Search)
}
