package lang

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	language "cloud.google.com/go/language/apiv1"
	"github.com/itsubaki/appstore-api/util"
	"google.golang.org/appengine"
	pb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func Annotate(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := language.NewClient(ctx)
	if err != nil {
		fmt.Fprint(w, err.Error())
		log.Fatalf(err.Error())
		return
	}

	/*
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprint(w, err.Error())
			log.Fatalf(err.Error())
			return
		}
		text := string(b)
	*/

	text := r.URL.Query().Get("text")
	pretty := r.URL.Query().Get("pretty")
	syntax, entities, sentiment := parse(r.URL.Query())

	resp, err := client.AnnotateText(ctx, &pb.AnnotateTextRequest{
		Document: &pb.Document{
			Source: &pb.Document_Content{
				Content: text,
			},
			Type: pb.Document_PLAIN_TEXT,
		},
		Features: &pb.AnnotateTextRequest_Features{
			ExtractSyntax:            syntax,
			ExtractEntities:          entities,
			ExtractDocumentSentiment: sentiment,
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

func parse(values url.Values) (bool, bool, bool) {
	syntax := false
	entities := false
	sentiment := false

	if qs := values.Get("syntax"); qs != "" {
		syntax, _ = strconv.ParseBool(qs)
	}

	if qs := values.Get("entities"); qs != "" {
		entities, _ = strconv.ParseBool(qs)
	}

	if qs := values.Get("sentiment"); qs != "" {
		sentiment, _ = strconv.ParseBool(qs)
	}

	return syntax, entities, sentiment
}
