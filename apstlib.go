package apstlib

import (
	"fmt"
	"net/http"

	"github.com/itsubaki/apstlib/ranking"
	"github.com/itsubaki/apstlib/review"
)

func AppEngine() {
	http.HandleFunc("/", root)

	http.HandleFunc("/ranking", ranking.Handle)

	http.HandleFunc("/capture/review", review.Capture)
	http.HandleFunc("/search/review", review.Search)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
