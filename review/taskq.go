package review

import (
	"golang.org/x/net/context"

	"github.com/itsubaki/apstlib/model"
	"google.golang.org/appengine/capability"
	"google.golang.org/appengine/delay"
	"google.golang.org/appengine/log"
)

var indexPutDelay = delay.Func("indexput", IndexDiffPut)
var storePutDelay = delay.Func("storeput", Insert)

func Taskq(ctx context.Context, name string, feed *model.ReviewFeed) {
	if !capability.Enabled(ctx, "taskqueue", "*") {
		log.Warningf(ctx, "taskqueue is currently unavailable.")
		return
	}

	indexPutDelay.Call(ctx, name, feed)
	storePutDelay.Call(ctx, name, feed)
}
