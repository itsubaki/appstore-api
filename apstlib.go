package apstlib

import (
	"fmt"
	"net/http"
)

func AppEngine() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	url := RankingURL(200, "", "grossing", "jp")
	fmt.Fprint(w, url)
}
