package lang

import (
	"fmt"
	"log"
	"net/http"

	"github.com/itsubaki/appstore-api/format"

	language "cloud.google.com/go/language/apiv1"
	"google.golang.org/appengine"
	pb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func Syntax(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := language.NewClient(ctx)
	if err != nil {
		fmt.Fprint(w, err.Error())
		log.Fatalf(err.Error())
		return
	}

	text := r.URL.Query().Get("text")
	pretty := r.URL.Query().Get("pretty")

	resp, err := client.AnalyzeSyntax(ctx, &pb.AnalyzeSyntaxRequest{
		Document: &pb.Document{
			Source: &pb.Document_Content{
				Content: text,
			},
			Type: pb.Document_PLAIN_TEXT,
		},
		EncodingType: pb.EncodingType_UTF8,
	})

	if err != nil {
		fmt.Fprint(w, err.Error())
		log.Fatalf(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	page, err := format.Json(resp, pretty)
	format.Print(ctx, w, page, err)
}
