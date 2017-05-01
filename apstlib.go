package apstlib

import (
	"fmt"
	"net/http"

	"github.com/itsubaki/apstlib/ranking"
	"github.com/itsubaki/apstlib/review"
)

func AppEngine() {
	http.HandleFunc("/", root)

	http.HandleFunc("/ranking", ranking.Latest)
	http.HandleFunc("/ranking/capture", ranking.Capture)
	http.HandleFunc("/ranking/search", ranking.Search)

	http.HandleFunc("/review/capture", review.Capture)
	http.HandleFunc("/review/search", review.Search)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
