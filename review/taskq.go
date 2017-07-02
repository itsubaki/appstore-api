package review

import (
	"golang.org/x/net/context"

	"github.com/itsubaki/apstweb/model"
	"google.golang.org/appengine/capability"
	"google.golang.org/appengine/delay"
	"google.golang.org/appengine/log"
)

var indexPutDelay = delay.Func("indexput", IndexDiffPut)
var storePutDelay = delay.Func("storeput", Insert)
var storeAppInfoDelay = delay.Func("store_app_info", InsertAppInfo)

func Taskq(ctx context.Context, id, appname string, feed *model.ReviewFeed) {
	if !capability.Enabled(ctx, "taskqueue", "*") {
		log.Warningf(ctx, "taskqueue is currently unavailable.")
		return
	}

	name := "Review_" + id
	indexPutDelay.Call(ctx, name, feed)
	storePutDelay.Call(ctx, name, feed)
	storeAppInfoDelay.Call(ctx, id, appname)
}
