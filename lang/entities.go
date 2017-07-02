package lang

import (
	"fmt"
	"log"
	"net/http"

	"github.com/itsubaki/apstweb/util"

	language "cloud.google.com/go/language/apiv1"
	"google.golang.org/appengine"
	pb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func Entities(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := language.NewClient(ctx)
	if err != nil {
		fmt.Fprint(w, err.Error())
		log.Fatalf(err.Error())
		return
	}

	text := r.URL.Query().Get("text")
	pretty := r.URL.Query().Get("pretty")

	resp, err := client.AnalyzeEntities(ctx, &pb.AnalyzeEntitiesRequest{
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
	page, err := util.Json(resp, pretty)
	util.Print(ctx, w, page, err)
}
