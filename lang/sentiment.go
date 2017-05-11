package lang

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/itsubaki/apstlib/util"

	language "cloud.google.com/go/language/apiv1"
	"google.golang.org/appengine"
	pb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func Sentiment(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := language.NewClient(ctx)
	if err != nil {
		fmt.Fprint(w, err.Error())
		log.Fatalf(err.Error())
		return
	}

	text := r.URL.Query().Get("text")
	pretty := r.URL.Query().Get("pretty")

	resp, err := client.AnalyzeSentiment(ctx, &pb.AnalyzeSentimentRequest{
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

	seed := text + "_pretty_" + pretty
	hash := sha256.Sum256([]byte(seed))
	key := string(hash[:16])
	if cached, hit := util.MemGet(ctx, key); hit {
		util.Print(ctx, w, cached, nil)
		return
	}

	page, err := util.Json(resp, pretty)
	util.Print(ctx, w, page, err)
	util.MemPut(ctx, key, page, 10*time.Minute)
}
